/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/gofiber/fiber/v2"
)

type DownloadHandler struct {
}

func (h *DownloadHandler) GetDownload(client *ytclient.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// return Render(ctx, layout.Index(client, *pageAttributes))
		return ctx.SendString("To do")
	}
}
