# goroutines

A child processes that can perform concurrent and/or asynchronous operations.
Created by using the `go` keyword, followed by a function call.

## Example(s)

```go
func main() {
    for i := 0; i < 20; i++ {
        // creates a child process that runs an anonymous function
        go func(num int) {
            fmt.Println(num)
        }(i) // pass i to the function to "fix" its value within lexical scope
    }
}
```

## Reference(s)

- [Learn Go with tests: Concurrency]()
