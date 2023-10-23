package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"portto-homework/internal/ethclient"
	"portto-homework/service/scanner-service/core"
	"portto-homework/service/scanner-service/repository"
	"portto-homework/service/scanner-service/repository/db"
	"sync"
	"syscall"
	"time"

	"portto-homework/internal/utils/logger"
	"portto-homework/service/scanner-service/config"
)

var (
	templateApp   *TemplateApp
	serverSetOnce sync.Once
	stop          = make(chan error, 1)
	wg            = &sync.WaitGroup{}
)

type TemplateApp struct {
	conf    config.ConfigSetup
	ethCore core.ETHCore
}

func newTemplateServer(ctx context.Context, conf config.ConfigSetup) {
	serverSetOnce.Do(func() {
		eth, err := ethclient.Connect(conf.ScannerConfig.Endpoint)
		if err != nil {
			panic(err)
		}

		repo := repository.NewDBRepo(repository.DBRepoIn{})

		job := core.New(core.CoreIn{
			Conf:            conf,
			EthClient:       eth,
			DB:              db.NewDBClient(),
			BlockRepo:       repo.BlocksRepo,
			TransactionRepo: repo.TransactionsRepo,
		})

		templateApp = &TemplateApp{
			conf:    conf,
			ethCore: job.ETHCore,
		}
	})
}

func Run() {
	// init config from ./config/scanner-service-config.yaml
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	// init logger
	if err := logger.InitSysLog(
		config.GetServerConfig().Name,
		config.GetServerConfig().Level); err != nil {
		panic(err)
	}

	// init db
	if err := db.NewDBClient().InitDBTable(); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	// init application server
	if newTemplateServer(ctx, config.GetConfig()); templateApp == nil {
		panic(errors.New("templateApp is nil"))
	}

	// graceful shutdown
	defer func() {
		templateApp.gracefulShutdownAPP(ctx, cancel)
	}()

	wg.Add(1)
	go func(w *sync.WaitGroup) {
		defer w.Done()
		defer templateApp.ethCore.Close()

		templateApp.ethCore.SubScribe(ctx)
		go func() {
			from := uint64(0)
			if templateApp.conf.ScannerConfig.ScanBlockFrom < 0 {
				from = templateApp.ethCore.GetSubscribeFrom() - uint64(-templateApp.conf.ScannerConfig.ScanBlockFrom)
			} else {
				from = uint64(templateApp.conf.ScannerConfig.ScanBlockFrom)
			}

			if err := templateApp.ethCore.ScanBlockDataFromNum(ctx, from); err != nil {
				logger.SysLog().Error(ctx, err.Error())
				stop <- err
			}
		}()

		<-ctx.Done()
	}(wg)

	// listen system signal
	wg.Add(1)
	go func(w *sync.WaitGroup) {
		defer w.Done()
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		logger.SysLog().Info(ctx, fmt.Sprintf("system stop event signal: %s", <-quit))
		stop <- nil
	}(wg)

	<-stop
}

func (a *TemplateApp) gracefulShutdownAPP(ctx context.Context, cancel context.CancelFunc) {
	cancel()
	c := make(chan struct{})
	go func() {
		wg.Wait()
		c <- struct{}{}
	}()

	select {
	case <-c:
		return
	case <-time.After(time.Duration(a.conf.ServerConfig.ShutdownTimeout) * time.Second):
		logger.SysLog().Error(ctx, "graceful shutdown timeout")
		return
	}
}
