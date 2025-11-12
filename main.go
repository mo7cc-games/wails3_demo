package main

import (
	"embed"
	"log/slog"
	"time"

	"app.local/app"
	"app.local/app/global"
	"app.local/app/utils/flog"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var FrontendDist embed.FS

func StartWailsApp() {
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

	global.WebWindow.Main.Window = global.WailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:             "Main",
		Width:            800,
		Height:           600,
		Frameless:        false,
		BackgroundColour: application.NewRGBA(0, 0, 0, 0),
		BackgroundType:   application.BackgroundTypeTransparent,
		URL:              "/",
	})
	global.WebWindow.Main.ListenWindowEvent()
}

func main() {
	global.SysInit() // 初始化全局系统

	StartWailsApp()

	go func() {
		for {
			global.WailsEvent.Time()
			time.Sleep(time.Second)
		}
	}()

	flog.AppLog.Info("main|WailsApp.Run", "启动 Wails 应用")
	// 运行 Wails 应用，并阻塞进程
	err := global.WailsApp.Run()
	if err != nil {
		flog.AppLog.Error("main|WailsApp.Run", err)
	}
}
