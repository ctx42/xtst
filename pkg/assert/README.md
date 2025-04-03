<!-- TOC -->
* [The `assert` package](#the-assert-package)
  * [Structure](#structure)
<!-- TOC -->

# The `assert` package

The `assert` package is a toolkit for Go testing that offers common assertions,
integrating well with the standard library. When writing tests, developers often
face a choice between using Go's standard `testing` package or packages like 
`assert`. The standard library requires verbose `if` statements for assertions, 
which can make tests harder to read. This package, on the other hand, provides 
one-line asserts, such as `assert.NoError`, which are more concise and clear. 
This simplicity helps quickly grasp the intent of each test, enhancing 
readability.

By making tests easier to write and read, this package hopes to encourage 
developers to invest more time in testing. Features like immediate feedback 
with easily readable output and a wide range of assertion functions lower the 
barrier to writing comprehensive tests. This can lead to better code coverage, 
as developers are more likely to write and maintain tests when the process is
straightforward and rewarding.

## Assertions

Most of the assertions are self-explanatory and I encourage you to see your
online [documentation](https://pkg.go.dev/github.com/ctx42/testing). Here we 
will highlight only the ones that we feel are interesting. 

## Asserting Structures

```go
type T struct {
    Int int
    Str string
}

have := T{Int: 1, Str: "abc"}
want := T{Int: 2, Str: "xyz"}

assert.Equal(want, have)
// Test Log:
// expected values to be equal:
//   trail: T.Int
//    want: 2
//    have: 1
//  ---
//   trail: T.Str
//    want: "xyz"
//    have: "abc"
```
