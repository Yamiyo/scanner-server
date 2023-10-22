package repository

import (
	"portto-homework/internal/utils/logger"
	"sync"
)

var (
	once sync.Once
	self *DBRepo
)

type DBRepoIn struct {
	Log logger.LoggerInterface
}

type DBRepo struct {
	in DBRepoIn
	DBRepoOut
}

type DBRepoOut struct {
	BlocksRepo       BlocksRepo
	TransactionsRepo TransactionsRepo
}

func NewDBRepo(in DBRepoIn) DBRepoOut {
	once.Do(func() {
		self = &DBRepo{
			in: in,
			DBRepoOut: DBRepoOut{
				BlocksRepo:       newBlocksRepo(in),
				TransactionsRepo: newTransactionsRepo(in),
			},
		}
	})

	return self.DBRepoOut
}
