/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(log *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.WithField("component", "middleware/logger")
		// log.Info("logger middleware enabled")
		ctx.Next()
	}
}

func requestLoggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the request body
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		bodyString := string(bodyBytes)

		// Process request
		c.Next()

		// Log request details
		// Include Request error message
		logger.WithFields(logrus.Fields{
			"status":       c.Writer.Status(),
			"method":       c.Request.Method,
			"path":         c.Request.URL.Path,
			"query_params": c.Request.URL.Query(),
			"ID":           c.Request.Header.Get("X-Request-ID"),
			"req_body":     bodyString,
			// "res_error_1":        c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"res_error_2": c.Errors.String(),
		}).Info("Request details")
	}
}
