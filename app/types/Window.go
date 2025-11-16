package types

/*
  存放窗口相关的 公共 类型
*/

type WindowStatus struct {
	WindowName   string  // 窗口名称
	Width        int     // 宽度
	Height       int     // 高度
	AvailWidth   int     // 可用宽度
	AvailHeight  int     // 可用高度
	ScreenLeft   int     // 屏幕左上角位置
	ScreenTop    int     // 屏幕左上角位置
	InnerWidth   int     // 内部宽度
	InnerHeight  int     // 内部高度
	Dpr          float32 // 当前 DPR
	IsFullscreen bool    // 是否全屏
	IsExtended   bool    // 是否是扩展屏幕
	IsMaximise   bool    // 是否最大化
	IsMinimise   bool    // 是否最小化
	IsFocused    bool    // 是否聚焦
}
