/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"fmt"

	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/gofiber/fiber/v2"
	"github.com/kkdai/youtube/v2"
)

func GetInfo(c *ytclient.Client) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		url := ctx.FormValue("video_tag")
		if url == "" {
			url = "https://www.youtube.com/watch?v=3WsEDZRif6U"
		}

		_, err := c.GetVideoInfo(url)
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("%v", err))
		}
		// c.JSON(http.StatusOK, *videoInfo.Formats)
		Formats := map[string][]youtube.Format{
			"Formats": c.Video.Formats,
		}
		return ctx.Render("index", Formats)
	}
}
