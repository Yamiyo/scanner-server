//go:generate mockgen -destination=./../mocks/mock_transactions_repo.go -source=transaction_repo.go TransactionsRepo
package repository

import (
	"context"

	"gorm.io/gorm"

	"portto-homework/internal/model/po"
	"portto-homework/internal/utils/logger"
)

type TransactionsRepo interface {
	CreateTransaction(ctx context.Context, db *gorm.DB, data []*po.Transaction) error
	CreateTransactionLogs(ctx context.Context, db *gorm.DB, data []*po.TransactionLog) error
}

type transactionsRepo struct {
	logger logger.LoggerInterface
	in     DBRepoIn
}

func newTransactionsRepo(in DBRepoIn) TransactionsRepo {
	return &transactionsRepo{
		in: in,
	}
}

func (repo *transactionsRepo) CreateTransaction(ctx context.Context, db *gorm.DB, data []*po.Transaction) error {
	return db.CreateInBatches(data, 100).Error
}

func (repo *transactionsRepo) CreateTransactionLogs(ctx context.Context, db *gorm.DB, data []*po.TransactionLog) error {
	return db.CreateInBatches(data, 100).Error
}
