/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrey4d/ytapiserver/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestApiServer_handlerHello(t *testing.T) {
	s := New(config.NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handlerHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello world!")
}
