// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package notice

// Row represents [Notice] row.
type Row struct {
	Name   string
	Format string
	Args   []any
}

// NewRow is constructor function for [Row].
func NewRow(name, format string, args ...any) Row {
	return Row{Name: name, Format: format, Args: args}
}

// Trail is convenience constructor function for [Row] with name `path`.
func Trail(name string) Row {
	return Row{Name: trail, Format: name}
}
