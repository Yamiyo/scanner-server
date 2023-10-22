package core

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
	"portto-homework/internal/ethclient"
	"portto-homework/internal/utils/logger"
	"portto-homework/service/scanner-service/config"
	"portto-homework/service/scanner-service/repository"
	"portto-homework/service/scanner-service/repository/db"
	"sync"
	"testing"
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
		wg:      &sync.WaitGroup{},
		parseCh: make(chan *types.Block, config.GetScannerConfig().PipelineNumber),
	}

	suite.ethCore.run(suite.ctx)
}

func (suite *ethCore_TestSuite) TearDownTest() {
	suite.ethCore.close()
}

func (suite *ethCore_TestSuite) TestETHCore_StorageBlockNumberList_OK() {
	latest, err := suite.ethCore.in.EthClient.BlockNumber(suite.ctx)

	suite.NoError(err)
	suite.NoError(suite.ethCore.ScanBlockDataFromNum(suite.ctx, latest-50))
}
