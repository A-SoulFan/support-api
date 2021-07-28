package logic

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/stroll/task"
	"asoul-fan-support/app/service/types"
	appErr "asoul-fan-support/lib/err"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type RandomGetLogic struct {
	ctx    context.Context
	svcCtx *service.Context
	dbCtx  *gorm.DB
}

func NewRandomGetLogic(ctx context.Context, svcCtx *service.Context) RandomGetLogic {
	return RandomGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		dbCtx:  svcCtx.Db.WithContext(ctx),
	}
}

func (r *RandomGetLogic) RandomGetStroll() (*types.StrollReply, error) {
	if stroll, err := task.RandomStroll(); err != nil {
		r.svcCtx.Logger.Error(err)
		return nil, appErr.NewError("暂时没有可以溜的数据哦，请稍后再试。")
	} else {
		return &types.StrollReply{
			Title:     stroll.Title,
			Cover:     stroll.Cover,
			BV:        stroll.BV,
			PlayUrl:   fmt.Sprintf("//player.bilibili.com/player.html?bvid=%s", stroll.BV),
			TargetUrl: stroll.TargetUrl,
			CreatedAt: stroll.CreatedAt,
		}, nil
	}
}
