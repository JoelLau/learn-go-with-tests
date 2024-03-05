# Methods

A method is a function with a receiver.
Defining one associates a function with the receiver's base type.

## Example

```go
// base type(?)
type Car struct {}

// method definition
func (c *Car) Honk() {
    fmt.Println("Honk Honk!")
}
```

## References

- [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#what-are-methods)
- [go.dev spec](https://go.dev/ref/spec#Method_declarations)
