# The `check` Package

The `check` package is designed for performing assertions in Go tests,
particularly as a foundational layer for the `assert` package. It provides
functions that return errors instead of boolean values, allowing callers to
adjust error messages to particular context, add more contextual information
about the check, improving assertion message comprehension.

## Example Usage

You use checks like any other function returning error.

```go
	have := errors.New("test error")

	err := NoError(have, WithPath("path"))

	fmt.Println(err)
	// Output:
	// expected error to be nil:
	//	path: path
	//	want: <nil>
	//	have: "test error"
```

The main purpose of returning an error from a check, instead of true false like 
it is in case of `assert` package is to give user ability to customize the 
message and/or add context. 

```go
have := errors.New("test error")

err := NoError(have, WithPath("path"))

err = notice.From(err, "prefix").Append("context", "wow")

fmt.Println(err)
// Output:
// [prefix] expected error to be nil:
//	   path: path
//	   want: <nil>
//	   have: "test error"
//	context: wow
```
