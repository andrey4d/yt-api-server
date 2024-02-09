/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/gofiber/fiber/v2"
)

func GetDescription(client *ytclient.Client) fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.SendString(client.Video.Description)
	}

}
