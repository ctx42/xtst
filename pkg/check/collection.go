// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"github.com/ctx42/xtst/internal"
	"github.com/ctx42/xtst/pkg/dump"
	"github.com/ctx42/xtst/pkg/notice"
)

// Len checks "have" has "want" elements. Returns nil if it has, otherwise it
// returns an error with a message indicating the expected and actual values.
func Len(want int, have any, opts ...Option) error {
	cnt, ok := internal.Len(have)
	if !ok {
		return notice.New("cannot execute len(%T)", have)
	}
	if want != cnt {
		ops := DefaultOptions().set(opts)
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
	ops := DefaultOptions().set(opts)
	dmp := dump.New(ops.DumpConfig)
	return notice.New("expected slice to have a value").
		Trail(ops.Trail).
		Want("%#v", want).
		Append("slice", "%s", dmp.DumpAny(bag))
}

// HasNo checks slice does not have "want" value. Returns nil if it doesn't,
// otherwise it returns an error with a message indicating the expected and
// actual values.
func HasNo[T comparable](want T, set []T, opts ...Option) error {
	for i, got := range set {
		if want == got {
			ops := DefaultOptions().set(opts)
			dmp := dump.New(ops.DumpConfig)
			return notice.New("expected slice not to have value").
				Trail(ops.Trail).
				Want("%#v", want).
				Append("index", "%d", i).
				Append("slice", "%s", dmp.DumpAny(set))
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
	ops := DefaultOptions().set(opts)
	dmp := dump.New(ops.DumpConfig)
	return val, notice.New("expected map to have a key").
		Trail(ops.Trail).
		Append("key", "%#v", key).
		Append("map", "%s", dmp.DumpAny(set))
}

// HasNoKey checks map has no key. Returns nil if it doesn't, otherwise it
// returns an error with a message indicating the expected and actual values.
func HasNoKey[K comparable, V any](key K, set map[K]V, opts ...Option) error {
	val, ok := set[key]
	if !ok {
		return nil
	}
	ops := DefaultOptions().set(opts)
	dmp := dump.New(ops.DumpConfig)
	return notice.New("expected map not to have a key").
		Trail(ops.Trail).
		Append("key", "%#v", key).
		Append("value", "%#v", val).
		Append("map", "%s", dmp.DumpAny(set))
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
	ops := DefaultOptions().set(opts)
	return notice.New("expected map to have a key with a value").
		Trail(ops.Trail).
		Append("key", "%#v", key).
		Want("%#v", want).
		Have("%#v", have)
}

// TODO(rz):
// // MapSubset checks the "have" is a subset of "want". In other words all keys
// // and values in "want" map must be in "have" map, and it is not an error when
// // "have" map has some other keys. Returns error if "want" is not a subset of
// // "have" with a message indicating the expected and actual values.
// func MapSubset[K cmp.Ordered, V any](want, have map[K]V, opts ...Option) error {
// 	return mapSubset(want, have, opts...)
// }
//
// // mapSubset checks the "have" map is a subset of "want" map. In other words
// // all keys and values in "want" map must be in "have" map, and it's not an
// // error when "have" map have some other keys.
// func mapSubset[K cmp.Ordered, V any](want, have map[K]V, opts ...Option) error {
// 	wKeys := make([]K, 0, len(want))
// 	for wKey := range want {
// 		wKeys = append(wKeys, wKey)
// 	}
// 	sort.Slice(wKeys, func(i, j int) bool { return wKeys[i] < wKeys[j] })
//
// 	var ers []error
// 	var missing []K
//     ops := DefaultOptions().set(opts)
// 	for _, wKey := range wKeys {
// 		hVal, exist := have[wKey]
// 		if !exist {
// 			missing = append(missing, wKey)
// 			continue
// 		}
// 		pth := ops.Trail + "[" + valToString(reflect.ValueOf(wKey)) + "]"
// 		if err := Equal(want[wKey], hVal, WithPath(pth)); err != nil {
// 			ers = append(ers, notice.Unwrap(err)...)
// 		}
// 	}
//
// 	var mKeys []string
// 	for _, wKey := range missing {
// 		mKeys = append(mKeys, valToString(reflect.ValueOf(wKey)))
// 	}
// 	sort.Strings(mKeys)
// 	if len(mKeys) > 0 {
// 		err := notice.New(`expected "have" map to have key(s)`).
// 			Append("missing key(s)", "%s", strings.Join(mKeys, ", "))
// 		ers = append(ers, err)
// 	}
// 	return errors.Join(ers...)
// }
//
// // MapsSubset checks all the "have" maps are subset of corresponding "want"
// // maps. In other words it iterates over "want" slice and checks map at the
// // same index is a subset of the map in "have" slice. Returns an error if any
// // of the "want" maps is not a subset of corresponding "have" map.
// func MapsSubset[K cmp.Ordered, V any](want, have []map[K]V, opts ...Option) error {
// 	ops := DefaultOptions().set(opts)
// 	if len(want) != len(have) {
// 		msg := "expected \"want\" and \"have\" to have the same number of " +
// 			"elements"
// 		return notice.New(msg).
// 			Trail(ops.Trail).
// 			Want("%d", len(want)).
// 			Have("%d", len(have))
// 	}
//
// 	var ers []error
// 	for i := range want {
// 		pth := fmt.Sprintf("[%d]map", i)
// 		if err := mapSubset(want[i], have[i], pth); err != nil {
// 			ers = append(ers, notice.Unwrap(err)...)
// 		}
// 	}
// 	return errors.Join(ers...)
// }

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

	ops := DefaultOptions().set(opts)
	dmp := dump.New(ops.DumpConfig)
	const hHeader = "expected \"want\" slice to be a subset of \"have\" slice"
	return notice.New(hHeader).
		Trail(ops.Trail).
		Append("missing values", "%s", dmp.DumpAny(missing))
}
