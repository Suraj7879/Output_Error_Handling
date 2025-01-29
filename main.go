package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/CodeYourFuture/immersive-go-course/projects/output-and-error-handling/fetcher"
)

func main() {
	f := fetcher.WeatherFetcher{}
	for {
		weather, err := f.Fetch("http://localhost:8080")

		if err != nil {
			if errors.Is(err, fetcher.ErrRetry) {
				continue
			} else {
				fmt.Fprintf(os.Stderr, "Error getting weather %v\n", err)
			}
		} else {
			fmt.Println("Weather data:", weather)
		}
	}
}
