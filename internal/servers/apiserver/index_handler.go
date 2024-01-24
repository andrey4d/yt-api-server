/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s ApiServer) handlerIndex(c *gin.Context) {
	Formats := map[string][]VideoFormat{
		"Formats": []VideoFormat{},
	}
	c.HTML(http.StatusOK, "index.html", Formats)
}
