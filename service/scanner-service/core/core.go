package core

import (
	"sync"

	"portto-homework/service/scanner-service/config"
	"portto-homework/service/scanner-service/repository"
	"portto-homework/service/scanner-service/repository/db"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	self *core
	once sync.Once
)

type core struct {
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
	ETHCore ETHCore
}

func New(in CoreIn) CoreOut {
	once.Do(func() {
		self = &core{
			in: in,
			CoreOut: CoreOut{
				ETHCore: newETHCore(in),
			},
		}
	})

	return self.CoreOut
}
