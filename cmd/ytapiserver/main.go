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
	logHandler := log.New()
	logger := logHandler.Logger
	logger.Info("config ", configPath)

	config := config.NewConfig()
	yaml.ParseFile(configPath, config)
	logHandler.SetLogLevel(config.LogLevel)

	apiServer := apiserver.New(config)
	apiServer.SetLogger(logger)

	if err := apiServer.Start(); err != nil {
		logger.Fatal(err)
	}

}
