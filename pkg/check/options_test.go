package check

import (
	"reflect"
	"testing"

	"github.com/ctx42/xtst/internal/affirm"
)

func Test_WithRoot(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := WithPath("pth")(ops)

	// --- Then ---
	affirm.Equal(t, "", ops.Path)
	affirm.Equal(t, "pth", have.Path)
}

func Test_DefaultOptions(t *testing.T) {
	// --- When ---
	have := DefaultOptions()

	// --- Then ---
	affirm.Equal(t, "", have.Path)
	affirm.Equal(t, 1, reflect.ValueOf(have).NumField())
}

func Test_Options_set(t *testing.T) {
	// --- Given ---
	ops := Options{}

	// --- When ---
	have := ops.set([]Option{WithPath("pth")})

	// --- Then ---
	affirm.Equal(t, "", ops.Path)
	affirm.Equal(t, "pth", have.Path)
}
