package main

import (
	"fmt"
	"time"
)

// RequestError is a custom error type that carries extra information
type RequestError struct {
	StatusCode int
	Message    string
	Timestamp  time.Time
}

// Implement the error interface by adding the Error() method
func (e *RequestError) Error() string {
	return fmt.Sprintf("[%v] HTTP %d: %s", e.Timestamp.Format("15:04:05"), e.StatusCode, e.Message)
}

func doRequest(url string) error {
	if url == "" {
		return &RequestError{
			StatusCode: 400,
			Message:    "Empty URL",
			Timestamp:  time.Now(),
		}
	}
	// Simulate server error
	return &RequestError{
		StatusCode: 500,
		Message:    "Internal Server Error",
		Timestamp:  time.Now(),
	}
}

func main() {
	fmt.Println("--- Custom Error Types ---")

	err := doRequest("")
	if err != nil {
		fmt.Println("Error:", err)

		// We can use type assertion to access specific fields if it's our custom type
		if reqErr, ok := err.(*RequestError); ok {
			fmt.Printf("Status Code: %d, Message: %s\n", reqErr.StatusCode, reqErr.Message)
		}
	}

	err = doRequest("https://example.com")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
