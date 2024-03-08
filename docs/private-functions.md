# Private Functions

In [Go](go-language.md), you can limit access to function by
changing the first character of a functions name to lower case.

This will limit access to the function to within the current [module](modules.md).


## Example(s)

```go
// this will be accessible to all who import this module
func PublicFunction() { ... }

// this is only accessible within this module
func PrivateFunction() { ... }
```

## Reference(s)

- [Learn Go with tests - Hello, World](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#one...last...refactor)
