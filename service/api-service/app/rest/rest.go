package rest

import (
	"context"
	"net/http"
	"portto-homework/service/api-service/core"
	"portto-homework/service/api-service/repository"
	"portto-homework/service/api-service/repository/db"
	"sync"

	"portto-homework/internal/utils/logger"
	"portto-homework/service/api-service/config"
	"portto-homework/service/api-service/controller/restctl"

	"github.com/gin-gonic/gin"
)

var (
	engine     *gin.Engine
	engineOnce sync.Once
)

var (
	self *restService
	once sync.Once
)

func NewRestService(ctx context.Context, conf config.ConfigSetup) ServiceInterface {
	once.Do(func() {
		repo := repository.NewDBRepo(repository.DBRepoIn{})

		co := core.New(core.CoreIn{
			Conf:      conf,
			DB:        db.NewDBClient(),
			BlockRepo: repo.BlocksRepo,
			TxnRepo:   repo.TransactionsRepo,
		})

		ctrl := restctl.New(restctl.RestCtrlIn{
			Conf:      conf,
			BlockCore: co.BlockCore,
			TxnCore:   co.TxnCore,
		})

		self = &restService{
			ResponseMiddleware: ctrl.MiddlewareCtrl,
			BlockCtrl:          ctrl.BlockCtrl,
			TxnCtrl:            ctrl.TxnCtrl,
		}
	})

	return self
}

type ServiceInterface interface {
	Run(ctx context.Context, stop chan error)
}

type restService struct {
	ResponseMiddleware restctl.ResponseMiddlewareInterface
	BlockCtrl          restctl.BlockCtrl
	TxnCtrl            restctl.TxnCtrl
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
func (s *restService) Run(ctx context.Context, stop chan error) {
	engine = s.newEngine()
	s.setRoutes(engine)
	defer logger.SysLog().Info(ctx, "api-service is shutdown")

	go func() {
		if err := engine.Run(config.GetGinConfig().HttpPort); err != nil && err != http.ErrServerClosed {
			logger.SysLog().Error(ctx, err.Error())
			stop <- err
		}
	}()

	<-ctx.Done()
}

func (s *restService) newEngine() *gin.Engine {
	return gin.New()
}

func (s *restService) setRoutes(engine *gin.Engine) {
	// set middlewares
	engine.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	// set router
	s.setPrivateRoutes(engine)
	s.setPublicRoutes(engine)
}

func (s *restService) setPublicRoutes(engine *gin.Engine) {
	publicRouteGroup := engine.Group("")

	// 設定 middleware
	publicRouteGroup.Use(
		s.ResponseMiddleware.Handle,
	)

	// 設定路由
	s.setAPIRoutes(publicRouteGroup)
}

func (s *restService) setPrivateRoutes(engine *gin.Engine) {
	privateRouteGroup := engine.Group("")

	privateRouteGroup.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
}
