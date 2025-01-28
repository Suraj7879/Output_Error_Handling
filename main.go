package main

import (
	"fmt"

	"github.com/CodeYourFuture/immersive-go-course/projects/output-and-error-handling/fetcher"
)

func main() {
	f := fetcher.WeatherFetcher{}
	for {
		weather, err := f.Fetch("http://localhost:8080")

		if err != nil {
			continue
		}
		fmt.Println("Weather data:", weather)
		// return
	}
	// response, err := http.Get("http://localhost:8080")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Errorf("error trying to read response: %w", err)
	// }
	// fmt.Println(string(body))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/CodeYourFuture/immersive-go-course/projects/output-and-error-handling/fetcher"
// )

// func main() {
// 	f := fetcher.WeatherFetcher{}
// 	const maxRetries = 3

// 	for attempt := 1; attempt <= maxRetries; attempt++ {
// 		weather, err := f.Fetch("http://localhost:8080")
// 		if err != nil {
// 			log.Printf("Attempt %d failed: %v", attempt, err)
// 			time.Sleep(1 * time.Second) // Wait before retrying
// 			continue
// 		}

// 		fmt.Println("Weather data:", weather)
// 		return // Exit on success
// 	}

// 	log.Fatal("Failed to fetch weather after multiple attempts")
// }
