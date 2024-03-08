# Receivers

Receiver functions are [Go](golang.md)'s equivalent of class methods
in Object-Oriented Languages. The **receiver** is the part between
the `func` keyword and the function name.

Receivers are function arguments that reference the type it is attached to.


## Example(s)

In this example, `b *Bitcoin` defines the reciever and its type.

```go
type Bitcoin int

func (b *Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}
```

## Reference(s)

**Context**: `receiver` has come up a in a few places in documentation.
(_e.g. https://go.dev/ref/spec#Method_values_)

- [Medium: What the heck are Receiver functions in Golang](https://medium.com/@adityaa803/wth-is-a-go-receiver-function-84da20653ca2)
