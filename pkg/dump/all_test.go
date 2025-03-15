// Copyright (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package dump

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

// ReadGolden is a helper return contents of a file at given path. The golden
// file may start with multiple comment lines which start with "# " followed by
// two empty lines. Everything after two empty lines is considered as golden
// file content.
//
// nolint: cyclop
func ReadGolden(t *testing.T, pth string) string {
	t.Helper()

	fil, err := os.Open(pth)
	if err != nil {
		t.Error(err)
		return ""
	}
	defer func() { _ = fil.Close() }()

	var data strings.Builder
	var line string
	var commentsCnt int   // Number of comment lines.
	var emptyCnt int      // Number of empty lines after comment lines.
	var commentsDone bool // Parsing comment lines done.

	rdr := bufio.NewReader(fil)
	for {
		if line, err = rdr.ReadString('\n'); err != nil {
			if !errors.Is(err, io.EOF) {
				break
			}
		}

		if !commentsDone {
			if strings.HasPrefix(line, "# ") {
				commentsCnt++
				continue
			}
			if commentsCnt == 0 {
				commentsDone = true
			} else {
				if commentsCnt > 0 && line == "\n" {
					emptyCnt++
				}
				if commentsCnt > 0 && emptyCnt == 2 {
					commentsDone = true
				}
				if commentsCnt > 0 && line != "\n" {
					t.Errorf("invalid golden file format: %s", pth)
					return ""
				}
				continue
			}
		}

		data.WriteString(line)
		if errors.Is(err, io.EOF) {
			err = nil
			break
		}
	}
	if err != nil {
		t.Error(err)
		return ""
	}

	return data.String()
}
