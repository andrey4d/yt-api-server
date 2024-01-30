/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"github.com/sirupsen/logrus"
)

func (s *ApiServer) SetLogger(logger *logrus.Logger) {
	logger.Info("apiServer set main logger")
	s.Logger = logger
}

func (s ApiServer) GetLogger() *logrus.Logger {
	return s.Logger
}
