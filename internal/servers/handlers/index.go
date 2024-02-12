/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/andrey4d/ytapiserver/templ/layout"
	"github.com/gofiber/fiber/v2"
)

type IndexHandler struct {
}

func (h *IndexHandler) GetIndex(client *ytclient.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		pageAttributes := GetPageAttributes()
		pageAttributes.PageTitle = "Youtube get video URLs."

		return Render(ctx, layout.Index(client, *pageAttributes))
	}
}
