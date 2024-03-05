# Iterfaces

Allow us to define functions that can be used with different types.
Decouples code while maintaining type-safety.


## Example(s)

```go
// define `Shape` interface
type Shape interface {

    // define a method that MUST be implemented
    Area() float64
}


// defines a `Triangle` struct
type Circle struct {
    Radius float64
}

// `Circle` counts as a `Shape` if it meets definition
func (c Circle) Area() float64 {
    return math.pi * c.Radius * c.Radius
}
```


## Reference(s)

- [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#refactor-1)
- [Interface Types](https://go.dev/ref/spec#Interface_types)
