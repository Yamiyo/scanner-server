package restctl

import (
	"sync"

	"portto-homework/service/api-service/config"
	"portto-homework/service/api-service/core"
)

var (
	self *restCtrl
	once sync.Once
)

type restCtrl struct {
	in RestCtrlIn
	RestCtrlOut
}

type RestCtrlIn struct {
	Conf      config.ConfigSetup
	BlockCore core.BlockCore
	TxnCore   core.TxnCore
}

type RestCtrlOut struct {
	MiddlewareCtrl ResponseMiddlewareInterface
	BlockCtrl      BlockCtrl
	TxnCtrl        TxnCtrl
}

// @title api-service: RESTFul API for portto homework.
// @version 1.0
// @description API Service provides web3 blocks & transactions info.

// @host 127.0.0.1:12345
// @accept json
// @produce json
// @query.collection.format multi
// @schemes https

// @in header

// @tag.name block
// @tag.description Get block info
// @tag.name transaction
// @tag.description Get transaction info
func New(in RestCtrlIn) RestCtrlOut {
	once.Do(func() {
		self = &restCtrl{
			in: in,
			RestCtrlOut: RestCtrlOut{
				MiddlewareCtrl: newResponseMiddleware(),
				BlockCtrl:      newBlockCtrl(in),
				TxnCtrl:        newTxnCtrl(in),
			},
		}
	})

	return self.RestCtrlOut
}
