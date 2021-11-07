package handler

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/keyval/logic"
	"asoul-fan-support/app/service/types"
	"asoul-fan-support/lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllHandler(svc *service.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lg := logic.NewMemberLogic(ctx, svc)
		if resp, err := lg.GetAll(); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

func GetExperienceListHandler(svc *service.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.MemberExperienceReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		lg := logic.NewMemberLogic(ctx, svc)
		if resp, err := lg.GetExperience(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}

func GetVideoListHandler(svc *service.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.MemberVideoReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		lg := logic.NewMemberLogic(ctx, svc)
		if resp, err := lg.GetVideos(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, SuccessResponse(resp))
		}
	}
}
