// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package dump.
//
// nolint: forbidigo
package dump

import (
	"bytes"
	"strings"
)

// printer represents code printer that is configuration aware.
type printer struct {
	cfg Config
	buf *strings.Builder
}

// newPrinter returns new printer configured by [Config].
func newPrinter(cfg Config) printer {
	return printer{cfg: cfg, buf: &strings.Builder{}}
}

// nli prints new line when not flat and at least one entry.
func (prn printer) nli(cnt int) printer {
	if !prn.cfg.Flat && cnt > 0 {
		prn.buf.WriteByte('\n')
	}
	return prn
}

// nl prints new line when not flat.
func (prn printer) nl() printer {
	if !prn.cfg.Flat {
		prn.buf.WriteByte('\n')
	}
	return prn
}

// comma prints comma when not flat and not last entry.
func (prn printer) comma(last bool) printer {
	if !(prn.cfg.Flat && last) {
		prn.buf.WriteByte(',')
	}
	return prn
}

// tab prints indentation with n spaces when not flat.
func (prn printer) tab(n int) printer {
	if prn.cfg.Flat {
		return prn
	}
	if n < 0 {
		n = 0
	}
	prn.buf.Write(bytes.Repeat([]byte{' '}, n*prn.cfg.TabWidth))
	return prn
}

// space writes a space when not compact.
func (prn printer) space() printer {
	if !prn.cfg.Compact {
		prn.buf.WriteByte(' ')
	}
	return prn
}

// sep writes separator space.
func (prn printer) sep(last bool) printer {
	if !prn.cfg.Compact && !last && prn.cfg.Flat {
		prn.buf.WriteByte(' ')
	}
	return prn
}

// write writes string to the builder.
func (prn printer) write(s string) printer {
	prn.buf.WriteString(s)
	return prn
}

// String returns built string.
func (prn printer) String() string { return prn.buf.String() }
