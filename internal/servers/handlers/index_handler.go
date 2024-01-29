/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)
}
