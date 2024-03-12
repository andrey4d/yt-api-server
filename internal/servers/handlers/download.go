/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"fmt"

	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/gofiber/fiber/v2"
)

type DownloadHandler struct {
}

func (h *DownloadHandler) GetDownload(client *ytclient.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		VideoTag := ctx.FormValue("Video")
		AudioTagg := ctx.FormValue("Audio")
		// return Render(ctx, layout.Index(client, *pageAttributes))
		out := fmt.Sprintf("Download Video TAG=%s Audio TAG-%s", VideoTag, AudioTagg)

		return ctx.SendString(out)
	}
}

// ffmpegVersionCmd := exec.Command("ffmpeg", "-y",
// "-i", videoFile.Name(),
// "-i", audioFile.Name(),
// "-c", "copy", // Just copy without re-encoding
// "-shortest", // Finish encoding when the shortest input stream ends
// destFile,
// "-loglevel", "warning",
