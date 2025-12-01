package main

import "fmt"

func runMapsDemo() {
	fmt.Println("\n--- MAP DEMO ---")

	// MAP
	scores := map[string]int{
		"Madhav": 90,
		"Rahul":  85,
		"Anita":  95,
	}
	fmt.Println("Scores map:", scores)

	// Access
	fmt.Println("Madhav's score:", scores["Madhav"])

	// Add
	scores["Riya"] = 88
	fmt.Println("After adding Riya:", scores)

	// Delete
	delete(scores, "Rahul")
	fmt.Println("After deleting Rahul:", scores)

	// Loop
	fmt.Println("Looping over map:")
	for name, score := range scores {
		fmt.Println(name, ":", score)
	}
}
