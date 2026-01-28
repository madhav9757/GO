# Closures

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

## Features Demonstrated

1.  **State Capture**: The `intSeq` function returns another function. This returned function "closes over" the variable `i` defined in `intSeq`.
2.  **Isolation**: calling `intSeq` creates a new instance of `i`, so different function calls have isolated state.
