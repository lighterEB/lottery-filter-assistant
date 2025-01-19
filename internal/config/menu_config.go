package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var (
	defaultMenuItem MenuConfig = MenuConfig{
		Items: []MenuItem{
			{
				Label:    "帮助",
				Shortcut: "",
				Type:     "",
				Submenu: []MenuItem{
					{
						Label:    "说明",
						Shortcut: "",
						Type:     "",
						Submenu:  nil,
					},
				},
			},
		},
	}
	menuConfigPath string = "config/menu.toml"
)

type MenuConfig struct {
	Items []MenuItem `toml:"items"`
}

type MenuItem struct {
	Label    string     `toml:"label"`
	Shortcut string     `toml:"shortcut,omitempty"`
	Type     string     `toml:"type,omitempty"`
	Submenu  []MenuItem `toml:"submenu, omitempty"`
}

func LoadMenuConfig() MenuConfig {
	var menuCfg MenuConfig
	if _,err := toml.DecodeFile(menuConfigPath, &menuCfg); err != nil {
		log.Printf("加载菜单配置失败: %v, 使用默认配置", err)
		return getDefaultMenuConfig()
	}
	return menuCfg
}

func getDefaultMenuConfig() MenuConfig{
	return MenuConfig{
		Items: defaultMenuItem.Items,
	}
}