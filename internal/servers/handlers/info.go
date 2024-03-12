/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"fmt"

	"github.com/andrey4d/ytapiserver/internal/ytclient"
	// "github.com/andrey4d/ytapiserver/templ/forms"
	"github.com/andrey4d/ytapiserver/internal/log"
	"github.com/andrey4d/ytapiserver/templ/layout"
	"github.com/gofiber/fiber/v2"
)

type InfoHandler struct {
	LogHandler *log.LogHandl
}

func (h *InfoHandler) GetInfo(client *ytclient.Client) fiber.Handler {

	pageAttributes := GetPageAttributes()
	pageAttributes.PageTitle = "Main page"

	return func(ctx *fiber.Ctx) error {
		url := ctx.FormValue("video_tag")
		h.LogHandler.LogModuleInfo("GetInfo()").Infoln(url)
		if url == "" {
			url = "3WsEDZRif6U"
		}

		_, err := client.GetVideoInfo(url)
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("%v", err))
		}
		pageAttributes.PageTitle = client.Video.Title

		return Render(ctx, layout.UrlLayout(client))
	}
}
