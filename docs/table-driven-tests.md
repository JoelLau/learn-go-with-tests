# Table Driven Tests

[Iterate](for-loops.md) through an [array](arrays.md) of test data,
running the same test on each element.

Useful when we want to re-use the same test with different parameters.

## Example(s)

```go
//
// In this example, shape is an interface that is
// implemented by Rectangle and Circle
//
func TestArea(t *testing.T) {

    // define a list of test cases
    testCases := []struct {
        shape Shape
        want  float64
    } {
        { Rectangle{12, 6}, 72.0 },
        { Circle{10}, 314.1592653589793 },
    }

    // iterate through each test cases
    for index, tt := range testCases {

        // run through each test case here
        have := tt.shape.Area()
        want := tt.want

        if have != want {
            t.Errorf("have %g want %g", have, want)
        }
    }
}
```

## Reference(s)

- [Learn Go with tests: Structs, methods & interfaces](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces#further-refactoring)
- [Go Wiki: TableDrivenTests](https://go.dev/wiki/TableDrivenTests)
