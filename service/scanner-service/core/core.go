package core

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"portto-homework/service/scanner-service/config"
	"portto-homework/service/scanner-service/repository"
	"portto-homework/service/scanner-service/repository/db"
	"sync"
)

var (
	self *mongoRepo
	once sync.Once
)

type mongoRepo struct {
	in CoreIn
	CoreOut
}

type CoreIn struct {
	Conf      config.ConfigSetup
	EthClient *ethclient.Client
	DB        *db.DBClient

	BlockRepo       repository.BlocksRepo
	TransactionRepo repository.TransactionsRepo
}

type CoreOut struct {
}

func New(in CoreIn) CoreOut {
	once.Do(func() {
		self = &mongoRepo{
			in:      in,
			CoreOut: CoreOut{},
		}
	})

	return self.CoreOut
}