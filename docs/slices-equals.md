# `slices.Equal()`

[Slice](slices.md)-specific alternative to [reflect.DeepEquals](reflect-deep-equals.md)
for shallow equality checks that are type safe.

**NOTE**: will only work with 1D arrays with elements that are comparable using 
the equality (`==`) operator.


## Example(s)

```go
a := int[5]{1, 2, 3, 4, 5}
b := "lorem ipsum"

slices.Equal(a, b) // will not compile!
```

```go
a := int[]{1, 2, 3, 4, 5}
b := int[]{1, 2, 3}

slices.Equal(a, b) // Output -> False
```

```go
a := int[]{1, 2, 3, 4, 5}
b := int[]{1, 2, 3, 4, 5}

slices.Equal(a, b) // Output -> True
```

## References

- [Learn Go with Tests: Arrays and Slices](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#write-the-minimal-amount-of-code-for-the-test-to-run-and-check-the-failing-test-output-2)
- [Official Docs: slices#Equal](https://pkg.go.dev/slices#Equal)
