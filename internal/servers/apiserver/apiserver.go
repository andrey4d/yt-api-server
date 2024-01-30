/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"github.com/andrey4d/ytapiserver/internal/config"
	"github.com/andrey4d/ytapiserver/internal/servers/handlers"
	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config *config.Config
	Logger *logrus.Logger
	app    fiber.App
	client *ytclient.Client
}

func New(config *config.Config) *ApiServer {

	viewEngine := html.New(config.Template_dir, config.Template_type)

	fiberConfig := fiber.Config{Views: viewEngine}

	return &ApiServer{
		config: config,
		app:    *fiber.New(fiberConfig),
		client: ytclient.New(),
	}
}

func (s *ApiServer) Start() error {

	if s.Logger == nil {
		s.Logger = logrus.New()
		s.Logger.Info("apiServer new default logger")
	}
	s.app.Use(logger.New())
	// s.router.Use(Logger(s.Logger))
	// s.router.Use(requestLoggingMiddleware(s.Logger))

	// s.router.Use(middlware.RequestIDMiddleware)

	s.ConfigureRouter()

	address := s.config.Address + ":" + s.config.Port
	s.Logger.Info("started api server at address ", address)
	return s.app.Listen(address)
}

func (s *ApiServer) ConfigureRouter() {
	s.app.Get("/hello", handlers.GetHello())
	s.app.Get("/", handlers.HandlerIndex)
	s.app.Post("/info/", handlers.GetInfo(s.client))
	s.app.Post("/url/", handlers.GetUrls(s.client))
	s.app.Get("/about", handlers.GetAbout())
}
