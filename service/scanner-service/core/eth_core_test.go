package core

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
	"math/big"
	"portto-homework/internal/ethclient"
	"portto-homework/internal/utils/logger"
	"portto-homework/service/scanner-service/config"
	"portto-homework/service/scanner-service/repository"
	"portto-homework/service/scanner-service/repository/db"
	"sync"
	"testing"
	"time"
)

func TestETHCore_TestSuite(t *testing.T) {
	if err := config.LoadConfig("./../../../conf.d/scanner-service-config.yaml"); err != nil {
		panic(err)
	}
	logger.InitSysLog("eth core", "debug")
	suite.Run(t, new(ethCore_TestSuite))
}

type ethCore_TestSuite struct {
	suite.Suite
	ctx     context.Context
	log     logger.LoggerInterface
	ethCore *ethCore
}

func (suite *ethCore_TestSuite) SetupTest() {
	eth, err := ethclient.Connect("https://data-seed-prebsc-2-s3.binance.org:8545/")
	if err != nil {
		panic(err)
	}

	repo := repository.NewDBRepo(repository.DBRepoIn{
		Log: logger.SysLog(),
	})

	suite.ctx = context.Background()
	suite.log = logger.SysLog()
	dbClient := db.NewDBClient()
	suite.NoError(dbClient.InitDBTable())

	suite.ethCore = &ethCore{
		in: CoreIn{
			Conf:            config.GetConfig(),
			EthClient:       eth,
			DB:              dbClient,
			BlockRepo:       repo.BlocksRepo,
			TransactionRepo: repo.TransactionsRepo,
		},
		realTimeWg: &sync.WaitGroup{},
		historyWg:  &sync.WaitGroup{},
		realTimeCh: make(chan *types.Block, config.GetScannerConfig().PipelineNumber),
		historyCh:  make(chan *types.Block, config.GetScannerConfig().PipelineNumber),
		subFrom:    0,
	}

	suite.ethCore.run(suite.ctx)
}

func (suite *ethCore_TestSuite) TearDownTest() {
	suite.ethCore.close()
}

func (suite *ethCore_TestSuite) TestETHCore_StorageBlockNumberList_OK() {
	num, err := suite.ethCore.in.EthClient.BlockNumber(suite.ctx)
	suite.NoError(err)

	latest := big.Int{}
	from := big.Int{}
	if suite.ethCore.in.Conf.ScannerConfig.ScanBlockFrom < 0 {
		from.Sub(latest.SetUint64(num), big.NewInt(int64(-suite.ethCore.in.Conf.ScannerConfig.ScanBlockFrom)))
	} else {
		from.SetUint64(uint64(suite.ethCore.in.Conf.ScannerConfig.ScanBlockFrom))
	}

	suite.NoError(suite.ethCore.ScanBlockDataFromNum(suite.ctx, from.Uint64()))
}

func (suite *ethCore_TestSuite) TestETHCore_Subscribe_OK() {
	ctx, cancel := context.WithCancel(suite.ctx)
	suite.ethCore.SubScribe(ctx)

	time.Sleep(20 * time.Second)
	cancel()
}
