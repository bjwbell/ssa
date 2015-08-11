// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var genericOps = []opData{
	// 2-input arithmetic
	// Types must be consistent with Go typing.  Add, for example, must take two values
	// of the same type and produces that same type.
	{name: "Add8"}, // arg0 + arg1
	{name: "Add16"},
	{name: "Add32"},
	{name: "Add64"},
	{name: "AddPtr"},
	// TODO: Add32F, Add64F, Add64C, Add128C

	{name: "Sub8"}, // arg0 - arg1
	{name: "Sub16"},
	{name: "Sub32"},
	{name: "Sub64"},
	// TODO: Sub32F, Sub64F, Sub64C, Sub128C

	{name: "Mul8"}, // arg0 * arg1
	{name: "Mul16"},
	{name: "Mul32"},
	{name: "Mul64"},
	{name: "MulPtr"}, // MulPtr is used for address calculations

	{name: "And8"}, // arg0 & arg1
	{name: "And16"},
	{name: "And32"},
	{name: "And64"},

	{name: "Or8"}, // arg0 | arg1
	{name: "Or16"},
	{name: "Or32"},
	{name: "Or64"},

	{name: "Xor8"}, // arg0 ^ arg1
	{name: "Xor16"},
	{name: "Xor32"},
	{name: "Xor64"},

	// For shifts, AxB means the shifted value has A bits and the shift amount has B bits.
	{name: "Lsh8x8"}, // arg0 << arg1
	{name: "Lsh8x16"},
	{name: "Lsh8x32"},
	{name: "Lsh8x64"},
	{name: "Lsh16x8"},
	{name: "Lsh16x16"},
	{name: "Lsh16x32"},
	{name: "Lsh16x64"},
	{name: "Lsh32x8"},
	{name: "Lsh32x16"},
	{name: "Lsh32x32"},
	{name: "Lsh32x64"},
	{name: "Lsh64x8"},
	{name: "Lsh64x16"},
	{name: "Lsh64x32"},
	{name: "Lsh64x64"},

	{name: "Rsh8x8"}, // arg0 >> arg1, signed
	{name: "Rsh8x16"},
	{name: "Rsh8x32"},
	{name: "Rsh8x64"},
	{name: "Rsh16x8"},
	{name: "Rsh16x16"},
	{name: "Rsh16x32"},
	{name: "Rsh16x64"},
	{name: "Rsh32x8"},
	{name: "Rsh32x16"},
	{name: "Rsh32x32"},
	{name: "Rsh32x64"},
	{name: "Rsh64x8"},
	{name: "Rsh64x16"},
	{name: "Rsh64x32"},
	{name: "Rsh64x64"},

	{name: "Rsh8Ux8"}, // arg0 >> arg1, unsigned
	{name: "Rsh8Ux16"},
	{name: "Rsh8Ux32"},
	{name: "Rsh8Ux64"},
	{name: "Rsh16Ux8"},
	{name: "Rsh16Ux16"},
	{name: "Rsh16Ux32"},
	{name: "Rsh16Ux64"},
	{name: "Rsh32Ux8"},
	{name: "Rsh32Ux16"},
	{name: "Rsh32Ux32"},
	{name: "Rsh32Ux64"},
	{name: "Rsh64Ux8"},
	{name: "Rsh64Ux16"},
	{name: "Rsh64Ux32"},
	{name: "Rsh64Ux64"},

	// (Left) rotates replace pattern matches in the front end
	// of (arg0 << arg1) ^ (arg0 >> (A-arg1))
	// where A is the bit width of arg0 and result.
	// Note that because rotates are pattern-matched from
	// shifts, that a rotate of arg1=A+k (k > 0) bits originated from
	//    (arg0 << A+k) ^ (arg0 >> -k) =
	//    0 ^ arg0>>huge_unsigned =
	//    0 ^ 0 = 0
	// which is not the same as a rotation by A+k
	//
	// However, in the specific case of k = 0, the result of
	// the shift idiom is the same as the result for the
	// rotate idiom, i.e., result=arg0.
	// This is different from shifts, where
	// arg0 << A is defined to be zero.
	//
	// Because of this, and also because the primary use case
	// for rotates is hashing and crypto code with constant
	// distance, rotate instructions are only substituted
	// when arg1 is a constant between 1 and A-1, inclusive.
	{name: "Lrot8"},
	{name: "Lrot16"},
	{name: "Lrot32"},
	{name: "Lrot64"},

	// 2-input comparisons
	{name: "Eq8"}, // arg0 == arg1
	{name: "Eq16"},
	{name: "Eq32"},
	{name: "Eq64"},
	{name: "EqPtr"},
	{name: "EqFat"}, // slice/interface; arg0 or arg1 is nil; other cases handled by frontend

	{name: "Neq8"}, // arg0 != arg1
	{name: "Neq16"},
	{name: "Neq32"},
	{name: "Neq64"},
	{name: "NeqPtr"},
	{name: "NeqFat"}, // slice/interface; arg0 or arg1 is nil; other cases handled by frontend

	{name: "Less8"}, // arg0 < arg1
	{name: "Less8U"},
	{name: "Less16"},
	{name: "Less16U"},
	{name: "Less32"},
	{name: "Less32U"},
	{name: "Less64"},
	{name: "Less64U"},

	{name: "Leq8"}, // arg0 <= arg1
	{name: "Leq8U"},
	{name: "Leq16"},
	{name: "Leq16U"},
	{name: "Leq32"},
	{name: "Leq32U"},
	{name: "Leq64"},
	{name: "Leq64U"},

	{name: "Greater8"}, // arg0 > arg1
	{name: "Greater8U"},
	{name: "Greater16"},
	{name: "Greater16U"},
	{name: "Greater32"},
	{name: "Greater32U"},
	{name: "Greater64"},
	{name: "Greater64U"},

	{name: "Geq8"}, // arg0 <= arg1
	{name: "Geq8U"},
	{name: "Geq16"},
	{name: "Geq16U"},
	{name: "Geq32"},
	{name: "Geq32U"},
	{name: "Geq64"},
	{name: "Geq64U"},

	// 1-input ops
	{name: "Not"}, // !arg0

	{name: "Neg8"}, // -arg0
	{name: "Neg16"},
	{name: "Neg32"},
	{name: "Neg64"},

	{name: "Com8"}, // ^arg0
	{name: "Com16"},
	{name: "Com32"},
	{name: "Com64"},

	// Data movement
	{name: "Phi"},  // select an argument based on which predecessor block we came from
	{name: "Copy"}, // output = arg0

	// constants.  Constant values are stored in the aux field.
	// booleans have a bool aux field, strings have a string aux
	// field, and so on.  All integer types store their value
	// in the AuxInt field as an int64 (including int, uint64, etc.).
	// For integer types smaller than 64 bits, only the low-order
	// bits of the AuxInt field matter.
	{name: "ConstBool"},
	{name: "ConstString"},
	{name: "ConstNil"},
	{name: "Const8"},
	{name: "Const16"},
	{name: "Const32"},
	{name: "Const64"},
	{name: "ConstPtr"}, // pointer-sized integer constant
	// TODO: Const32F, ...

	// Constant-like things
	{name: "Arg"}, // memory input to the function.

	// The address of a variable.  arg0 is the base pointer (SB or SP, depending
	// on whether it is a global or stack variable).  The Aux field identifies the
	// variable.  It will be either an *ExternSymbol (with arg0=SB), *ArgSymbol (arg0=SP),
	// or *AutoSymbol (arg0=SP).
	{name: "Addr"}, // Address of a variable.  Arg0=SP or SB.  Aux identifies the variable.

	{name: "SP"},   // stack pointer
	{name: "SB"},   // static base pointer (a.k.a. globals pointer)
	{name: "Func"}, // entry address of a function

	// Memory operations
	{name: "Load"},  // Load from arg0.  arg1=memory
	{name: "Store"}, // Store arg1 to arg0.  arg2=memory.  Returns memory.
	{name: "Move"},  // arg0=destptr, arg1=srcptr, arg2=mem, auxint=size.  Returns memory.
	{name: "Zero"},  // arg0=destptr, arg1=mem, auxint=size. Returns memory.

	// Function calls.  Arguments to the call have already been written to the stack.
	// Return values appear on the stack.  The method receiver, if any, is treated
	// as a phantom first argument.
	{name: "ClosureCall"}, // arg0=code pointer, arg1=context ptr, arg2=memory.  Returns memory.
	{name: "StaticCall"},  // call function aux.(*gc.Sym), arg0=memory.  Returns memory.

	// Conversions: signed extensions, zero (unsigned) extensions, truncations
	{name: "SignExt8to16"},
	{name: "SignExt8to32"},
	{name: "SignExt8to64"},
	{name: "SignExt16to32"},
	{name: "SignExt16to64"},
	{name: "SignExt32to64"},
	{name: "ZeroExt8to16"},
	{name: "ZeroExt8to32"},
	{name: "ZeroExt8to64"},
	{name: "ZeroExt16to32"},
	{name: "ZeroExt16to64"},
	{name: "ZeroExt32to64"},
	{name: "Trunc16to8"},
	{name: "Trunc32to8"},
	{name: "Trunc32to16"},
	{name: "Trunc64to8"},
	{name: "Trunc64to16"},
	{name: "Trunc64to32"},

	// Automatically inserted safety checks
	{name: "IsNonNil"},   // arg0 != nil
	{name: "IsInBounds"}, // 0 <= arg0 < arg1

	// Indexing operations
	{name: "ArrayIndex"},   // arg0=array, arg1=index.  Returns a[i]
	{name: "PtrIndex"},     // arg0=ptr, arg1=index. Computes ptr+sizeof(*v.type)*index, where index is extended to ptrwidth type
	{name: "OffPtr"},       // arg0 + auxint (arg0 and result are pointers)
	{name: "StructSelect"}, // arg0=struct, auxint=field offset.  Returns field at that offset (size=size of result type)

	// Slices
	{name: "SliceMake"}, // arg0=ptr, arg1=len, arg2=cap
	{name: "SlicePtr"},  // ptr(arg0)
	{name: "SliceLen"},  // len(arg0)
	{name: "SliceCap"},  // cap(arg0)

	// Strings
	{name: "StringMake"}, // arg0=ptr, arg1=len
	{name: "StringPtr"},  // ptr(arg0)
	{name: "StringLen"},  // len(arg0)

	// Interfaces
	{name: "ITab"}, // arg0=interface, returns itable field

	// Spill&restore ops for the register allocator.  These are
	// semantically identical to OpCopy; they do not take/return
	// stores like regular memory ops do.  We can get away without memory
	// args because we know there is no aliasing of spill slots on the stack.
	{name: "StoreReg"},
	{name: "LoadReg"},

	// Used during ssa construction.  Like Copy, but the arg has not been specified yet.
	{name: "FwdRef"},
}

//     kind           control    successors
//   ------------------------------------------
//     Exit        return mem                []
//    Plain               nil            [next]
//       If   a boolean Value      [then, else]
//     Call               mem  [nopanic, panic]  (control opcode should be OpCall or OpStaticCall)

var genericBlocks = []blockData{
	{name: "Exit"},  // no successors.  There should only be 1 of these.
	{name: "Dead"},  // no successors; determined to be dead but not yet removed
	{name: "Plain"}, // a single successor
	{name: "If"},    // 2 successors, if control goto Succs[0] else goto Succs[1]
	{name: "Call"},  // 2 successors, normal return and panic
	// TODO(khr): BlockPanic for the built-in panic call, has 1 edge to the exit block
}

func init() {
	archs = append(archs, arch{"generic", genericOps, genericBlocks, nil})
}
