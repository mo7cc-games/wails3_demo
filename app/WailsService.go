package app

import (
	"app.local/app/global"
	"app.local/app/types"
	"app.local/app/utils/flog"
	"github.com/m-startgo/go-utils/mjson"
	"github.com/m-startgo/go-utils/mmath"
)

/*
 这里的方法是暴露给前端调用的，禁止被其它 go 模块直接引入
*/

type WailsService struct{}

// 加法器
func (s *WailsService) Add(n1 float64, n2 float64) string {
	a := mmath.NewFromFloat(n1)
	b := mmath.NewFromFloat(n2)
	result := a.Add(b)
	return result.String()
}

// Action 通信
func (s *WailsService) Action(opt types.ActionOpt) {
	global.WailsEvent.Action(opt)
}

// 暴露给前端的用户获取当前窗口信息的方法
func (s *WailsService) GetWindowInfo(WindowName string) global.WebviewWindow {
	w := global.GetWebviewWindow(WindowName)
	if w == nil {
		// 未知窗口，这里应该打印警告
		go flog.AppLog.Warn("WailsService.GetWindowInfo", "未知窗口："+WindowName)
		return global.WebviewWindow{}
	}
	return *w
}

// 获取所有窗口的信息
func (s *WailsService) GetAllWindowInfo() global.WebWindowType {
	return global.WebWindow
}

// 打开测试窗口
func (s *WailsService) OpenTestWindow() {
	global.OpenTestWindow()
}

// 打开悬浮球窗口
func (s *WailsService) OpenBallWindow() {
	global.BallWindow.Open()
}

// 放大悬浮球
func (s *WailsService) BallWindowZoomIn() {
	global.BallWindow.BallWindowZoomIn()
}

// 还原悬浮球
func (s *WailsService) BallWindowReset() {
	global.BallWindow.BallWindowReset()
}

func (s *WailsService) GetBallWindowZoomDir() global.BallWindowType {
	return global.BallWindow
}

// 从前端传来的窗口变化的信息
func (s *WailsService) WindowChange(opt types.WindowStatus) {
	w := global.GetWebviewWindow(opt.WindowName)
	if w == nil {
		// 未知窗口，这里应该打印警告
		go flog.AppLog.Warn("WailsService.WindowChange", "未知窗口："+opt.WindowName)
		return
	}
	w.NowDpr = opt.Dpr
	if w.OrgDpr == 0 {
		w.OrgDpr = opt.Dpr
	}

	go flog.AppLog.Debug("WailsService.WindowChange|w", mjson.IndentJson(w))
	go flog.AppLog.Debug("WailsService.WindowChange|opt", mjson.IndentJson(opt))
}
