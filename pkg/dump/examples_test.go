package dump_test

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ctx42/xtst/internal/types"
	"github.com/ctx42/xtst/pkg/dump"
)

func ExampleDump_DumpAny() {
	val := types.TA{
		Dur: 3,
		Int: 42,
		Loc: types.WAW,
		Str: "abc",
		Tim: time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		TAp: nil,
	}

	have := dump.DefaultDump().DumpAny(val)

	fmt.Println(have)
	// Output:
	// {
	// 	Int: 42,
	// 	Str: "abc",
	// 	Tim: "2000-01-02T03:04:05Z",
	// 	Dur: "3ns",
	// 	Loc: "Europe/Warsaw",
	// 	TAp: nil,
	// }
}

func ExampleDump_DumpAny_flatCompact() {
	val := map[string]any{
		"int": 42,
		"loc": types.WAW,
		"nil": nil,
	}

	cfg := dump.NewConfig(dump.Flat)
	have := dump.New(cfg).DumpAny(val)

	fmt.Println(have)
	// Output:
	// map[string]any{"int": 42, "loc": "Europe/Warsaw", "nil": nil}
}

func ExampleDump_DumpAny_customTimeFormat() {
	val := map[time.Time]int{time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC): 42}

	cfg := dump.NewConfig(dump.Flat, dump.TimeFormat(time.Kitchen))
	have := dump.New(cfg).DumpAny(val)

	fmt.Println(have)
	// Output:
	// map[time.Time]int{"3:04AM": 42}
}

func ExampleDump_DumpAny_customDumper() {
	var i int
	dumper := func(dmp dump.Dump, lvl int, val reflect.Value) string {
		switch val.Kind() {
		case reflect.Int:
			return fmt.Sprintf("%X", val.Int())
		default:
			panic("unexpected kind")
		}
	}

	cfg := dump.NewConfig(dump.Flat, dump.Compact, dump.WithDumper(i, dumper))

	have := dump.New(cfg).DumpAny(42)

	fmt.Println(have)
	// Output:
	// 2A
}

func ExampleDump_DumpAny_recursive() {
	type Node struct {
		Value    int
		Children []*Node
	}

	val := &Node{
		Value: 1,
		Children: []*Node{
			{
				Value:    2,
				Children: nil,
			},
			{
				Value: 3,
				Children: []*Node{
					{
						Value:    4,
						Children: nil,
					},
				},
			},
		},
	}

	have := dump.DefaultDump().DumpAny(val)
	fmt.Println(have)
	// Output:
	// {
	// 	Value: 1,
	// 	Children: []*dump_test.Node{
	// 		{
	// 			Value: 2,
	// 			Children: nil,
	// 		},
	// 		{
	// 			Value: 3,
	// 			Children: []*dump_test.Node{
	// 				{
	// 					Value: 4,
	// 					Children: nil,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
