//go:generate mockgen -destination=./../mocks/mock_transaction_repo.go -source=transaction_repo.go TransactionsRepo
package repository

import (
	"context"

	"gorm.io/gorm"

	"portto-homework/internal/model/po"
	"portto-homework/internal/utils/logger"
)

type TransactionsRepo interface {
	GetTransactionsByBlockNum(ctx context.Context, db *gorm.DB, blockNum uint64) ([]*po.Transaction, error)
	GetTransaction(ctx context.Context, db *gorm.DB, txHash string) (*po.Transaction, error)
	GetTransactionLogs(ctx context.Context, db *gorm.DB, txHash string) ([]*po.TransactionLog, error)
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

func (repo *transactionsRepo) GetTransactionsByBlockNum(ctx context.Context, db *gorm.DB, blockNum uint64) ([]*po.Transaction, error) {
	result := make([]*po.Transaction, 0)

	if err := db.Table("transactions").
		Where("`block_num` = ? ", blockNum).
		Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (repo *transactionsRepo) GetTransaction(ctx context.Context, db *gorm.DB, txHash string) (*po.Transaction, error) {
	result := &po.Transaction{}

	if err := db.Table("transactions").
		Where("`tx_hash` = ? ", txHash).
		Find(result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (repo *transactionsRepo) GetTransactionLogs(ctx context.Context, db *gorm.DB, txHash string) ([]*po.TransactionLog, error) {
	result := make([]*po.TransactionLog, 0)

	if err := db.Table("transaction_logs").
		Where("`tx_hash` = ? ", txHash).
		Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
