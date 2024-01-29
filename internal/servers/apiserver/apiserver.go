/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"github.com/andrey4d/ytapiserver/internal/config"
	"github.com/andrey4d/ytapiserver/internal/servers/handlers"
	"github.com/andrey4d/ytapiserver/internal/ytclient"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *config.Config
	Logger *logrus.Logger
	router *gin.Engine
	client *ytclient.Client
}

func New(config *config.Config) *ApiServer {

	return &ApiServer{
		config: config,
		router: gin.Default(),
		client: ytclient.New(),
	}
}

func (s *ApiServer) Start() error {

	if s.Logger == nil {
		s.Logger = logrus.New()
		s.Logger.Info("apiServer new default logger")
	}

	s.router.LoadHTMLGlob("web/templates/*.html")

	s.ConfigureRouter()

	address := s.config.Address + ":" + s.config.Port
	s.Logger.Info("started api server at address ", address)
	return s.router.Run(address)
}

func (s ApiServer) ConfigureRouter() {
	s.router.GET("/hello", s.handlerHello())
	s.router.GET("/", handlers.HandlerIndex)
	s.router.POST("/info/", handlers.GetInfo(s.client))
	s.router.POST("/url/", handlers.GetUrls(s.client))
	s.router.GET("/about", handlers.GetAbout())
}
