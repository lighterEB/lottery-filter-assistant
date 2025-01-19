package initialize

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/BurntSushi/toml"
)

// 窗口配置结构
type WindowConfig struct {
	Window struct {
		Title     string  `toml:"title"`
		Width     float32 `toml: width`
		Height    float32 `toml: height`
		MinWidth  float32 `toml: min_width`
		MinHeight float32 `toml: min_height`
	} `toml:"window"`
}

// 初始化应用
func InitApp() (fyne.App, fyne.Window) {
	// 读取窗口配置
	var config WindowConfig
	if _, err := toml.DecodeFile("config/window.toml", &config); err != nil {
		log.Fatal(err)
	}

	app := app.New()
	w := app.NewWindow(config.Window.Title)

	// 设置窗口大小
	w.SetFixedSize(false)
	w.Resize(fyne.NewSize(config.Window.Width, config.Window.Height))

	return app, w
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
