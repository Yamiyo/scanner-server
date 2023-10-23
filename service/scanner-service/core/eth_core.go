package core

import (
	"context"
	"math/big"
	"sync"
	"time"

	"portto-homework/internal/model/po"
	"portto-homework/internal/utils/logger"

	"github.com/ethereum/go-ethereum/core/types"
)

type ETHCore interface {
	GetSubscribeFrom() uint64
	SubScribe(ctx context.Context)
	ScanBlockDataFromNum(ctx context.Context, num uint64) error
	Close()
}

type ethCore struct {
	in         CoreIn
	realTimeCh chan *types.Block
	historyCh  chan *types.Block
	realTimeWg *sync.WaitGroup
	historyWg  *sync.WaitGroup

	subFrom uint64
}

func newETHCore(in CoreIn) ETHCore {
	eth := &ethCore{
		in:         in,
		realTimeCh: make(chan *types.Block, in.Conf.ScannerConfig.PipelineNumber),
		historyCh:  make(chan *types.Block, in.Conf.ScannerConfig.PipelineNumber),
		realTimeWg: &sync.WaitGroup{},
		historyWg:  &sync.WaitGroup{},
	}

	eth.run(context.Background())

	return eth
}

func (core *ethCore) Close() {
	// close realTimeCh
	close(core.realTimeCh)
	// close historyCh
	close(core.historyCh)

	// wait for all goroutine done
	core.realTimeWg.Wait()
	logger.SysLog().Info(context.Background(), "real time pipeline shutdown done")
	core.historyWg.Wait()
	logger.SysLog().Info(context.Background(), "history pipeline shutdown done")
}

func (core *ethCore) GetSubscribeFrom() uint64 {
	return core.subFrom
}

// SubScribe subscribe function need to be called before ScanBlockDataFromNum
func (core *ethCore) SubScribe(ctx context.Context) {
	// get start subscribe block number
	from, err := core.in.EthClient.BlockNumber(ctx)
	if err != nil {
		logger.SysLog().Error(ctx, err.Error())
		panic(err)
	}
	core.subFrom = from

	core.realTimeWg.Add(1)
	go func() {
		core.realTimeWg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			// get latest block number
			latest, err := core.in.EthClient.BlockNumber(ctx)
			if err != nil {
				logger.SysLog().Error(ctx, err.Error())
				return
			}

			for ; from < latest; from++ {
				block, err := core.in.EthClient.BlockByNumber(ctx, new(big.Int).SetUint64(from))
				if err != nil {
					return
				}

				core.realTimeCh <- block
			}

			time.Sleep(time.Duration(core.in.Conf.ScannerConfig.ScanInterval) * time.Second)
		}
	}()
}

func (core *ethCore) ScanBlockDataFromNum(ctx context.Context, num uint64) error {
	// get latest block number
	if core.subFrom == 0 {
		// get latest block number
		latest, err := core.in.EthClient.BlockNumber(ctx)
		if err != nil {
			return err
		}
		core.subFrom = latest
	}

	for i := num; i < core.subFrom; i++ {
		block, err := core.in.EthClient.BlockByNumber(ctx, new(big.Int).SetUint64(i))
		if err != nil {
			return err
		}

		core.historyCh <- block
		select {
		case <-ctx.Done():
			return nil
		default:
		}
	}

	return nil
}

func (core *ethCore) run(ctx context.Context) {
	// create multi goroutine to parse block real time data
	for i := 0; i < core.in.Conf.ScannerConfig.PipelineNumber; i++ {
		core.realTimeWg.Add(1)
		go core.getBlockFromChan(ctx, core.realTimeWg, core.realTimeCh)
	}

	// create multi goroutine to parse block history data
	for i := 0; i < core.in.Conf.ScannerConfig.PipelineNumber; i++ {
		core.historyWg.Add(1)
		go core.getBlockFromChan(ctx, core.historyWg, core.historyCh)
	}
}

func (core *ethCore) getBlockFromChan(ctx context.Context, wg *sync.WaitGroup, ch chan *types.Block) {
	defer wg.Done()
	db := core.in.DB.Session()
	for block := range ch {
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
