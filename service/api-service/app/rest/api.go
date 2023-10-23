package rest

import (
	"github.com/gin-gonic/gin"
)

func (s *restService) setAPIRoutes(parentRouteGroup *gin.RouterGroup) {
	router := parentRouteGroup.Group("")

	// BlockCtrl
	block := router.Group("/blocks")
	block.GET("", s.BlockCtrl.GetBlockLatestN)
	block.GET("/:num", s.BlockCtrl.GetBlockInfo)

	// TxnCtrl
	txn := router.Group("/transactions")
	txn.GET("/:txn_hash", s.TxnCtrl.GetTxnInfo)
}
