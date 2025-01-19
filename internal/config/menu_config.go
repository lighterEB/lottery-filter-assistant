package config

import (
	"LotteryFilterAssistant/internal/constants"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	defaultMenuItem MenuConfig = MenuConfig{
		Menu: struct {
			Items []MenuItem `toml:"items"`
		}{
			Items: []MenuItem{
				{
					Label:    "帮助",
					Shortcut: "",
					Type:     "",
					Submenu: []MenuItem{
						{
							Label: "说明",
						},
						{
							Label: "退出",
						},
					},
				},
			},
		},
	}
)

type MenuConfig struct {
	Menu struct {
		Items []MenuItem `toml:"items"`
	} `toml:"menu"`
}

type MenuItem struct {
	Label    string     `toml:"label"`
	Shortcut string     `toml:"shortcut,omitempty"`
	Type     string     `toml:"type,omitempty"`
	Submenu  []MenuItem `toml:"submenu,omitempty"`
}

func LoadMenuConfig() MenuConfig {
	var menuCfg MenuConfig
	if _, err := toml.DecodeFile(constants.CONFIG_PATH, &menuCfg); err != nil {
		log.Printf("加载菜单配置失败: %v, 使用默认配置", err)
		return getDefaultMenuConfig()
	}
	return menuCfg
}

func getDefaultMenuConfig() MenuConfig {
	return MenuConfig{
		Menu: struct {
			Items []MenuItem `toml:"items"`
		}{
			Items: defaultMenuItem.Menu.Items,
		},
	}
}
