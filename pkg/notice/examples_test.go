// Copyright (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package notice

import (
	"fmt"
)

func ExampleNew() {
	msg := New("expected values to be equal").
		Want("%s", "abc").
		Have("%s", "xyz")

	fmt.Println(msg)
	// Output:
	// expected values to be equal:
	//	want: abc
	//	have: xyz
}

func ExampleNew_formated() {
	msg := New("expected %s to be equal", "values").
		Want("%s", "abc").
		Have("%s", "xyz")

	fmt.Println(msg)
	// Output:
	// expected values to be equal:
	//	want: abc
	//	have: xyz
}

func ExampleFrom() {
	var err error
	err = New("expected values to be equal").
		Want("%s", "abc").
		Have("%s", "xyz")

	msg := From(err, "optional prefix").
		Append("my", "%s", "value")

	fmt.Println(msg)
	// Output:
	// [optional prefix] expected values to be equal:
	//	want: abc
	//	have: xyz
	//	  my: value
}

func ExampleNotice_SetHeader() {
	msg := New("expected %s to be equal", "values").
		Want("%s", "abc").
		Have("%s", "xyz")

	_ = msg.SetHeader("some other %s", "header")

	fmt.Println(msg)
	// Output:
	// some other header:
	//	want: abc
	//	have: xyz
}

func ExampleNotice_Append() {
	msg := New("expected %s to be equal", "values").
		Want("%s", "abc").
		Have("%s", "xyz").
		Append("name", "%d", 5)

	fmt.Println(msg)
	// Output:
	// expected values to be equal:
	//	want: abc
	//	have: xyz
	//	name: 5
}

func ExampleNotice_AppendRow() {
	row0 := NewRow("number", "%d", 5)
	row1 := NewRow("string", "%s", "abc")

	msg := New("expected %s to be equal", "values").
		Want("%s", "abc").
		Have("%s", "xyz").
		AppendRow(row0, row1)

	fmt.Println(msg)
	// Output:
	// expected values to be equal:
	//	  want: abc
	//	  have: xyz
	//	number: 5
	//	string: abc
}

func ExampleNotice_Prepend() {
	msg := New("expected %s to be equal", "values").
		Path("Struct.Path").
		Want("%s", "abc").
		Have("%s", "xyz").
		Prepend("name", "%d", 5)

	fmt.Println(msg)
	// Output:
	// expected values to be equal:
	//	path: Struct.Path
	//	name: 5
	//	want: abc
	//	have: xyz
}

func ExampleNotice_Path() {
	msg := New("expected %s to be equal", "values").
		Path("Struct.Path").
		Want("%s", "abc").
		Have("%s", "xyz")

	fmt.Println(msg)
	// Output:
	// expected values to be equal:
	//	path: Struct.Path
	//	want: abc
	//	have: xyz
}

func ExampleLines() {
	lines := Lines(2, "line1\nline2\nline3")

	fmt.Println(lines)
	// Output:
	//		>| line1
	//		>| line2
	//		>| line3
}
