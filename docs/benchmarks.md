# Benchmarks

[Go](go-language.md) provides in-built features to test the performance of your code.
To write a benchmark test, we write a [`for` loop](for-loop.md) that loops `b.N` (from `b *testing.B`) times.
The contents of the loop are executed and measured before the results are reported.

To run the benchmark, use the `-bench` flag on the [test command](test-cmd.md): `go test -bench=.`


## Example

We run the contents of the for loop `b.N` number of times
```go
// repeat_test.go

import "testing"

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```

Benchmark results: The code was executed _10,000,000_ times and took _136_ nanoseconds to run on average.
```plaintext
goos: darwin
goarch: amd64
pkg: github.com/quii/learn-go-with-tests/for/v4
10000000           136 ns/op
PASS
```

## Reference

- [Learn Go with tests: Iteration#Benchmarking](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration#benchmarking)
