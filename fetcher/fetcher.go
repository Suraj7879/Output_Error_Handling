package fetcher

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

var ErrRetry = errors.New("should retry")

type WeatherFetcher struct {
	client http.Client
}

func (w *WeatherFetcher) Fetch(url string) (string, error) {
	response, err := w.client.Get(url)

	if err != nil {
		return "", fmt.Errorf("couldn't make http request: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return "", fmt.Errorf("failed to read the body response: %w", err)
	}
	return string(body), nil
}
