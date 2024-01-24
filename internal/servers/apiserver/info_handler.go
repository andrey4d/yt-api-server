/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/andrey4d/ytapiserver/internal/youtuber"
	"github.com/gin-gonic/gin"
	"github.com/kkdai/youtube/v2"
)

type VideoFormat struct {
	Itag          int    `json:"itag,omitempty"`
	FPS           int    `json:"fps,omitempty"`
	VideoQuality  string `json:"video_quality,omitempty"`
	AudioQuality  string `json:"audio_quality,omitempty"`
	AudioChannels int    `json:"audio_channels,omitempty"`
	Language      string `json:"language,omitempty"`
	Size          int64  `json:"size,omitempty"`
	Bitrate       int    `json:"bitrate,omitempty"`
	MimeType      string `json:"mime_type,omitempty"`
}

type VideoInfo struct {
	ID          string         `json:"id,omitempty"`
	Title       string         `json:"title,omitempty"`
	Author      string         `json:"author,omitempty"`
	Duration    string         `json:"duration,omitempty"`
	Description string         `json:"description,omitempty"`
	Formats     *[]VideoFormat `json:"formats,omitempty"`
}

func (s ApiServer) handlerInfo(c *gin.Context) {
	url := "https://www.youtube.com/watch?v=3WsEDZRif6U"

	videoInfo, err := getVideoInfo(url)
	if err != nil {
		s.logger.Warn(err)
		c.String(http.StatusInternalServerError, "%v", err)
	}

	s.logger.Infof("++ handlerInfo %s\n", url)
	c.JSON(http.StatusOK, *videoInfo.Formats)
}

func getVideoInfo(url string) (*VideoInfo, error) {
	video, err := youtuber.GetDownloader().GetVideo(url)
	if err != nil {
		return nil, err
	}

	return &VideoInfo{
		Title:       video.Title,
		Author:      video.Author,
		Duration:    video.Duration.String(),
		Description: video.Description,
		Formats:     getFormats(video),
	}, nil
}

func getFormats(video *youtube.Video) *[]VideoFormat {

	formats := []VideoFormat{}

	for _, format := range video.Formats {
		bitrate := format.AverageBitrate
		if bitrate == 0 {
			// Some formats don't have the average bitrate
			bitrate = format.Bitrate
		}

		size := format.ContentLength
		if size == 0 {
			// Some formats don't have this information
			size = int64(float64(bitrate) * video.Duration.Seconds() / 8)
		}

		formats = append(formats, VideoFormat{
			Itag:          format.ItagNo,
			FPS:           format.FPS,
			VideoQuality:  format.QualityLabel,
			AudioQuality:  strings.ToLower(strings.TrimPrefix(format.AudioQuality, "AUDIO_QUALITY_")),
			AudioChannels: format.AudioChannels,
			Size:          size,
			Bitrate:       bitrate,
			MimeType:      format.MimeType,
			Language:      format.LanguageDisplayName(),
		})
	}

	return &formats
}

func formatsToStrings(formats *[]VideoFormat) []string {
	out := []string{}
	for _, format := range *formats {
		sf := fmt.Sprintf("%s %s %s %s %s %s %s %s %s\n",
			strconv.Itoa(format.Itag),
			strconv.Itoa(format.FPS),
			format.VideoQuality,
			format.AudioQuality,
			strconv.Itoa(format.AudioChannels),
			fmt.Sprintf("%0.1f", float64(format.Size)/1024/1024),
			strconv.Itoa(format.Bitrate),
			format.MimeType,
			format.Language)
		out = append(out, sf)
	}
	return out
}
