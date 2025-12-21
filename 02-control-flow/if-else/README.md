# If-Else Statements in Go

This example demonstrates conditional statements in Go using `if`, `else if`, and `else`.

## Concepts Covered

### Basic If-Else
Go's conditional statements work similarly to other languages, but without parentheses around the condition:

```go
if age < 18 {
    fmt.Println("You are a minor.")
} else if age >= 18 && age < 60 {
    fmt.Println("You are an adult.")
} else {
    fmt.Println("You are a senior citizen.")
}
```

### If with Short Statement
Go allows you to execute a short statement before the condition. The variable declared is only available within the if-else block:

```go
if greeting := "Hello Go!"; len(greeting) > 5 {
    fmt.Println(greeting)
}
```

This is particularly useful for error handling and checking values without polluting the outer scope.

## Running the Code

```bash
go run if-else.go
```

## Expected Output

```
You are an adult.
Hello Go!
```

## Key Points

- No parentheses required around conditions
- Braces `{}` are mandatory, even for single statements
- Short statement syntax keeps variables scoped to the if block
- Logical operators: `&&` (AND), `||` (OR), `!` (NOT)