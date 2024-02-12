/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"fmt"

	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/gofiber/fiber/v2"
)

type UrlsHandler struct{}

func (h *UrlsHandler) GetUrls(c *ytclient.Client) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		tag := ctx.FormValue("films")

		format, err := c.GetFormat(tag)
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("%v", err))
		}

		link_text := fmt.Sprintf("Quality=%s, audio channels=%d, video=%s", format.QualityLabel, format.AudioChannels, format.MimeType)

		return ctx.SendString(fmt.Sprintf("<a href='%s'>%s</a>", format.URL, link_text))
	}
}
