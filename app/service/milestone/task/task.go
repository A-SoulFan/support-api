package task

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/milestone/model"
	"context"
	"sync"
	"time"
)

const (
	defaultCacheNumber = 1000
)

type taskRebuildCache struct {
	svcCtx         *service.Context
	cacheData      []*model.Milestone
	milestoneModel model.MilestoneModel
	lock           sync.Mutex
}

var (
	once sync.Once
	_tr  *taskRebuildCache
)

func Register(svc *service.Context) {
	once.Do(func() {
		_tr = &taskRebuildCache{
			svcCtx:         svc,
			cacheData:      make([]*model.Milestone, 0, defaultCacheNumber),
			milestoneModel: model.NewMilestoneModel(svc.Db.WithContext(context.TODO())),
		}

		_tr.init()
	})
}

func FindCacheAllByTimestampDesc(startTimestamp, limit uint) []*model.Milestone {
	return _tr.findCacheAllByTimestampDesc(startTimestamp, limit)
}

func (tr *taskRebuildCache) init() {
	if err := tr.RebuildCache(); err != nil {
		panic(err)
	}

	ticker(tr, tr.svcCtx)
}

func (tr *taskRebuildCache) findCacheAllByTimestampDesc(startTimestamp, limit uint) []*model.Milestone {
	if k := search(tr.cacheData, startTimestamp); k < 0 {
		return nil
	} else if (len(tr.cacheData) - k) < int(limit) {
		return nil
	} else {
		_list := make([]*model.Milestone, 0, limit)
		for _, milestone := range tr.cacheData[k:(k + int(limit))] {
			_m := *milestone
			_list = append(_list, &_m)
		}
		return _list
	}
}

func (tr *taskRebuildCache) RebuildCache() error {
	tr.lock.Lock()
	defer tr.lock.Unlock()
	if list, err := tr.milestoneModel.FindAllByTimestamp(uint(time.Now().UnixNano()/1e6), defaultCacheNumber, "DESC"); err != nil {
		return err
	} else {
		tr.cacheData = list
	}
	return nil
}

func search(milestones []*model.Milestone, startTimestamp uint) int {
	for i := 0; i < len(milestones); i++ {
		if milestones[i].Timestamp < startTimestamp {
			return i
		} else if milestones[i].Timestamp > startTimestamp {
			return -1
		}
	}
	return -1
}

func ticker(tr *taskRebuildCache, svc *service.Context) {
	ticker := time.NewTicker(5 * time.Minute)

	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := tr.RebuildCache(); err != nil {
					svc.Logger.Error(err)
				}
				//svc.Logger.Info("milestone.task.rebuildCache: successfully.")
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(ticker)
}
