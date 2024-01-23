/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"io"
	"net/http"
)

func (s ApiServer) handlerHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("Hello ", r.Method, r.Host)
		io.WriteString(w, "Hello world!")
	}
}
