package global

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

/*
悬浮球的代码
*/

type BallWindowType struct {
	ZoomInW  int    // 放大后的宽度
	ZoomInH  int    // 放大后的高度
	OrgW     int    // 原始宽度
	OrgH     int    // 原始高度
	OrgY     int    // 原始位置Y
	OrgX     int    // 原始位置X
	ScreenW  int    // 屏幕宽度
	ScreenH  int    // 屏幕高度
	ZoomDir  string // 放大方向
	IsPinTop bool   // 是否置顶
}

var BallWindow = BallWindowType{
	ZoomInW: 300,
	ZoomInH: 500,
	OrgW:    100,
	OrgH:    100,
}

func (ball *BallWindowType) Open() {
	if WebWindow.Ball.IsRunner {
		WebWindow.Ball.Window.Show()
		WebWindow.Ball.Window.Focus()
		return
	}

	ball.SetStartSize() // 未运行的时候设置初始位置

	WebWindow.Ball.Window = WailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:            "Ball",
		Width:           ball.OrgW,
		Height:          ball.OrgH,
		InitialPosition: application.WindowXY,
		X:               ball.OrgX,
		Y:               ball.OrgY,
		BackgroundType:  application.BackgroundTypeTransparent,
		URL:             "/#/ball",
		StartState:      application.WindowStateNormal,
		DisableResize:   true,
		Windows: application.WindowsWindow{
			DisableFramelessWindowDecorations: true, // 禁用窗口装饰
			HiddenOnTaskbar:                   true, // 任务栏隐藏
		},
	})
	WebWindow.Ball.EnableFrameless = true // 启用无边框模式

	// 启用监听
	WebWindow.Ball.ListenWindowEvent()

	Tray.OpenBallWindowItem.SetHidden(true)
	Tray.CloseBallWindowItem.SetHidden(false)
	ball.SetPinTop() // 默认置顶
}

func (ball *BallWindowType) SetPinTop() {
	if WebWindow.Ball.Window != nil {
		WebWindow.Ball.Window.SetAlwaysOnTop(true) // 永远置顶
		ball.IsPinTop = true
		ball.ShowSetTopBtn()
	}
}

func (ball *BallWindowType) UnsetPinTop() {
	if WebWindow.Ball.Window != nil {
		WebWindow.Ball.Window.SetAlwaysOnTop(false) // 取消置顶
		ball.IsPinTop = false
		ball.ShowSetTopBtn()
	}
}

func (ball *BallWindowType) ShowSetTopBtn() {
	if ball.IsPinTop {
		Tray.BallSetTopBtn.SetHidden(true)
		Tray.BallUnsetTopBtn.SetHidden(false)
		ContextMenu.BallSetTopBtn.SetHidden(true)
		ContextMenu.BallUnsetTopBtn.SetHidden(false)
	} else {
		Tray.BallSetTopBtn.SetHidden(false)
		Tray.BallUnsetTopBtn.SetHidden(true)
		ContextMenu.BallSetTopBtn.SetHidden(false)
		ContextMenu.BallUnsetTopBtn.SetHidden(true)
	}
}

func (ball *BallWindowType) Close() {
	if WebWindow.Ball.IsRunner {
		WebWindow.Ball.Window.Close()
	}
	Tray.OpenBallWindowItem.SetHidden(false)
	Tray.CloseBallWindowItem.SetHidden(true)

	Tray.BallSetTopBtn.SetHidden(true)
	Tray.BallUnsetTopBtn.SetHidden(true)
	ContextMenu.BallSetTopBtn.SetHidden(true)
	ContextMenu.BallUnsetTopBtn.SetHidden(true)
}

