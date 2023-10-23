//go:generate mockgen -destination=./../mocks/mock_block_repo.go -source=blocks_repo.go BlocksRepo
package repository

import (
	"context"

	"gorm.io/gorm"

	"portto-homework/internal/model/po"
	"portto-homework/internal/utils/logger"
)

type BlocksRepo interface {
	CreateBlock(ctx context.Context, db *gorm.DB, data *po.Block) error
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

func (repo *blocksRepo) CreateBlock(ctx context.Context, db *gorm.DB, data *po.Block) error {
	return db.Create(data).Error
}
