package apod

import (
	"encoding/json"
	"image"
	"image/jpeg"
	"net/http"
)

type APODAnswer struct {
	Date, Explanation, HDUrl, MediaType, ServiceVersion, Title, Url string
}

type Client struct {
	key string
}

func NewClient(key string) Client {
	return Client{key}
}

func (c *Client) GetImage() (im image.Image, err error) {
	response, err := http.Get("https://api.nasa.gov/planetary/nasa?hd&api_key=" + c.key)
	if err != nil {
		return
	}

	var m APODAnswer

	jsonDecoder := json.NewDecoder(response.Body)
	jsonDecoder.Decode(&m)

	response.Body.Close()

	response, err = http.Get(m.HDUrl)
	if err != nil {
		return
	}

	defer response.Body.Close()

	im, err = jpeg.Decode(response.Body)

	return
}
