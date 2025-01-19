package initialize

import (
	"LotteryFilterAssistant/internal/config"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func InitWindow(window fyne.Window) {
	windowConfig := config.LoadWindowConfig()
	// 窗口居中
	window.CenterOnScreen()
	// 设置标题
	window.SetTitle(windowConfig.Title)
	window.Resize(fyne.NewSize(windowConfig.Width, windowConfig.Height))
	bg := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	bg.SetMinSize(fyne.NewSize(windowConfig.MinWidth, windowConfig.MinHeight))
	window.SetContent(bg)

}
