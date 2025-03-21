package check

import (
	"fmt"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
	"github.com/ctx42/xtst/internal/types"
)

func Test_Count(t *testing.T) {
	t.Run("error unsupported what type", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := Count(1, 123, "ab cd ef", opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected argument \"what\" to be string got int:\n" +
			"\ttrail: type.field"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error unsupported where type", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := Count(1, "ab", 123, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "unsupported \"where\" type: int:\n" +
			"\ttrail: type.field"
		affirm.Equal(t, wMsg, err.Error())
	})

	t.Run("error count with option", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := Count(2, "a", "abc abc anc", opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected string to contain substrings:\n" +
			"\t     trail: type.field\n" +
			"\twant count: 2\n" +
			"\thave count: 3\n" +
			"\t      what: \"a\"\n" +
			"\t     where: \"abc abc anc\""
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_Count_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		count int
		what  any
		where any
	}{
		{"one", 1, "ab", "ab cd ef"},
		{"multiple", 2, "ab", "ab cd ab"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			err := Count(tc.count, tc.what, tc.where)

			// --- Then ---
			affirm.Nil(t, err)
		})
	}
}

func Test_Count_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		wantCnt int
		haveCnt int
		what    any
		where   any
	}{
		{"not existing", 1, 0, "gh", "ab cd ef"},
		{"existing with wrong count", 2, 1, "ab", "ab cd ef"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			err := Count(tc.wantCnt, tc.what, tc.where)

			// --- Then ---
			affirm.NotNil(t, err)
			wMsg := "expected string to contain substrings:\n" +
				"\twant count: %d\n" +
				"\thave count: %d\n" +
				"\t      what: %q\n" +
				"\t     where: %q"
			wMsg = fmt.Sprintf(wMsg, tc.wantCnt, tc.haveCnt, tc.what, tc.where)
			affirm.Equal(t, wMsg, err.Error())
		})
	}
}

func Test_SameType(t *testing.T) {
	t.Run("error with option", func(t *testing.T) {
		// --- Given ---
		opt := WithTrail("type.field")

		// --- When ---
		err := SameType(42, 4.2, opt)

		// --- Then ---
		affirm.NotNil(t, err)
		wMsg := "expected same types:\n" +
			"\ttrail: type.field\n" +
			"\t want: int\n" +
			"\t have: float64"
		affirm.Equal(t, wMsg, err.Error())
	})
}

func Test_SameType_success_tabular(t *testing.T) {
	var ptr *types.TPtr
	var itf types.TItf
	itf = &types.TPtr{}

	tt := []struct {
		testN string

		val0 any
		val1 any
	}{
		{"int", 42, 44},
		{"float64", 42.0, 44.0},
		{"bool", true, false},
		{"nil ptr 0", ptr, &types.TPtr{}},
		{"nil ptr 1", &types.TPtr{}, ptr},
		{"nil itf 0", itf, &types.TPtr{}},
		{"nil itf 1", &types.TPtr{}, itf},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			err := SameType(tc.val0, tc.val1)

			// --- Then ---
			affirm.Nil(t, err)
		})
	}
}

func Test_SameType_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		val0 any
		val1 any
		wMsg string
	}{
		{
			"different types",
			42,
			42.0,
			"expected same types:\n\twant: int\n\thave: float64",
		},
		{
			"different ptr types",
			&types.TPtr{},
			&types.TVal{},
			"expected same types:\n\twant: *types.TPtr\n\thave: *types.TVal",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			err := SameType(tc.val0, tc.val1)

			// --- Then ---
			affirm.NotNil(t, err)
			affirm.Equal(t, tc.wMsg, err.Error())
		})
	}
}
