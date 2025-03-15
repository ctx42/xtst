package must

import (
	"fmt"
)

func ExampleFirst() {
	type Row struct{ Name string }

	// Query to database returning rows.
	query := func() ([]Row, error) { return []Row{{"a"}, {"b"}}, nil }

	have := First(query())

	fmt.Println(have)
	// Output:
	// {a}
}

func ExampleSingle() {
	type Row struct{ Name string }

	// Query to database returning rows.
	query := func() ([]Row, error) { return []Row{{"a"}}, nil }

	// Will panic if database returned more than one error.
	have := Single(query())

	fmt.Println(have)
	// Output:
	// {a}
}
