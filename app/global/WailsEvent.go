package global

/*
  这里的方法是传递数据给前端的。所以它被全局到处调用
*/

import (
	"app.local/app/types"
	"github.com/m-startgo/go-utils/mtime"
)

type WailsEventType struct{}

var WailsEvent WailsEventType

func (e WailsEventType) Time() {
	now := mtime.NowDefaultString()
	WailsApp.Event.Emit("Time", now)
}

// 专门用于事件通信，会执行前端的  Events.On('Action')
func (e WailsEventType) Action(opt types.ActionOpt) {
	WailsApp.Event.Emit("Action", opt)
}

func (e WailsEventType) WindowChange(opt types.ActionOpt) {
	WailsApp.Event.Emit("WindowChange", opt)
}

func (e WailsEventType) BallWindowZoom(opt BallWindowType) {
	WailsApp.Event.Emit("BallWindowZoom", opt)
}
