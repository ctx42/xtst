// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"github.com/ctx42/testing/internal"
	"github.com/ctx42/testing/pkg/notice"
)

// Len checks "have" has "want" elements. Returns nil if it has, otherwise it
// returns an error with a message indicating the expected and actual values.
func Len(want int, have any, opts ...Option) error {
	cnt, ok := internal.Len(have)
	if !ok {
		return notice.New("cannot execute len(%T)", have)
	}
	if want != cnt {
		ops := DefaultOptions(opts...)
		msg := notice.New("expected %T length", have).
			Trail(ops.Trail).
			Want("%d", want).
			Have("%d", cnt)
		return msg
	}
	return nil
}

// Has checks slice has "want" value. Returns nil if it does, otherwise it
// returns an error with a message indicating the expected and actual values.
func Has[T comparable](want T, bag []T, opts ...Option) error {
	for _, got := range bag {
		if want == got {
			return nil
		}
	}
	ops := DefaultOptions(opts...)
	return notice.New("expected slice to have a value").
		Trail(ops.Trail).
		Want("%#v", want).
		Append("slice", "%s", ops.Dumper.Any(bag))
}

// HasNo checks slice does not have "want" value. Returns nil if it doesn't,
// otherwise it returns an error with a message indicating the expected and
// actual values.
func HasNo[T comparable](want T, set []T, opts ...Option) error {
	for i, got := range set {
		if want == got {
			ops := DefaultOptions(opts...)
			return notice.New("expected slice not to have value").
				Trail(ops.Trail).
				Want("%#v", want).
				Append("index", "%d", i).
				Append("slice", "%s", ops.Dumper.Any(set))
		}
	}
	return nil
}

// HasKey checks map has a key. If the key exists it returns its value and nil,
// otherwise it returns zero value and an error with a message indicating the
// expected and actual values.
func HasKey[K comparable, V any](key K, set map[K]V, opts ...Option) (V, error) {
	val, ok := set[key]
	if ok {
		return val, nil
	}
	ops := DefaultOptions(opts...)
	return val, notice.New("expected map to have a key").
		Trail(ops.Trail).
		Append("key", "%#v", key).
		Append("map", "%s", ops.Dumper.Any(set))
}

// HasNoKey checks map has no key. Returns nil if it doesn't, otherwise it
// returns an error with a message indicating the expected and actual values.
func HasNoKey[K comparable, V any](key K, set map[K]V, opts ...Option) error {
	val, ok := set[key]
	if !ok {
		return nil
	}
	ops := DefaultOptions(opts...)
	return notice.New("expected map not to have a key").
		Trail(ops.Trail).
		Append("key", "%#v", key).
		Append("value", "%#v", val).
		Append("map", "%s", ops.Dumper.Any(set))
}

// HasKeyValue checks map has a key with given value. Returns nil if it doesn't,
// otherwise it returns an error with a message indicating the expected and
// actual values.
func HasKeyValue[K, V comparable](key K, want V, set map[K]V, opts ...Option) error {
	have, err := HasKey(key, set, opts...)
	if err != nil {
		return err
	}
	if want == have {
		return nil
	}
	ops := DefaultOptions(opts...)
	return notice.New("expected map to have a key with a value").
		Trail(ops.Trail).
		Append("key", "%#v", key).
		Want("%#v", want).
		Have("%#v", have)
}

// SliceSubset checks the "have" is a subset "want". In other words all values
// in "want" slice must be in "have" slice. Returns nil if it does, otherwise
// returns an error with a message indicating the expected and actual values.
func SliceSubset[V comparable](want, have []V, opts ...Option) error {
	var missing []V
	for _, wantVal := range want {
		found := false
		for _, haveVal := range have {
			if wantVal == haveVal {
				found = true
				break
			}
		}
		if !found {
			missing = append(missing, wantVal)
		}
	}
	if len(missing) == 0 {
		return nil
	}

	ops := DefaultOptions(opts...)
	const hHeader = "expected \"want\" slice to be a subset of \"have\" slice"
	return notice.New(hHeader).
		Trail(ops.Trail).
		Append("missing values", "%s", ops.Dumper.Any(missing))
}
