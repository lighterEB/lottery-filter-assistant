package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
)

type MenuHandler struct {
	window fyne.Window
}

func NewMenuHandler(window fyne.Window) *MenuHandler{
	return &MenuHandler{
		window: window,
	}
}

// 菜单动作
func (h *MenuHandler) HandleMenuAction(label string) func() {
	return func() {
		switch label {
		case "退出":
			h.window.Close()
		case "关于":
			dialog.ShowInformation("关于", "彩票过滤助手", h.window)
		default:
			dialog.ShowInformation("提示","功能开发中..."+label, h.window)
		}
	}
}

// 解析快捷键
func (h *MenuHandler) ParseShortcut(shortcut string) fyne.Shortcut{
	switch shortcut {
	case "Ctrl+I":
		return &desktop.CustomShortcut{KeyName: fyne.KeyI, Modifier: fyne.KeyModifierControl}
	case "Ctrl+E":
		return &desktop.CustomShortcut{KeyName: fyne.KeyE, Modifier: fyne.KeyModifierControl}
	case "Ctrl+Q":
		return &desktop.CustomShortcut{KeyName: fyne.KeyQ, Modifier: fyne.KeyModifierControl}
	default:
		return nil

	}
}