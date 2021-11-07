package handler

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/banner/logic"
	"asoul-fan-support/app/service/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBannerListHandler(svc *service.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.BannerListReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		lg := logic.NewBannerListLogic(ctx, svc)
		if resp, err := lg.GetList(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

type ASFResponse struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func ErrorResponse(err error) *ASFResponse {
	return &ASFResponse{Code: -1, ErrMsg: err.Error()}
}

func SuccessResponse(data interface{}) *ASFResponse {
	return &ASFResponse{Code: 0, ErrMsg: "ok", Data: data}
}
