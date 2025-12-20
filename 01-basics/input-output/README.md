# Input and Output Program in Go

This Go program demonstrates **basic input and output operations** using the `fmt` package.
It takes user input from the keyboard and displays it on the console.

---

## 📂 File

- `input_output.go` – Contains a function that accepts user input and prints it

---

## 🧾 Program Code

```go
package main

import "fmt"

func input_output() {
	var name string
	var age int

	// Input
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	// Output
	fmt.Println("\n--- User Details ---")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
```

---

## ⚙️ Explanation

package main
Required for executable Go programs.

import "fmt"
Used for input (Scan) and output (Print, Println).

fmt.Scan(&variable)
Takes input from the user and stores it in the given variable.

& (address operator)
Needed because Scan modifies the variable’s value.

---

## ⚠️ Important Note

This file does NOT contain func main(), so it will not run by itself.

To execute this function, you must call it from a main() function.

## ✅ Correct Way to Run

Create or modify a file like this:

```
package main

func main() {
	input_output()
}
```

Then run:

```
go run input_output.go main.go
```

---

## ▶️ Sample Output

```
Enter your name: Madhav
Enter your age: 22

--- User Details ---
Name: Madhav
Age: 22
```
----

## 🔑 Key Points

fmt.Scan() reads input from standard input

Go does not allow unused functions

Every executable program needs exactly one main() function