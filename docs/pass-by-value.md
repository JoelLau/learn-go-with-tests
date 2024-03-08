# Pass By Reference

When calling functions and methods, **ALL** arguments are copied by default,
including the base class. It is important that we use [pointers](pointers.md)
when writing changes so that we modify the correct instance of a variable.

## Examples

Updates will fail here because `Deposit` and `Balance` are
referencing different copies of `Bank`.

```go
type Bank struct{
    balance int
}

func (b Bank) Deposit(amt int) {
    b.balance += amt
    return
}

func (b Bank) Balance() int {
    return b.balance
}

bank := Bank{}
b.Deposit(5)
b.Balance() // Output -> 0

```

Use [pointers](pointers.md) to reference the same copy.

```go
type Bank struct{
    balance int
}

func (b *Bank) Deposit(amt int) {
    b.balance += amt
    return
}

func (b *Bank) Balance() int {
    return b.balance
}

bank := Bank{}
b.Deposit(5)
b.Balance() // Output -> 5

```
## References

-[Learn Go with tests - Pointers & errors](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#thats-not-quite-right)

