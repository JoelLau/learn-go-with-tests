# Structs

Is a named collection of fields we can store data;
A way of creating custom data types.

## Example(s)

```go
// Instead of this...
rectangeWidth  := 5
rectangeHeight := 5

// We can define a struct
type Rectangle struct {
    Width  int
    Height in,
}

// Initialising a struct
rectangle := Rectangle {
    Width: 5,
    Height: 5,
}

// Accessing struct data
fmt.Println(rectangle.Width)  // Output -> 5
fmt.Println(rectangle.Height) // Output -> 5
```

## References

- [Learn Go with tests: Structs, methods & interfaces](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#refactor)
