/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"net/http"

	"github.com/andrey4d/ytapiserver/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	config      *config.Config
	logger      *logrus.Logger
	router      *chi.Mux
	middlewares *chi.Middlewares
}

func New(config *config.Config) *ApiServer {
	return &ApiServer{
		config: config,
		router: chi.NewRouter(),
	}
}

func (s *ApiServer) Start() error {

	if s.logger == nil {
		s.logger = logrus.New()
		s.logger.Info("apiServer new default logger")
	}

	address := s.config.Address + ":" + s.config.Port
	s.logger.Info("started api server at address ", address)

	// TODO: middleware
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(s.GetMwLogger(s.logger))
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.URLFormat)

	s.ConfigureRouter()

	return http.ListenAndServe(address, s.router)
}

func (s *ApiServer) ConfigureRouter() {

	s.router.HandleFunc("/hello", s.handlerHello())
	s.router.HandleFunc("/info", s.handlerInfo())
}
