package main

import (
	"asoul-fan-support/app/handler"
	"asoul-fan-support/app/service"
	"asoul-fan-support/app/service/config"
	"asoul-fan-support/app/service/stroll/task"
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
}

func initRegister(svc *service.Context) {
	task.Register(svc)
}
