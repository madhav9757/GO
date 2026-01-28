package main

import "fmt"

// Go uses Composition via Struct Embedding, not Inheritance

type Author struct {
	Name string
	Bio  string
}

func (a Author) WriteBio() {
	fmt.Printf("Author: %s - Bio: %s\n", a.Name, a.Bio)
}

// Book embeds Author
type Book struct {
	Author // Anonymous field (Embedding)
	Title  string
	Year   int
}

// Method name collision (Shadowing)
func (b Book) WriteBio() {
	fmt.Printf("Book Bio: '%s' by %s\n", b.Title, b.Name)
}

func main() {
	fmt.Println("--- Struct Embedding ---")

	auth := Author{
		Name: "J.K. Rowling",
		Bio:  "Famous fantasy author",
	}

	b := Book{
		Author: auth,
		Title:  "Harry Potter",
		Year:   1997,
	}

	// 1. Promoted Fields
	// We can access b.Author.Name or just b.Name
	fmt.Println("Book Title:", b.Title)
	fmt.Println("Author Name (via promotion):", b.Name)
	fmt.Println("Author Name (explicitly):", b.Author.Name)

	// 2. Promoted Methods
	// Since Book has its own WriteBio, it shadows Author.WriteBio
	fmt.Println("\nCalling WriteBio:")
	b.WriteBio()        // Calls Book's version
	b.Author.WriteBio() // Calls embedded Author's version

	// 3. Embedding Interface-like behavior
	// This is the core of Go's way of sharing logic between types.
}
