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
