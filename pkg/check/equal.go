// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"sort"
	"time"

	"github.com/ctx42/testing/pkg/notice"
)

// Equal recursively checks both values are equal. Returns nil if they are,
// otherwise it returns an error with a message indicating the expected and
// actual values.
//
// nolint: cyclop, gocognit
func Equal(want, have any, opts ...Option) error {
	wVal := reflect.ValueOf(want)
	hVal := reflect.ValueOf(have)
	return wrap(deepEqual(wVal, hVal, opts...))
}

// deepEqual is the internal comparison function which is called recursively.
//
// nolint: gocognit, cyclop
func deepEqual(wVal, hVal reflect.Value, opts ...Option) error {
	ops := DefaultOptions(opts...)

	if i := slices.Index(ops.SkipTrails, ops.Trail); i >= 0 {
		ops.Trail += " <skipped>"
		ops.logTrail()
		return nil
	}

	if !wVal.IsValid() && !hVal.IsValid() {
		ops.logTrail()
		return nil
	}

	if !wVal.IsValid() || !hVal.IsValid() {
		var wItf, hItf any
		if wVal.IsValid() {
			wItf = wVal.Interface()
		}
		if hVal.IsValid() {
			hItf = hVal.Interface()
		}
		ops.logTrail()
		return equalError(wItf, hItf, ops)
	}

	wType := wVal.Type()
	hType := hVal.Type()
	if wType != hType {
		ops.logTrail()
		return equalError(wVal.Interface(), hVal.Interface(), ops)
	}

	if chk, ok := ops.TrailCheckers[ops.Trail]; ok {
		ops.logTrail()
		return chk(wVal.Interface(), hVal.Interface(), WithOptions(ops))
	}

	if chk, ok := ops.TypeCheckers[wType]; ok {
		ops.logTrail()
		return chk(wVal.Interface(), hVal.Interface(), opts...)
	}

	switch knd := wVal.Kind(); knd {
	case reflect.Ptr:
		if wType == typTimeLocPtr && hType == typTimeLocPtr {
			ops.logTrail()
			wZone := wVal.Interface().(*time.Location) // nolint: forcetypeassert
			hZone := hVal.Interface().(*time.Location) // nolint: forcetypeassert
			return Zone(wZone, hZone, WithOptions(ops))
		}

		if wVal.IsNil() && hVal.IsNil() {
			ops.logTrail()
			return nil
		}
		if wVal.IsNil() || hVal.IsNil() {
			ops.logTrail()
			wItf := wVal.Interface()
			hItf := hVal.Interface()
			return equalError(wItf, hItf, ops)
		}

		return deepEqual(wVal.Elem(), hVal.Elem(), WithOptions(ops))

	case reflect.Struct:
		wTyp := wVal.Type()
		hTyp := hVal.Type()
		if wTyp == typTime && hTyp == typTime {
			ops.logTrail()
			return Time(wVal.Interface(), hVal.Interface(), opts...)
		}
		if wTyp == typTimeLoc && hTyp == typTimeLoc {
			ops.logTrail()
			wZone := wVal.Interface().(time.Location) // nolint: forcetypeassert
			hZone := hVal.Interface().(time.Location) // nolint: forcetypeassert
			return Zone(&wZone, &hZone, opts...)
		}
		typeName := wVal.Type().Name()

		sOps := ops
		trail := ops.structTrail(typeName, "")
		sOps.Trail = trail

		var ers []error
		for i := 0; i < wVal.NumField(); i++ {
			wfVal := wVal.Field(i)
			hfVal := hVal.Field(i)
			if !(wfVal.IsValid() && wfVal.CanInterface()) {
				continue
			}
			wSF := wVal.Type().Field(i)
			trail = sOps.structTrail("", wSF.Name)
			iOps := sOps
			iOps.Trail = trail
			if err := deepEqual(wfVal, hfVal, WithOptions(iOps)); err != nil {
				ers = append(ers, notice.Unwrap(err)...)
			}
		}
		return errors.Join(ers...)

	case reflect.Slice, reflect.Array:
		if wVal.Len() != hVal.Len() {
			ops.logTrail()
			wItf := wVal.Interface()
			hItf := hVal.Interface()
			return equalError(wItf, hItf, ops).
				Prepend("have len", "%d", hVal.Len()).
				Prepend("want len", "%d", wVal.Len())
		}
		if knd == reflect.Slice && wVal.Pointer() == hVal.Pointer() {
			ops.logTrail()
			return nil
		}
		var ers []error
		for i := 0; i < wVal.Len(); i++ {
			wiVal := wVal.Index(i)
			hiVal := hVal.Index(i)
			iOps := ops
			trail := ops.arrTrail(knd.String(), i)
			iOps.Trail = trail
			if err := deepEqual(wiVal, hiVal, WithOptions(iOps)); err != nil {
				ers = append(ers, notice.Unwrap(err)...)
			}
		}
		return errors.Join(ers...)

	case reflect.Map:
		if wVal.Len() != hVal.Len() {
			ops.logTrail()
			return equalError(wVal.Interface(), hVal.Interface(), ops).
				Prepend("have len", "%d", hVal.Len()).
				Prepend("want len", "%d", wVal.Len())
		}
		if wVal.Pointer() == hVal.Pointer() {
			ops.logTrail()
			return nil
		}

		keys := wVal.MapKeys()
		sort.Slice(keys, func(i, j int) bool {
			return valToString(keys[i]) < valToString(keys[j])
		})

		var ers []error
		for _, key := range keys {
			wkVal := wVal.MapIndex(key)
			hkVal := hVal.MapIndex(key)
			kOps := ops
			trail := ops.mapTrail(valToString(key))
			kOps.Trail = trail
			if !hkVal.IsValid() {
				hItf := hVal.Interface()
				err := equalError(hItf, nil, kOps)
				ers = append(ers, notice.Unwrap(err)...)
				continue
			}
			if err := deepEqual(wkVal, hkVal, WithOptions(kOps)); err != nil {
				ers = append(ers, notice.Unwrap(err)...)
			}
		}
		return errors.Join(ers...)

	case reflect.Interface:
		wElem := wVal.Elem()
		hElem := hVal.Elem()
		return deepEqual(wElem, hElem, WithOptions(ops))

	case reflect.Bool:
		ops.logTrail()
		if wVal.Bool() == hVal.Bool() {
			return nil
		}
		return equalError(wVal.Interface(), hVal.Interface(), ops)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ops.logTrail()
		if wVal.Int() == hVal.Int() {
			return nil
		}
		return equalError(wVal.Interface(), hVal.Interface(), ops)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		ops.logTrail()
		if wVal.Uint() == hVal.Uint() {
			return nil
		}
		return equalError(wVal.Interface(), hVal.Interface(), ops)

	case reflect.Float32, reflect.Float64:
		ops.logTrail()
		if wVal.Float() == hVal.Float() {
			return nil
		}
		return equalError(wVal.Interface(), hVal.Interface(), ops)

	case reflect.Complex64, reflect.Complex128:
		ops.logTrail()
		if wVal.Complex() == hVal.Complex() {
			return nil
		}
		return equalError(wVal.Interface(), hVal.Interface(), ops)

	case reflect.String:
		ops.logTrail()
		if wVal.String() == hVal.String() {
			return nil
		}
		return equalError(wVal.Interface(), hVal.Interface(), ops)

	case reflect.Chan, reflect.Func:
		ops.logTrail()
		if wVal.Pointer() == hVal.Pointer() {
			return nil
		}
		return equalError(wVal.Interface(), hVal.Interface(), ops)

	default:
		ops.logTrail()
		// For types, we haven't explicitly handled, use DeepEqual.
		if reflect.DeepEqual(wVal.Interface(), hVal.Interface()) {
			return nil
		}
		return equalError(wVal.Interface(), hVal.Interface(), ops)
	}
}

// equalError returns error for not equal values.
func equalError(want, have any, ops Options) *notice.Notice {
	wTyp, hTyp := fmt.Sprintf("%T", want), fmt.Sprintf("%T", have)
	if wTyp == hTyp {
		wTyp, hTyp = "", ""
	}

	msg := notice.New("expected values to be equal").
		Trail(ops.Trail)

	if b, ok := want.(byte); ok && isPrintableChar(b) {
		_ = msg.Want("%#v ('%s')", want, string(b))
	} else {
		_ = msg.Want("%s", ops.Dumper.Any(want))
	}

	if b, ok := have.(byte); ok && isPrintableChar(b) {
		_ = msg.Have("%#v ('%s')", have, string(b))
	} else {
		_ = msg.Have("%s", ops.Dumper.Any(have))
	}

	if wTyp != "" {
		// nolint: govet
		_ = msg.
			Append("want type", "%s", wTyp).
			Append("have type", "%s", hTyp)
	}
	return msg
}
