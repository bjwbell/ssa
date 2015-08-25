// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// checkFunc checks invariants of f.
func checkFunc(f *Func) {
	blockMark := make([]bool, f.NumBlocks())
	valueMark := make([]bool, f.NumValues())

	for _, b := range f.Blocks {
		if blockMark[b.ID] {
			f.Fatalf("block %s appears twice in %s!", b, f.Name)
		}
		blockMark[b.ID] = true
		if b.Func != f {
			f.Fatalf("%s.Func=%s, want %s", b, b.Func.Name, f.Name)
		}

		for i, c := range b.Succs {
			for j, d := range b.Succs {
				if i != j && c == d {
					f.Fatalf("%s.Succs has duplicate block %s", b, c)
				}
			}
		}
		// Note: duplicate successors are hard in the following case:
		//      if(...) goto x else goto x
		//   x: v = phi(a, b)
		// If the conditional is true, does v get the value of a or b?
		// We could solve this other ways, but the easiest is just to
		// require (by possibly adding empty control-flow blocks) that
		// all successors are distinct.  They will need to be distinct
		// anyway for register allocation (duplicate successors implies
		// the existence of critical edges).

		for _, p := range b.Preds {
			var found bool
			for _, c := range p.Succs {
				if c == b {
					found = true
					break
				}
			}
			if !found {
				f.Fatalf("block %s is not a succ of its pred block %s", b, p)
			}
		}

		switch b.Kind {
		case BlockExit:
			if len(b.Succs) != 0 {
				f.Fatalf("exit block %s has successors", b)
			}
			if b.Control == nil {
				f.Fatalf("exit block %s has no control value", b)
			}
			if !b.Control.Type.IsMemory() {
				f.Fatalf("exit block %s has non-memory control value %s", b, b.Control.LongString())
			}
		case BlockDead:
			if len(b.Succs) != 0 {
				f.Fatalf("dead block %s has successors", b)
			}
			if len(b.Preds) != 0 {
				f.Fatalf("dead block %s has predecessors", b)
			}
			if len(b.Values) != 0 {
				f.Fatalf("dead block %s has values", b)
			}
			if b.Control != nil {
				f.Fatalf("dead block %s has a control value", b)
			}
		case BlockPlain:
			if len(b.Succs) != 1 {
				f.Fatalf("plain block %s len(Succs)==%d, want 1", b, len(b.Succs))
			}
			if b.Control != nil {
				f.Fatalf("plain block %s has non-nil control %s", b, b.Control.LongString())
			}
		case BlockIf:
			if len(b.Succs) != 2 {
				f.Fatalf("if block %s len(Succs)==%d, want 2", b, len(b.Succs))
			}
			if b.Control == nil {
				f.Fatalf("if block %s has no control value", b)
			}
			if !b.Control.Type.IsBoolean() {
				f.Fatalf("if block %s has non-bool control value %s", b, b.Control.LongString())
			}
		case BlockCall:
			if len(b.Succs) != 2 {
				f.Fatalf("call block %s len(Succs)==%d, want 2", b, len(b.Succs))
			}
			if b.Control == nil {
				f.Fatalf("call block %s has no control value", b)
			}
			if !b.Control.Type.IsMemory() {
				f.Fatalf("call block %s has non-memory control value %s", b, b.Control.LongString())
			}
		}
		if len(b.Succs) > 2 && b.Likely != BranchUnknown {
			f.Fatalf("likeliness prediction %d for block %s with %d successors: %s", b.Likely, b, len(b.Succs))
		}

		for _, v := range b.Values {
			for _, arg := range v.Args {
				if arg == nil {
					f.Fatalf("value %v has nil arg", v.LongString())
				}
			}

			if valueMark[v.ID] {
				f.Fatalf("value %s appears twice!", v.LongString())
			}
			valueMark[v.ID] = true

			if v.Block != b {
				f.Fatalf("%s.block != %s", v, b)
			}
			if v.Op == OpPhi && len(v.Args) != len(b.Preds) {
				f.Fatalf("phi length %s does not match pred length %d for block %s", v.LongString(), len(b.Preds), b)
			}

			if v.Op == OpAddr {
				if len(v.Args) == 0 {
					f.Fatalf("no args for OpAddr %s", v.LongString())
				}
				if v.Args[0].Op != OpSP && v.Args[0].Op != OpSB {
					f.Fatalf("bad arg to OpAddr %v", v)
				}
			}

			// TODO: check for cycles in values
			// TODO: check type
		}
	}

	// Check to make sure all Blocks referenced are in the function.
	if !blockMark[f.Entry.ID] {
		f.Fatalf("entry block %v is missing", f.Entry)
	}
	for _, b := range f.Blocks {
		for _, c := range b.Preds {
			if !blockMark[c.ID] {
				f.Fatalf("predecessor block %v for %v is missing", c, b)
			}
		}
		for _, c := range b.Succs {
			if !blockMark[c.ID] {
				f.Fatalf("successor block %v for %v is missing", c, b)
			}
		}
	}

	if len(f.Entry.Preds) > 0 {
		f.Fatalf("entry block %s of %s has predecessor(s) %v", f.Entry, f.Name, f.Entry.Preds)
	}

	// Check to make sure all Values referenced are in the function.
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			for i, a := range v.Args {
				if !valueMark[a.ID] {
					f.Fatalf("%v, arg %d of %v, is missing", a, i, v)
				}
			}
		}
		if b.Control != nil && !valueMark[b.Control.ID] {
			f.Fatalf("control value for %s is missing: %v", b, b.Control)
		}
	}
	for _, id := range f.bid.free {
		if blockMark[id] {
			f.Fatalf("used block b%d in free list", id)
		}
	}
	for _, id := range f.vid.free {
		if valueMark[id] {
			f.Fatalf("used value v%d in free list", id)
		}
	}
}
