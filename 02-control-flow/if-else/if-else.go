package main

import "fmt"

func main() {
    age := 20

    if age < 18 {
        fmt.Println("You are a minor.")
    } else if age >= 18 && age < 60 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a senior citizen.")
    }

    // Using a short statement in if
    if greeting := "Hello Go!"; len(greeting) > 5 {
        fmt.Println(greeting)
    }
}
