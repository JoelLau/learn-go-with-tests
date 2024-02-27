# Arrays

Homogenous collection of data of fixed length.
For dynamic length, see [Slices](slices.md)


**NOTE**: it is recommended to use slices wherever possible.


## Examples

Fixed length array

```go
sample := int[5]{1, 2, 3, 4, 5}

# will not be accepted by the below
#   func Sum(num int[3]) {...}
#   -> i.e. Sum(sample) will not compile
```

## References

- [Learn Go with Tests: Arrays and Slices](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices)
