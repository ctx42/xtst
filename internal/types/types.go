// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package types provides example types used in tests.
package types

import (
	"errors"
	"fmt"
	"time"
)

// WAW represents Europe/Warsaw timezone.
var WAW *time.Location

func init() {
	var err error
	WAW, err = time.LoadLocation("Europe/Warsaw")
	if err != nil {
		panic(err)
	}
}

// /////////////////////////////////////////////////////////////////////////////

// IntType is type alias used in tests.
type IntType int

// /////////////////////////////////////////////////////////////////////////////

// TItf is an interface used in tests.
type TItf interface{ AAA() string }

// /////////////////////////////////////////////////////////////////////////////

// TInt is type used in tests.
type TInt struct{ V int }

// NewTInt returns error when v is not 42.
func NewTInt(v int) (*TInt, error) {
	if v != 42 {
		return nil, errors.New("not cool")
	}
	return &TInt{V: v}, nil
}

// /////////////////////////////////////////////////////////////////////////////

// TA is an example type with fields of different types.
type TA struct {
	Int int
	Str string
	Tim time.Time
	Dur time.Duration
	Loc *time.Location
	TAp *TA

	private int // nolint: unused
}

// /////////////////////////////////////////////////////////////////////////////

// TVal implements [TItf], has value receiver methods.
type TVal struct{ Val string } // nolint: errname

// AAA implements TItf.
func (typ TVal) AAA() string { return typ.Val }

// Error implements builtin in Error interface.
func (typ TVal) Error() string { return typ.Val }

// /////////////////////////////////////////////////////////////////////////////

// TPtr implements [TItf], has pointer receiver methods.
type TPtr struct{ Val string } // nolint: errname

// AAA implements TItf.
func (typ *TPtr) AAA() string { return typ.Val }

// Variadic1 is variadic function with one additional argument.
func (typ *TPtr) Variadic1(str string, i ...int) string {
	return fmt.Sprintf("%s %s %v", typ.Val, str, i)
}

// Error implements builtin in error interface.
func (typ *TPtr) Error() string { return typ.Val }

// PS adds two numbers.
func (typ *TPtr) PS(a, b string) string { return a + typ.Val + b }

// /////////////////////////////////////////////////////////////////////////////

// T1 represents nested structure.
type T1 struct {
	Int int
	T1  *T1
}

// /////////////////////////////////////////////////////////////////////////////
