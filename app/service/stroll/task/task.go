package task

import (
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/stroll/model"
	"context"
	"errors"
	"math/rand"
	"sync"
	"time"
)

type timerQuery struct {
	svcCtx        *service.Context
	candidateList []*model.Stroll
	strollModel   model.StrollModel
}

var (
	once sync.Once
	_tq  *timerQuery
)

func Register(svc *service.Context) {
	once.Do(func() {
		_tq = &timerQuery{
			svcCtx:        svc,
			candidateList: make([]*model.Stroll, 0, 10),
			strollModel:   model.NewStrollModel(svc.Db.WithContext(context.TODO())),
		}

		_tq.init()
	})
}

func RandomStroll() (model.Stroll, error) {
	if _tq == nil {
		panic("not registered RandomStroll task.")
	}

	return _tq.randomStroll()
}

func (tq *timerQuery) init() {
	if err := tq.generateCandidateList(); err != nil {
		panic(err)
	}

	ticker(tq, tq.svcCtx)
}

func (tq *timerQuery) randomStroll() (model.Stroll, error) {
	l := len(tq.candidateList)
	if l != 0 {
		if stroll := tq.candidateList[rand.Intn(l)]; stroll != nil {
			return *stroll, nil
		}
	}
	return model.Stroll{}, errors.New("candidate list is empty. ")
}

func (tq *timerQuery) generateCandidateList() error {
	var (
		maxId        uint
		err          error
		randomIdList = make([]uint, 0, 10)
		strollList   = make([]*model.Stroll, 0, 20)
	)

	if maxId, err = tq.strollModel.FindMaxId(); err != nil {
		return err
	}

	for i := 0; i < 20; i++ {
		randId := rand.Intn(int(maxId)) + 1
		randomIdList = append(randomIdList, uint(randId))
	}

	if _list, err := tq.strollModel.FindAllByIds(randomIdList); err != nil {
		return err
	} else {
		strollList = append(strollList, _list...)
	}

	tq.candidateList = strollList
	return nil
}

func ticker(tq *timerQuery, svc *service.Context) {
	ticker := time.NewTicker(5 * time.Second)

	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := tq.generateCandidateList(); err != nil {
					svc.Logger.Error(err)
				}
				//svc.Logger.Info("stroll.task.timerQuery: successfully generated candidate list.")
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(ticker)
}
