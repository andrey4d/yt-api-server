/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAbout(t *testing.T) {
	tests := []struct {
		name         string
		route        string
		expectedCode int
	}{
		{
			name:         "About",
			route:        "/about",
			expectedCode: 200,
		},
		{
			name:         "About wrong",
			route:        "/info",
			expectedCode: 404,
		},
	}

	app := fiber.New()
	app.Get("/about", GetAbout())

	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.route, nil)
		resp, _ := app.Test(req, 1)
		assert.Equalf(t, tt.expectedCode, resp.StatusCode, tt.name)
	}
}
