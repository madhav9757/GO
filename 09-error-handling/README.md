# 09 - Error Handling in Go

In Go, error handling is a core part of the language design. Unlike languages that use exceptions (`try-catch`), Go uses explicit return values to handle errors.

## Key Concepts

### 1. The `error` Interface

In Go, an error is any value that implements the built-in `error` interface:

```go
type error interface {
    Error() string
}
```

### 2. Explicit Error Checking

Functions that can fail should return an `error` as their last result.

```go
val, err := someFunction()
if err != nil {
    // Handle error
    return err
}
// Use val
```

### 3. Creating Errors

- `errors.New("message")`: For simple text-based errors.
- `fmt.Errorf("error: %v", details)`: For formatted errors with context.

### 4. Custom Errors

You can define your own error types by implementing the `Error()` method on a struct. This is useful for providing extra data about the error.

### 5. Panic and Recover

- `panic`: Stops the normal execution of the goroutine and starts panicking.
- `recover`: A built-in function that regains control of a panicking goroutine. Usually used within `defer`.

## Files in this module

- `basics/errors.go`: Basic error handling patterns.
- `custom/custom-errors.go`: Defining and using custom error types.
- `panic-recover/panic-recover.go`: Understanding when and how to use panic/recover.
