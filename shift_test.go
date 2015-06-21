// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"testing"
)

func TestShiftConstAMD64(t *testing.T) {
	c := NewConfig("amd64", DummyFrontend{t})
	fun := makeConstShiftFunc(c, 18, OpLsh, TypeUInt64)
	checkOpcodeCounts(t, fun.f, map[Op]int{OpAMD64SHLQconst: 1, OpAMD64CMPQconst: 0, OpAMD64ANDQconst: 0})
	fun = makeConstShiftFunc(c, 66, OpLsh, TypeUInt64)
	checkOpcodeCounts(t, fun.f, map[Op]int{OpAMD64SHLQconst: 0, OpAMD64CMPQconst: 0, OpAMD64ANDQconst: 0})
	fun = makeConstShiftFunc(c, 18, OpRsh, TypeUInt64)
	checkOpcodeCounts(t, fun.f, map[Op]int{OpAMD64SHRQconst: 1, OpAMD64CMPQconst: 0, OpAMD64ANDQconst: 0})
	fun = makeConstShiftFunc(c, 66, OpRsh, TypeUInt64)
	checkOpcodeCounts(t, fun.f, map[Op]int{OpAMD64SHRQconst: 0, OpAMD64CMPQconst: 0, OpAMD64ANDQconst: 0})
	fun = makeConstShiftFunc(c, 18, OpRsh, TypeInt64)
	checkOpcodeCounts(t, fun.f, map[Op]int{OpAMD64SARQconst: 1, OpAMD64CMPQconst: 0})
	fun = makeConstShiftFunc(c, 66, OpRsh, TypeInt64)
	checkOpcodeCounts(t, fun.f, map[Op]int{OpAMD64SARQconst: 1, OpAMD64CMPQconst: 0})
}

func makeConstShiftFunc(c *Config, amount int64, op Op, typ Type) fun {
	ptyp := &TypeImpl{Size_: 8, Ptr: true, Name: "ptr"}
	fun := Fun(c, "entry",
		Bloc("entry",
			Valu("mem", OpArg, TypeMem, 0, ".mem"),
			Valu("FP", OpFP, TypeUInt64, 0, nil),
			Valu("argptr", OpOffPtr, ptyp, 8, nil, "FP"),
			Valu("resptr", OpOffPtr, ptyp, 16, nil, "FP"),
			Valu("load", OpLoad, typ, 0, nil, "argptr", "mem"),
			Valu("c", OpConst, TypeUInt64, amount, nil),
			Valu("shift", op, typ, 0, nil, "load", "c"),
			Valu("store", OpStore, TypeMem, 0, nil, "resptr", "shift", "mem"),
			Exit("store")))
	Compile(fun.f)
	return fun
}
