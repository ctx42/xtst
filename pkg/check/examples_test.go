// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"errors"
	"fmt"

	"github.com/ctx42/xtst/pkg/notice"
)

func ExampleError() {
	err := Error(nil)

	fmt.Println(err)
	// Output:
	// expected non-nil error
}

func ExampleNoError() {
	have := errors.New("test error")

	err := NoError(have)

	fmt.Println(err)
	// Output:
	// expected error to be nil:
	//	want: <nil>
	//	have: "test error"
}

func ExampleNoError_withPath() {
	have := errors.New("test error")

	err := NoError(have, WithPath("path"))

	fmt.Println(err)
	// Output:
	// expected error to be nil:
	//	path: path
	//	want: <nil>
	//	have: "test error"
}

func ExampleNoError_changeMessage() {
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
}
