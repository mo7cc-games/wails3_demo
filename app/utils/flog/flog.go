package flog

import (
	"time"

	"github.com/m-startgo/go-utils/mcycle"
	"github.com/m-startgo/go-utils/mlog"
)

var AppLog *mlog.Logger

func LogInit() {
	AppLog = mlog.New(mlog.Config{
		Path: "./logs/app",
		Name: "app",
	})

	cy := mcycle.New(mcycle.Options{
		Task: func() {
			AppLog.Info("flog.LogInit", "开始清理日志文件")
			AppLog.Clear(mlog.ClearOpt{}) // 清理日志文件
		},
		Interval: time.Hour * 24,
	})
	cy.Start()
}
