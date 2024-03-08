# `t.Fatalf`

Reports error messages when running [go tests commands](go-test-cmd.md)
Stops current test execution if test fails.
To continue execution when a test fails, see [t.Fatalf](t-fatalf.md)


## Example(s)

```go
package somepackage

import "testing"


func TestSomePackage(t *testing) {
    value, err := SomeFunction()

    if err != nil {
        // exiting early here because value will be wrong if there's an error
        t.Fatalf("received an expected error %s", err.Error())
    }

    if value != "expected value" {
        // continue testing
        t.Errorf("did not get expected value. wanted %s, got %s instead", value, "expected value")
    }
}
```

## Reference(s)

- [Learn Go with tests: Pointers & errors](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#write-the-test-first-3)
