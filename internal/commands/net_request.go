package commands

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func GetResource[T any](url string) (T, error) {
	var zero T
	responseBody := []byte{}
	cachedVal, exists := globalConfig.Cache.Get(url)
	if exists {
		responseBody = cachedVal
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return zero, err
		}
		defer resp.Body.Close()
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return zero, err
		}
		responseBody = bytes
	}
	file, _ := os.Create("json.json")
	if file != nil {
		file.Write(responseBody)
	}
	decoded := new(T)
	err := json.Unmarshal(responseBody, decoded)
	if err != nil {
		return zero, err
	}
	return *decoded, nil
}
