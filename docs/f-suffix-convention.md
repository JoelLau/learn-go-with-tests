# `f` Suffix Convention

The `f` stands for *format*.
Often used when printing or creating strings and expecting placeholders.

## Example(s)

print

```go
package main

import "fmt"


func main() {
    // print with newline
    fmt.Println("this prints a single line")

    // print without newline, and with formatting
    word := "foo"
    fmt.Printf("this prints the word: %s", word)
}
```


forming strings with sprint

```go
name := "slim shady"
sentence := fmt.Sprintf("hi, my name is %s", name) // Output -> "hi, my name is slim shady"
```


## Reference(s)

- Joel's observations 😅
- [t-errorf](t-errorf.md)
- [t-fatalf](t-fatalf.md)
