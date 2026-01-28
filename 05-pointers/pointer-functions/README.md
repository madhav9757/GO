# Pointers and Functions

In Go, every argument is **passed by value**. This means Go makes a copy of the argument and passes it to the function.

## Pass by Value vs. Pass by Pointer

### 1. Pass by Value (Default)
When you pass a standard variable, the function receives a copy. Any changes made inside the function do not affect the original variable.

```go
func increment(val int) {
    val++
}
```

### 2. Pass by Pointer
If you want a function to modify the original value, you must pass its **memory address** (a pointer).

```go
func increment(ptr *int) {
    *ptr++
}
```

## When to use Pointers in Functions?

1.  **Modification:** When the function needs to update the state of the passed argument.
2.  **Performance:** When passing very large structs or arrays. Passing a pointer (8 bytes on 64-bit systems) is much faster than copying a massive object.
3.  **Consistency:** Sometimes used to indicate that a value can be `nil` (meaning "no value" or "not provided").

## Safety Note
Always check for `nil` if you expect a pointer might be empty:

```go
func doSomething(ptr *int) {
    if ptr == nil {
        return
    }
    // safely use *ptr
}
```
