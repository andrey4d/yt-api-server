/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandlerIndex(c *fiber.Ctx) error {
	return c.Render("index", nil)
}
