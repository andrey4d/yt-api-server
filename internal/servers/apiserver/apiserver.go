/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

var apiServer *ApiServer

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func GetApiServer(config *Config) *ApiServer {
	if apiServer != nil {
		return apiServer
	}
	return New(config)
}

func (s ApiServer) Start() error {

	s.logger.Info(s.config)

	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	address := s.config.Address + ":" + s.config.Port
	s.logger.Info("started api server at address ", address)

	s.ConfigureRouter()

	return http.ListenAndServe(address, s.router)
}

func (s ApiServer) ConfigureRouter() {
	s.router.HandleFunc("/hello", s.handlerHello())
	s.router.HandleFunc("/info", s.handlerInfo())
}
