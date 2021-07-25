package logic

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/milestone/model"
	"asoul-fan-support/app/service/types"
	appErr "asoul-fan-support/lib/err"
	"context"
	"gorm.io/gorm"
)

type NextGroupLogic struct {
	ctx    context.Context
	svcCtx *service.Context
	dbCtx  *gorm.DB
}

func NewGroupLogic(ctx context.Context, svcCtx *service.Context) NextGroupLogic {
	return NextGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		dbCtx:  svcCtx.Db.WithContext(ctx),
	}
}

func (ng *NextGroupLogic) NextGroup(req types.NextGroupReq) (*types.PaginationList, error) {
	var (
		list    []*model.Milestone
		nextKey *uint
		err     error
	)

	if list, err = model.NewMilestoneModel(ng.dbCtx).FindAllByTimestamp(req.NextKey, req.Size+uint(1), "DESC"); err != nil {
		ng.svcCtx.Logger.Error(err)
		return nil, appErr.NewError("服务器异常，请稍后再试")
	}

	if len(list) > int(req.Size) {
		nextKey = &list[len(list)-1].Timestamp
		list = list[0 : len(list)-1]
	}

	return &types.PaginationList{
		List:    toReply(list),
		NextKey: nextKey,
	}, nil
}

func toReply(list []*model.Milestone) []*types.NextGroupReply {
	_list := make([]*types.NextGroupReply, 0, len(list))
	for _, m := range list {
		_list = append(_list, &types.NextGroupReply{
			Title:     m.Title,
			Subtitled: m.Subtitled,
			Type:      m.Type,
			Content:   m.Content,
			TargetUrl: m.TargetUrl,
			Timestamp: m.Timestamp,
		})
	}
	return _list
}
