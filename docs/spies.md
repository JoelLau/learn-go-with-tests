# Spies

[Spying](#) is a testing technique that makes use of [mocking](mocking.md).
By implementing an [interface](interfaces.md), we replace parts of an
application and track metadata associated with it. One common thing to
track is the number of times a function has been called.


## Example(s)

```go
// --- main ---

type Performer() interface {
    Perform()
}

type Clown struct {}

func (_ Clown) Perform {}

type Application struct {
    P Perform
}

func main() {
    app := Application{ P: Clown{} }
    app.Perform()
}

// --- test ---

type PerformerSpy struct {
    Calls int
}

type (p PerformerSpy) Perform {
    p.Calls++
}


func TestApplication(t *testing.T) {
    spy := PerformerSpy{}
    app := Application { P: spy }
    app.Perform()

    if p.Calls < 1 {
        t.Errorf("expected `Perform` to be called at least once")
    }
}
```

## Reference(s)

- [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking#write-the-test-first-2)
