# Basic Functions

Functions are central in Go. Here's how to declare and use them.

## Syntax

```go
func name(parameter1 type, parameter2 type) returnType {
    // body
    return value
}
```

## Features Demonstrated

1.  **Standard Function**: `add(a int, b int) int`
    -   Takes two integers and returns one.
2.  **Type Omission**: `subtract(a, b int) int`
    -   When parameters share a type, you can omit the type for all but the last.
3.  **Named Return Values**: `calculate(a, b int) (sum int, diff int)`
    -   Return values can be named. They are initialized to zero values.
    -   A `return` without arguments returns the current values of the named return parameters (naked return).
