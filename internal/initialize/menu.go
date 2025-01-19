package initialize

import (
	"LotteryFilterAssistant/internal/ui/menu"
	"log"

	"fyne.io/fyne/v2"
	"github.com/BurntSushi/toml"
)

type MenuConfig struct {
	Menu struct {
		Title string     `toml:"title"`
		Items []MenuItem `toml:"items"`
	} `toml:"menu"`
}

type MenuItem struct {
	Label    string     `toml:"label"`
	Shortcut string     `toml:"shortcut,omitempty"`
	Type     string     `toml:"type,omitempty"`
	Submenu  []MenuItem `toml:"submenu, omitempty"`
}

func InitMainMenu(window fyne.Window) {
	var config MenuConfig
	if _, err := toml.DecodeFile("config/menu.toml", &config); err != nil {
		log.Fatal(err)
	}

	handler := menu.NewMenuHandler(window)
	mainMenu := createMainMenu(config.Menu.Items, handler)
	window.SetMainMenu(mainMenu)

}

func createMainMenu(items []MenuItem, handler *menu.MenuHandler) *fyne.MainMenu{
	var menuItems []*fyne.Menu

	for _, item := range items {
		menuItems = append(menuItems, createMenu(item, handler))
	}
	return fyne.NewMainMenu(menuItems...)
}

func createMenu(item MenuItem, handler *menu.MenuHandler) *fyne.Menu {
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