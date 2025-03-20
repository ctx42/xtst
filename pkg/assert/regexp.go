package assert

import (
	"github.com/ctx42/xtst/pkg/check"
	"github.com/ctx42/xtst/pkg/tester"
)

// Regexp asserts that "want" regexp matches a string representation of "have.
// Returns true if it is, otherwise marks the test as failed, writes error
// message to test log and returns false.
//
// The "want" can be either regular expression string or instance of
// [regexp.Regexp]. The [fmt.Sprint] s used to get string representation of
// have argument.
func Regexp(t tester.T, want, have any, opts ...check.Option) bool {
	t.Helper()
	if e := check.Regexp(want, have, opts...); e != nil {
		t.Error(e)
		return false
	}
	return true
}
