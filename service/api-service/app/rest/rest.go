package rest

import (
	"context"
	"net/http"
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
		ctrl := restctl.New(restctl.RestCtrlIn{
			Conf: conf,
		})

		self = &restService{
			ResponseMiddleware: ctrl.ModdlewareCtrl,
		}
	})

	return self
}

type ServiceInterface interface {
	Run(ctx context.Context, stop chan error)
}

type restService struct {
	ResponseMiddleware restctl.ResponseMiddlewareInterface
}

func (s *restService) Run(ctx context.Context, stop chan error) {
	engine = s.newEngine()
	s.setRoutes(engine)

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
