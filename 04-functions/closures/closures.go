package main

import "fmt"

// intSeq returns a function that yields the next integer in a sequence.
// The returned function closes over the variable i.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// nextInt captures its own i value
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// newInts has its own separate i value
	newInts := intSeq()
	fmt.Println(newInts())
}
