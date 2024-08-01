package pkg

import (
	"io"
	"net/http"
)

func CallURLGet(url string) (result string, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	result = string(body)
	return
}
