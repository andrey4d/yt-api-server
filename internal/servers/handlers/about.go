/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type About struct {
	msg string
}

func GetAbout() gin.HandlerFunc {
	about := About{msg: "About"}
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"about": about.msg})
	}

}
