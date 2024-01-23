/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func (s *ApiServer) SetLogger(logger *logrus.Logger) {
	logger.Info("apiServer set main logger")
	s.logger = logger
}

func (s ApiServer) GetLogger() *logrus.Logger {
	return s.logger
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
