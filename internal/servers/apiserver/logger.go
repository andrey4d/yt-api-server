/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"github.com/sirupsen/logrus"
)

func (s *ApiServer) SetLogger(logger *logrus.Logger) {

	s.logHandler.Logger = logger
	s.logHandler.LogModuleInfo("apiServer").Info("set main logger")
}
