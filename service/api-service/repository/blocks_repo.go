package repository

import (
	"context"
	"gorm.io/gorm"
	"portto-homework/internal/model/po"
	"portto-homework/internal/utils/logger"
)

type BlocksRepo interface {
	GetBlockLatestN(ctx context.Context, db *gorm.DB, n int) ([]*po.Block, error)
	GetBlockByNum(ctx context.Context, db *gorm.DB, num uint64) (*po.Block, error)
}

type blocksRepo struct {
	logger logger.LoggerInterface
	in     DBRepoIn
}

func newBlocksRepo(in DBRepoIn) BlocksRepo {
	return &blocksRepo{
		in: in,
	}
}

func (repo *blocksRepo) GetBlockLatestN(ctx context.Context, db *gorm.DB, n int) ([]*po.Block, error) {
	result := make([]*po.Block, 0, n)

	if err := db.Table("blocks").Order("num desc").Limit(n).Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (repo *blocksRepo) GetBlockByNum(ctx context.Context, db *gorm.DB, num uint64) (*po.Block, error) {
	result := &po.Block{}
	if err := db.Table("blocks").
		Where("`num` = ? ", num).
		First(result).Error; err != nil {
		return result, err
	}

	return result, nil
}
