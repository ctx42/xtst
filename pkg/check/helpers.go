package check

import (
	"sync"

	"github.com/ctx42/xtst/pkg/dump"
)

// defaultDump is a function returning default value dumper.
var defaultDump = sync.OnceValue(func() dump.Dump {
	cfg := dump.NewConfig(dump.Flat, dump.Compact)
	return dump.New(cfg)
})

// Dump dumps given value as a string.
func Dump(value any) string { return defaultDump().DumpAny(value) }
