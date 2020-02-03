package util

import (
tml "github.com/BurntSushi/toml"
)

func InitConfig(path string) *Config {

	var cfg Config
	if _, err := tml.DecodeFile(path, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}
