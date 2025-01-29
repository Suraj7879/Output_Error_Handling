package fetcher

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
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

	switch response.StatusCode {
	case http.StatusOK:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return "", fmt.Errorf("error trying to read the response %w", err)
		}

		return string(body), nil

	case http.StatusTooManyRequests:
		err := handle429(response.Header.Get("retry-after"))
		if err != nil {
			return "", fmt.Errorf("error Handling 'too many requests' response: %w", err)
		}

		return "", ErrRetry

	default:
		errorDescription := convertHTTPResponseToDescription(response)
		return "", fmt.Errorf("unexpected response from server: %s", errorDescription)
	}
}

func handle429(retryAfterHeader string) error {
	delay, err := parseDelay(retryAfterHeader)
	if err != nil {
		return err
	}

	if delay > 1*time.Second {
		fmt.Fprintf(os.Stderr, "Server reported its receiving too many requests - waiting %s seconds before retrying", delay)
	}

	if delay > 5*time.Second {
		return fmt.Errorf("giving up request: server told us it's going to be too busy for requests for more than the next 5 seconds")
	}

	time.Sleep(delay)

	return nil
}

func parseDelay(retryAfterHeader string) (time.Duration, error) {
	waitFor, err := strconv.Atoi(retryAfterHeader)
	if err == nil {
		return time.Duration(waitFor) * time.Second, nil
	}

	waitUntil, err := http.ParseTime(retryAfterHeader)
	if err == nil {
		return time.Until(waitUntil), nil
	}

	return -1, fmt.Errorf("couldn't parse retry-after header as an integer number of seconds or a date. Value was: %q", retryAfterHeader)

}

func convertHTTPResponseToDescription(response *http.Response) string {
	var bodyString string
	body, err := io.ReadAll(response.Body)
	if err == nil {
		bodyString = string(body)
	} else {
		bodyString = "<error reading body>"
	}

	return fmt.Sprintf("Status Code %s, Body: %s", response.Status, bodyString)
}