// 设置初始位置初始化悬浮球的位置和大小
func (ball *BallWindowType) SetStartSize() {
	Screen := WailsApp.Screen.GetAll()
	if len(Screen) == 0 {
		return
	}
	ScreenSize := Screen[0].WorkArea
	ball.ScreenW = ScreenSize.Width
	ball.ScreenH = ScreenSize.Height

	// 右下角，离边缘50px
	ball.OrgX = ball.ScreenW - ball.OrgW - 50
	ball.OrgY = ball.ScreenH - ball.OrgH - 50

	if WebWindow.Ball.IsRunner {
		WebWindow.Ball.Window.SetSize(ball.OrgW, ball.OrgH)
		WebWindow.Ball.Window.SetPosition(ball.OrgX, ball.OrgY)
	}

	ball.ZoomDir = "Reset"
}

// 放大悬浮球
func (ball *BallWindowType) BallWindowZoomIn() {
	if WebWindow.Ball.IsRunner {
		ball.UpdateOrgBounds() // 放大之前，先更新原始位置和大小
		if ball.ZoomDir != "Reset" {
			return // 已经放大了，不能重复放大
		}

		X := ball.OrgX
		Y := ball.OrgY
		ScreenW := ball.ScreenW
		ScreenH := ball.ScreenH

		var xDir string
		var yDir string

		// 判断在屏幕的哪个象限，然后放大后调整位置
		if X <= ScreenW/2-ball.ZoomInW/2 {
			// 左半屏 向右展开了 ball.ZoomInW
			xDir = "Right"
		} else {
			// 右半屏 向左展开 ball.ZoomInW
			X = X + ball.OrgW - ball.ZoomInW + 1
			xDir = "Left"
		}

		if Y <= ScreenH/2-ball.ZoomInH/2 {
			// 上半屏，向下展开 ball.ZoomInH
			yDir = "Down"
		} else {
			// 下半屏，向上展开 ball.ZoomInH
			Y = Y + ball.OrgH - ball.ZoomInH + 1
			yDir = "Up"
		}

		WebWindow.Ball.Window.SetSize(ball.ZoomInW, ball.ZoomInH)
		WebWindow.Ball.Window.SetPosition(X, Y)

		var Dir string
		if xDir == "Right" && yDir == "Down" {
			Dir = "RightDown"
			// 不变
		} else if xDir == "Right" && yDir == "Up" {
			Dir = "RightUp"
		} else if xDir == "Left" && yDir == "Down" {
			Dir = "LeftDown"
		} else if xDir == "Left" && yDir == "Up" {
			Dir = "LeftUp"
		}
		ball.ZoomDir = Dir
		WailsEvent.BallWindowZoom(*ball)
	}
}

// 还原始大小
func (ball *BallWindowType) BallWindowReset() {
	if WebWindow.Ball.IsRunner {
		ball.UpdateOrgBounds() // 缩小之前，先更新和大小
		if ball.ZoomDir == "Reset" {
			return // 已经还原了，不能重复还原
		}

		X := ball.OrgX
		Y := ball.OrgY

		if ball.ZoomDir == "RightUp" {
			// 向右向上 放大 ，则 缩小 为 向左向下
			Y = Y + (ball.ZoomInH - ball.OrgH)
		}
		if ball.ZoomDir == "LeftDown" {
			// 向左向下 放大 ，则 缩小 为 向右向上
			X = X + (ball.ZoomInW - ball.OrgW)
		}
		if ball.ZoomDir == "LeftUp" {
			// 向左向上 放大 ，则 缩小 为 向右向下
			X = X + (ball.ZoomInW - ball.OrgW)
			Y = Y + (ball.ZoomInH - ball.OrgH)
		}

		WebWindow.Ball.Window.SetSize(ball.OrgW, ball.OrgH)
		WebWindow.Ball.Window.SetPosition(X, Y)

		ball.OrgX = X
		ball.OrgY = Y
		ball.ZoomDir = "Reset"
		WailsEvent.BallWindowZoom(*ball)
	}
}

func (ball *BallWindowType) UpdateOrgBounds() {
	if WebWindow.Ball.IsRunner {
		rect := WebWindow.Ball.Window.Bounds()
		ball.OrgX = rect.X
		ball.OrgY = rect.Y
	}
}
