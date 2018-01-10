package main

import (
	"fmt"
	"github.com/ndoering/nasa/apod"
	"image/jpeg"
	"os"
)

func main() {
	c := apod.NewClient("DEMO_KEY")

	im, _ := c.GetImage()

	file, err := os.Create("/home/shogun/Downloads/test.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	m := jpeg.Options{100}
	jpeg.Encode(file, im, &m)
}
