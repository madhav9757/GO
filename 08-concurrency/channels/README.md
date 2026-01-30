# Channels

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

> "Do not communicate by sharing memory; instead, share memory by communicating."

## 📌 Usage

```go
ch := make(chan int) // Unbuffered
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```

## 🔹 Unbuffered Channels vs Buffered Channels

- **Unbuffered (`make(chan int)`)**: Synchronous. Sending blocks until receiving happens, and vice-versa. Great for synchronization.
- **Buffered (`make(chan int, 5)`)**: Asynchronous (up to a limit). Sending only blocks when the buffer is full. Receiving only blocks when the buffer is empty.

## 🔸 Closing Channels

A sender can `close(ch)` to indicate that no more values will be sent. Receivers can test whether a channel has been closed:

```go
v, ok := <-ch
if !ok {
    // Channel closed
}
```

## 🔀 Select Statement

The `select` statement lets a goroutine wait on multiple communication operations. It blocks until one of its cases can run, then it executes that case.
