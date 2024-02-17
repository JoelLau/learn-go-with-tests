# `for` loops

Like other c-based languages, [`for` loops](#) in [go](go-language.md)
start with the `for` keyword and 3 parts separated by a semi-colon (`;`).

In order, they are each for:
1. initialising variables scoped to the for loop
1. the condition at which to keep looping for if evaluated to `True`
1. modifications to be applied after the contents of the loop


## Example

```go
// prints 0 - 4, each on a newline 
for i := 0 ; i < 5 ; i++ {
    fmt.Println(i)
}
```
