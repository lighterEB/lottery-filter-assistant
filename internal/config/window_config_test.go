package config

import (
	"fmt"
	"testing"
)

func TestWindowConfig(t *testing.T) {
	wCfg := LoadWindowConfig()
	fmt.Printf("%+v\n", wCfg)
}