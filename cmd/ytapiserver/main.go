/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package main

import (
	"flag"
	"github.com/andrey4d/ytapiserver/internal/log"

	"github.com/andrey4d/ytapiserver/internal/servers/apiserver"
	"github.com/andrey4d/ytapiserver/internal/yaml"
)

var configPath string
var LogLevel string

func init() {
	flag.StringVar(&configPath, "config", "configs/apiserver.yaml", "path to config file")
}

func main() {
	flag.Parse()
	log.Info("config ", configPath)

	config := apiserver.NewConfig()
	yaml.ParseFile(configPath, config)

	apiServer := apiserver.New(config)
	if err := apiServer.Start(); err != nil {
		log.Fatal(err)
	}

}
