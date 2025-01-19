package initialize

import (
	"LotteryFilterAssistant/internal/config"
	"LotteryFilterAssistant/internal/ui/menu"

	"fyne.io/fyne/v2"
)

func InitMainMenu(window fyne.Window) {
	menuConfig := config.LoadMenuConfig()
	handler := menu.NewMenuHandler(window)
	mainMenu := createMainMenu(menuConfig.Menu.Items, handler)
	window.SetMainMenu(mainMenu)
}

func createMainMenu(items []config.MenuItem, handler *menu.MenuHandler) *fyne.MainMenu {
	var menuItems []*fyne.Menu

	for _, item := range items {
		menuItems = append(menuItems, createMenu(item, handler))
	}
	return fyne.NewMainMenu(menuItems...)
}

func createMenu(item config.MenuItem, handler *menu.MenuHandler) *fyne.Menu {
	var items []*fyne.MenuItem

	if len(item.Submenu) > 0 {
		for _, subItem := range item.Submenu {
			if subItem.Type == "separator" {
				items = append(items, fyne.NewMenuItemSeparator())
				continue
			}

			menuItem := fyne.NewMenuItem(subItem.Label, handler.HandleMenuAction(subItem.Label))
			if subItem.Shortcut != "" {
				menuItem.Shortcut = handler.ParseShortcut(subItem.Shortcut)
			}
			items = append(items, menuItem)
		}
	}
	return fyne.NewMenu(item.Label, items...)
}
