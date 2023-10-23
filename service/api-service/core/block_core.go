package core

import (
	"context"

	"portto-homework/internal/model/bo"
)

type BlockCore interface {
	GetBlockLatestN(ctx context.Context, n int) ([]*bo.GetBlockLatestNResp, error)
	GetBlockInfo(ctx context.Context, num uint64) (*bo.GetBlockInfoResp, error)
}

type blockCore struct {
	in CoreIn
}

func newBlockCore(in CoreIn) BlockCore {
	return &blockCore{
		in: in,
	}
}

func (core *blockCore) GetBlockLatestN(ctx context.Context, n int) ([]*bo.GetBlockLatestNResp, error) {
	db := core.in.DB.Session()
	data, err := core.in.BlockRepo.GetBlockLatestN(ctx, db, n)
	if err != nil {
		return nil, err
	}

	result := make([]*bo.GetBlockLatestNResp, 0, len(data))
	for _, v := range data {
		result = append(result, &bo.GetBlockLatestNResp{
			Num:        v.Num,
			Hash:       v.Hash,
			Time:       v.Time,
			ParentHash: v.ParentHash,
		})
	}
	return result, nil
}

func (core *blockCore) GetBlockInfo(ctx context.Context, num uint64) (*bo.GetBlockInfoResp, error) {
	db := core.in.DB.Session()
	block, err := core.in.BlockRepo.GetBlockByNum(ctx, db, num)
	if err != nil {
		return nil, err
	}

	txns, err := core.in.TxnRepo.GetTransactionsByBlockNum(ctx, db, block.Num)
	if err != nil {
		return nil, err
	}

	txnsHash := make([]string, 0, len(txns))
	for _, txn := range txns {
		txnsHash = append(txnsHash, txn.TxHash)
	}
	return &bo.GetBlockInfoResp{
		Num:          block.Num,
		Hash:         block.Hash,
		Time:         block.Time,
		ParentHash:   block.ParentHash,
		Transactions: txnsHash,
	}, nil
}
