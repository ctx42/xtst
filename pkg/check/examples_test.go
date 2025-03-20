// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check_test

import (
	"errors"
	"fmt"

	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/notice"
)

func ExampleError() {
	err := check.Error(nil)

	fmt.Println(err)
	// Output:
	// expected non-nil error
}

func ExampleNoError() {
	have := errors.New("test error")

	err := check.NoError(have)

	fmt.Println(err)
	// Output:
	// expected error to be nil:
	//	want: <nil>
	//	have: "test error"
}

func ExampleNoError_withTrail() {
	have := errors.New("test error")

	err := check.NoError(have, check.WithTrail("type.field"))

	fmt.Println(err)
	// Output:
	// expected error to be nil:
	//	trail: type.field
	//	 want: <nil>
	//	 have: "test error"
}

func ExampleNoError_changeMessage() {
	have := errors.New("test error")

	err := check.NoError(have, check.WithTrail("type.field"))

	err = notice.From(err, "prefix").Append("context", "wow")

	fmt.Println(err)
	// Output:
	// [prefix] expected error to be nil:
	//	  trail: type.field
	//	   want: <nil>
	//	   have: "test error"
	//	context: wow
}
