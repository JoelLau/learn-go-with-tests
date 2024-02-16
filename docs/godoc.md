# godoc

Command line utility for viewing [go](go-language.md) documentation.
Used to be included in [go cli](go-cli.md) before it was removed in version 1.14.


## Installation

NOTE: ensure that go binaries are in PATH (see [golang-setup](golang-setup.md))

```bash
go install golang.org/x/tools/cmd/godoc@latest
```

## Usage

start http server

```bash
godoc -http :8000
```

visit `localhost:8000` on any browser


## Reference

- [Learn Go with tests: Hello, World #Go doc](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#go-doc)
