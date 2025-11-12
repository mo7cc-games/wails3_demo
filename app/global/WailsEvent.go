package global

/*
这里的方法是传递数据给前端的。所以它会 go 被全局到处调用

*/

import (
	"github.com/m-startgo/go-utils/mtime"
)

type WailsEventType struct{}

var WailsEvent WailsEventType

func (e WailsEventType) Time() {
	now := mtime.NowDefaultString()
	WailsApp.Event.Emit("Time", now)
}

// 专门用于事件通信，会执行前端的  Events.On('Action')
func (e WailsEventType) Action(WindowName string, ActionName string) {
	WailsApp.Event.Emit("Action", WindowName, ActionName)
}

// 窗口变化事件
type WindowChangeParam struct {
	WindowName string
	Action     string
}

func (e WailsEventType) WindowChange(param WindowChangeParam) {
	WailsApp.Event.Emit("WindowChange", param)
}
