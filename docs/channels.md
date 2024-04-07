# Channels

[Channels](#) are thread safe streams that allow us to send and receive messages.
These messages can be of any type and must be declared. Use the `<-` operator
to send and receive messages.

Channels are especially useful during concurrent operations via [goroutines]()
that often result in race conditions. By using a channel, we can receive and process
messages in a predictable manner.


## Example(s)

```go
// custom type
type result struct {
    string
    bool
}

// creating a channel
c := make(chan result)

// send a new result object to channel
c <- result {"asdf", true}

// receive sent result, store in `r`
r := <- c
```

## Reference(s)

- [Learn Go with tests: Concurrency]()
