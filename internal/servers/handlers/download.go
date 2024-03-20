/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"fmt"
	"strconv"

	"github.com/andrey4d/ytapiserver/internal/ytclient"
	"github.com/gofiber/fiber/v2"
)

type DownloadHandler struct {
}

func (h *DownloadHandler) GetDownload(client *ytclient.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		VideoTag := ctx.FormValue("Video")
		AudioTag := ctx.FormValue("Audio")

		vURL, err := h.GetFormatURL(client, VideoTag)
		if err != nil {
			return err
		}

		aURL, err := h.GetFormatURL(client, AudioTag)
		if err != nil {
			return err
		}

		// return Render(ctx, layout.Index(client, *pageAttributes))
		out := fmt.Sprintf("Url=%s,<br> Url=%s", vURL, aURL)

		return ctx.SendString(out)
	}
}

func (h *DownloadHandler) GetFormatURL(client *ytclient.Client, iTag string) (string, error) {
	tag, err := strconv.Atoi(iTag)
	if err != nil {
		return "", err
	}
	f := client.Video.Formats.Itag(tag)
	if len(f) != 1 {
		return "", fmt.Errorf("youtube FormatList len not 1: %d", len(f))
	}
	return f[0].URL, nil
}

// ffmpegVersionCmd := exec.Command("ffmpeg", "-y",
// "-i", videoFile.Name(),
// "-i", audioFile.Name(),
// "-c", "copy", // Just copy without re-encoding
// "-shortest", // Finish encoding when the shortest input stream ends
// destFile,
// "-loglevel", "warning",
