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

type DescriptionHandler struct{}

func (h *DescriptionHandler) GetDescription(client *ytclient.Client) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		attributes := layout.DescriptionAttributes{
			Author:      client.Video.Author,
			Description: client.Video.Description,
		}
		return Render(ctx, layout.DescriptionContent(attributes))
	}

}
