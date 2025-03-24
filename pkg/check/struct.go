// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package check

// TODO(rz): make very detailed code review of this file.

// // structEqual is entrypoint for checking structs equivalency. Returns nil if
// // they are, otherwise it returns an error with a message indicating the
// // expected and actual values.
// func structEqual(want, have any, opts ...Option) error {
// 	wVal := reflect.ValueOf(want)
// 	hVal := reflect.ValueOf(have)
//
// 	// If "want" is a pointer type it must not be nil.
// 	if wVal.Kind() == reflect.Ptr && wVal.IsNil() {
// 		ops := DefaultOptions(opts...)
// 		return notice.New("expected not nil struct").
// 			Trail(ops.Trail).
// 			Append("which", "want").
// 			Want("<not-nil>").
// 			Have("<nil>")
// 	}
//
// 	// If "have" is a pointer type it must not be nil.
// 	if hVal.Kind() == reflect.Ptr && hVal.IsNil() {
// 		ops := DefaultOptions(opts...)
// 		return notice.New("expected not nil struct").
// 			Trail(ops.Trail).
// 			Append("which", "have").
// 			Want("<not-nil>").
// 			Have("<nil>")
// 	}
//
// 	// Check if wVal and hVal are of the same type
// 	if deRef(wVal).Type() != deRef(hVal).Type() {
// 		ops := DefaultOptions(opts...)
// 		return notice.New("expected struct of the same type").
// 			Trail(ops.Trail).
// 			Append("want type", "%s", wVal.Type().String()).
// 			Append("have type", "%s", hVal.Type().String())
// 	}
//
// 	wTyp := reflect.TypeOf(want)
// 	hTyp := reflect.TypeOf(have)
// 	if wTyp == typTime && hTyp == typTime {
// 		wTime := want.(time.Time) // nolint: forcetypeassert
// 		hTime := have.(time.Time) // nolint: forcetypeassert
// 		return Time(wTime, hTime, opts...)
// 	}
//
// 	if wTyp == typTimeLocPtr && hTyp == typTimeLocPtr {
// 		wTime := want.(*time.Location) // nolint: forcetypeassert
// 		hTime := have.(*time.Location) // nolint: forcetypeassert
// 		return Zone(wTime, hTime, opts...)
// 	}
//
// 	typeName := deRef(wVal).Type().Name()
// 	// TODO(rz): structPath better name? Better approach.
// 	ops := DefaultOptions(opts...).structTrail(typeName, "")
// 	if err := structEq(wVal, hVal, WithOptions(ops)); err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// // structEq checks both structures have fields with equal values. Returns nil
// // if they are, otherwise it returns an error with a message indicating the
// // expected and actual values.
// func structEq(want, have reflect.Value, opts ...Option) error {
// 	want = deRef(want)
// 	have = deRef(have)
// 	if want.Kind() != reflect.Struct || have.Kind() != reflect.Struct {
// 		ops := DefaultOptions(opts...)
// 		return notice.New("expected arguments to be structs").
// 			Trail(ops.Trail).
// 			Append("want type", "%s", typeString(want)).
// 			Append("have type", "%s", typeString(have))
// 	}
//
// 	wType := want.Type()
// 	var ers []error
// 	for i := 0; i < wType.NumField(); i++ {
// 		wFld := want.Field(i)
// 		hFld := have.Field(i)
// 		sf := wType.Field(i)
// 		ops := DefaultOptions(opts...).structTrail("", sf.Name)
//
// 		if sf.Anonymous && sf.Type.Kind() == reflect.Struct {
// 			err := structEq(wFld, hFld, WithOptions(ops))
// 			ers = append(ers, notice.Unwrap(err)...)
// 			continue
// 		}
// 		if !(wFld.IsValid() && wFld.CanInterface()) {
// 			continue
// 		}
// 		err := fieldEq(wFld, hFld, WithOptions(ops))
// 		ers = append(ers, notice.Unwrap(err)...)
// 	}
// 	return errors.Join(ers...)
// }
//
// // structEqual checks both fields contain equal values returns error if they
// // are.
// func fieldEq(wVal, hVal reflect.Value, opts ...Option) error {
// 	// TODO(rz):
// 	// if _, ok := ops.Skip[ops.Path]; ok {
// 	// 	return nil
// 	// }
//
// 	wTyp := wVal.Type()
//
// 	// Case when both fields are instances of [time.Location].
// 	if wTyp == reflect.TypeOf(&time.Location{}) {
// 		w := wVal.Interface().(*time.Location) // nolint: forcetypeassert
// 		h := hVal.Interface().(*time.Location) // nolint: forcetypeassert
// 		ops := DefaultOptions(opts...)
// 		return Zone(w, h, WithOptions(ops))
// 	}
//
// 	// Case when one or both fields are nil.
// 	if wVal.Kind() == reflect.Ptr && wVal.IsNil() {
// 		ops := DefaultOptions(opts...).logTrail() // TODO(rz): test trail log.
// 		if wVal.Kind() == reflect.Ptr && hVal.IsNil() {
// 			return nil
// 		}
// 		return notice.New("expected struct field to be nil").
// 			Trail(ops.Trail).
// 			Want("<nil>").
// 			Have("<not-nil>")
// 	}
//
// 	ops := DefaultOptions(opts...)
//
// 	wVal = deRef(wVal)
// 	hVal = deRef(hVal)
//
// 	// Case when both fields are [time.Time].
// 	if wTyp == typTime {
// 		w := wVal.Interface().(time.Time) // nolint: forcetypeassert
// 		h := hVal.Interface().(time.Time) // nolint: forcetypeassert
// 		// 	TODO(rz): should be configurable Time / TimeExactly.
// 		return Time(w, h, WithOptions(ops))
// 	}
//
// 	if wVal.Kind() == reflect.Func {
// 		return same(wVal.Interface(), hVal.Interface(), opts...)
// 	}
//
// 	if wVal.Kind() == reflect.Struct {
// 		return structEq(wVal, hVal, WithOptions(ops))
// 	}
//
// 	wItf := wVal.Interface()
// 	hItf := hVal.Interface()
// 	return Equal(wItf, hItf, WithOptions(ops))
// }
