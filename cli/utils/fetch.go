package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchAndUnmarshal(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}
