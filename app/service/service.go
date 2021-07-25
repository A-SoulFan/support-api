package service

import (
	"asoul-fan-support/app/service/config"
	"asoul-fan-support/lib/cache"
	"asoul-fan-support/lib/cache/driver"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Context struct {
	Logger *zap.SugaredLogger
	Config config.Config
	Db     *gorm.DB
	Cache  cache.Interface
}

func NewServiceContext(c config.Config) *Context {
	return &Context{
		Logger: newLogger(c),
		Config: c,
		Db:     newGormDbConn(c),
		Cache:  newCache(c),
	}
}

func (svcCtx *Context) IsDevEnvironment() bool {
	return strings.ToUpper(svcCtx.Config.App.Env) == "DEV"
}

func (svcCtx *Context) Stop() error {
	return svcCtx.Logger.Sync()
}

func newGormDbConn(c config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 nil,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func newCache(c config.Config) cache.Interface {
	return driver.NewGoCache(5*time.Minute, 6*time.Minute)
}

func newLogger(c config.Config) *zap.SugaredLogger {
	if logger, err := zap.NewProduction(); err != nil {
		panic(err)
	} else {
		return logger.Sugar()
	}
}
