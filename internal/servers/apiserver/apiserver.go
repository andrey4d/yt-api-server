/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"github.com/andrey4d/ytapiserver/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *config.Config
	logger *logrus.Logger
	router *gin.Engine
}

func New(config *config.Config) *ApiServer {

	return &ApiServer{
		config: config,
		router: gin.Default(),
	}
}

func (s *ApiServer) Start() error {

	if s.logger == nil {
		s.logger = logrus.New()
		s.logger.Info("apiServer new default logger")
	}

	// TODO: middleware
	// s.router.Use(middleware.RequestID)
	// s.router.Use(middleware.RealIP)
	// s.router.Use(s.GetMwLogger(s.logger))
	// s.router.Use(middleware.Recoverer)
	// s.router.Use(middleware.URLFormat)

	s.ConfigureRouter()

	address := s.config.Address + ":" + s.config.Port
	s.logger.Info("started api server at address ", address)
	return s.router.Run(address)
}

func (s *ApiServer) ConfigureRouter() {

	s.router.GET("/hello", s.handlerHello())
	s.router.GET("/info", s.handlerInfo)
}
