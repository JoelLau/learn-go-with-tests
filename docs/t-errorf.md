# `t.Errorf`

Reports error messages when running [go tests commands](go-test-cmd.md)
Continues running execution even after being called.
To stop execution when a test fails, see [t.Fatalf](t-fatalf.md)


## Example(s)

```go
package main

import "testing"

func TestHello(t *testing.T) {
    have := Hello()
    want := "Hello, World!"
    
    if have != want {
        t.Errorf("have '%s' want '%s'", have, want)
    }
}
```

## Reference(s)

- [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#how-to-test)
