package core

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"portto-homework/internal/model/po"
	"portto-homework/internal/utils/logger"
	"sync"
)

type ETHCore interface {
	ScanBlockDataFromNum(ctx context.Context, num uint64) error
}

type ethCore struct {
	in      CoreIn
	parseCh chan *types.Block
	wg      *sync.WaitGroup
}

func newETHCore(in CoreIn) ETHCore {
	eth := &ethCore{
		in:      in,
		parseCh: make(chan *types.Block, in.Conf.ScannerConfig.PipelineNumber),
		wg:      &sync.WaitGroup{},
	}

	eth.run(context.Background())

	return eth
}

func (core *ethCore) close() {
	// close parseCh
	close(core.parseCh)

	// wait for all goroutine done
	core.wg.Wait()
}

func (core *ethCore) ScanBlockDataFromNum(ctx context.Context, num uint64) error {
	// get latest block number
	latest, err := core.in.EthClient.BlockNumber(ctx)
	if err != nil {
		return err
	}

	for i := num; i < latest; i++ {
		block, err := core.in.EthClient.BlockByNumber(ctx, new(big.Int).SetUint64(i))
		if err != nil {
			return err
		}

		core.parseCh <- block
	}

	return nil
}

func (core *ethCore) run(ctx context.Context) {

	for i := 0; i < core.in.Conf.ScannerConfig.PipelineNumber; i++ {
		core.wg.Add(1)
		go func() {
			defer core.wg.Done()
			db := core.in.DB.Session()
			for block := range core.parseCh {
				if block == nil {
					continue
				}

				transactions, logs, err := core.getTransactionsByBlock(ctx, block)
				if err != nil {
					logger.SysLog().Error(ctx, err.Error())
					return
				}

				if err := core.in.BlockRepo.CreateBlock(ctx, db, &po.Block{
					Num:        block.Number().Uint64(),
					Hash:       block.Hash().String(),
					ParentHash: block.ParentHash().String(),
					Time:       block.Time(),
				}); err != nil {
					logger.SysLog().Error(ctx, err.Error())
					return
				}

				if err := core.in.TransactionRepo.CreateTransaction(ctx, db, transactions); err != nil {
					logger.SysLog().Error(ctx, err.Error())
					return
				}

				if err := core.in.TransactionRepo.CreateTransactionLogs(ctx, db, logs); err != nil {
					logger.SysLog().Error(ctx, err.Error())
					return
				}
			}
		}()
	}
}

func (core *ethCore) getTransactionsByBlock(ctx context.Context, block *types.Block) ([]*po.Transaction, []*po.TransactionLog, error) {
	transactions := make([]*po.Transaction, 0, len(block.Transactions()))
	transactionLogs := make([]*po.TransactionLog, 0)
	for index, transaction := range block.Transactions() {
		txn, _, err := core.in.EthClient.TransactionByHash(ctx, transaction.Hash())
		if err != nil {
			return nil, nil, err
		}

		receipt, err := core.in.EthClient.TransactionReceipt(ctx, transaction.Hash())
		if err != nil {
			return nil, nil, err
		}

		var to *string
		if txn.To() != nil {
			toStr := txn.To().String()
			to = &toStr
		}
		from, err := core.in.EthClient.TransactionSender(ctx, transaction, block.Hash(), uint(index))
		if err != nil {
			return nil, nil, err
		}

		for _, log := range receipt.Logs {
			logJson, err := log.MarshalJSON()
			if err != nil {
				return nil, nil, err
			}
			transactionLogs = append(transactionLogs, &po.TransactionLog{
				BlockNum: block.Number().Uint64(),
				TxHash:   txn.Hash().String(),
				Index:    log.Index,
				Log:      logJson,
			})
		}

		transactions = append(transactions, &po.Transaction{
			BlockNum: block.Number().Uint64(),
			TxHash:   txn.Hash().String(),
			From:     from.String(),
			To:       to,
			Value:    txn.Value().String(),
			Nonce:    txn.Nonce(),
			Data:     txn.Data(),
		})
	}

	return transactions, transactionLogs, nil
}
