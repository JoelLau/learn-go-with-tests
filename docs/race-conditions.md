# Race Conditions

A race condition is a scenario that occurs when multiple processes attempt to
access the same resource, resulting in an error. In [go](go-language.md), this
usually happens when using [goroutines](goroutines.md).

To detect these, use [race detector]() to analyse the code.

Code that is free from race conditions are said to be "[thread safe](thread-safe.md)"


## Example(s)

```go
results := make(map[string]int)

for i := 0; i < 50; i++ {
    go func(num int) {
        // WARN: mutating a map (data type) is not thread safe!
        results["asdf"] = num
    }(i)
}
```


## Reference(s)

- [Learn Go with tests: Concurency]()
