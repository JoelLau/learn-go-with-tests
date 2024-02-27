# Slices

Like [arrays](arrays.md), but with dynamic length.
They can be "sliced" by using the following syntax `slice[low:high]`.

**NOTE**: "sliced" slices are still linked to the original slice.
It is recommended to use [copy](copy.md) to separate them to prevent
accidental modifications.

## Example(s)

```go
// initialisation
example = int[]{1, 5, 10}

// slice examples
fmt.Println(example[1:]) // Output -> 5, 10
fmt.Println(example[:1]) // Output -> 1, 5
fmt.Println(example[:])  // Output -> 1, 5, 10
```

## References

- [Learn Go with Tests: Arrays and Slices](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices)
