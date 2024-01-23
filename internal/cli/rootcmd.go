/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package cli

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Run: run,
}

func Execute() {
	rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		logrus.Fatal("url not define")
	}
	video, format, err := getVideoWithFormat(args[0])
	if err != nil {
		logrus.Fatal(err)
	}

	url, err := downloader.GetStreamURL(video, format)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(url)
}
