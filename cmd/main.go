package main

import (
	"LotteryFilterAssistant/internal/initialize"
)

func main() {
	// 初始化应用和窗口
	_, window := initialize.InitApp()

	// 初始化布局
	initialize.InitLayout(window)
	initialize.InitMainMenu(window)

	window.ShowAndRun()
}
