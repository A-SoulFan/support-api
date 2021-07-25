package handler

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/milestone/logic"
	"asoul-fan-support/app/service/types"
	"asoul-fan-support/lib/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MilestoneNextGroup(svc *service.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.NextGroupReq
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.NewServerErrorResponse(err))
			return
		}

		lg := logic.NewGroupLogic(ctx, svc)
		if resp, err := lg.NextGroup(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewServerErrorResponse(err))
			return
		} else {
			ctx.JSON(http.StatusOK, response.NewSuccessJsonResponse(resp))
		}
	}
}
