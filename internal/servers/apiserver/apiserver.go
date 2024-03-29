/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"os"

	"github.com/andrey4d/ytapiserver/internal/config"
	"github.com/andrey4d/ytapiserver/internal/log"
	"github.com/andrey4d/ytapiserver/internal/servers/handlers"
	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config     *config.Config
	app        fiber.App
	client     *ytclient.Client
	logHandler *log.LogHandl
}

func New(config *config.Config) *ApiServer {

	return &ApiServer{
		config:     config,
		app:        *fiber.New(),
		client:     ytclient.New(),
		logHandler: &log.LogHandl{},
	}
}

func (s *ApiServer) Start() error {

	s.defaultLogger()
	s.ConfigureMiddleware()
	s.ConfigureRouter()
	s.StaticResources()

	address := s.config.Address + ":" + s.config.Port
	s.logHandler.LogModuleInfo("start()").Info("started api server at address ", address)

	return s.app.Listen(address)
}

func (s *ApiServer) ConfigureRouter() {
	s.app.Get("/about", handlers.GetAbout())

	handlerIndex := handlers.IndexHandler{}
	handlerInfo := handlers.InfoHandler{LogHandler: s.logHandler}
	handlerUrls := handlers.UrlsHandler{}
	handlerDescription := handlers.DescriptionHandler{}
	handlerDownload := handlers.DownloadHandler{}

	s.app.Get("/", handlerIndex.GetIndex(s.client))
	s.app.Get("/info/", handlerInfo.GetInfo(s.client))
	s.app.Get("/url/", handlerUrls.GetUrls(s.client))

	s.app.Get("/download/", handlerDownload.GetDownload(s.client))
	s.app.Get("/description/", handlerDescription.GetDescription(s.client))

}

func (s *ApiServer) initHandlers() {

}

func (s *ApiServer) ConfigureMiddleware() {
	s.app.Use(requestid.New())
	s.app.Use(logger.New(logger.Config{
		Format: "${time} | ${locals:requestid} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${query} | ${error}\n",
	}))
}

func (s *ApiServer) defaultLogger() {
	if s.logHandler.Logger == nil {
		s.logHandler.Logger = logrus.New()
		s.logHandler.LogModuleInfo("apiServer").Info("new logger by default")
	}
}

func (s *ApiServer) StaticResources() {
	resources_path, err := os.Getwd()
	if err != nil {
		s.logHandler.Logger.Fatal(err)
	}
	s.app.Static("/static", resources_path+"/web/static")
}
