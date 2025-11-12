package app

import (
	"app.local/app/global"
	"app.local/app/utils/flog"
	"github.com/m-startgo/go-utils/mjson"
	"github.com/m-startgo/go-utils/mmath"
	"github.com/wailsapp/wails/v3/pkg/application"
)

/*
这里的方法只能被前端调用，也就是不会被其它 go 模块引入
*/

type WailsService struct{}

func (s *WailsService) Add(n1 float64, n2 float64) string {
	a := mmath.NewFromFloat(n1)
	b := mmath.NewFromFloat(n2)
	result := a.Add(b)
	return result.String()
}

// 专门用于 Pinia 的 Action 通信，被前端调用
func (s *WailsService) Action(WindowName string, ActionName string) {
	global.WailsEvent.Action(WindowName, ActionName)
}

// 专门用于 Pinia 的 Action 通信，被前端调用
func (s *WailsService) GetWindowInfo(WindowName string) global.WebviewWindow {
	w := global.GetWebviewWindow(WindowName)
	if w == nil {
		// 未知窗口，这里应该打印警告
		flog.AppLog.Warn("WailsService.GetWindowInfo", "未知窗口："+WindowName)
		return global.WebviewWindow{}
	}
	return *w
}

func (s *WailsService) OpenTestWindow() {
	if global.WebWindow.Test.IsRunner {
		// About 已经在运行，忽略此次请求
		flog.AppLog.Warn("WailsService.OpenAboutWindow", "窗口已经在运行！")
		return
	}
	global.WebWindow.Test.Window = global.WailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:           "Test",
		Width:          800,
		Height:         600,
		BackgroundType: application.BackgroundTypeTransparent,
		URL:            "/#/test",
	})
	global.WebWindow.Test.EnableFrameless = true // 启用无边框模式

	// 启用监听
	global.WebWindow.Test.ListenWindowEvent()
}

type DprChangeOpt struct {
	WindowName string
	NewDpr     float32
}

// 从前端拿到窗口变化的数据
func (s *WailsService) DprChange(opt DprChangeOpt) {
	w := global.GetWebviewWindow(opt.WindowName)
	if w == nil {
		// 未知窗口，这里应该打印警告
		flog.AppLog.Warn("WailsService.DprChange", "未知窗口："+opt.WindowName)
		return
	}
	if w.OrgDpr == 0 {
		w.OrgDpr = opt.NewDpr
	}
	w.NowDpr = opt.NewDpr
	flog.AppLog.Debug("WailsService.DprChange", opt.WindowName, opt.NewDpr)
}

type WindowChangeOpt struct {
	WindowName   string  // 窗口名称
	Width        int     // 宽度
	Height       int     // 高度
	AvailWidth   int     // 可用宽度
	AvailHeight  int     // 可用高度
	ScreenLeft   int     // 屏幕左上角位置
	ScreenTop    int     // 屏幕左上角位置
	InnerWidth   int     // 内部宽度
	InnerHeight  int     // 内部高度
	NowDpr       float32 // 当前 DPR
	IsFullscreen bool    // 是否全屏
	IsExtended   bool    // 是否是扩展屏幕
	IsMaximise   bool    // 是否最大化
	IsMinimise   bool    // 是否最小化
	IsFocused    bool    // 是否聚焦
}

func (s *WailsService) WindowChange(opt WindowChangeOpt) {
	w := global.GetWebviewWindow(opt.WindowName)
	if w == nil {
		// 未知窗口，这里应该打印警告
		flog.AppLog.Warn("WailsService.WindowChange", "未知窗口："+opt.WindowName)
		return
	}
	w.NowDpr = opt.NowDpr
	if w.OrgDpr == 0 {
		w.OrgDpr = opt.NowDpr
	}

	flog.AppLog.Debug("WailsService.WindowChange|w", mjson.IndentJson(w))
	flog.AppLog.Debug("WailsService.WindowChange|opt", mjson.IndentJson(opt))
}
