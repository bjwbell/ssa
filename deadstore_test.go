// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"testing"
)

func TestDeadStore(t *testing.T) {
	c := NewConfig("amd64", DummyFrontend{t})
	ptrType := &TypeImpl{Size_: 8, Ptr: true, Name: "testptr"} // dummy for testing
	fun := Fun(c, "entry",
		Bloc("entry",
			Valu("start", OpArg, TypeMem, 0, ".mem"),
			Valu("v", OpConst, TypeBool, 0, true),
			Valu("addr1", OpAddr, ptrType, 0, nil),
			Valu("addr2", OpAddr, ptrType, 0, nil),
			Valu("store1", OpStore, TypeMem, 0, nil, "addr1", "v", "start"),
			Valu("store2", OpStore, TypeMem, 0, nil, "addr2", "v", "store1"),
			Valu("store3", OpStore, TypeMem, 0, nil, "addr1", "v", "store2"),
			Goto("exit")),
		Bloc("exit",
			Exit("store3")))

	CheckFunc(fun.f)
	dse(fun.f)
	CheckFunc(fun.f)

	v := fun.values["store1"]
	if v.Op != OpCopy {
		t.Errorf("dead store not removed")
	}
}
func TestDeadStorePhi(t *testing.T) {
	// make sure we don't get into an infinite loop with phi values.
	c := NewConfig("amd64", DummyFrontend{t})
	ptrType := &TypeImpl{Size_: 8, Ptr: true, Name: "testptr"} // dummy for testing
	fun := Fun(c, "entry",
		Bloc("entry",
			Valu("start", OpArg, TypeMem, 0, ".mem"),
			Valu("v", OpConst, TypeBool, 0, true),
			Valu("addr", OpAddr, ptrType, 0, nil),
			Goto("loop")),
		Bloc("loop",
			Valu("phi", OpPhi, TypeMem, 0, nil, "start", "store"),
			Valu("store", OpStore, TypeMem, 0, nil, "addr", "v", "phi"),
			If("v", "loop", "exit")),
		Bloc("exit",
			Exit("store")))

	CheckFunc(fun.f)
	dse(fun.f)
	CheckFunc(fun.f)
}

func TestDeadStoreTypes(t *testing.T) {
	// Make sure a narrow store can't shadow a wider one.  We test an even
	// stronger restriction, that one store can't shadow another unless the
	// types of the address fields are identical (where identicalness is
	// decided by the CSE pass).
	c := NewConfig("amd64", DummyFrontend{t})
	t1 := &TypeImpl{Size_: 8, Ptr: true, Name: "t1"}
	t2 := &TypeImpl{Size_: 4, Ptr: true, Name: "t2"}
	fun := Fun(c, "entry",
		Bloc("entry",
			Valu("start", OpArg, TypeMem, 0, ".mem"),
			Valu("v", OpConst, TypeBool, 0, true),
			Valu("addr1", OpAddr, t1, 0, nil),
			Valu("addr2", OpAddr, t2, 0, nil),
			Valu("store1", OpStore, TypeMem, 0, nil, "addr1", "v", "start"),
			Valu("store2", OpStore, TypeMem, 0, nil, "addr2", "v", "store1"),
			Goto("exit")),
		Bloc("exit",
			Exit("store2")))

	CheckFunc(fun.f)
	cse(fun.f)
	dse(fun.f)
	CheckFunc(fun.f)

	v := fun.values["store1"]
	if v.Op == OpCopy {
		t.Errorf("store %s incorrectly removed", v)
	}
}
