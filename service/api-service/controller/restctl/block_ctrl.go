//go:generate mockgen -destination=./../../mocks/mock_block_ctrl.go -source=block_ctrl.go BlockCtrl
package restctl

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"portto-homework/internal/constant"
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

// GetBlockLatestN godoc
// @Summary Get latest N blocks
// @Description Get latest N blocks
// @Tags block
// @Accept json
// @Produce json
// @Param Content-Type header string false "Request data format. Example: `application/json;charset=utf-8`"
// @Param ChainID header string false "request chain ID with UUID format. Example: `5d714332-60b4-451d-b45e-539f7b77f562`"
// @Param limit query string true "Input limit. Example: 25"
// @Success 200 {object} []*bo.GetBlockLatestNResp
// @Router /blocks [get]
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

// GetBlockInfo godoc
// @Summary Get blocks info by block number
// @Description Get blocks info by block number
// @Tags block
// @Accept json
// @Produce json
// @Param Content-Type header string false "Request data format. Example: `application/json;charset=utf-8`"
// @Param ChainID header string false "request chain ID with UUID format. Example: `5d714332-60b4-451d-b45e-539f7b77f562`"
// @Param num path string true "Input num. Example: `34449189`"
// @Success 200 {object} *bo.GetBlockInfoResp
// @Router /blocks/:num [get]
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
