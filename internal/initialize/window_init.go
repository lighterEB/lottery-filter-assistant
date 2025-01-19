package initialize

import (
	"LotteryFilterAssistant/internal/config"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func InitWindow(app fyne.App, winCfg *config.WindowConfig) fyne.Window {
	window := app.NewWindow(winCfg.Window.Title)
	// 窗口居中
	window.CenterOnScreen()
	// 设置标题
	window.Resize(fyne.NewSize(winCfg.Window.Width, winCfg.Window.Height))
	bg := canvas.NewRectangle(color.NRGBA{0, 0, 0, 0})
	bg.SetMinSize(fyne.NewSize(winCfg.Window.MinWidth, winCfg.Window.MinHeight))
	window.SetContent(bg)

	return window
}
