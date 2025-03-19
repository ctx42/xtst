package playground

import (
	"fmt"
	"testing"

	"github.com/ctx42/xtst/internal/diff"
	"github.com/ctx42/xtst/pkg/dump"
)

func Test_Name(t *testing.T) {
	// --- Given ---
	type Node struct {
		Value    int
		Children []*Node
	}

	want := &Node{
		Value: 1,
		Children: []*Node{
			{Value: 2, Children: nil},
			{
				Value: 3,
				Children: []*Node{
					{Value: 4, Children: nil},
				},
			},
		},
	}

	have := &Node{
		Value: 1,
		Children: []*Node{
			{Value: 2, Children: nil},
			{
				Value: 3,
				Children: []*Node{
					{Value: 5, Children: nil},
				},
			},
		},
	}

	wantD := dump.DefaultDump().DumpAny(want)
	haveD := dump.DefaultDump().DumpAny(have)

	// --- When ---
	edits := diff.Strings(wantD, haveD)

	// --- Then ---
	// fmt.Println(edits)
	d, _ := diff.ToUnified("want", "have", wantD, edits, 2)
	fmt.Println(d) // TODO(rz):
}
