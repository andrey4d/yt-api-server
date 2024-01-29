/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s ApiServer) handlerHello() gin.HandlerFunc {

	return func(c *gin.Context) {
		s.Logger.Info("Hello ", c.Request.Method)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	}
}
