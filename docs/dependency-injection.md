# Dependency Injection (DI)

To pass a dependency as a function argument.


## Example(s)

This function has a dependency on `fmt.Printf`
```go
func Greet(name string) {
    fmt.Printf("Hello, %s", name)
}

func main() {
    Greet("World")
}
```

We can give pass `io.writer` as a dependency instead
```go
func Greet(writer io.writer, name string) {
    fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
    Greet(os.Stdout, "World")
}
```


## Reference(s)

- [Learn Go with tests - Dependency Injection](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection)
