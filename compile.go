// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"fmt"
	"log"
)

// Compile is the main entry point for this package.
// Compile modifies f so that on return:
//   · all Values in f map to 0 or 1 assembly instructions of the target architecture
//   · the order of f.Blocks is the order to emit the Blocks
//   · the order of b.Values is the order to emit the Values in each Block
//   · f has a non-nil regAlloc field
func Compile(f *Func) {
	// TODO: debugging - set flags to control verbosity of compiler,
	// which phases to dump IR before/after, etc.
	fmt.Printf("compiling %s\n", f.Name)

	// hook to print function & phase if panic happens
	phaseName := "init"
	defer func() {
		if phaseName != "" {
			fmt.Printf("panic during %s while compiling %s\n", phaseName, f.Name)
		}
	}()

	// Run all the passes
	printFunc(f)
	checkFunc(f)
	for _, p := range passes {
		phaseName = p.name
		fmt.Printf("  pass %s begin\n", p.name)
		p.fn(f)
		fmt.Printf("  pass %s end\n", p.name)
		printFunc(f)
		checkFunc(f)
	}

	// Squash error printing defer
	phaseName = ""
}

type pass struct {
	name string
	fn   func(*Func)
}

// list of passes for the compiler
var passes = [...]pass{
	{"phielim", phielim},
	{"copyelim", copyelim},
	{"opt", opt},
	{"generic cse", cse},
	{"generic deadcode", deadcode},
	{"fuse", fuse},
	{"lower", lower},
	{"lowered cse", cse},
	{"lowered deadcode", deadcode},
	{"critical", critical}, // remove critical edges
	{"layout", layout},     // schedule blocks
	{"schedule", schedule}, // schedule values
	{"regalloc", regalloc},
	{"stackalloc", stackalloc},
	{"cgen", cgen},
}

// Double-check phase ordering constraints.
// This code is intended to document the ordering requirements
// between different phases.  It does not override the passes
// list above.
type constraint struct {
	a, b string // a must come before b
}

var passOrder = [...]constraint{
	// don't layout blocks until critical edges have been removed
	{"critical", "layout"},
	// regalloc requires the removal of all critical edges
	{"critical", "regalloc"},
	// regalloc requires all the values in a block to be scheduled
	{"schedule", "regalloc"},
	// stack allocation requires register allocation
	{"regalloc", "stackalloc"},
	// code generation requires stack allocation
	{"stackalloc", "cgen"},
}

func init() {
	for _, c := range passOrder {
		a, b := c.a, c.b
		i := -1
		j := -1
		for k, p := range passes {
			if p.name == a {
				i = k
			}
			if p.name == b {
				j = k
			}
		}
		if i < 0 {
			log.Panicf("pass %s not found", a)
		}
		if j < 0 {
			log.Panicf("pass %s not found", b)
		}
		if i >= j {
			log.Panicf("passes %s and %s out of order", a, b)
		}
	}
}
