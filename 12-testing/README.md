# 🧪 Module 12: Testing and Benchmarking in Go

Testing is a first-class citizen in Go. The standard library provides the `testing` package, and the `go test` command is all you need to run tests, benchmarks, and examples.

## 1. Unit Testing

Located in `./unit-tests`.

Unit tests in Go are functions starting with `Test` in files ending with `_test.go`.

### Key Concepts:

- **Table-Driven Tests**: The idiomatic way to test multiple scenarios by looping over a slice of structs.
- **Subtests**: Using `t.Run()` to group related test cases.

### How to Run:

```bash
go test -v ./unit-tests/...
```

## 2. Benchmarking

Located in `./benchmark-tests`.

Benchmarks are functions starting with `Benchmark` and use the `*testing.B` type. They run the code `b.N` times to measure performance.

### How to Run:

```bash
go test -v -bench . ./benchmark-tests/...
```

In this example, we compare the performance of recursive Fibonacci vs. iterative Fibonacci. You'll notice the iterative version is significantly faster!

## 3. Mocking with Interfaces

Located in `./mock-tests`.

Go doesn't have a built-in mocking framework because it's rarely needed if you use **Interfaces**. You can simply implement the interface in your test file with dummy data or custom logic.

### How to Run:

```bash
go test -v ./mock-tests/...
```

## Summary of Commands

- `go test ./...` - Run all tests in the module.
- `go test -v ./...` - Run all tests with verbose output.
- `go test -bench . ./...` - Run all benchmarks.
- `go test -cover ./...` - Run tests and show code coverage.
