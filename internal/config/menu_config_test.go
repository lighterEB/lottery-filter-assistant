package config

import "testing"

func TestMenu(t *testing.T) {
	menuCfg := LoadMenuConfig()
	println(&menuCfg)
}