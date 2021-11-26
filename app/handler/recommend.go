package handler

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/recommend_slice/logic"
	"asoul-fan-support/lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecommendHandler(svc *service.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lg := logic.NewRecommendSliceLogic(svc)
		if resp, err := lg.Handle(); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}
