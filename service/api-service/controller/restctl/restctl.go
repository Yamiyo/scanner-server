package restctl

import (
	"sync"

	"portto-homework/service/api-service/config"
)

var (
	self *mongoRepo
	once sync.Once
)

type mongoRepo struct {
	in RestCtrlIn
	RestCtrlOut
}

type RestCtrlIn struct {
	Conf config.ConfigSetup
}

type RestCtrlOut struct {
	ModdlewareCtrl ResponseMiddlewareInterface
}

func New(in RestCtrlIn) RestCtrlOut {
	once.Do(func() {
		self = &mongoRepo{
			in: in,
			RestCtrlOut: RestCtrlOut{
				ModdlewareCtrl: newResponseMiddleware(),
			},
		}
	})

	return self.RestCtrlOut
}
