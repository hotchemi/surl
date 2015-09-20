package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "surl"
	app.Author = "Shintaro Katafuchi"
	app.Email = "hot.chemistry21@gmail.com"
	app.Version = "0.1.0"
	app.Usage = "Shorten an url in your clipboard."
	app.Action = shorten
	app.Run(os.Args)
}

func shorten(context *cli.Context) {





}