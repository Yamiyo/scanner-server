package restctl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portto-homework/internal/constant"
)

type TxnCtrl interface {
	GetTxnInfo(ctx *gin.Context)
}

type txnCtrl struct {
	in RestCtrlIn
}

func newTxnCtrl(in RestCtrlIn) TxnCtrl {
	return &txnCtrl{
		in: in,
	}
}

// GetTxnInfo godoc
// @Summary Get transaction info by txn hash
// @Description Get transaction info by txn hash
// @Tags txn
// @Accept json
// @Produce json
// @Param Content-Type header string false "Request data format. Example: `application/json;charset=utf-8`"
// @Param ChainID header string false "request chain ID with UUID format. Example: `5d714332-60b4-451d-b45e-539f7b77f562`"
// @Param txn_hash path string true "Input txn_hash. Example: `0x5510c3187af8f24ecfbd42af5c72cc76d070bcb4ab3ac98d5b4e12a15b04dda9`"
// @Success 200 {object} *bo.GetTransactionInfoResp
// @Router /transactions/:txn_hash [get]
func (ctrl *txnCtrl) GetTxnInfo(ctx *gin.Context) {
	txnHash := ctx.Param("txn_hash")
	if txnHash == "" {
		SetResp(ctx, http.StatusBadRequest, constant.InvalidateParameter, "txn_hash is empty")
		return
	}

	data, err := ctrl.in.TxnCore.GetTxnInfo(ctx, txnHash)
	if err != nil {
		SetResp(ctx, http.StatusInternalServerError, constant.InternalServerError, err)
		return
	}
	SetResp(ctx, http.StatusOK, "", data)
}
