package check

import (
	"fmt"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
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
