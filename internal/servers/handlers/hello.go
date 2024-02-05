/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"github.com/andrey4d/ytapiserver/internal/log"
	"github.com/andrey4d/ytapiserver/templ/view"
	"github.com/gofiber/fiber/v2"
)

func GetHello(logger *log.LogHandl) fiber.Handler {
	entry := logger.GetLogger().WithField("module", "GetHello()")
	message := "Hello World!"

	return func(c *fiber.Ctx) error {
		userAgent := c.GetReqHeaders()["User-Agent"]
		entry.Info(c.OriginalURL(), " | ", userAgent)
		pageAttributes := *GetPageAttributes()
		pageAttributes.PageTitle = "Hello page"
		return Render(c, view.Hello(message, pageAttributes))
	}
}
