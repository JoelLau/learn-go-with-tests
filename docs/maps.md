# Map (Data Type)

A [map](#) is a collection of key-value pairs.


## Example(s)

```go
// creates a map where keys are *string* and values are *int*
exampleMap := map[string]int{}

// stores *32* (int) under the key "*this is a key*"
exampleMap["this is a key"] := 32

// retrieve value
fmt.Println(exampleMap["this is a key"]) // Output -> 32
```
```go
// FOOTGUN ALERT: this creates a `nil` pointer!
var badMap = map[string]int

// this is ok (**with** curly braces)
ok := map[string]int{}

// this is also ok
alsoOk := make(map[string]int)

```

## See Also

- [what types of values can be keys?](acceptable-key-types.md).

## Reference(s)

- [Learn Go with tests: Maps](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps#write-the-test-first)
