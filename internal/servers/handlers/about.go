/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type About struct {
	msg string
}

func GetAbout() fiber.Handler {
	about := About{msg: "About"}
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"about": about.msg})
	}

}
