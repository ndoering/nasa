package apod

import (
	"encoding/json"
	"image"
	"image/jpeg"
	"net/http"
)

type ConnectionError struct {
	msg  string
	code int
}

func (c *ConnectionError) Error() string {
	return c.msg
}

func newConnectionError(message string, code int) error {
	return &ConnectionError{message, code}
}

type APODAnswer struct {
	Date, Explanation, HDUrl, MediaType, ServiceVersion, Title, Url string
}

type Client struct {
	key string
	hd  bool
}

func NewClient(key string, hd bool) Client {
	return Client{key, hd}
}

func (c *Client) GetImage() (im image.Image, err error) {

	url, err := getImageUrl(c.key, c.hd)
	if err != nil {
		return
	}

	response, err := http.Get(url)
	if err != nil {
		return
	}

	defer response.Body.Close()

	im, err = jpeg.Decode(response.Body)

	return
}

func getImageUrl(key string, hd bool) (url string, err error) {
	response, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=" + key)
	if err != nil {
		response.Body.Close()
		return
	}

	if response.StatusCode != 200 {
		err = newConnectionError("Access not allowed. Please check your API key.", response.StatusCode)
		response.Body.Close()
		return
	}

	var m APODAnswer

	jsonDecoder := json.NewDecoder(response.Body)
	jsonDecoder.Decode(&m)

	response.Body.Close()

	if hd == true {
		url = m.HDUrl
	} else {
		url = m.Url
	}

	return
}
