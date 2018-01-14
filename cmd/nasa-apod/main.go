package main

import (
	"flag"
	"fmt"
	"github.com/ndoering/nasa/apod"
	"image/jpeg"
	"os"
)

func main() {
	var apiKey string
	var hd bool
	flag.StringVar(&apiKey, "key", "DEMO_KEY", "The API_KEY to use with NASA API")
	flag.BoolVar(&hd, "hd", false, "Download the HD image")

	flag.Parse()

	c := apod.NewClient(apiKey, hd)

	im, _ := c.GetImage()

	file, err := os.Create("test.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	m := jpeg.Options{100}
	jpeg.Encode(file, im, &m)
}
