# Struct Basics in Go

Structs are the primary way to define custom data types in Go. They allow you to group related data fields together.

## 📝 Key Concepts

### 1. Defining a Struct
Use the `type` and `struct` keywords:
```go
type Person struct {
    Name string
    Age  int
}
```

### 2. Initialization
- **Field Names:** `p := Person{Name: "Alice", Age: 30}` (Recommended)
- **Positional:** `p := Person{"Alice", 30}`
- **Zero Value:** `var p Person` (Fields initialized to their zero values)

### 3. Pointers to Structs
Go provides a convenient way to work with pointers to structs. You can access fields directly without explicit dereferencing.
```go
ptr := &Person{Name: "Bob"}
fmt.Println(ptr.Name) // Same as (*ptr).Name
```

### 4. Anonymous Structs
Useful for one-off data structures, common in configuration or testing.
```go
config := struct {
    Port int
}{
    Port: 8080,
}
```

### 5. Nested Structs
You can use one struct as a field in another.
```go
type Employee struct {
    Info    Person
    Address string
}
```

## 🚀 How to Run
```bash
go run structs.go
```
