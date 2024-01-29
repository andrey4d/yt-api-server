/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"fmt"
	"net/http"

	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/gin-gonic/gin"
)

func GetUrls(c *ytclient.Client) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tag := ctx.PostForm("films")

		format, err := c.GetFormat(tag)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "%v", err)
		}

		link_text := fmt.Sprintf("Quality=%s, audio channels=%d, video=%s", format.QualityLabel, format.AudioChannels, format.MimeType)

		ctx.String(http.StatusOK, "<a href='%s'>%s</a>", format.URL, link_text)
	}
}
