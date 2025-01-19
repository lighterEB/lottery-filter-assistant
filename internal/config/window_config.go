package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// 定义默认值
var (
	defaultWidth     float32 = 1024
	defaultHeight    float32 = 800
	defaultMinWidth  float32 = 800
	defaultMinHeight float32 = 640
	defaultTitle     string  = "彩票过滤助手"
	// 配置文件默认路径
	windowConfigPath string = "config/window.toml"
)

// 窗口配置结构
type WindowConfig struct {
	AppId     string  `toml:"app_id"`
	Title     string  `toml:"title"`
	Width     float32 `toml:"width"`
	Height    float32 `toml:"height"`
	MinWidth  float32 `toml:"min_width"`
	MinHeight float32 `toml:"min_height"`
}

// 加载窗口配置
func LoadWindowConfig() WindowConfig {
	var wincfg WindowConfig
	if _, err := toml.DecodeFile(windowConfigPath, &wincfg); err != nil {
		log.Printf("加载配置出错：%v, 使用默认配置", err)
		return getDefaultConfig()
	}
	return wincfg
}

// 获取默认配置
func getDefaultConfig() WindowConfig {
	return WindowConfig{
		Width:     defaultWidth,
		Height:    defaultHeight,
		MinWidth:  defaultMinWidth,
		MinHeight: defaultMinHeight,
		Title:     defaultTitle,
	}
}
