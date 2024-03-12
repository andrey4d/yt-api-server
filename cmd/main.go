/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package main

import (
	"flag"

	"github.com/andrey4d/ytapiserver/internal/config"
	"github.com/andrey4d/ytapiserver/internal/log"

	"github.com/andrey4d/ytapiserver/internal/servers/apiserver"
	"github.com/andrey4d/ytapiserver/internal/yaml"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "configs/apiserver.yaml", "path to config file")
}

func main() {

	flag.Parse()
	LogHandler := log.New()
	LogHandler.LogModuleInfo("apiServer").Info("config ", configPath)

	config := config.NewConfig()
	yaml.ParseFile(configPath, config, LogHandler)
	LogHandler.LogModuleInfo("apiServer").Info("load config ", config)

	LogHandler.SetLogLevel(config.LogLevel)

	apiServer := apiserver.New(config)
	apiServer.SetLogger(LogHandler.Logger)

	if err := apiServer.Start(); err != nil {
		LogHandler.LogModuleInfo("apiServer").Fatal(err)
	}

}
