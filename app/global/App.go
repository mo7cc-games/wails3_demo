package global

import (
	"app.local/app/utils/flog"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// 全局 Wails 应用实例
var WailsApp *application.App

func SysInit() {
	flog.LogInit() // 初始化日志系统
}
