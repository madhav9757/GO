# Concurrency Patterns

Go's concurrency primitives allow for powerful design patterns.

## 🏭 Worker Pool

A worker pool limits the number of goroutines running simultaneously. This is useful when you have a large number of tasks but limited resources (CPU, Memory, Network connections).

1.  Create a pool of worker goroutines.
2.  Send tasks to a shared `jobs` channel.
3.  Workers pull from `jobs`, process, and send to `results`.

## ⏱️ Context

The `context` package is standard for managing request lifecycles, cancellations, and timeouts across API boundaries and processes.

- **Cancellation**: If a user cancels a request, all goroutines working on it should stop to save resources.
- **Timeouts**: Stop an operation if it takes too long.
- **Values**: Pass request-scoped values (Auth tokens, Trace IDs).

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()

select {
case <-time.After(2 * time.Second):
    fmt.Println("Overslept")
case <-ctx.Done():
    fmt.Println(ctx.Err()) // "context deadline exceeded"
}
```
