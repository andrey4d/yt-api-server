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

type FormatsByType map[string][]YtFormat
type YtVideo *youtube.Video
type YtFormat *youtube.Format

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

func (c *Client) GetAudioFormatsOnly() *[]youtube.Format {
	outFormats := []youtube.Format{}
	for _, format := range c.Video.Formats {
		if format.AudioChannels == 0 {
			outFormats = append(outFormats, format)
		}
	}
	return &outFormats
}

func (c *Client) GetVideoFormatsOnly() *[]youtube.Format {
	outFormats := []youtube.Format{}
	for _, format := range c.Video.Formats {
		if format.QualityLabel == "" {
			outFormats = append(outFormats, format)
		}
	}
	return &outFormats
}

func (c *Client) GetVideoAudioFormatsOnly() *[]youtube.Format {
	outFormats := []youtube.Format{}
	for _, format := range c.Video.Formats {
		if format.QualityLabel != "" && format.AudioChannels != 0 {

			outFormats = append(outFormats, format)
		}
	}

	return &outFormats
}

// func GetFormatsType(client *ytclient.Client) *ytclient.FormatsByType{

//   formatsByType := ytclient.FormatsByType{
// 	"video": []ytclient.YtFormat{},
// 	"audio": []ytclient.YtFormat{},
// 	"video-audio": []ytclient.YtFormat{},
// 	}

//   for _, format  := range client.Video.Formats {
// 	if format.AudioChannels == 0 {
// 	  formatsByType["audio"] = append(formatsByType["audio"], &format)
// 	}
//   }
// ffmpeg -y -i https://rr2---sn-oxu8pnpvo-nq8e.googlevideo.com/videoplayback?expire=1709927450&ei=uhfrZb6IDbuFhcIP_ca-6Ak&ip=129.159.138.23&id=o-ANZOAGIDaVO-KQpV0P7St2OcQ8pels7wfv6-EqBFuxrC&itag=139&source=youtube&requiressl=yes&xpc=EgVo2aDSNQ%3D%3D&mh=dG&mm=31%2C29&mn=sn-oxu8pnpvo-nq8e%2Csn-ua87zn7l&ms=au%2Crdu&mv=u&mvi=2&pl=23&vprv=1&svpuc=1&mime=audio%2Fmp4&gir=yes&clen=5176732&dur=848.828&lmt=1686670171495702&mt=1709905572&fvip=3&keepalive=yes&fexp=24007246&c=ANDROID&txp=5432434&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cxpc%2Cvprv%2Csvpuc%2Cmime%2Cgir%2Cclen%2Cdur%2Clmt&sig=AJfQdSswRQIgewcOWPE6wrcd0EWTq-FXFONdG4uUttuF4tIDRjCAj9wCIQCsKYSNJecaB0f6UViy-68sznWmsaOHBp-fFee3RJq2aQ%3D%3D&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl&lsig=APTiJQcwRgIhAKqYDQ_yM9rG3DjeoEGZ8Gnx2w8v1x-0CTpBxMWaAv_CAiEApAI5XZBRfb7qTJvfnrDTevtbH0VnFJWPIGNjSyejCAk%3D \
// -i https://rr2---sn-oxu8pnpvo-nq8e.googlevideo.com/videoplayback?expire=1709927450&ei=uhfrZb6IDbuFhcIP_ca-6Ak&ip=129.159.138.23&id=o-ANZOAGIDaVO-KQpV0P7St2OcQ8pels7wfv6-EqBFuxrC&itag=160&source=youtube&requiressl=yes&xpc=EgVo2aDSNQ%3D%3D&mh=dG&mm=31%2C29&mn=sn-oxu8pnpvo-nq8e%2Csn-ua87zn7l&ms=au%2Crdu&mv=u&mvi=2&pl=23&vprv=1&svpuc=1&mime=video%2Fmp4&gir=yes&clen=1361709&dur=848.699&lmt=1686670413323662&mt=1709905572&fvip=3&keepalive=yes&fexp=24007246&c=ANDROID&txp=5432434&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cxpc%2Cvprv%2Csvpuc%2Cmime%2Cgir%2Cclen%2Cdur%2Clmt&sig=AJfQdSswRAIgRGXKhxlEaYiwjB4m9Fee8_1P2Y3qYV4lFTNj_YXCQ_kCIGvP1LGpEcjSNQgqIEtqznb5E7REL2QMt1T21J2sHjSr&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl&lsig=APTiJQcwRgIhAKqYDQ_yM9rG3DjeoEGZ8Gnx2w8v1x-0CTpBxMWaAv_CAiEApAI5XZBRfb7qTJvfnrDTevtbH0VnFJWPIGNjSyejCAk%3D \
// -c copy -shortest out.mp4
// return &formatsByType
// }
