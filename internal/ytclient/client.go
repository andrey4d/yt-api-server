/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package ytclient

import (
	"errors"
	"strconv"

	"github.com/kkdai/youtube/v2"
)

type Client struct {
	YtClient youtube.Client
	Video    *youtube.Video
}

func New() *Client {
	return &Client{
		YtClient: youtube.Client{},
		Video:    &youtube.Video{},
	}
}

func (c *Client) GetVideoInfo(url string) (*youtube.Video, error) {
	video, err := c.YtClient.GetVideo(url)
	if err != nil {
		return nil, err
	}
	c.Video = video
	return video, nil
}

func (c *Client) GetFormat(idTag string) (*youtube.Format, error) {
	tag, err := strconv.Atoi(idTag)
	if err != nil {
		return nil, err
	}

	for _, format := range c.Video.Formats {
		if format.ItagNo == tag {
			return &format, nil
		}
	}

	return nil, errors.New("video format not found")
}

func (c *Client) GetUrl(tag string) (string, error) {
	format, err := c.GetFormat(tag)
	if err != nil {
		return "", err
	}

	url, err := c.YtClient.GetStreamURL(c.Video, format)
	if err != nil {
		return "", err
	}
	return url, nil
}
