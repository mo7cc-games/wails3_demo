package global

/*
和 wails App 全局状态相关的函数

*/

import (
	"log"
	"time"

	"app.local/app/utils/flog"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type SetTrayMenuOpt struct {
	Icon []byte
}

type TrayType struct {
	OpenBallWindowItem  *application.MenuItem
	CloseBallWindowItem *application.MenuItem
	BallSetTopBtn       *application.MenuItem
	BallUnsetTopBtn     *application.MenuItem
}

type ContextMenuType struct {
	BallSetTopBtn   *application.MenuItem
	BallUnsetTopBtn *application.MenuItem
}

// 全局 Wails 应用实例
var WailsApp *application.App

func SysInit() {
	go flog.LogInit() // 初始化日志系统
}

func BeforeRun() {
	// 初始化悬浮球为关闭状态
	BallWindow.Close()

	// 启动时钟事件
	go func() {
		for {
			WailsEvent.Time()
			time.Sleep(time.Second)
		}
	}()
}

var Tray TrayType

func SetTrayMenu(opt SetTrayMenuOpt) {
	trayMenu := application.NewMenu()

	trayMenu.Add("打开主窗口").OnClick(func(ctx *application.Context) {
		OpenMainWindow()
	})

	trayMenu.AddSeparator()
	trayMenu.Add("隐藏全部窗口").OnClick(func(ctx *application.Context) {
		WebWindow.HideAllWindows()
	})
	trayMenu.Add("还原全部窗口").OnClick(func(ctx *application.Context) {
		WebWindow.ShowAllWindows()
	})

	trayMenu.AddSeparator()
	Tray.BallSetTopBtn = trayMenu.Add("悬浮球置顶").OnClick(func(ctx *application.Context) {
		BallWindow.SetPinTop()
	})
	Tray.BallUnsetTopBtn = trayMenu.Add("悬浮球置顶取消").OnClick(func(ctx *application.Context) {
		BallWindow.UnsetPinTop()
	})
	Tray.OpenBallWindowItem = trayMenu.Add("显示悬浮球").OnClick(func(ctx *application.Context) {
		BallWindow.Open()
	})
	Tray.CloseBallWindowItem = trayMenu.Add("关闭悬浮球").OnClick(func(ctx *application.Context) {
		BallWindow.Close()
	})

	trayMenu.AddSeparator()
	trayMenu.Add("退出").OnClick(func(ctx *application.Context) {
		WailsApp.Quit()
	})

	tray := WailsApp.SystemTray.New()
	tray.SetIcon(opt.Icon)
	tray.SetMenu(trayMenu)
	go flog.AppLog.Info("main|SetTrayMenu", "设置托盘区")
}

var ContextMenu ContextMenuType

func SetContextMenu() {
	/*
		使用方式
		style="--custom-contextmenu: editor-menu; --custom-contextmenu-data: '墨七的数据'"
	*/
	editorMenu := application.NewContextMenu("editor-menu")
	editorMenu.Add("Cut").OnClick(func(ctx *application.Context) {
		data := ctx.ContextMenuData()
		log.Println("Cut 菜单被点击", data)
	})
	editorMenu.Add("Copy").OnClick(func(ctx *application.Context) {
		data := ctx.ContextMenuData()
		log.Println("Copy 菜单被点击", data)
	})
	editorMenu.Add("Paste").OnClick(func(ctx *application.Context) {
		data := ctx.ContextMenuData()
		log.Println("Paste 菜单被点击", data)
	})

	/*
		使用方式
		style="--custom-contextmenu: ball-menu; --custom-contextmenu-data: '墨七的数据'"
	*/
	ballMenu := application.NewContextMenu("ball-menu")
	ballMenu.Add("隐藏全部窗口").OnClick(func(ctx *application.Context) {
		WebWindow.HideAllWindows()
		if WebWindow.Ball.IsRunner {
			WebWindow.Ball.Window.Show()
		}
	})
	ballMenu.Add("显示全部窗口").OnClick(func(ctx *application.Context) {
		WebWindow.ShowAllWindows()
	})

	ContextMenu.BallSetTopBtn = ballMenu.Add("置顶").OnClick(func(ctx *application.Context) {
		BallWindow.SetPinTop()
	})
	ContextMenu.BallUnsetTopBtn = ballMenu.Add("取消置顶").OnClick(func(ctx *application.Context) {
		BallWindow.UnsetPinTop()
	})

	ballMenu.AddSeparator()
	ballMenu.Add("关闭悬浮球").OnClick(func(ctx *application.Context) {
		BallWindow.Close()
	})
}
