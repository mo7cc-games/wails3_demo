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

//go:embed build/appicon.png
var IconFS embed.FS

func GetIcon() []byte {
	data, _ := IconFS.ReadFile("build/appicon.png")
	return data
}

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

	// 创建右键菜单
	global.WebWindow.Main.Window = global.WailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:           "Main",
		Width:          800,
		Height:         600,
		BackgroundType: application.BackgroundTypeTransparent,
		URL:            "/",
	})
	global.WebWindow.Main.EnableFrameless = true // 启用无边框模式

	global.WebWindow.Main.ListenWindowEvent()

	flog.AppLog.Info("main|StartWailsApp", "启动Wails应用")

	go func() {
		for {
			global.WailsEvent.Time()
			time.Sleep(time.Second)
		}
	}()
}

func SetTrayMenu() {
	iconBytes := GetIcon()

	trayMenu := application.NewMenu()

	trayMenu.Add("还原窗口").OnClick(func(ctx *application.Context) {
		global.WebWindow.Main.Window.Show()
		global.WebWindow.Main.Window.Restore()
	})
	trayMenu.AddSeparator()
	trayMenu.Add("最小化到托盘区").OnClick(func(ctx *application.Context) {
		global.WebWindow.Main.Window.Hide()
	})
	trayMenu.AddSeparator()
	trayMenu.Add("退出").OnClick(func(ctx *application.Context) {
		global.WailsApp.Quit()
	})

	tray := global.WailsApp.SystemTray.New()
	tray.SetIcon(iconBytes)
	tray.SetMenu(trayMenu)
	flog.AppLog.Info("main|SetTrayMenu", "设置托盘区")
}

func main() {
	global.SysInit() // 初始化全局系统

	StartWailsApp()

	go SetTrayMenu()

	// 运行 Wails 应用，并阻塞进程
	err := global.WailsApp.Run()
	if err != nil {
		flog.AppLog.Error("main|WailsApp.Run", err)
	}
}
