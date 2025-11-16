package global

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

/*
 打开窗口相关的函数
*/

func OpenMainWindow() {
	if WebWindow.Main.IsRunner {
		WebWindow.Main.Window.Show()
		WebWindow.Main.Window.Focus()
		return
	}

	WebWindow.Main.Window = WailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:            "Main",
		Width:           800,
		Height:          600,
		InitialPosition: application.WindowCentered,
		BackgroundType:  application.BackgroundTypeTransparent,
		URL:             "/",
		Windows: application.WindowsWindow{
			DisableFramelessWindowDecorations: true, // 禁用窗口装饰
		},
	})
	WebWindow.Main.EnableFrameless = true // 启用无边框模式

	WebWindow.Main.ListenWindowEvent()
}

func OpenTestWindow() {
	if WebWindow.Test.IsRunner {
		WebWindow.Test.Window.Show()
		WebWindow.Test.Window.Focus()
		return
	}
	WebWindow.Test.Window = WailsApp.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:            "Test",
		Width:           600,
		Height:          400,
		InitialPosition: application.WindowCentered,
		BackgroundType:  application.BackgroundTypeTransparent,
		URL:             "/#/test",
		Windows: application.WindowsWindow{
			DisableFramelessWindowDecorations: true, // 禁用窗口装饰
		},
	})
	WebWindow.Test.EnableFrameless = true // 启用无边框模式
	// 启用监听
	WebWindow.Test.ListenWindowEvent()
}
