# Basic Pointers in Go

A pointer is a variable that stores the memory address of another value.

## Syntax

### 1. The Address-of Operator (`&`)
Use `&` to get the memory address of a variable.

```go
x := 10
address := &x // address is a pointer to x
```

### 2. The Dereference Operator (`*`)
Use `*` to "dereference" a pointer, which means accessing the value stored at that address.

```go
value := *address // value is 10
```

### 3. Pointer Types
The type of a pointer is `*T`, where `T` is the type of the value it points to.
- `*int` is a pointer to an integer.
- `*string` is a pointer to a string.

## The `nil` Pointer
If a pointer is declared but not initialized, its value is `nil`. Attempting to dereference a `nil` pointer will cause a **runtime panic**.

```go
var p *int
fmt.Println(p) // <nil>
// fmt.Println(*p) // This would CRASH the program!
```

## The `new()` Function
The `new(T)` built-in function allocates memory for a variable of type `T` and returns a pointer to it (`*T`). The memory is initialized to the type's zero value.

```go
p := new(int) // p is of type *int
fmt.Println(*p) // 0
```

## Why use Pointers?
- **Avoid Copying:** When passing large data to functions.
- **Shared State:** When multiple parts of a program need to access or modify the same data.
