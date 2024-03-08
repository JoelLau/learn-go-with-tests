# `&` Operator

Used to show the memory address of a variable. Useful for checking
if multiple variables are referencing the same piece of data. (See
[pointers](pointers.md))


## Examples

```go
x := "hello"
fmt.Printf(" x = '%s'\n", x)  // Output -> " x = 'hello'"
fmt.Printf("&x = '%p'\n", &x) // Output -> "&x = '0x14000102020'"
```


## References

- [Learn Go with tests - Pointers & errors](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#thats-not-quite-right)
