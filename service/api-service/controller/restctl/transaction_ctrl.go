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
