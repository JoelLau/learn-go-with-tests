# Learn Go with tests

Code written while following along [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests)

## Usage

Run all tests in module

```shell
go test -v ./...
```

Run all tests in workspace ([source](https://github.com/golang/go/issues/50745#issuecomment-1402920852))

```shell
go list -f '{{.Dir}}' -m | xargs go test -v -cover
```


Run godoc server

```shell
# visit localhost:8080 on any browser
godoc -http :8080
```
