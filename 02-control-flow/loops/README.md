# Loops in Go

This example demonstrates the various ways to use loops in Go. Unlike many languages, Go has only one looping construct: the `for` loop.

## Concepts Covered

### Traditional For Loop
The classic three-component loop with initialization, condition, and post statement:

```go
for i := 1; i <= 5; i++ {
    fmt.Println("For loop iteration:", i)
}
```

### While-Style Loop
Go doesn't have a `while` keyword, but you can create while-style loops using `for` with only a condition:

```go
n := 1
for n <= 5 {
    fmt.Println("While-like loop:", n)
    n++
}
```

### Infinite Loop
Omit all three components to create an infinite loop. Use `break` to exit:

```go
count := 1
for {
    if count > 3 {
        break
    }
    fmt.Println("Infinite loop with break:", count)
    count++
}
```

## Running the Code

```bash
go run loops.go
```

## Expected Output

```
For loop iteration: 1
For loop iteration: 2
For loop iteration: 3
For loop iteration: 4
For loop iteration: 5
While-like loop: 1
While-like loop: 2
While-like loop: 3
While-like loop: 4
While-like loop: 5
Infinite loop with break: 1
Infinite loop with break: 2
Infinite loop with break: 3
```

## Key Points

- Go has only `for` loops (no `while` or `do-while`)
- `break` exits the loop immediately
- `continue` skips to the next iteration
- Variables declared in the initialization are scoped to the loop