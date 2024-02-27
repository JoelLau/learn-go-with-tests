# `reflect.DeepEqual()`

Checks if _any_ 2 variables are the same.
Useful for comparing pointers, arrays and slices but is **not type safe!**
_(will check at runtime, not compile time)_


## Examples

Comparing slices
```go
a := []int{1, 2, 3}
b := []int{1, 2, 3, 4, 5}

reflect.DeepEqual(a, b) // Output -> false
```

Comparing slice to string will still compile
```go
a := []int{1, 2, 3}
b := "lorem ipsum"

reflect.DeepEqual(a, b) // Output -> false
```

## References

- [Learn Go with Tests: Arrays and Slices](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#write-the-minimal-amount-of-code-for-the-test-to-run-and-check-the-failing-test-output-2)
- [reflect#DeepEqual](https://pkg.go.dev/reflect#DeepEqual)
