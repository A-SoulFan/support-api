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
	"strings"
)

const (
	memberListKey          = "member_list"
	memberExperiencePrefix = "member_experience_"
	memberVideoPrefix      = "member_videos_"
)

type MemberLogic struct {
	ctx    context.Context
	svcCtx *service.Context
	dbCtx  *gorm.DB
}

func NewMemberLogic(ctx context.Context, svcCtx *service.Context) MemberLogic {
	return MemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		dbCtx:  svcCtx.Db.WithContext(ctx),
	}
}

func (m *MemberLogic) GetAll() (*types.MemberAll, error) {
	val, err := model.NewDefaultKeyValueModel(m.dbCtx).FindOneByKey(memberListKey)
	if err != nil {
		m.svcCtx.Logger.Error(err)
		return nil, err
	}

	if val == nil {
		return nil, appErr.NewError("获取数据失败")
	}

	var list []interface{}
	if err := json.Unmarshal(val.Value, &list); err != nil {
		m.svcCtx.Logger.Error(err)
		return nil, err
	}

	return &types.MemberAll{MemberList: list}, nil
}

func (m *MemberLogic) GetExperience(req types.MemberExperienceReq) (*types.MemberExperienceResp, error) {
	queryKey := fmt.Sprintf("%s%s", memberExperiencePrefix, strings.ToLower(req.MemberName))

	val, err := model.NewDefaultKeyValueModel(m.dbCtx).FindOneByKey(queryKey)
	if err != nil {
		m.svcCtx.Logger.Error(err)
		return nil, err
	}

	if val == nil {
		return nil, appErr.NewError("获取数据失败")
	}

	var list []interface{}
	if err := json.Unmarshal(val.Value, &list); err != nil {
		m.svcCtx.Logger.Error(err)
		return nil, err
	}

	return &types.MemberExperienceResp{
		MemberName: req.MemberName,
		TotalCount: len(list),
		TotalPage:  1,
		VideoList:  list,
	}, nil
}

func (m *MemberLogic) GetVideos(req types.MemberVideoReq) (*types.MemberExperienceResp, error) {
	queryKey := fmt.Sprintf("%s%s", memberVideoPrefix, strings.ToLower(req.MemberName))

	val, err := model.NewDefaultKeyValueModel(m.dbCtx).FindOneByKey(queryKey)
	if err != nil {
		m.svcCtx.Logger.Error(err)
		return nil, err
	}

	if val == nil {
		return nil, appErr.NewError("获取数据失败")
	}

	var list []interface{}
	if err := json.Unmarshal(val.Value, &list); err != nil {
		m.svcCtx.Logger.Error(err)
		return nil, err
	}

	return &types.MemberExperienceResp{
		MemberName: req.MemberName,
		TotalCount: len(list),
		TotalPage:  1,
		VideoList:  list,
	}, nil
}
