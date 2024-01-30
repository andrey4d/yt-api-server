/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package log

import (
	"github.com/sirupsen/logrus"
)

type LogHandl struct {
	Logger *logrus.Logger
}

func New() *LogHandl {
	logger := logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logger.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	return &LogHandl{
		Logger: logger,
	}
}

func (l *LogHandl) GetLogger() *logrus.Logger {
	if l.Logger != nil {
		l.LogModuleInfo("log handler").Info("log reuse")
		return l.Logger
	}
	return New().Logger
}

func (l *LogHandl) SetLogLevel(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	l.Logger.SetLevel(level)
	return nil
}

func (l *LogHandl) LogModuleInfo(moduleName string) *logrus.Entry {
	return l.Logger.WithField("module", moduleName)
}
