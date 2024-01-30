/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package log

import (
	"github.com/sirupsen/logrus"
)

type LogHandler struct {
	Logger *logrus.Logger
}

var logHandler LogHandler

func New() *LogHandler {
	logger := logrus.New()
	entry := logger.WithFields(logrus.Fields{"module": "apiServer"})
	_ = entry
	return &LogHandler{
		Logger: logger,
	}
}

func (l *LogHandler) GetLogger() *logrus.Logger {
	if l.Logger != nil {
		logrus.Info("log reuse")
		return l.Logger
	}
	return New().Logger
}

func (l *LogHandler) SetLogLevel(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	l.Logger.SetLevel(level)
	return nil
}
