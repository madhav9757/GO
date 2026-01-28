# Variadic Functions

Variadic functions can be called with any number of trailing arguments. For example, `fmt.Println` is a common variadic function.

## Syntax

```go
func functionName(args ...Type) {
    // args is a slice of Type inside the function
}
```

## Features Demonstrated

1.  **Defining**: `func sum(nums ...int)` defines a function that takes zero or more `int` arguments.
2.  **Using**: Inside the function, the param `nums` is equivalent to `[]int`.
3.  **Slice Expansion**: If you already have a slice, apply it to a variadic function using `func(slice...)`.
