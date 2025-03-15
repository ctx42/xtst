# Notice Package

The `notice` package provides a set of utilities for building structured 
assertion messages. It's designed to create easy to read and understand error 
messages with a header and contextual rows. The package supports fluent 
interfaces for building messages and includes helper functions for formatting 
and unwrapping errors.

## Usage

Creating a basic message

```go
msg := New("expected %s to be equal", "values").
    Want("%s", "abc").
    Have("%s", "xyz")

fmt.Println(msg.Error())
```

Output:

```text
expected values to be equal:
  want: abc
  have: xyz
```

For more examples see [examples_test.go](examples_test.go) file.

## Formatting Lines

```go
lines := notice.Lines(1, "line1\nline2\nline3")

fmt.Println(lines)
```

Output:

```text
  >| line1
  >| line2
  >| line3
```

