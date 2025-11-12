package global

import (
	"app.local/app/utils/flog"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

// 定义有哪些窗口
type WebWindowType struct {
	Main WebviewWindow
	Test WebviewWindow
}

// GetWebviewWindow 返回指定名字的 WebviewWindow 指针（不区分大小写）。
func GetWebviewWindow(name string) *WebviewWindow {
	switch name {
	case "Main", "main":
		return &WebWindow.Main
	case "Test", "test":
		return &WebWindow.Test
	default:
		return nil
	}
}

// 用于存储全局窗口实例
var WebWindow WebWindowType

// 定义窗口结构体
type WebviewWindow struct {
	Window          *application.WebviewWindow
	Name            string  // 窗口名称
	IsListen        bool    // 是否已启用监听事件
	IsRunner        bool    // 是否正在运行
	IsFullscreen    bool    // 是否全屏
	NowDpr          float32 // 当前窗口的 DPR 值
	OrgDpr          float32 // 原始 DPR 值
	EnableFrameless bool    // 启用无边框模式
}

func (s *WebviewWindow) ListenWindowEvent() {
	if s.IsListen {
		// 避免重复监听
		return
	}
	s.IsListen = true
	s.Name = s.Window.Name()
	// 窗口正在关闭
	s.Window.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		s.IsRunner = false
		s.IsListen = false // 重置监听状态
		flog.AppLog.Debug("global.ListenWindowEvent", s.Name, "窗口正在关闭")
	})
	// 窗口运行时准备就绪
	s.Window.OnWindowEvent(events.Common.WindowRuntimeReady, func(e *application.WindowEvent) {
		s.IsRunner = true
		flog.AppLog.Debug("global.ListenWindowEvent", s.Name, "窗口运行时准备就绪")
	})

	// 窗口显示
	s.Window.OnWindowEvent(events.Common.WindowShow, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowShow",
		})
	})

	// 窗口进入全屏
	s.Window.OnWindowEvent(events.Common.WindowFullscreen, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowFullscreen",
		})
	})

	// 窗口退出全屏
	s.Window.OnWindowEvent(events.Common.WindowUnFullscreen, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowUnFullscreen",
		})
	})

	// 窗口已移动
	s.Window.OnWindowEvent(events.Common.WindowDidMove, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowDidMove",
		})
	})
	// 窗口已调整大小
	s.Window.OnWindowEvent(events.Common.WindowDidResize, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowDidResize",
		})
	})

	// 窗口 DPI 改变
	s.Window.OnWindowEvent(events.Common.WindowDPIChanged, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowDPIChanged",
		})
	})

	// 窗口被隐藏
	s.Window.OnWindowEvent(events.Common.WindowHide, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowHide",
		})
	})

	// 窗口最大化
	s.Window.OnWindowEvent(events.Common.WindowMaximise, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowMaximise",
		})
	})

	// 窗口最小化
	s.Window.OnWindowEvent(events.Common.WindowMinimise, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowMinimise",
		})
	})

	// 窗口退出最大化
	s.Window.OnWindowEvent(events.Common.WindowUnMaximise, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowUnMaximise",
		})
	})

	// 窗口退出最小化
	s.Window.OnWindowEvent(events.Common.WindowUnMinimise, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowUnMinimise",
		})
	})

	// 窗口恢复（从最小化/最大化状态）
	s.Window.OnWindowEvent(events.Common.WindowRestore, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowRestore",
		})
	})

	// 窗口切换为无边框模式
	s.Window.OnWindowEvent(events.Common.WindowToggleFrameless, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowToggleFrameless",
		})
	})

	// 窗口获得焦点
	s.Window.OnWindowEvent(events.Common.WindowFocus, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowFocus",
		})
	})
	// 窗口失去焦点
	s.Window.OnWindowEvent(events.Common.WindowLostFocus, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowLostFocus",
		})
	})

	// 窗口缩放
	s.Window.OnWindowEvent(events.Common.WindowZoom, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowZoom",
		})
	})

	// 窗口放大
	s.Window.OnWindowEvent(events.Common.WindowZoomIn, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowZoomIn",
		})
	})

	// 窗口缩小
	s.Window.OnWindowEvent(events.Common.WindowZoomOut, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowZoomOut",
		})
	})

	// 窗口缩放重置
	s.Window.OnWindowEvent(events.Common.WindowZoomReset, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowZoomReset",
		})
	})

	// 文件被拖放到窗口的指定区域
	s.Window.OnWindowEvent(events.Common.WindowDropZoneFilesDropped, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowDropZoneFilesDropped",
		})
	})

	// 文件被拖放到窗口
	s.Window.OnWindowEvent(events.Common.WindowFilesDropped, func(e *application.WindowEvent) {
		WailsEvent.WindowChange(WindowChangeParam{
			WindowName: s.Name,
			Action:     "WindowFilesDropped",
		})
	})
}
