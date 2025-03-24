// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/ctx42/xtst/pkg/dump"
	"github.com/ctx42/xtst/pkg/notice"
)

// TODO(rz): track (log) all the fields (trails).
// TODO(rz): make very detailed code review of this file.

// Equal recursively checks both values are equal. Returns nil if they are,
// otherwise it returns an error with a message indicating the expected and
// actual values.
//
// nolint: cyclop, gocognit
func Equal(want, have any, opts ...Option) error {
	return deepEqual(reflect.ValueOf(want), reflect.ValueOf(have), opts...)
}

// deepEqual is the internal recursive comparison function.
func deepEqual(a, b reflect.Value, opts ...Option) error {
	if !a.IsValid() && !b.IsValid() {
		return nil
	}

	if !a.IsValid() || !b.IsValid() {
		var aItf, bItf any
		if a.IsValid() {
			aItf = a.Interface()
		}
		if b.IsValid() {
			bItf = b.Interface()
		}
		ops := DefaultOptions(opts...)
		return equalError(aItf, bItf, ops)
	}

	aType := a.Type()
	bType := b.Type()
	if aType != bType {
		ops := DefaultOptions(opts...)
		return equalError(a.Interface(), b.Interface(), ops)
	}

	switch a.Kind() {
	case reflect.Ptr:
		if aType == typTimeLocPtr && bType == typTimeLocPtr {
			aZone := a.Interface().(*time.Location)
			bZone := b.Interface().(*time.Location)
			return Zone(aZone, bZone, opts...)
		}

		if a.IsNil() && b.IsNil() {
			DefaultOptions(opts...).logTrail()
			return nil
		}
		if a.IsNil() || b.IsNil() {
			ops := DefaultOptions(opts...).logTrail()
			wItf := a.Interface()
			hItf := b.Interface()
			return equalError(wItf, hItf, ops)
		}

		// TODO(rz):
		// aPtr := a.Pointer()
		// bPtr := b.Pointer()

		ops := DefaultOptions(opts...)
		// TODO(rz):
		// if aPtr == bPtr {
		// 	return nil
		// }
		return deepEqual(a.Elem(), b.Elem(), WithOptions(ops))

	case reflect.Struct:
		aTyp := a.Type()
		bTyp := b.Type()
		if aTyp == typTime && bTyp == typTime {
			return Time(a.Interface(), b.Interface(), opts...)
		}
		// TODO(rz): what if someone tries to compare time.Location instead of
		//  *time.Location?
		// if aTyp == typTimeLoc && bTyp == typTimeLoc {
		// 	aZone := a.Interface().(time.Location)
		// 	bZone := b.Interface().(time.Location)
		// 	return Zone(&aZone, &bZone, opts...)
		// }
		typeName := a.Type().Name()
		sOps := DefaultOptions(opts...).structTrail(typeName, "")

		var ers error
		for i := 0; i < a.NumField(); i++ {
			aVal := a.Field(i)
			bVal := b.Field(i)
			if !(aVal.IsValid() && aVal.CanInterface()) {
				continue
			}
			sf := a.Type().Field(i)
			iOps := sOps.structTrail("", sf.Name)
			iOps.skipType = true
			if err := deepEqual(aVal, bVal, WithOptions(iOps)); err != nil {
				ers = errors.Join(ers, err)
			}
		}
		return ers

	case reflect.Slice, reflect.Array:
		if a.Len() != b.Len() {
			ops := DefaultOptions(opts...)
			return equalError(a.Interface(), b.Interface(), ops)
		}
		if a.Pointer() == b.Pointer() {
			return nil
		}
		var ers error
		for i := 0; i < a.Len(); i++ {
			aVal := a.Index(i)
			bVal := b.Index(i)
			iOps := DefaultOptions(opts...).arrTrail(i)
			if err := deepEqual(aVal, bVal, WithOptions(iOps)); err != nil {
				ers = errors.Join(ers, err)
			}
		}
		return ers

	case reflect.Map:
		if a.Len() != b.Len() {
			ops := DefaultOptions(opts...)
			return equalError(a.Interface(), b.Interface(), ops)
		}
		if a.Pointer() == b.Pointer() {
			return nil
		}

		keys := a.MapKeys()
		sort.Slice(keys, func(i, j int) bool {
			return valToString(keys[i]) < valToString(keys[j])
		})

		var ers error
		for _, key := range keys {
			aVal := a.MapIndex(key)
			bVal := b.MapIndex(key)
			kOps := DefaultOptions(opts...).mapTrail(valToString(key))
			if !bVal.IsValid() {
				aItf := b.Interface()
				ers = errors.Join(ers, equalError(aItf, nil, kOps))
				continue
			}
			if err := deepEqual(aVal, bVal, WithOptions(kOps)); err != nil {
				ers = errors.Join(ers, err)
			}
		}
		return ers

	case reflect.Interface:
		ops := DefaultOptions(opts...).logTrail()
		if a.IsNil() && b.IsNil() {
			return nil
		}
		if a.IsNil() || b.IsNil() {
			return equalError(a.Interface(), b.Interface(), ops)
		}
		aElem := a.Elem()
		bElem := b.Elem()
		return deepEqual(aElem, bElem, WithOptions(ops))

	case reflect.Bool:
		ops := DefaultOptions(opts...).logTrail()
		if a.Bool() == b.Bool() {
			return nil
		}
		return equalError(a.Interface(), b.Interface(), ops)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ops := DefaultOptions(opts...).logTrail()
		if a.Int() == b.Int() {
			return nil
		}
		return equalError(a.Interface(), b.Interface(), ops)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		ops := DefaultOptions(opts...).logTrail()
		if a.Uint() == b.Uint() {
			return nil
		}
		return equalError(a.Interface(), b.Interface(), ops)

	case reflect.Float32, reflect.Float64:
		ops := DefaultOptions(opts...).logTrail()
		if a.Float() == b.Float() {
			return nil
		}
		return equalError(a.Interface(), b.Interface(), ops)

	case reflect.String:
		ops := DefaultOptions(opts...).logTrail()
		if a.String() == b.String() {
			return nil
		}
		return equalError(a.Interface(), b.Interface(), ops)

	case reflect.Chan, reflect.Func:
		ops := DefaultOptions(opts...).logTrail()
		if a.Pointer() == b.Pointer() {
			return nil
		}
		return equalError(a.Interface(), b.Interface(), ops)

	default:
		ops := DefaultOptions(opts...).logTrail()
		// For types, we haven't explicitly handled, use DeepEqual.
		if reflect.DeepEqual(a.Interface(), b.Interface()) {
			return nil
		}
		return equalError(a.Interface(), b.Interface(), ops)
	}
}

// equalError returns error for not equal values.
func equalError(want, have any, ops Options) error {
	wTyp, hTyp := fmt.Sprintf("%T", want), fmt.Sprintf("%T", have)
	if wTyp == hTyp {
		wTyp, hTyp = "", ""
	}

	msg := notice.New("expected values to be equal").
		Trail(ops.Trail)

	if b, ok := want.(byte); ok && isPrintableChar(b) {
		_ = msg.Want("%#v ('%s')", want, string(b))
	} else {
		_ = msg.Want("%s", dump.New(ops.DumpCfg).Dump(1, want))
	}

	if b, ok := have.(byte); ok && isPrintableChar(b) {
		_ = msg.Have("%#v ('%s')", have, string(b))
	} else {
		_ = msg.Have("%s", dump.New(ops.DumpCfg).Dump(1, have))
	}

	if wTyp != "" {
		// nolint: govet
		_ = msg.
			Append("want type", "%s", wTyp).
			Append("have type", "%s", hTyp)
	}
	return msg
}
