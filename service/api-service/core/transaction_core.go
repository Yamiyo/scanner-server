package core

import (
	"context"
	"portto-homework/internal/model/bo"
)

type TxnCore interface {
	GetTxnInfo(ctx context.Context, txnHash string) (*bo.GetTransactionInfoResp, error)
}

type txnCore struct {
	in CoreIn
}

func newTxnCore(in CoreIn) TxnCore {
	return &txnCore{
		in: in,
	}
}

func (core *txnCore) GetTxnInfo(ctx context.Context, txnHash string) (*bo.GetTransactionInfoResp, error) {
	db := core.in.DB.Session()
	txn, err := core.in.TxnRepo.GetTransaction(ctx, db, txnHash)
	if err != nil {
		return nil, err
	}

	txnLogs, err := core.in.TxnRepo.GetTransactionLogs(ctx, db, txnHash)
	if err != nil {
		return nil, err
	}

	logs := make([]*bo.TransactionLog, 0, len(txnLogs))
	for _, txnLog := range txnLogs {
		logs = append(logs, &bo.TransactionLog{
			Index: txnLog.Index,
			Data:  string(txnLog.Log),
		})
	}

	return &bo.GetTransactionInfoResp{
		BlockNum: txn.BlockNum,
		TxHash:   txn.TxHash,
		From:     txn.From,
		To:       txn.To,
		Value:    txn.Value,
		Nonce:    txn.Nonce,
		Data:     txn.Data,
		Logs:     logs,
	}, nil
}
