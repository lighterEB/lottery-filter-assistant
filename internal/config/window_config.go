package config

import (
	"LotteryFilterAssistant/internal/constants"
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
)

// 窗口配置结构
type WindowConfig struct {
	Window struct {
		Title     string  `toml:"title"`
		Width     float32 `toml:"width"`
		Height    float32 `toml:"height"`
		MinWidth  float32 `toml:"min_width"`
		MinHeight float32 `toml:"min_height"`
	} `toml:"window"`
}

// 加载窗口配置
func LoadWindowConfig() *WindowConfig {
	var wincfg WindowConfig
	if _, err := toml.DecodeFile(constants.CONFIG_PATH, &wincfg); err != nil {
		log.Printf("加载配置出错：%v, 使用默认配置", err)
		return getDefaultConfig()
	}
	return &wincfg
}

// 获取默认配置
func getDefaultConfig() *WindowConfig {
	return &WindowConfig{
		Window: struct {
			Title     string  `toml:"title"`
			Width     float32 `toml:"width"`
			Height    float32 `toml:"height"`
			MinWidth  float32 `toml:"min_width"`
			MinHeight float32 `toml:"min_height"`
		}{
			Width:     defaultWidth,
			Height:    defaultHeight,
			MinWidth:  defaultMinWidth,
			MinHeight: defaultMinHeight,
			Title:     defaultTitle,
		},
	}
}
