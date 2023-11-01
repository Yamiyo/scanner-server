//go:generate mockgen -destination=./../../mocks/mock_middleware_ctrl.go -source=middleware_response.go ResponseMiddlewareInterface
package restctl

import (
	"time"

	"portto-homework/internal/constant"
	"portto-homework/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func newResponseMiddleware() *responseMiddleware {
	return &responseMiddleware{}
}

type ResponseMiddlewareInterface interface {
	Handle(ctx *gin.Context)
}

type responseMiddleware struct{}

type meta struct {
	RequestID   string    `json:"request_id" example:"5d714332-60b4-451d-b45e-539f7b77f562"`
	UserID      string    `json:"user_id" example:"5d714332-60b4-451d-b45e-539f7b77f562"`
	RequestTime time.Time `json:"request_time" example:"2023-11-01T07:55:48.51208251Z"`
	Times       float64   `json:"times" example:"0.22385625"`
}

type resp struct {
	Meta        meta        `json:"meta"`
	MessageCode string      `json:"msg_code" example:"invalidate_parameter"`
	Data        interface{} `json:"data"`
}

type respWithPagination struct {
	resp
	Pagination interface{} `json:"pagination"`
}

func (b *responseMiddleware) Handle(ctx *gin.Context) {
	now := time.Now()
	if chainID := ctx.GetString(constant.App_ChainID); chainID == "" {
		ctx.Set(constant.App_ChainID, uuid.New().String())
	}

	ctx.Next()

	switch ctx.GetString(Resp_Format) {
	case RespFormat_Standard:
		ctx.JSON(
			ctx.GetInt(Resp_Status),
			resp{
				Meta: meta{
					RequestID:   ctx.GetString(constant.App_ChainID),
					RequestTime: now,
					Times:       time.Since(now).Seconds(),
				},
				MessageCode: ctx.GetString(Resp_MessageCode),
				Data:        ctx.MustGet(Resp_Data),
			},
		)

	case RespFormat_Pagination:
		ctx.JSON(
			ctx.GetInt(Resp_Status),
			respWithPagination{
				resp: resp{
					Meta: meta{
						RequestID: ctx.GetString(constant.App_ChainID),
					},
					MessageCode: ctx.GetString(Resp_MessageCode),
					Data:        ctx.MustGet(Resp_Data),
				},
				Pagination: ctx.MustGet(Resp_Pagination),
			},
		)

	default:
	}
}

const (
	RespFormat_Standard   = "RespFormat_Standard"
	RespFormat_Pagination = "RespFormat_Pagination"

	Resp_Format      = "Resp_Format"
	Resp_Data        = "Resp_Data"
	Resp_Status      = "Resp_Status"
	Resp_MessageCode = "Resp_MessageCode"
	Resp_Pagination  = "Resp_Pagination"
)

func SetResp(ctx *gin.Context, statusCode int, msgCode string, data interface{}) {
	ctx.Set(Resp_Format, RespFormat_Standard)
	ctx.Set(Resp_Status, statusCode)
	ctx.Set(Resp_MessageCode, msgCode)
	ctx.Set(Resp_Data, data)
}

func SetRespWithPagination(ctx *gin.Context, statusCode int, msgCode string, pagination *model.PaginationResp, data interface{}) {
	ctx.Set(Resp_Format, RespFormat_Pagination)
	ctx.Set(Resp_Status, statusCode)
	ctx.Set(Resp_MessageCode, msgCode)
	ctx.Set(Resp_Data, data)
	ctx.Set(Resp_Pagination, pagination)
}
