package initialize

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// 初始化应用
func InitApp() {
	app := app.New()
	w := app.NewWindow("")
	InitLayout(w)
	InitWindow(w)
	InitMainMenu(w)
	w.ShowAndRun()
}

// 创建一个保持纵横比的容器
func NewAspectContainer(content fyne.CanvasObject, aspect float32) *fyne.Container {
	return container.NewWithoutLayout(content)
}

// 初始化布局
func InitLayout(window fyne.Window) fyne.CanvasObject {
	// 创建主容器
	mainContainer := container.NewBorder(nil, nil, nil, nil)

	// 设置内容自适应
	window.SetContent(mainContainer)

	return mainContainer
}
