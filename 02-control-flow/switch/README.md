# Switch Statements in Go

This example demonstrates the `switch` statement in Go, which provides a cleaner alternative to long if-else chains.

## Concepts Covered

### Basic Switch
Go's switch statements are more powerful than in many other languages. Cases automatically break (no fall-through by default):

```go
day := 3

switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
default:
    fmt.Println("Invalid day")
}
```

### Switch Without Expression
You can use switch without an expression, making it work like an if-else chain with cleaner syntax:

```go
num := 10
switch {
case num < 0:
    fmt.Println("Negative number")
case num == 0:
    fmt.Println("Zero")
case num > 0:
    fmt.Println("Positive number")
}
```

This is particularly useful when you have complex conditions or different variables in each case.

## Running the Code

```bash
go run switch.go
```

## Expected Output

```
Wednesday
Positive number
```

## Key Points

- No automatic fall-through (unlike C/Java) - each case breaks automatically
- Use `fallthrough` keyword if you need fall-through behavior
- Can switch on any comparable type (not just integers)
- Switch without an expression is equivalent to `switch true`
- Multiple values per case: `case 1, 2, 3:`
- No need for `break` statements

## Additional Features

Go's switch is more flexible than many languages:
- Can have initialization statement: `switch x := getValue(); x {}`
- Type switches for interface types
- Cases are evaluated top to bottom, and only the first match executes