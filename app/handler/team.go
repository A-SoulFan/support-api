package handler

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/keyval/logic"
	"asoul-fan-support/app/service/types"
	"asoul-fan-support/lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTeamVideoListHandler(svc *service.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TeamVideosReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		lg := logic.NewTeamLogic(ctx, svc)
		if resp, err := lg.GetVideos(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

func GetTeamEventListHandler(svc *service.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TeamEventsReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		lg := logic.NewTeamLogic(ctx, svc)
		if resp, err := lg.GetEvents(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}
