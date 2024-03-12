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

func (c *Client) GetVideoFormatsOnly() *[]youtube.Format {
	outFormats := []youtube.Format{}
	for _, format := range c.Video.Formats {
		if format.AudioChannels == 0 {
			outFormats = append(outFormats, format)
		}
	}
	return &outFormats
}

func (c *Client) GetAudioFormatsOnly() *[]youtube.Format {
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

// VIDEO="https://rr3---sn-hxb54vo-v8cs.googlevideo.com/videoplayback?expire=1710236502&ei=9s7vZcKjEsrHv_IP9ciqsAs&ip=188.162.12.26&id=o-AFuZrwckept0_9n1Cxvg8dTA5m1Ms55ndiTohS1gDjNP&itag=160&source=youtube&requiressl=yes&xpc=EgVo2aDSNQ%3D%3D&mh=dG&mm=31%2C29&mn=sn-hxb54vo-v8cs%2Csn-n8v7znz7&ms=au%2Crdu&mv=m&mvi=3&pl=23&initcwndbps=656250&vprv=1&svpuc=1&mime=video%2Fmp4&gir=yes&clen=1361709&dur=848.699&lmt=1686670413323662&mt=1710214628&fvip=11&keepalive=yes&fexp=24007246&beids=24350321&c=ANDROID&txp=5432434&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cxpc%2Cvprv%2Csvpuc%2Cmime%2Cgir%2Cclen%2Cdur%2Clmt&sig=AJfQdSswRAIgGnhGIyMd96BjRDn7pOMQVfodVpwSGgytXBCQZWDqMhMCIEHv69OuBEb_tjoqtAu6bAMxtqmcc3X4NMvsEKZl3KOu&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=APTiJQcwRAIgNrl9pEf-52Cfj1cYYoUILDTNIfimzI1vvOywwyBqjG0CICf_IGlEf9mbxRPIjGTB94DXmZ2D0fJ5psJ2w91LkGM0"
// AUDIO="https://rr3---sn-hxb54vo-v8cs.googlevideo.com/videoplayback?expire=1710236502&ei=9s7vZcKjEsrHv_IP9ciqsAs&ip=188.162.12.26&id=o-AFuZrwckept0_9n1Cxvg8dTA5m1Ms55ndiTohS1gDjNP&itag=139&source=youtube&requiressl=yes&xpc=EgVo2aDSNQ%3D%3D&mh=dG&mm=31%2C29&mn=sn-hxb54vo-v8cs%2Csn-n8v7znz7&ms=au%2Crdu&mv=m&mvi=3&pl=23&initcwndbps=656250&vprv=1&svpuc=1&mime=audio%2Fmp4&gir=yes&clen=5176732&dur=848.828&lmt=1686670171495702&mt=1710214628&fvip=11&keepalive=yes&fexp=24007246&beids=24350321&c=ANDROID&txp=5432434&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cxpc%2Cvprv%2Csvpuc%2Cmime%2Cgir%2Cclen%2Cdur%2Clmt&sig=AJfQdSswRAIgaIWgafGhd6asU7OGCDAkwoEVC79MYtrp9dwdLiDV0dECIGj3PhSbll0MYbfpUycnT8_UmffJNk4ub9e1-QhHpvgQ&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=APTiJQcwRAIgNrl9pEf-52Cfj1cYYoUILDTNIfimzI1vvOywwyBqjG0CICf_IGlEf9mbxRPIjGTB94DXmZ2D0fJ5psJ2w91LkGM0"
// ffmpeg -y -i "${VIDEO}" -i "${AUDIO}" -c copy -loglevel warning -shortest out.mp4
// ffmpeg -y -i "${VIDEO}" -i "${AUDIO}" -f matroska -loglevel warning -shortest pipe:1
// return &formatsByType
// }
