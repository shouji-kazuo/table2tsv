package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	tApp := cli.App{
		Name:    "table2tsv",
		Usage:   "htmlのテーブルをtsvで吐く",
		Version: "1.0.0",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Kazuo SHOUJI",
				Email: "shouji_kazuo@kurusugawa.jp",
			},
		},
		Copyright: "Kurusugawa Computer Inc.",
		Action: func(aContext *cli.Context) error {
			return nil
		},
	}

	if tError := tApp.Run(os.Args); tError != nil {
		fmt.Fprintln(os.Stderr, tError.Error())
	}
}
