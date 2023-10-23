package core

import (
	"sync"

	"portto-homework/service/api-service/config"
	"portto-homework/service/api-service/repository"
	"portto-homework/service/api-service/repository/db"
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
	Conf config.ConfigSetup
	DB   *db.DBClient

	BlockRepo repository.BlocksRepo
	TxnRepo   repository.TransactionsRepo
}

type CoreOut struct {
	BlockCore BlockCore
	TxnCore   TxnCore
}

func New(in CoreIn) CoreOut {
	once.Do(func() {
		self = &core{
			in: in,
			CoreOut: CoreOut{
				BlockCore: newBlockCore(in),
				TxnCore:   newTxnCore(in),
			},
		}
	})

	return self.CoreOut
}
