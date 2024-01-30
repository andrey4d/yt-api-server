/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5/middleware"
)

func (s *ApiServer) SetLogger(logger *logrus.Logger) {
	logger.Info("apiServer set main logger")
	s.Logger = logger
}

func (s ApiServer) GetLogger() *logrus.Logger {
	return s.Logger
}

func (s ApiServer) GetMwLogger(log *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log.WithField("component", "middleware/logger")

		log.Info("logger middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.WithFields(
				logrus.Fields{
					"method":      r.Method,
					"path":        r.URL.Path,
					"remote_addr": r.RemoteAddr,
					"user_agent":  r.UserAgent(),
					"request_id":  middleware.GetReqID(r.Context()),
				},
			)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()

			defer func() {
				entry.Infof("request completed (status: %d bytes: %d duration: %s)\n", ww.Status(), ww.BytesWritten(), time.Since(t1).String())
			}()
			// Next middleware
			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}

}

// entry := log.WithFields(
// 	logrus.Fields{
// 		"method":      ctx.Request.Method,
// 		"path":        ctx.Request.URL.Path,
// 		"remote_addr": ctx.Request.RemoteAddr,
// 		"user_agent":  ctx.Request.UserAgent(),
// 		"request_id":  ctx.Request.Header.Values("X-Request-ID"),
// 	},
// )
