/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

type Config struct {
	Port     string `yaml:"port"`
	Address  string `yaml:"address"`
	LogLevel string `yaml:"log_level"`
}

var config *Config

func NewConfig() *Config {
	return &Config{
		Port:     "8080",
		Address:  "127.0.0.1",
		LogLevel: "debug",
	}
}

func GetConfig() *Config {
	if config != nil {
		return config
	}
	return NewConfig()
}

func (c Config) GetLogLevel() string {
	return c.LogLevel
}
