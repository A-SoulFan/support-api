package handler

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/stroll/logic"
	"asoul-fan-support/lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RandomStrollHandler(svc *service.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lg := logic.NewRandomGetLogic(ctx, svc)
		if resp, err := lg.RandomGetStroll(); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}

func LastUpdateTimeHandler(svc *service.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lg := logic.NewRandomGetLogic(ctx, svc)
		if resp, err := lg.LastUpdateTime(); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}
