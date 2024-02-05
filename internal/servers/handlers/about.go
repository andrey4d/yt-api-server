/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"github.com/andrey4d/ytapiserver/templ/view"
	"github.com/gofiber/fiber/v2"
)

func GetAbout() fiber.Handler {
	about := []string{"Project for", "Download video", "From Youtube"}

	return func(c *fiber.Ctx) error {
		pageAttributes := GetPageAttributes()
		pageAttributes.PageTitle = "About this project."
		return Render(c, view.About(about, *pageAttributes))
	}

}
