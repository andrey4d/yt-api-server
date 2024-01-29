/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"net/http"

	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/gin-gonic/gin"
	"github.com/kkdai/youtube/v2"
)

func GetInfo(c *ytclient.Client) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		url := ctx.PostForm("video_tag")
		if url == "" {
			url = "https://www.youtube.com/watch?v=3WsEDZRif6U"
		}

		_, err := c.GetVideoInfo(url)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "%v", err)
		}
		// c.JSON(http.StatusOK, *videoInfo.Formats)
		Formats := map[string][]youtube.Format{
			"Formats": c.Video.Formats,
		}
		ctx.HTML(http.StatusOK, "index.html", Formats)
	}
}
