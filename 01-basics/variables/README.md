# Go Variables Program

This is a simple Go program demonstrating **variables in Go**. It covers different ways to declare and use variables, multiple variable declarations, short-hand declarations, and zero values.

---

## 📂 Files

- `variables.go` – Main program file demonstrating variable usage in Go.

---

## 📝 Program Features

1. **Variable Declaration with Type**
   ```go
   var a int = 10
   var name string = "Madhav"
   ```
2. **Variable Declaration without Type (Type Inference)**

   ```bash
   var b = 20
   var city = "Delhi"
   ```

3. **Short-hand Declaration (Inside Functions Only)**

   ```bash
   c := 30
   isGo := true
   ```

4. **Multiple Variable Declaration**

   ```bash
   var x, y, z int = 1, 2, 3
   p, q := 5, "Go"
   ```

5. **Multiple Variable Declaration**
   - Demonstrates default values for uninitialized variables:
   ```bash
    var n int // 0
    var s string // ""
    var f float64 // 0
    var flag bool // false
   ```

---

# ⚡ How to Run

1. Make sure Go is installed.
2. Open terminal and navigate to the folder containing variables.go.
3. Run the program:
   ```bash
   go run variables.go
   ```
4. go run variables.go
   ```bash
   a = 10
   name = Madhav
   b = 20
   city = Delhi
   c = 30
   isGo = true
   x, y, z = 1 2 3
   p, q = 5 Go

    Zero values:
    int: 0
    string:
    float: 0
    bool: false
    ```

--- 

## 🔑 Notes
- Go is strongly typed.
- Short-hand := cannot be used outside functions.
- Unused variables will throw a compile-time error.
- Uninitialized variables automatically get zero values.

