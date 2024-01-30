/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"github.com/andrey4d/ytapiserver/internal/config"
	"github.com/andrey4d/ytapiserver/internal/log"
	"github.com/andrey4d/ytapiserver/internal/servers/handlers"
	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/gofiber/template/html/v2"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config     *config.Config
	app        fiber.App
	client     *ytclient.Client
	logHandler *log.LogHandler
}

func New(config *config.Config) *ApiServer {

	viewEngine := html.New(config.Template_dir, config.Template_type)

	fiberConfig := fiber.Config{Views: viewEngine}

	return &ApiServer{
		config:     config,
		app:        *fiber.New(fiberConfig),
		client:     ytclient.New(),
		logHandler: &log.LogHandler{},
	}
}

func (s *ApiServer) Start() error {

	if s.logHandler.Logger == nil {
		s.logHandler.Logger = logrus.New()
		s.logHandler.LogModuleInfo("apiServer").Info("new logger by default")
	}

	s.app.Use(requestid.New())

	s.app.Use(logger.New(logger.Config{
		Format: "${time} | ${locals:requestid} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${query} | ${error}\n",
	}))

	// s.router.Use(Logger(s.Logger))
	// s.router.Use(requestLoggingMiddleware(s.Logger))

	// s.router.Use(middlware.RequestIDMiddleware)

	s.ConfigureRouter()

	address := s.config.Address + ":" + s.config.Port
	s.logHandler.LogModuleInfo("start()").Info("started api server at address ", address)
	return s.app.Listen(address)
}

func (s *ApiServer) ConfigureRouter() {
	s.app.Get("/hello", handlers.GetHello(s.logHandler))
	s.app.Get("/", handlers.HandlerIndex)
	s.app.Post("/info/", handlers.GetInfo(s.client))
	s.app.Post("/url/", handlers.GetUrls(s.client))
	s.app.Get("/about", handlers.GetAbout())

}
