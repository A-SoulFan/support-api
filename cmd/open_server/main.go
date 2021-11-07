package main

import (
	"asoul-fan-support/app/handler"
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/config"
	milestoneTask "asoul-fan-support/app/service/milestone/task"
	recommendSliceTask "asoul-fan-support/app/service/recommend_slice/task"
	strollTask "asoul-fan-support/app/service/stroll/task"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func main() {
	c := loadConfig()

	svc := service.NewServiceContext(c)
	defer svc.Stop()

	r := gin.Default()
	initRouters(r, svc)
	initRegister(svc)

	_ = r.Run(c.App.Port)
}

func loadConfig() config.Config {
	var (
		path    = "config/config.json"
		c       = config.Config{}
		content []byte
		err     error
	)

	if content, err = ioutil.ReadFile(path); err != nil {
		log.Fatalf("error: config file %s, %s", path, err.Error())
	}

	if err = json.Unmarshal(content, &c); err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	return c
}

func initRouters(r *gin.Engine, svc *service.Context) {
	// 随机溜
	r.GET("/api/stroll/random", handler.RandomStrollHandler(svc))

	// 大事件
	r.GET("/api/milestone/next-group", handler.MilestoneNextGroup(svc))

	// 下列为新人指南相关API 暂时实现在这里
	// 注意 response request 风格均不相同

	// 头部图片
	r.GET("/asf/mobile/headpicture", handler.GetBannerListHandler(svc))

	// 团队成员
	r.GET("/asf/mobile/member/all", handler.GetAllHandler(svc))
	// 团队个人经历
	r.GET("/asf/mobile/member/experience", handler.GetExperienceListHandler(svc))
	// 个人作品
	r.GET("/asf/mobile/member/videos", handler.GetVideoListHandler(svc))
	// 团队作品
	r.GET("/asf/mobile/team/videos", handler.GetTeamVideoListHandler(svc))
	// 团队事件
	r.GET("/asf/mobile/team/events", handler.GetTeamEventListHandler(svc))
}

func initRegister(svc *service.Context) {
	strollTask.Register(svc)
	milestoneTask.Register(svc)
	recommendSliceTask.Register(svc)
}
