package main

import (
	"embed"
	"log/slog"

	"app.local/app"
	"app.local/app/global"
	"app.local/app/utils/flog"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var FrontendDist embed.FS

//go:embed build/appicon.png
var IconFS embed.FS

func getIcon() []byte {
	data, _ := IconFS.ReadFile("build/appicon.png")
	return data
}

func startWailsApp() {
	global.WailsApp = application.New(application.Options{
		Name:        "LoneStarEngineer",
		Description: "《孤星工程师》",
		Services: []application.Service{
			application.NewService(&app.WailsService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(FrontendDist),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		LogLevel: slog.LevelWarn,
	})

	go flog.AppLog.Info("main.startWailsApp", "启动Wails应用")
}

func main() {
	global.SysInit() // 初始化全局系统

	startWailsApp() // 启动 Wails 应用

	// 配置上下文菜单
	global.SetContextMenu()

	// 配置托盘区菜单
	global.SetTrayMenu(
		global.SetTrayMenuOpt{
			Icon: getIcon(),
		},
	)

	global.OpenMainWindow() // 打开主窗口

	// 运行前的准备工作
	global.BeforeRun()

	// 运行 Wails 应用，并阻塞进程
	err := global.WailsApp.Run()
	if err != nil {
		go flog.AppLog.Error("main|WailsApp.Run", err)
	}
}
