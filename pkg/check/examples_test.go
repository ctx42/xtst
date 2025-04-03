// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check_test

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ctx42/testing/pkg/check"
	"github.com/ctx42/testing/pkg/notice"
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
	//   want: <nil>
	//   have: "test error"
}

func ExampleNoError_withTrail() {
	have := errors.New("test error")

	err := check.NoError(have, check.WithTrail("type.field"))

	fmt.Println(err)
	// Output:
	// expected error to be nil:
	//   trail: type.field
	//    want: <nil>
	//    have: "test error"
}

func ExampleNoError_changeMessage() {
	have := errors.New("test error")

	err := check.NoError(have, check.WithTrail("type.field"))

	err = notice.From(err, "prefix").Append("context", "wow")

	fmt.Println(err)
	// Output:
	// [prefix] expected error to be nil:
	//     trail: type.field
	//      want: <nil>
	//      have: "test error"
	//   context: wow
}

func ExampleEqual_structs() {
	type T struct {
		Int int
		Str string
	}

	have := T{Int: 1, Str: "abc"}
	want := T{Int: 2, Str: "xyz"}

	err := check.Equal(want, have)

	fmt.Println(err)
	// Output:
	// expected values to be equal:
	//   trail: T.Int
	//    want: 2
	//    have: 1
	//  ---
	//   trail: T.Str
	//    want: "xyz"
	//    have: "abc"
}

func ExampleEqual_recursiveStructs() {
	type T struct {
		Int  int
		Next *T
	}

	have := T{1, &T{2, &T{3, &T{42, nil}}}}
	want := T{1, &T{2, &T{3, &T{4, nil}}}}

	err := check.Equal(want, have)

	fmt.Println(err)
	// Output:
	// expected values to be equal:
	//   trail: T.Next.Next.Next.Int
	//    want: 4
	//    have: 42
}

func ExampleEqual_maps() {
	type T struct {
		Str string
	}

	want := map[int]T{1: {Str: "abc"}, 2: {Str: "xyz"}}
	have := map[int]T{1: {Str: "abc"}, 3: {Str: "xyz"}}

	err := check.Equal(want, have)

	fmt.Println(err)
	// Output:
	// expected values to be equal:
	//       trail: map[2]
	//        want:
	//              map[int]check_test.T{
	//                1: {
	//                  Str: "abc",
	//                },
	//                3: {
	//                  Str: "xyz",
	//                },
	//              }
	//        have: nil
	//   want type: map[int]check_test.T
	//   have type: <nil>
}

func ExampleEqual_arrays() {
	want := [...]int{1, 2, 3}
	have := [...]int{1, 2, 3, 4}

	err := check.Equal(want, have)

	fmt.Println(err)
	// Output:
	// expected values to be equal:
	//        want:
	//              [3]int{
	//                1,
	//                2,
	//                3,
	//              }
	//        have:
	//              [4]int{
	//                1,
	//                2,
	//                3,
	//                4,
	//              }
	//   want type: [3]int
	//   have type: [4]int
}

func ExampleEqual_slices() {
	want := []int{1, 2, 3}
	have := []int{1, 2, 3, 4}

	err := check.Equal(want, have)

	fmt.Println(err)
	// Output:
	// expected values to be equal:
	//   want len: 3
	//   have len: 4
	//       want:
	//             []int{
	//               1,
	//               2,
	//               3,
	//             }
	//       have:
	//             []int{
	//               1,
	//               2,
	//               3,
	//               4,
	//             }
}

func ExampleEqual_customTrailCheckers() {
	type T struct {
		Str string
		Any []any
	}

	chk := func(want, have any, opts ...check.Option) error {
		wVal := want.(float64)
		hVal := want.(float64)
		return check.Epsilon(wVal, 0.01, hVal, opts...)
	}
	opt := check.WithTrailChecker("T.Any[1]", chk)

	want := T{Str: "abc", Any: []any{1, 2.123, "abc"}}
	have := T{Str: "abc", Any: []any{1, 2.124, "abc"}}

	err := check.Equal(want, have, opt)

	fmt.Println(err)
	// Output:
	//  <nil>
}

func ExampleEqual_listVisitedTrails() {
	type T struct {
		Int  int
		Next *T
	}

	have := T{1, &T{2, &T{3, &T{42, nil}}}}
	want := T{1, &T{2, &T{3, &T{42, nil}}}}
	trails := make([]string, 0)

	err := check.Equal(want, have, check.WithTrailLog(&trails))

	fmt.Println(err)
	fmt.Println(strings.Join(trails, "\n"))
	// Output:
	// <nil>
	// T.Int
	// T.Next.Int
	// T.Next.Next.Int
	// T.Next.Next.Next.Int
	// T.Next.Next.Next.Next
}

func ExampleEqual_skipTrails() {
	type T struct {
		Int  int
		Next *T
	}

	have := T{1, &T{2, &T{3, &T{42, nil}}}}
	want := T{1, &T{2, &T{8, &T{42, nil}}}}
	trails := make([]string, 0)

	err := check.Equal(
		want,
		have,
		check.WithTrailLog(&trails),
		check.WithSkipTrail("T.Next.Next.Int"),
	)

	fmt.Println(err)
	fmt.Println(strings.Join(trails, "\n"))
	// Output:
	// <nil>
	// T.Int
	// T.Next.Int
	// T.Next.Next.Int <skipped>
	// T.Next.Next.Next.Int
	// T.Next.Next.Next.Next
}

func ExampleJSON() {
	want := `{"A": 1, "B": 2}`
	have := `{"A": 1, "B": 3}`

	err := check.JSON(want, have)

	fmt.Println(err)
	// Output:
	// expected JSON strings to be equal:
	//   want: {"A":1,"B":2}
	//   have: {"A":1,"B":3}
}

func ExampleTime() {
	want := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	have := time.Date(2025, 1, 1, 0, 1, 1, 0, time.UTC)

	err := check.Time(want, have)

	fmt.Println(err)
	// Output:
	//  expected equal dates:
	//   want: 2025-01-01T00:00:00Z
	//   have: 2025-01-01T00:01:01Z
	//   diff: -1m1s
}
