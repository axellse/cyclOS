package config

import (
	_ "embed"
	"encoding/json"
)

type DisplayConfig struct {
	//HideDrawing will disable the backlight when BeginDisplay is called, enabling it again when Display is called, hiding any tearing during drawing.
	HideDrawing bool
}

type Config struct {
	//Display is passed to the display driver
	Display DisplayConfig
}

//go:embed config.json
var configRaw []byte

func LoadConfig() Config {
	config := Config{}
	err := json.Unmarshal(configRaw, &config)
	if err != nil {
		panic(err)
	}

	return config
}