# Multiple Return Values

Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

## Syntax

```go
func functionName() (type1, type2) {
    return value1, value2
}
```

## Features Demonstrated

1.  **Multiple Returns**: `vals()` returns two `int` values.
2.  **Blank Identifier**: `_` is used to ignore return values you don't need.
3.  **Swapping**: `swap()` demonstrates clearly how multiple returns can simplify logic like swapping variables.
