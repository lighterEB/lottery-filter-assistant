package main

import (
	"LotteryFilterAssistant/internal/constants"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 创建应用实例
	myApp := app.New()
	// 创建主窗口
	window := myApp.NewWindow("彩票分析助手")
	// 设置窗口大小
	window.Resize(fyne.NewSize(1024, 768))

	// 创建主要标签页
	tabs := container.NewDocTabs(
		// 数据管理标签及其子标签
		container.NewTabItem(constants.TabDataManage, container.NewAppTabs(
			container.NewTabItem(constants.TabDataImport, widget.NewLabel(constants.TabDataImport)),
			container.NewTabItem(constants.TabDataMaintain, widget.NewLabel(constants.TabDataMaintain)),
		)),

		// 过滤条件标签及其子标签
		container.NewTabItem(constants.TabFilterSetting, container.NewAppTabs(
			container.NewTabItem(constants.TabConditionSetting, widget.NewLabel(constants.TabConditionSetting)),
			container.NewTabItem(constants.TabConditionManage, widget.NewLabel(constants.TabConditionManage)),
		)),

		// 过滤分析标签及其子标签
		container.NewTabItem(constants.TabFilterAnalyze, container.NewAppTabs(
			container.NewTabItem(constants.TabRealTimeAnalyze, widget.NewLabel(constants.TabRealTimeAnalyze)),
			container.NewTabItem(constants.TabBatchAnalyze, widget.NewLabel(constants.TabBatchAnalyze)),
		)),

		// 统计报表标签及其子标签
		container.NewTabItem(constants.TabReport, container.NewAppTabs(
			container.NewTabItem(constants.TabDataStatistics, widget.NewLabel(constants.TabDataStatistics)),
			container.NewTabItem(constants.TabChartDisplay, widget.NewLabel(constants.TabChartDisplay)),
		)),
	)

	// 设置标签位置为左侧
	tabs.SetTabLocation(container.TabLocationLeading)

    // 禁用标签关闭功能
    tabs.CloseIntercept = func(*container.TabItem) {}

	// 设置主窗口内容
	window.SetContent(tabs)

	// 显示并运行
	window.ShowAndRun()
}
