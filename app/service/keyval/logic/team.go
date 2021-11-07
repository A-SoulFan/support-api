package logic

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/keyval/model"
	"asoul-fan-support/app/service/types"
	appErr "asoul-fan-support/lib/err"
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

const (
	teamVideos       = "team_videos"
	teamEventsPrefix = "team_events_"
)

type TeamLogic struct {
	ctx    context.Context
	svcCtx *service.Context
	dbCtx  *gorm.DB
}

func NewTeamLogic(ctx context.Context, svcCtx *service.Context) TeamLogic {
	return TeamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		dbCtx:  svcCtx.Db.WithContext(ctx),
	}
}

func (t *TeamLogic) GetVideos(req types.TeamVideosReq) (*types.TeamVideosResp, error) {
	val, err := model.NewDefaultKeyValueModel(t.dbCtx).FindOneByKey(teamVideos)
	if err != nil {
		t.svcCtx.Logger.Error(err)
		return nil, err
	}

	if val == nil {
		return nil, appErr.NewError("获取数据失败")
	}

	var list []interface{}
	if err := json.Unmarshal(val.Value, &list); err != nil {
		t.svcCtx.Logger.Error(err)
		return nil, err
	}

	return &types.TeamVideosResp{
		TotalCount: len(list),
		TotalPage:  1,
		VideoList:  list,
	}, nil
}

func (t *TeamLogic) GetEvents(req types.TeamEventsReq) (*types.TeamEventsResp, error) {
	queryKey := fmt.Sprintf("%s%s", teamEventsPrefix, req.Year)
	val, err := model.NewDefaultKeyValueModel(t.dbCtx).FindOneByKey(queryKey)
	if err != nil {
		t.svcCtx.Logger.Error(err)
		return nil, err
	}

	if val == nil {
		return nil, appErr.NewError("获取数据失败")
	}

	var list []interface{}
	if err := json.Unmarshal(val.Value, &list); err != nil {
		t.svcCtx.Logger.Error(err)
		return nil, err
	}

	return &types.TeamEventsResp{
		TotalCount: len(list),
		TotalPage:  1,
		EventList:  list,
	}, nil
}
