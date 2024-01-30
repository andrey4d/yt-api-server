/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetHello() fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello world!",
		})
	}
}
