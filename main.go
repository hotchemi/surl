package main

import (
	"os"

	"fmt"
	"github.com/atotto/clipboard"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

func main() {
	app := cli.NewApp()
	app.Name = "surl"
	app.Author = "Shintaro Katafuchi"
	app.Email = "hot.chemistry21@gmail.com"
	app.Version = "0.1.0"
	app.Usage = "Shorten an url in your clipboard."
	app.Action = action
	app.Run(os.Args)
}

func action(context *cli.Context) {
	text := readClipboard()
	res := httpGet(text)
	json := unmarshal(res)
	if json.Errormessage != "" {
		fmt.Fprintln(os.Stderr, json.Errormessage)
		os.Exit(1)
	}
	clipboard.WriteAll(json.Shorturl)
	fmt.Println("Write the value to clipboard: " + json.Shorturl)
}

func readClipboard() string {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return text
}

func httpGet(url string) []byte {
	response, _ := http.Get(buildUrl(url))
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return body
}

func buildUrl(url string) string {
	baseUrl := "http://is.gd/create.php?format=simple&format=json&url="
	return baseUrl + url
}

func unmarshal(b []byte) data {
	var d data
	json.Unmarshal(b, &d)
	return d
}

type data struct {
	Shorturl string
	Errormessage string
}