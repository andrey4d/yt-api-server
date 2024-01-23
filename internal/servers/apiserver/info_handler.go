/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package apiserver

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/andrey4d/ytapiserver/internal/youtuber"
	"github.com/kkdai/youtube/v2"
)

type VideoFormat struct {
	Itag          int
	FPS           int
	VideoQuality  string
	AudioQuality  string
	AudioChannels int
	Language      string
	Size          int64
	Bitrate       int
	MimeType      string
}

type VideoInfo struct {
	ID          string
	Title       string
	Author      string
	Duration    string
	Description string
	Formats     *[]VideoFormat
}

func (s ApiServer) handlerInfo() http.HandlerFunc {

	url := "https://www.youtube.com/watch?v=3WsEDZRif6U"
	video, err := youtuber.GetDownloader().GetVideo(url)
	if err != nil {
		s.logger.Fatal(err)
	}
	videoInfo := VideoInfo{
		Title:       video.Title,
		Author:      video.Author,
		Duration:    video.Duration.String(),
		Description: video.Description,
		Formats:     getFormats(video),
	}

	s.logger.Infof("%+v\n", videoInfo.Formats)

	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Infof("Info %s %s\n", r.Method, r.Host)

		for _, format := range *videoInfo.Formats {
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
			io.WriteString(w, sf)
		}

	}
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
