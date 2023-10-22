package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"portto-homework/service/api-service/repository/db"
	"sync"
	"syscall"
	"time"

	"portto-homework/internal/utils/logger"
	"portto-homework/service/api-service/app/rest"
	"portto-homework/service/api-service/config"
)

var (
	templateApp   *TemplateApp
	serverSetOnce sync.Once
	stop          = make(chan error, 1)
	wg            = &sync.WaitGroup{}
)

type TemplateApp struct {
	RestServer rest.ServiceInterface
	conf       config.ConfigSetup
}

func newTemplateServer(ctx context.Context, conf config.ConfigSetup) {
	serverSetOnce.Do(func() {
		templateApp = &TemplateApp{
			RestServer: rest.NewRestService(ctx, conf),
			conf:       conf,
		}
	})
}

func Run() {
	// init config from ./config/api-service-config.yaml
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

	// run  restful api service
	wg.Add(1)
	go func(w *sync.WaitGroup) {
		defer w.Done()
		templateApp.RestServer.Run(ctx, stop)
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
