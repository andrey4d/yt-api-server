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

	s.router.LoadHTMLGlob("web/templates/*.html")

	s.ConfigureRouter()

	address := s.config.Address + ":" + s.config.Port
	s.logger.Info("started api server at address ", address)
	return s.router.Run(address)
}

func (s *ApiServer) ConfigureRouter() {
	s.router.GET("/hello", s.handlerHello())
	s.router.GET("/", s.handlerIndex)
	s.router.POST("/info/", s.handlerInfo)
}
