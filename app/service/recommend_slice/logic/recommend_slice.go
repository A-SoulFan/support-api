package logic

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/recommend_slice/task"
	"context"
)

type RecommendSliceLogic struct {
	ctx    context.Context
	svcCtx *service.Context
}

func NewRecommendSliceLogic(svc *service.Context) *RecommendSliceLogic {
	return &RecommendSliceLogic{svcCtx: svc}
}

func (rm *RecommendSliceLogic) Handle() ([]task.Video, error) {
	return task.Hot(20), nil
}
