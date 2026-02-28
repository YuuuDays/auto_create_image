package config

import (
	"encoding/json"
	"os"
)

// Config は設定を管理
type Config struct {
	PromptOrder []string `json:"prompt_order"`
}

// Load は設定ファイルを読み込む
func Load(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
