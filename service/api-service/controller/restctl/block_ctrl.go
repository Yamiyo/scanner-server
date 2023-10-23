package restctl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portto-homework/internal/constant"
	"strconv"
)

type BlockCtrl interface {
	GetBlockLatestN(ctx *gin.Context)
	GetBlockInfo(ctx *gin.Context)
}

type blockCtrl struct {
	in RestCtrlIn
}

func newBlockCtrl(in RestCtrlIn) BlockCtrl {
	return &blockCtrl{
		in: in,
	}
}

func (ctrl *blockCtrl) GetBlockLatestN(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		SetResp(ctx, http.StatusBadRequest, constant.InvalidateParameter, "limit is invalid")
		return
	}

	data, err := ctrl.in.BlockCore.GetBlockLatestN(ctx, limit)
	if err != nil {
		SetResp(ctx, http.StatusInternalServerError, constant.InternalServerError, err)
		return
	}
	SetResp(ctx, http.StatusOK, "", data)
}

func (ctrl *blockCtrl) GetBlockInfo(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("num"), 10, 64)
	if err != nil {
		SetResp(ctx, http.StatusBadRequest, constant.InvalidateParameter, err)
		return
	}

	data, err := ctrl.in.BlockCore.GetBlockInfo(ctx, id)
	if err != nil {
		SetResp(ctx, http.StatusInternalServerError, constant.InternalServerError, err)
		return
	}
	SetResp(ctx, http.StatusOK, "", data)
}
