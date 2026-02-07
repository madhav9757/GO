package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// User struct with JSON tags
type User struct {
	ID        int      `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	IsActive  bool     `json:"is_active"`
	Languages []string `json:"languages,omitempty"`
}

func main() {
	// --- 1. Marshaling (Go Struct -> JSON) ---
	user1 := User{
		ID:        1,
		Username:  "jdoe",
		Email:     "john@example.com",
		IsActive:  true,
		Languages: []string{"Go", "Python", "JavaScript"},
	}

	fmt.Println("--- Marshaling (Struct to JSON) ---")
	// json.MarshalIndent makes the output human-readable
	jsonData, err := json.MarshalIndent(user1, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
	} else {
		fmt.Println(string(jsonData))

		// Saving JSON to a file
		os.WriteFile("user.json", jsonData, 0644)
	}

	// --- 2. Unmarshaling (JSON -> Go Struct) ---
	fmt.Println("\n--- Unmarshaling (JSON to Struct) ---")
	rawJSON := `{"id": 2, "username": "asmith", "email": "alice@example.com", "is_active": false}`
	var user2 User

	if err := json.Unmarshal([]byte(rawJSON), &user2); err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
	} else {
		fmt.Printf("Unmarshaled User: %+v\n", user2)
	}

	// --- 3. Working with Generic Maps (Unknown JSON structure) ---
	fmt.Println("\n--- Unmarshaling to Map (Dynamic JSON) ---")
	dynamicJSON := `{"name": "Project X", "version": 1.2, "tags": ["beta", "confidential"]}`
	var result map[string]interface{}

	if err := json.Unmarshal([]byte(dynamicJSON), &result); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Name: %v\n", result["name"])
		fmt.Printf("Version: %v\n", result["version"])
	}
}
