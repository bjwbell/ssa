// Code generated from gen/MIPS.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "math"
import "cmd/internal/obj"

var _ = math.MinInt8 // in case not otherwise used
var _ = obj.ANOP     // in case not otherwise used
func rewriteValueMIPS(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValueMIPS_OpAdd16(v)
	case OpAdd32:
		return rewriteValueMIPS_OpAdd32(v)
	case OpAdd32F:
		return rewriteValueMIPS_OpAdd32F(v)
	case OpAdd32withcarry:
		return rewriteValueMIPS_OpAdd32withcarry(v)
	case OpAdd64F:
		return rewriteValueMIPS_OpAdd64F(v)
	case OpAdd8:
		return rewriteValueMIPS_OpAdd8(v)
	case OpAddPtr:
		return rewriteValueMIPS_OpAddPtr(v)
	case OpAddr:
		return rewriteValueMIPS_OpAddr(v)
	case OpAnd16:
		return rewriteValueMIPS_OpAnd16(v)
	case OpAnd32:
		return rewriteValueMIPS_OpAnd32(v)
	case OpAnd8:
		return rewriteValueMIPS_OpAnd8(v)
	case OpAndB:
		return rewriteValueMIPS_OpAndB(v)
	case OpAtomicAdd32:
		return rewriteValueMIPS_OpAtomicAdd32(v)
	case OpAtomicAnd8:
		return rewriteValueMIPS_OpAtomicAnd8(v)
	case OpAtomicCompareAndSwap32:
		return rewriteValueMIPS_OpAtomicCompareAndSwap32(v)
	case OpAtomicExchange32:
		return rewriteValueMIPS_OpAtomicExchange32(v)
	case OpAtomicLoad32:
		return rewriteValueMIPS_OpAtomicLoad32(v)
	case OpAtomicLoadPtr:
		return rewriteValueMIPS_OpAtomicLoadPtr(v)
	case OpAtomicOr8:
		return rewriteValueMIPS_OpAtomicOr8(v)
	case OpAtomicStore32:
		return rewriteValueMIPS_OpAtomicStore32(v)
	case OpAtomicStorePtrNoWB:
		return rewriteValueMIPS_OpAtomicStorePtrNoWB(v)
	case OpAvg32u:
		return rewriteValueMIPS_OpAvg32u(v)
	case OpBitLen32:
		return rewriteValueMIPS_OpBitLen32(v)
	case OpClosureCall:
		return rewriteValueMIPS_OpClosureCall(v)
	case OpCom16:
		return rewriteValueMIPS_OpCom16(v)
	case OpCom32:
		return rewriteValueMIPS_OpCom32(v)
	case OpCom8:
		return rewriteValueMIPS_OpCom8(v)
	case OpConst16:
		return rewriteValueMIPS_OpConst16(v)
	case OpConst32:
		return rewriteValueMIPS_OpConst32(v)
	case OpConst32F:
		return rewriteValueMIPS_OpConst32F(v)
	case OpConst64F:
		return rewriteValueMIPS_OpConst64F(v)
	case OpConst8:
		return rewriteValueMIPS_OpConst8(v)
	case OpConstBool:
		return rewriteValueMIPS_OpConstBool(v)
	case OpConstNil:
		return rewriteValueMIPS_OpConstNil(v)
	case OpConvert:
		return rewriteValueMIPS_OpConvert(v)
	case OpCtz32:
		return rewriteValueMIPS_OpCtz32(v)
	case OpCvt32Fto32:
		return rewriteValueMIPS_OpCvt32Fto32(v)
	case OpCvt32Fto64F:
		return rewriteValueMIPS_OpCvt32Fto64F(v)
	case OpCvt32to32F:
		return rewriteValueMIPS_OpCvt32to32F(v)
	case OpCvt32to64F:
		return rewriteValueMIPS_OpCvt32to64F(v)
	case OpCvt64Fto32:
		return rewriteValueMIPS_OpCvt64Fto32(v)
	case OpCvt64Fto32F:
		return rewriteValueMIPS_OpCvt64Fto32F(v)
	case OpDiv16:
		return rewriteValueMIPS_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueMIPS_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueMIPS_OpDiv32(v)
	case OpDiv32F:
		return rewriteValueMIPS_OpDiv32F(v)
	case OpDiv32u:
		return rewriteValueMIPS_OpDiv32u(v)
	case OpDiv64F:
		return rewriteValueMIPS_OpDiv64F(v)
	case OpDiv8:
		return rewriteValueMIPS_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueMIPS_OpDiv8u(v)
	case OpEq16:
		return rewriteValueMIPS_OpEq16(v)
	case OpEq32:
		return rewriteValueMIPS_OpEq32(v)
	case OpEq32F:
		return rewriteValueMIPS_OpEq32F(v)
	case OpEq64F:
		return rewriteValueMIPS_OpEq64F(v)
	case OpEq8:
		return rewriteValueMIPS_OpEq8(v)
	case OpEqB:
		return rewriteValueMIPS_OpEqB(v)
	case OpEqPtr:
		return rewriteValueMIPS_OpEqPtr(v)
	case OpGeq16:
		return rewriteValueMIPS_OpGeq16(v)
	case OpGeq16U:
		return rewriteValueMIPS_OpGeq16U(v)
	case OpGeq32:
		return rewriteValueMIPS_OpGeq32(v)
	case OpGeq32F:
		return rewriteValueMIPS_OpGeq32F(v)
	case OpGeq32U:
		return rewriteValueMIPS_OpGeq32U(v)
	case OpGeq64F:
		return rewriteValueMIPS_OpGeq64F(v)
	case OpGeq8:
		return rewriteValueMIPS_OpGeq8(v)
	case OpGeq8U:
		return rewriteValueMIPS_OpGeq8U(v)
	case OpGetClosurePtr:
		return rewriteValueMIPS_OpGetClosurePtr(v)
	case OpGreater16:
		return rewriteValueMIPS_OpGreater16(v)
	case OpGreater16U:
		return rewriteValueMIPS_OpGreater16U(v)
	case OpGreater32:
		return rewriteValueMIPS_OpGreater32(v)
	case OpGreater32F:
		return rewriteValueMIPS_OpGreater32F(v)
	case OpGreater32U:
		return rewriteValueMIPS_OpGreater32U(v)
	case OpGreater64F:
		return rewriteValueMIPS_OpGreater64F(v)
	case OpGreater8:
		return rewriteValueMIPS_OpGreater8(v)
	case OpGreater8U:
		return rewriteValueMIPS_OpGreater8U(v)
	case OpHmul32:
		return rewriteValueMIPS_OpHmul32(v)
	case OpHmul32u:
		return rewriteValueMIPS_OpHmul32u(v)
	case OpInterCall:
		return rewriteValueMIPS_OpInterCall(v)
	case OpIsInBounds:
		return rewriteValueMIPS_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValueMIPS_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValueMIPS_OpIsSliceInBounds(v)
	case OpLeq16:
		return rewriteValueMIPS_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueMIPS_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueMIPS_OpLeq32(v)
	case OpLeq32F:
		return rewriteValueMIPS_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValueMIPS_OpLeq32U(v)
	case OpLeq64F:
		return rewriteValueMIPS_OpLeq64F(v)
	case OpLeq8:
		return rewriteValueMIPS_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueMIPS_OpLeq8U(v)
	case OpLess16:
		return rewriteValueMIPS_OpLess16(v)
	case OpLess16U:
		return rewriteValueMIPS_OpLess16U(v)
	case OpLess32:
		return rewriteValueMIPS_OpLess32(v)
	case OpLess32F:
		return rewriteValueMIPS_OpLess32F(v)
	case OpLess32U:
		return rewriteValueMIPS_OpLess32U(v)
	case OpLess64F:
		return rewriteValueMIPS_OpLess64F(v)
	case OpLess8:
		return rewriteValueMIPS_OpLess8(v)
	case OpLess8U:
		return rewriteValueMIPS_OpLess8U(v)
	case OpLoad:
		return rewriteValueMIPS_OpLoad(v)
	case OpLsh16x16:
		return rewriteValueMIPS_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueMIPS_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValueMIPS_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValueMIPS_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueMIPS_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueMIPS_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValueMIPS_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValueMIPS_OpLsh32x8(v)
	case OpLsh8x16:
		return rewriteValueMIPS_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueMIPS_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValueMIPS_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValueMIPS_OpLsh8x8(v)
	case OpMIPSADD:
		return rewriteValueMIPS_OpMIPSADD(v)
	case OpMIPSADDconst:
		return rewriteValueMIPS_OpMIPSADDconst(v)
	case OpMIPSAND:
		return rewriteValueMIPS_OpMIPSAND(v)
	case OpMIPSANDconst:
		return rewriteValueMIPS_OpMIPSANDconst(v)
	case OpMIPSCMOVZ:
		return rewriteValueMIPS_OpMIPSCMOVZ(v)
	case OpMIPSCMOVZzero:
		return rewriteValueMIPS_OpMIPSCMOVZzero(v)
	case OpMIPSLoweredAtomicAdd:
		return rewriteValueMIPS_OpMIPSLoweredAtomicAdd(v)
	case OpMIPSLoweredAtomicStore:
		return rewriteValueMIPS_OpMIPSLoweredAtomicStore(v)
	case OpMIPSMOVBUload:
		return rewriteValueMIPS_OpMIPSMOVBUload(v)
	case OpMIPSMOVBUreg:
		return rewriteValueMIPS_OpMIPSMOVBUreg(v)
	case OpMIPSMOVBload:
		return rewriteValueMIPS_OpMIPSMOVBload(v)
	case OpMIPSMOVBreg:
		return rewriteValueMIPS_OpMIPSMOVBreg(v)
	case OpMIPSMOVBstore:
		return rewriteValueMIPS_OpMIPSMOVBstore(v)
	case OpMIPSMOVBstorezero:
		return rewriteValueMIPS_OpMIPSMOVBstorezero(v)
	case OpMIPSMOVDload:
		return rewriteValueMIPS_OpMIPSMOVDload(v)
	case OpMIPSMOVDstore:
		return rewriteValueMIPS_OpMIPSMOVDstore(v)
	case OpMIPSMOVFload:
		return rewriteValueMIPS_OpMIPSMOVFload(v)
	case OpMIPSMOVFstore:
		return rewriteValueMIPS_OpMIPSMOVFstore(v)
	case OpMIPSMOVHUload:
		return rewriteValueMIPS_OpMIPSMOVHUload(v)
	case OpMIPSMOVHUreg:
		return rewriteValueMIPS_OpMIPSMOVHUreg(v)
	case OpMIPSMOVHload:
		return rewriteValueMIPS_OpMIPSMOVHload(v)
	case OpMIPSMOVHreg:
		return rewriteValueMIPS_OpMIPSMOVHreg(v)
	case OpMIPSMOVHstore:
		return rewriteValueMIPS_OpMIPSMOVHstore(v)
	case OpMIPSMOVHstorezero:
		return rewriteValueMIPS_OpMIPSMOVHstorezero(v)
	case OpMIPSMOVWload:
		return rewriteValueMIPS_OpMIPSMOVWload(v)
	case OpMIPSMOVWreg:
		return rewriteValueMIPS_OpMIPSMOVWreg(v)
	case OpMIPSMOVWstore:
		return rewriteValueMIPS_OpMIPSMOVWstore(v)
	case OpMIPSMOVWstorezero:
		return rewriteValueMIPS_OpMIPSMOVWstorezero(v)
	case OpMIPSMUL:
		return rewriteValueMIPS_OpMIPSMUL(v)
	case OpMIPSNEG:
		return rewriteValueMIPS_OpMIPSNEG(v)
	case OpMIPSNOR:
		return rewriteValueMIPS_OpMIPSNOR(v)
	case OpMIPSNORconst:
		return rewriteValueMIPS_OpMIPSNORconst(v)
	case OpMIPSOR:
		return rewriteValueMIPS_OpMIPSOR(v)
	case OpMIPSORconst:
		return rewriteValueMIPS_OpMIPSORconst(v)
	case OpMIPSSGT:
		return rewriteValueMIPS_OpMIPSSGT(v)
	case OpMIPSSGTU:
		return rewriteValueMIPS_OpMIPSSGTU(v)
	case OpMIPSSGTUconst:
		return rewriteValueMIPS_OpMIPSSGTUconst(v)
	case OpMIPSSGTUzero:
		return rewriteValueMIPS_OpMIPSSGTUzero(v)
	case OpMIPSSGTconst:
		return rewriteValueMIPS_OpMIPSSGTconst(v)
	case OpMIPSSGTzero:
		return rewriteValueMIPS_OpMIPSSGTzero(v)
	case OpMIPSSLL:
		return rewriteValueMIPS_OpMIPSSLL(v)
	case OpMIPSSLLconst:
		return rewriteValueMIPS_OpMIPSSLLconst(v)
	case OpMIPSSRA:
		return rewriteValueMIPS_OpMIPSSRA(v)
	case OpMIPSSRAconst:
		return rewriteValueMIPS_OpMIPSSRAconst(v)
	case OpMIPSSRL:
		return rewriteValueMIPS_OpMIPSSRL(v)
	case OpMIPSSRLconst:
		return rewriteValueMIPS_OpMIPSSRLconst(v)
	case OpMIPSSUB:
		return rewriteValueMIPS_OpMIPSSUB(v)
	case OpMIPSSUBconst:
		return rewriteValueMIPS_OpMIPSSUBconst(v)
	case OpMIPSXOR:
		return rewriteValueMIPS_OpMIPSXOR(v)
	case OpMIPSXORconst:
		return rewriteValueMIPS_OpMIPSXORconst(v)
	case OpMod16:
		return rewriteValueMIPS_OpMod16(v)
	case OpMod16u:
		return rewriteValueMIPS_OpMod16u(v)
	case OpMod32:
		return rewriteValueMIPS_OpMod32(v)
	case OpMod32u:
		return rewriteValueMIPS_OpMod32u(v)
	case OpMod8:
		return rewriteValueMIPS_OpMod8(v)
	case OpMod8u:
		return rewriteValueMIPS_OpMod8u(v)
	case OpMove:
		return rewriteValueMIPS_OpMove(v)
	case OpMul16:
		return rewriteValueMIPS_OpMul16(v)
	case OpMul32:
		return rewriteValueMIPS_OpMul32(v)
	case OpMul32F:
		return rewriteValueMIPS_OpMul32F(v)
	case OpMul32uhilo:
		return rewriteValueMIPS_OpMul32uhilo(v)
	case OpMul64F:
		return rewriteValueMIPS_OpMul64F(v)
	case OpMul8:
		return rewriteValueMIPS_OpMul8(v)
	case OpNeg16:
		return rewriteValueMIPS_OpNeg16(v)
	case OpNeg32:
		return rewriteValueMIPS_OpNeg32(v)
	case OpNeg32F:
		return rewriteValueMIPS_OpNeg32F(v)
	case OpNeg64F:
		return rewriteValueMIPS_OpNeg64F(v)
	case OpNeg8:
		return rewriteValueMIPS_OpNeg8(v)
	case OpNeq16:
		return rewriteValueMIPS_OpNeq16(v)
	case OpNeq32:
		return rewriteValueMIPS_OpNeq32(v)
	case OpNeq32F:
		return rewriteValueMIPS_OpNeq32F(v)
	case OpNeq64F:
		return rewriteValueMIPS_OpNeq64F(v)
	case OpNeq8:
		return rewriteValueMIPS_OpNeq8(v)
	case OpNeqB:
		return rewriteValueMIPS_OpNeqB(v)
	case OpNeqPtr:
		return rewriteValueMIPS_OpNeqPtr(v)
	case OpNilCheck:
		return rewriteValueMIPS_OpNilCheck(v)
	case OpNot:
		return rewriteValueMIPS_OpNot(v)
	case OpOffPtr:
		return rewriteValueMIPS_OpOffPtr(v)
	case OpOr16:
		return rewriteValueMIPS_OpOr16(v)
	case OpOr32:
		return rewriteValueMIPS_OpOr32(v)
	case OpOr8:
		return rewriteValueMIPS_OpOr8(v)
	case OpOrB:
		return rewriteValueMIPS_OpOrB(v)
	case OpRound32F:
		return rewriteValueMIPS_OpRound32F(v)
	case OpRound64F:
		return rewriteValueMIPS_OpRound64F(v)
	case OpRsh16Ux16:
		return rewriteValueMIPS_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueMIPS_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueMIPS_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueMIPS_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueMIPS_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueMIPS_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueMIPS_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueMIPS_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueMIPS_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueMIPS_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueMIPS_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueMIPS_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueMIPS_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueMIPS_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueMIPS_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueMIPS_OpRsh32x8(v)
	case OpRsh8Ux16:
		return rewriteValueMIPS_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueMIPS_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueMIPS_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueMIPS_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueMIPS_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueMIPS_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueMIPS_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueMIPS_OpRsh8x8(v)
	case OpSelect0:
		return rewriteValueMIPS_OpSelect0(v)
	case OpSelect1:
		return rewriteValueMIPS_OpSelect1(v)
	case OpSignExt16to32:
		return rewriteValueMIPS_OpSignExt16to32(v)
	case OpSignExt8to16:
		return rewriteValueMIPS_OpSignExt8to16(v)
	case OpSignExt8to32:
		return rewriteValueMIPS_OpSignExt8to32(v)
	case OpSignmask:
		return rewriteValueMIPS_OpSignmask(v)
	case OpSlicemask:
		return rewriteValueMIPS_OpSlicemask(v)
	case OpSqrt:
		return rewriteValueMIPS_OpSqrt(v)
	case OpStaticCall:
		return rewriteValueMIPS_OpStaticCall(v)
	case OpStore:
		return rewriteValueMIPS_OpStore(v)
	case OpSub16:
		return rewriteValueMIPS_OpSub16(v)
	case OpSub32:
		return rewriteValueMIPS_OpSub32(v)
	case OpSub32F:
		return rewriteValueMIPS_OpSub32F(v)
	case OpSub32withcarry:
		return rewriteValueMIPS_OpSub32withcarry(v)
	case OpSub64F:
		return rewriteValueMIPS_OpSub64F(v)
	case OpSub8:
		return rewriteValueMIPS_OpSub8(v)
	case OpSubPtr:
		return rewriteValueMIPS_OpSubPtr(v)
	case OpTrunc16to8:
		return rewriteValueMIPS_OpTrunc16to8(v)
	case OpTrunc32to16:
		return rewriteValueMIPS_OpTrunc32to16(v)
	case OpTrunc32to8:
		return rewriteValueMIPS_OpTrunc32to8(v)
	case OpXor16:
		return rewriteValueMIPS_OpXor16(v)
	case OpXor32:
		return rewriteValueMIPS_OpXor32(v)
	case OpXor8:
		return rewriteValueMIPS_OpXor8(v)
	case OpZero:
		return rewriteValueMIPS_OpZero(v)
	case OpZeroExt16to32:
		return rewriteValueMIPS_OpZeroExt16to32(v)
	case OpZeroExt8to16:
		return rewriteValueMIPS_OpZeroExt8to16(v)
	case OpZeroExt8to32:
		return rewriteValueMIPS_OpZeroExt8to32(v)
	case OpZeromask:
		return rewriteValueMIPS_OpZeromask(v)
	}
	return false
}
func rewriteValueMIPS_OpAdd16(v *Value) bool {
	// match: (Add16 x y)
	// cond:
	// result: (ADD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAdd32(v *Value) bool {
	// match: (Add32 x y)
	// cond:
	// result: (ADD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAdd32F(v *Value) bool {
	// match: (Add32F x y)
	// cond:
	// result: (ADDF x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADDF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAdd32withcarry(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Add32withcarry <t> x y c)
	// cond:
	// result: (ADD c (ADD <t> x y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		c := v.Args[2]
		v.reset(OpMIPSADD)
		v.AddArg(c)
		v0 := b.NewValue0(v.Pos, OpMIPSADD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpAdd64F(v *Value) bool {
	// match: (Add64F x y)
	// cond:
	// result: (ADDD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADDD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAdd8(v *Value) bool {
	// match: (Add8 x y)
	// cond:
	// result: (ADD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAddPtr(v *Value) bool {
	// match: (AddPtr x y)
	// cond:
	// result: (ADD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAddr(v *Value) bool {
	// match: (Addr {sym} base)
	// cond:
	// result: (MOVWaddr {sym} base)
	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpMIPSMOVWaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueMIPS_OpAnd16(v *Value) bool {
	// match: (And16 x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAnd32(v *Value) bool {
	// match: (And32 x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAnd8(v *Value) bool {
	// match: (And8 x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAndB(v *Value) bool {
	// match: (AndB x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAtomicAdd32(v *Value) bool {
	// match: (AtomicAdd32 ptr val mem)
	// cond:
	// result: (LoweredAtomicAdd ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSLoweredAtomicAdd)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicAnd8(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	types := &b.Func.Config.Types
	_ = types
	// match: (AtomicAnd8 ptr val mem)
	// cond: !config.BigEndian
	// result: (LoweredAtomicAnd (AND <types.UInt32Ptr> (MOVWconst [^3]) ptr) 		(OR <types.UInt32> (SLL <types.UInt32> (ZeroExt8to32 val) 			(SLLconst <types.UInt32> [3] 				(ANDconst  <types.UInt32> [3] ptr))) 		(NORconst [0] <types.UInt32> (SLL <types.UInt32> 			(MOVWconst [0xff]) (SLLconst <types.UInt32> [3] 				(ANDconst <types.UInt32> [3] ptr))))) mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(!config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicAnd)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, types.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = ^3
		v0.AddArg(v1)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSOR, types.UInt32)
		v3 := b.NewValue0(v.Pos, OpMIPSSLL, types.UInt32)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v4.AddArg(val)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v5.AuxInt = 3
		v6 := b.NewValue0(v.Pos, OpMIPSANDconst, types.UInt32)
		v6.AuxInt = 3
		v6.AddArg(ptr)
		v5.AddArg(v6)
		v3.AddArg(v5)
		v2.AddArg(v3)
		v7 := b.NewValue0(v.Pos, OpMIPSNORconst, types.UInt32)
		v7.AuxInt = 0
		v8 := b.NewValue0(v.Pos, OpMIPSSLL, types.UInt32)
		v9 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v9.AuxInt = 0xff
		v8.AddArg(v9)
		v10 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v10.AuxInt = 3
		v11 := b.NewValue0(v.Pos, OpMIPSANDconst, types.UInt32)
		v11.AuxInt = 3
		v11.AddArg(ptr)
		v10.AddArg(v11)
		v8.AddArg(v10)
		v7.AddArg(v8)
		v2.AddArg(v7)
		v.AddArg(v2)
		v.AddArg(mem)
		return true
	}
	// match: (AtomicAnd8 ptr val mem)
	// cond: config.BigEndian
	// result: (LoweredAtomicAnd (AND <types.UInt32Ptr> (MOVWconst [^3]) ptr) 		(OR <types.UInt32> (SLL <types.UInt32> (ZeroExt8to32 val) 			(SLLconst <types.UInt32> [3] 				(ANDconst  <types.UInt32> [3] 					(XORconst <types.UInt32> [3] ptr)))) 		(NORconst [0] <types.UInt32> (SLL <types.UInt32> 			(MOVWconst [0xff]) (SLLconst <types.UInt32> [3] 				(ANDconst <types.UInt32> [3] 					(XORconst <types.UInt32> [3] ptr)))))) mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicAnd)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, types.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = ^3
		v0.AddArg(v1)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSOR, types.UInt32)
		v3 := b.NewValue0(v.Pos, OpMIPSSLL, types.UInt32)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v4.AddArg(val)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v5.AuxInt = 3
		v6 := b.NewValue0(v.Pos, OpMIPSANDconst, types.UInt32)
		v6.AuxInt = 3
		v7 := b.NewValue0(v.Pos, OpMIPSXORconst, types.UInt32)
		v7.AuxInt = 3
		v7.AddArg(ptr)
		v6.AddArg(v7)
		v5.AddArg(v6)
		v3.AddArg(v5)
		v2.AddArg(v3)
		v8 := b.NewValue0(v.Pos, OpMIPSNORconst, types.UInt32)
		v8.AuxInt = 0
		v9 := b.NewValue0(v.Pos, OpMIPSSLL, types.UInt32)
		v10 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v10.AuxInt = 0xff
		v9.AddArg(v10)
		v11 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v11.AuxInt = 3
		v12 := b.NewValue0(v.Pos, OpMIPSANDconst, types.UInt32)
		v12.AuxInt = 3
		v13 := b.NewValue0(v.Pos, OpMIPSXORconst, types.UInt32)
		v13.AuxInt = 3
		v13.AddArg(ptr)
		v12.AddArg(v13)
		v11.AddArg(v12)
		v9.AddArg(v11)
		v8.AddArg(v9)
		v2.AddArg(v8)
		v.AddArg(v2)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpAtomicCompareAndSwap32(v *Value) bool {
	// match: (AtomicCompareAndSwap32 ptr old new_ mem)
	// cond:
	// result: (LoweredAtomicCas ptr old new_ mem)
	for {
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpMIPSLoweredAtomicCas)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicExchange32(v *Value) bool {
	// match: (AtomicExchange32 ptr val mem)
	// cond:
	// result: (LoweredAtomicExchange ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSLoweredAtomicExchange)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicLoad32(v *Value) bool {
	// match: (AtomicLoad32 ptr mem)
	// cond:
	// result: (LoweredAtomicLoad ptr mem)
	for {
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSLoweredAtomicLoad)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicLoadPtr(v *Value) bool {
	// match: (AtomicLoadPtr ptr mem)
	// cond:
	// result: (LoweredAtomicLoad  ptr mem)
	for {
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSLoweredAtomicLoad)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicOr8(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	types := &b.Func.Config.Types
	_ = types
	// match: (AtomicOr8 ptr val mem)
	// cond: !config.BigEndian
	// result: (LoweredAtomicOr (AND <types.UInt32Ptr> (MOVWconst [^3]) ptr) 		(SLL <types.UInt32> (ZeroExt8to32 val) 			(SLLconst <types.UInt32> [3] 				(ANDconst <types.UInt32> [3] ptr))) mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(!config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicOr)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, types.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = ^3
		v0.AddArg(v1)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSSLL, types.UInt32)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v3.AddArg(val)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v4.AuxInt = 3
		v5 := b.NewValue0(v.Pos, OpMIPSANDconst, types.UInt32)
		v5.AuxInt = 3
		v5.AddArg(ptr)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v.AddArg(v2)
		v.AddArg(mem)
		return true
	}
	// match: (AtomicOr8 ptr val mem)
	// cond: config.BigEndian
	// result: (LoweredAtomicOr (AND <types.UInt32Ptr> (MOVWconst [^3]) ptr) 		(SLL <types.UInt32> (ZeroExt8to32 val) 			(SLLconst <types.UInt32> [3] 				(ANDconst <types.UInt32> [3] 					(XORconst <types.UInt32> [3] ptr)))) mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicOr)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, types.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = ^3
		v0.AddArg(v1)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSSLL, types.UInt32)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v3.AddArg(val)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v4.AuxInt = 3
		v5 := b.NewValue0(v.Pos, OpMIPSANDconst, types.UInt32)
		v5.AuxInt = 3
		v6 := b.NewValue0(v.Pos, OpMIPSXORconst, types.UInt32)
		v6.AuxInt = 3
		v6.AddArg(ptr)
		v5.AddArg(v6)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v.AddArg(v2)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpAtomicStore32(v *Value) bool {
	// match: (AtomicStore32 ptr val mem)
	// cond:
	// result: (LoweredAtomicStore ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSLoweredAtomicStore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicStorePtrNoWB(v *Value) bool {
	// match: (AtomicStorePtrNoWB ptr val mem)
	// cond:
	// result: (LoweredAtomicStore  ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSLoweredAtomicStore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAvg32u(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Avg32u <t> x y)
	// cond:
	// result: (ADD (SRLconst <t> (SUB <t> x y) [1]) y)
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADD)
		v0 := b.NewValue0(v.Pos, OpMIPSSRLconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpMIPSSUB, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpBitLen32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (BitLen32 <t> x)
	// cond:
	// result: (SUB (MOVWconst [32]) (CLZ <t> x))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpMIPSSUB)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 32
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCLZ, t)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpClosureCall(v *Value) bool {
	// match: (ClosureCall [argwid] entry closure mem)
	// cond:
	// result: (CALLclosure [argwid] entry closure mem)
	for {
		argwid := v.AuxInt
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSCALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpCom16(v *Value) bool {
	// match: (Com16 x)
	// cond:
	// result: (NORconst [0] x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCom32(v *Value) bool {
	// match: (Com32 x)
	// cond:
	// result: (NORconst [0] x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCom8(v *Value) bool {
	// match: (Com8 x)
	// cond:
	// result: (NORconst [0] x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpConst16(v *Value) bool {
	// match: (Const16 [val])
	// cond:
	// result: (MOVWconst [val])
	for {
		val := v.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS_OpConst32(v *Value) bool {
	// match: (Const32 [val])
	// cond:
	// result: (MOVWconst [val])
	for {
		val := v.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS_OpConst32F(v *Value) bool {
	// match: (Const32F [val])
	// cond:
	// result: (MOVFconst [val])
	for {
		val := v.AuxInt
		v.reset(OpMIPSMOVFconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS_OpConst64F(v *Value) bool {
	// match: (Const64F [val])
	// cond:
	// result: (MOVDconst [val])
	for {
		val := v.AuxInt
		v.reset(OpMIPSMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS_OpConst8(v *Value) bool {
	// match: (Const8 [val])
	// cond:
	// result: (MOVWconst [val])
	for {
		val := v.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS_OpConstBool(v *Value) bool {
	// match: (ConstBool [b])
	// cond:
	// result: (MOVWconst [b])
	for {
		b := v.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValueMIPS_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// cond:
	// result: (MOVWconst [0])
	for {
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueMIPS_OpConvert(v *Value) bool {
	// match: (Convert x mem)
	// cond:
	// result: (MOVWconvert x mem)
	for {
		x := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSMOVWconvert)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpCtz32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Ctz32 <t> x)
	// cond:
	// result: (SUB (MOVWconst [32]) (CLZ <t> (SUBconst <t> [1] (AND <t> x (NEG <t> x)))))
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpMIPSSUB)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 32
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCLZ, t)
		v2 := b.NewValue0(v.Pos, OpMIPSSUBconst, t)
		v2.AuxInt = 1
		v3 := b.NewValue0(v.Pos, OpMIPSAND, t)
		v3.AddArg(x)
		v4 := b.NewValue0(v.Pos, OpMIPSNEG, t)
		v4.AddArg(x)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpCvt32Fto32(v *Value) bool {
	// match: (Cvt32Fto32 x)
	// cond:
	// result: (TRUNCFW x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSTRUNCFW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCvt32Fto64F(v *Value) bool {
	// match: (Cvt32Fto64F x)
	// cond:
	// result: (MOVFD x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVFD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCvt32to32F(v *Value) bool {
	// match: (Cvt32to32F x)
	// cond:
	// result: (MOVWF x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVWF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCvt32to64F(v *Value) bool {
	// match: (Cvt32to64F x)
	// cond:
	// result: (MOVWD x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVWD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCvt64Fto32(v *Value) bool {
	// match: (Cvt64Fto32 x)
	// cond:
	// result: (TRUNCDW x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSTRUNCDW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCvt64Fto32F(v *Value) bool {
	// match: (Cvt64Fto32F x)
	// cond:
	// result: (MOVDF x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVDF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpDiv16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div16 x y)
	// cond:
	// result: (Select1 (DIV (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, MakeTuple(types.Int32, types.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv16u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div16u x y)
	// cond:
	// result: (Select1 (DIVU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, MakeTuple(types.UInt32, types.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div32 x y)
	// cond:
	// result: (Select1 (DIV x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, MakeTuple(types.Int32, types.Int32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv32F(v *Value) bool {
	// match: (Div32F x y)
	// cond:
	// result: (DIVF x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSDIVF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpDiv32u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div32u x y)
	// cond:
	// result: (Select1 (DIVU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, MakeTuple(types.UInt32, types.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv64F(v *Value) bool {
	// match: (Div64F x y)
	// cond:
	// result: (DIVD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSDIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpDiv8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div8 x y)
	// cond:
	// result: (Select1 (DIV (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, MakeTuple(types.Int32, types.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv8u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div8u x y)
	// cond:
	// result: (Select1 (DIVU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, MakeTuple(types.UInt32, types.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Eq16 x y)
	// cond:
	// result: (SGTUconst [1] (XOR (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, types.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEq32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Eq32 x y)
	// cond:
	// result: (SGTUconst [1] (XOR x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, types.UInt32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Eq32F x y)
	// cond:
	// result: (FPFlagTrue (CMPEQF x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQF, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Eq64F x y)
	// cond:
	// result: (FPFlagTrue (CMPEQD x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQD, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Eq8 x y)
	// cond:
	// result: (SGTUconst [1] (XOR (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, types.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEqB(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (EqB x y)
	// cond:
	// result: (XORconst [1] (XOR <types.Bool> x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, types.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEqPtr(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (EqPtr x y)
	// cond:
	// result: (SGTUconst [1] (XOR x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, types.UInt32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq16 x y)
	// cond:
	// result: (XORconst [1] (SGT (SignExt16to32 y) (SignExt16to32 x)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, types.Bool)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v2.AddArg(x)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq16U x y)
	// cond:
	// result: (XORconst [1] (SGTU (ZeroExt16to32 y) (ZeroExt16to32 x)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, types.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v2.AddArg(x)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq32 x y)
	// cond:
	// result: (XORconst [1] (SGT y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, types.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Geq32F x y)
	// cond:
	// result: (FPFlagTrue (CMPGEF x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGEF, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq32U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq32U x y)
	// cond:
	// result: (XORconst [1] (SGTU y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, types.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Geq64F x y)
	// cond:
	// result: (FPFlagTrue (CMPGED x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGED, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq8 x y)
	// cond:
	// result: (XORconst [1] (SGT (SignExt8to32 y) (SignExt8to32 x)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, types.Bool)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v2.AddArg(x)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq8U x y)
	// cond:
	// result: (XORconst [1] (SGTU (ZeroExt8to32 y) (ZeroExt8to32 x)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, types.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v2.AddArg(x)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGetClosurePtr(v *Value) bool {
	// match: (GetClosurePtr)
	// cond:
	// result: (LoweredGetClosurePtr)
	for {
		v.reset(OpMIPSLoweredGetClosurePtr)
		return true
	}
}
func rewriteValueMIPS_OpGreater16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater16 x y)
	// cond:
	// result: (SGT (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpGreater16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater16U x y)
	// cond:
	// result: (SGTU (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpGreater32(v *Value) bool {
	// match: (Greater32 x y)
	// cond:
	// result: (SGT x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpGreater32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Greater32F x y)
	// cond:
	// result: (FPFlagTrue (CMPGTF x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTF, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGreater32U(v *Value) bool {
	// match: (Greater32U x y)
	// cond:
	// result: (SGTU x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpGreater64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Greater64F x y)
	// cond:
	// result: (FPFlagTrue (CMPGTD x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTD, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGreater8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater8 x y)
	// cond:
	// result: (SGT (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpGreater8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater8U x y)
	// cond:
	// result: (SGTU (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpHmul32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Hmul32 x y)
	// cond:
	// result: (Select0 (MULT x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSMULT, MakeTuple(types.Int32, types.Int32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpHmul32u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Hmul32u x y)
	// cond:
	// result: (Select0 (MULTU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSMULTU, MakeTuple(types.UInt32, types.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpInterCall(v *Value) bool {
	// match: (InterCall [argwid] entry mem)
	// cond:
	// result: (CALLinter [argwid] entry mem)
	for {
		argwid := v.AuxInt
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSCALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpIsInBounds(v *Value) bool {
	// match: (IsInBounds idx len)
	// cond:
	// result: (SGTU len idx)
	for {
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpMIPSSGTU)
		v.AddArg(len)
		v.AddArg(idx)
		return true
	}
}
func rewriteValueMIPS_OpIsNonNil(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (IsNonNil ptr)
	// cond:
	// result: (SGTU ptr (MOVWconst [0]))
	for {
		ptr := v.Args[0]
		v.reset(OpMIPSSGTU)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpIsSliceInBounds(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (IsSliceInBounds idx len)
	// cond:
	// result: (XORconst [1] (SGTU idx len))
	for {
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, types.Bool)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq16 x y)
	// cond:
	// result: (XORconst [1] (SGT (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, types.Bool)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq16U x y)
	// cond:
	// result: (XORconst [1] (SGTU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, types.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq32 x y)
	// cond:
	// result: (XORconst [1] (SGT x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, types.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Leq32F x y)
	// cond:
	// result: (FPFlagTrue (CMPGEF y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGEF, TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq32U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq32U x y)
	// cond:
	// result: (XORconst [1] (SGTU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, types.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Leq64F x y)
	// cond:
	// result: (FPFlagTrue (CMPGED y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGED, TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq8 x y)
	// cond:
	// result: (XORconst [1] (SGT (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, types.Bool)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq8U x y)
	// cond:
	// result: (XORconst [1] (SGTU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, types.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLess16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less16 x y)
	// cond:
	// result: (SGT (SignExt16to32 y) (SignExt16to32 x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpLess16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less16U x y)
	// cond:
	// result: (SGTU (ZeroExt16to32 y) (ZeroExt16to32 x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpLess32(v *Value) bool {
	// match: (Less32 x y)
	// cond:
	// result: (SGT y x)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpLess32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Less32F x y)
	// cond:
	// result: (FPFlagTrue (CMPGTF y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTF, TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLess32U(v *Value) bool {
	// match: (Less32U x y)
	// cond:
	// result: (SGTU y x)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpLess64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Less64F x y)
	// cond:
	// result: (FPFlagTrue (CMPGTD y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTD, TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLess8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less8 x y)
	// cond:
	// result: (SGT (SignExt8to32 y) (SignExt8to32 x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpLess8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less8U x y)
	// cond:
	// result: (SGTU (ZeroExt8to32 y) (ZeroExt8to32 x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpLoad(v *Value) bool {
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpMIPSMOVBUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && isSigned(t))
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpMIPSMOVBload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && !isSigned(t))
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpMIPSMOVBUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && isSigned(t))
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpMIPSMOVHload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && !isSigned(t))
	// result: (MOVHUload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpMIPSMOVHUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) || isPtr(t))
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is32BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpMIPSMOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (MOVFload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpMIPSMOVFload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpMIPSMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpLsh16x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh16x16 <t> x y)
	// cond:
	// result: (CMOVZ (SLL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh16x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh16x32 <t> x y)
	// cond:
	// result: (CMOVZ (SLL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueMIPS_OpLsh16x64(v *Value) bool {
	// match: (Lsh16x64 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SLLconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 _ (Const64 [c]))
	// cond: uint32(c) >= 16
	// result: (MOVWconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpLsh16x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh16x8 <t> x y)
	// cond:
	// result: (CMOVZ (SLL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh32x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh32x16 <t> x y)
	// cond:
	// result: (CMOVZ (SLL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh32x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh32x32 <t> x y)
	// cond:
	// result: (CMOVZ (SLL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueMIPS_OpLsh32x64(v *Value) bool {
	// match: (Lsh32x64 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SLLconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 _ (Const64 [c]))
	// cond: uint32(c) >= 32
	// result: (MOVWconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpLsh32x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh32x8 <t> x y)
	// cond:
	// result: (CMOVZ (SLL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh8x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh8x16 <t> x y)
	// cond:
	// result: (CMOVZ (SLL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh8x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh8x32 <t> x y)
	// cond:
	// result: (CMOVZ (SLL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueMIPS_OpLsh8x64(v *Value) bool {
	// match: (Lsh8x64 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SLLconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 _ (Const64 [c]))
	// cond: uint32(c) >= 8
	// result: (MOVWconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpLsh8x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh8x8 <t> x y)
	// cond:
	// result: (CMOVZ (SLL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpMIPSADD(v *Value) bool {
	// match: (ADD x (MOVWconst [c]))
	// cond:
	// result: (ADDconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (ADD (MOVWconst [c]) x)
	// cond:
	// result: (ADDconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (ADD x (NEG y))
	// cond:
	// result: (SUB x y)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSNEG {
			break
		}
		y := v_1.Args[0]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADD (NEG y) x)
	// cond:
	// result: (SUB x y)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSNEG {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSADDconst(v *Value) bool {
	// match: (ADDconst [off1] (MOVWaddr [off2] {sym} ptr))
	// cond:
	// result: (MOVWaddr [off1+off2] {sym} ptr)
	for {
		off1 := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpMIPSMOVWaddr)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		return true
	}
	// match: (ADDconst [0] x)
	// cond:
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(c+d))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c + d))
		return true
	}
	// match: (ADDconst [c] (ADDconst [d] x))
	// cond:
	// result: (ADDconst [int64(int32(c+d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (SUBconst [d] x))
	// cond:
	// result: (ADDconst [int64(int32(c-d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSAND(v *Value) bool {
	b := v.Block
	_ = b
	// match: (AND x (MOVWconst [c]))
	// cond:
	// result: (ANDconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (AND (MOVWconst [c]) x)
	// cond:
	// result: (ANDconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (AND x x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (AND (SGTUconst [1] x) (SGTUconst [1] y))
	// cond:
	// result: (SGTUconst [1] (OR <x.Type> x y))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSGTUconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSSGTUconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSOR, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (AND (SGTUconst [1] y) (SGTUconst [1] x))
	// cond:
	// result: (SGTUconst [1] (OR <x.Type> x y))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSGTUconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		y := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSSGTUconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		x := v_1.Args[0]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSOR, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSANDconst(v *Value) bool {
	// match: (ANDconst [0] _)
	// cond:
	// result: (MOVWconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (ANDconst [-1] x)
	// cond:
	// result: x
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [c&d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c & d
		return true
	}
	// match: (ANDconst [c] (ANDconst [d] x))
	// cond:
	// result: (ANDconst [c&d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSCMOVZ(v *Value) bool {
	b := v.Block
	_ = b
	// match: (CMOVZ _ b (MOVWconst [0]))
	// cond:
	// result: b
	for {
		b := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpMIPSMOVWconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = b.Type
		v.AddArg(b)
		return true
	}
	// match: (CMOVZ a _ (MOVWconst [c]))
	// cond: c!=0
	// result: a
	for {
		a := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpMIPSMOVWconst {
			break
		}
		c := v_2.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}
	// match: (CMOVZ a (MOVWconst [0]) c)
	// cond:
	// result: (CMOVZzero a c)
	for {
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		c := v.Args[2]
		v.reset(OpMIPSCMOVZzero)
		v.AddArg(a)
		v.AddArg(c)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSCMOVZzero(v *Value) bool {
	// match: (CMOVZzero _ (MOVWconst [0]))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (CMOVZzero a (MOVWconst [c]))
	// cond: c!=0
	// result: a
	for {
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSLoweredAtomicAdd(v *Value) bool {
	// match: (LoweredAtomicAdd ptr (MOVWconst [c]) mem)
	// cond: is16Bit(c)
	// result: (LoweredAtomicAddconst [c] ptr mem)
	for {
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpMIPSLoweredAtomicAddconst)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSLoweredAtomicStore(v *Value) bool {
	// match: (LoweredAtomicStore ptr (MOVWconst [0]) mem)
	// cond:
	// result: (LoweredAtomicStorezero ptr mem)
	for {
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPSLoweredAtomicStorezero)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBUload(v *Value) bool {
	// match: (MOVBUload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVBUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2) && !isSigned(x.Type)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2) && !isSigned(x.Type)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBUreg(v *Value) bool {
	b := v.Block
	_ = b
	// match: (MOVBUreg x:(MOVBUload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUload {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg x:(MOVBUreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg <t> x:(MOVBload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBUload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpMIPSMOVBload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		ptr := x.Args[0]
		mem := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBUreg (ANDconst [c] x))
	// cond:
	// result: (ANDconst [c&0xff] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0xff
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [int64(uint8(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBload(v *Value) bool {
	// match: (MOVBload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVBload  [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2) && isSigned(x.Type)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2) && isSigned(x.Type)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBreg(v *Value) bool {
	b := v.Block
	_ = b
	// match: (MOVBreg x:(MOVBload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBload {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg <t> x:(MOVBUload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVBload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		ptr := x.Args[0]
		mem := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVBreg (ANDconst [c] x))
	// cond: c & 0x80 == 0
	// result: (ANDconst [c&0x7f] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x80 == 0) {
			break
		}
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0x7f
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [int64(int8(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBstore(v *Value) bool {
	// match: (MOVBstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVWconst [0]) mem)
	// cond:
	// result: (MOVBstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBreg x) mem)
	// cond:
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBUreg x) mem)
	// cond:
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVBUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHreg x) mem)
	// cond:
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHUreg x) mem)
	// cond:
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVWreg x) mem)
	// cond:
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBstorezero(v *Value) bool {
	// match: (MOVBstorezero [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVBstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstorezero [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVDload(v *Value) bool {
	// match: (MOVDload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVDload  [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off] {sym} ptr (MOVDstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVDstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVDstore(v *Value) bool {
	// match: (MOVDstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVFload(v *Value) bool {
	// match: (MOVFload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVFload  [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVFload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVFload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFload [off] {sym} ptr (MOVFstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVFstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVFstore(v *Value) bool {
	// match: (MOVFstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVFstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHUload(v *Value) bool {
	// match: (MOVHUload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVHUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2) && !isSigned(x.Type)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVHstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2) && !isSigned(x.Type)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHUreg(v *Value) bool {
	b := v.Block
	_ = b
	// match: (MOVHUreg x:(MOVBUload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUload {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVHUload {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVBUreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVHUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg <t> x:(MOVHload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVHUload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpMIPSMOVHload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		ptr := x.Args[0]
		mem := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHUload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVHUreg (ANDconst [c] x))
	// cond:
	// result: (ANDconst [c&0xffff] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0xffff
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [int64(uint16(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHload(v *Value) bool {
	// match: (MOVHload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVHload  [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2) && isSigned(x.Type)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVHstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2) && isSigned(x.Type)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHreg(v *Value) bool {
	b := v.Block
	_ = b
	// match: (MOVHreg x:(MOVBload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBload {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUload {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHload _ _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVHload {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHreg _))
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVHreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg <t> x:(MOVHUload [off] {sym} ptr mem))
	// cond: x.Uses == 1 && clobber(x)
	// result: @x.Block (MOVHload <t> [off] {sym} ptr mem)
	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpMIPSMOVHUload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		ptr := x.Args[0]
		mem := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	// match: (MOVHreg (ANDconst [c] x))
	// cond: c & 0x8000 == 0
	// result: (ANDconst [c&0x7fff] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x8000 == 0) {
			break
		}
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0x7fff
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [int64(int16(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHstore(v *Value) bool {
	// match: (MOVHstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVHstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVWconst [0]) mem)
	// cond:
	// result: (MOVHstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPSMOVHstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHreg x) mem)
	// cond:
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHUreg x) mem)
	// cond:
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVWreg x) mem)
	// cond:
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHstorezero(v *Value) bool {
	// match: (MOVHstorezero [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVHstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstorezero [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWload(v *Value) bool {
	// match: (MOVWload [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVWload  [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off] {sym} ptr (MOVWstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWreg(v *Value) bool {
	// match: (MOVWreg x)
	// cond: x.Uses == 1
	// result: (MOVWnop x)
	for {
		x := v.Args[0]
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWnop)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [c])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWstore(v *Value) bool {
	// match: (MOVWstore [off1] {sym} x:(ADDconst [off2] ptr) val mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWconst [0]) mem)
	// cond:
	// result: (MOVWstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPSMOVWstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWreg x) mem)
	// cond:
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWstorezero(v *Value) bool {
	// match: (MOVWstorezero [off1] {sym} x:(ADDconst [off2] ptr) mem)
	// cond: (is16Bit(off1+off2) || x.Uses == 1)
	// result: (MOVWstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstorezero [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMUL(v *Value) bool {
	// match: (MUL (MOVWconst [0]) _)
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (MUL _ (MOVWconst [0]))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (MUL (MOVWconst [1]) x)
	// cond:
	// result: x
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MUL x (MOVWconst [1]))
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MUL (MOVWconst [-1]) x)
	// cond:
	// result: (NEG x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		x := v.Args[1]
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}
	// match: (MUL x (MOVWconst [-1]))
	// cond:
	// result: (NEG x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}
	// match: (MUL (MOVWconst [c]) x)
	// cond: isPowerOfTwo(int64(uint32(c)))
	// result: (SLLconst [log2(int64(uint32(c)))] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: isPowerOfTwo(int64(uint32(c)))
	// result: (SLLconst [log2(int64(uint32(c)))] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}
	// match: (MUL (MOVWconst [c]) (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(c)*int32(d))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c) * int32(d))
		return true
	}
	// match: (MUL (MOVWconst [d]) (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [int64(int32(c)*int32(d))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c) * int32(d))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSNEG(v *Value) bool {
	// match: (NEG (MOVWconst [c]))
	// cond:
	// result: (MOVWconst [int64(int32(-c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(-c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSNOR(v *Value) bool {
	// match: (NOR x (MOVWconst [c]))
	// cond:
	// result: (NORconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSNORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (NOR (MOVWconst [c]) x)
	// cond:
	// result: (NORconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSNORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSNORconst(v *Value) bool {
	// match: (NORconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [^(c|d)])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = ^(c | d)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSOR(v *Value) bool {
	b := v.Block
	_ = b
	// match: (OR x (MOVWconst [c]))
	// cond:
	// result: (ORconst  [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (OR (MOVWconst [c]) x)
	// cond:
	// result: (ORconst  [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (OR x x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (OR (SGTUzero x) (SGTUzero y))
	// cond:
	// result: (SGTUzero (OR <x.Type> x y))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSGTUzero {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSSGTUzero {
			break
		}
		y := v_1.Args[0]
		v.reset(OpMIPSSGTUzero)
		v0 := b.NewValue0(v.Pos, OpMIPSOR, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (OR (SGTUzero y) (SGTUzero x))
	// cond:
	// result: (SGTUzero (OR <x.Type> x y))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSGTUzero {
			break
		}
		y := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSSGTUzero {
			break
		}
		x := v_1.Args[0]
		v.reset(OpMIPSSGTUzero)
		v0 := b.NewValue0(v.Pos, OpMIPSOR, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSORconst(v *Value) bool {
	// match: (ORconst [0] x)
	// cond:
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ORconst [-1] _)
	// cond:
	// result: (MOVWconst [-1])
	for {
		if v.AuxInt != -1 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = -1
		return true
	}
	// match: (ORconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [c|d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c | d
		return true
	}
	// match: (ORconst [c] (ORconst [d] x))
	// cond:
	// result: (ORconst [c|d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGT(v *Value) bool {
	// match: (SGT (MOVWconst [c]) x)
	// cond:
	// result: (SGTconst  [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSSGTconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SGT x (MOVWconst [0]))
	// cond:
	// result: (SGTzero x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSSGTzero)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTU(v *Value) bool {
	// match: (SGTU (MOVWconst [c]) x)
	// cond:
	// result: (SGTUconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SGTU x (MOVWconst [0]))
	// cond:
	// result: (SGTUzero x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSSGTUzero)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTUconst(v *Value) bool {
	// match: (SGTUconst [c] (MOVWconst [d]))
	// cond: uint32(c)>uint32(d)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(c) > uint32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (MOVWconst [d]))
	// cond: uint32(c)<=uint32(d)
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(c) <= uint32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTUconst [c] (MOVBUreg _))
	// cond: 0xff < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVBUreg {
			break
		}
		if !(0xff < uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (MOVHUreg _))
	// cond: 0xffff < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVHUreg {
			break
		}
		if !(0xffff < uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (ANDconst [m] _))
	// cond: uint32(m) < uint32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		m := v_0.AuxInt
		if !(uint32(m) < uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (SRLconst _ [d]))
	// cond: uint32(d) <= 31 && 1<<(32-uint32(d)) <= uint32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSRLconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(d) <= 31 && 1<<(32-uint32(d)) <= uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTUzero(v *Value) bool {
	// match: (SGTUzero (MOVWconst [d]))
	// cond: uint32(d) != 0
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(d) != 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUzero (MOVWconst [d]))
	// cond: uint32(d) == 0
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(d) == 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTconst(v *Value) bool {
	// match: (SGTconst [c] (MOVWconst [d]))
	// cond: int32(c) > int32(d)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(c) > int32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVWconst [d]))
	// cond: int32(c) <= int32(d)
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(c) <= int32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVBreg _))
	// cond: 0x7f < int32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVBreg {
			break
		}
		if !(0x7f < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVBreg _))
	// cond: int32(c) <= -0x80
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVBreg {
			break
		}
		if !(int32(c) <= -0x80) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVBUreg _))
	// cond: 0xff < int32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVBUreg {
			break
		}
		if !(0xff < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVBUreg _))
	// cond: int32(c) < 0
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVBUreg {
			break
		}
		if !(int32(c) < 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVHreg _))
	// cond: 0x7fff < int32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVHreg {
			break
		}
		if !(0x7fff < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVHreg _))
	// cond: int32(c) <= -0x8000
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVHreg {
			break
		}
		if !(int32(c) <= -0x8000) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVHUreg _))
	// cond: 0xffff < int32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVHUreg {
			break
		}
		if !(0xffff < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVHUreg _))
	// cond: int32(c) < 0
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVHUreg {
			break
		}
		if !(int32(c) < 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (ANDconst [m] _))
	// cond: 0 <= int32(m) && int32(m) < int32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int32(m) && int32(m) < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (SRLconst _ [d]))
	// cond: 0 <= int32(c) && uint32(d) <= 31 && 1<<(32-uint32(d)) <= int32(c)
	// result: (MOVWconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSRLconst {
			break
		}
		d := v_0.AuxInt
		if !(0 <= int32(c) && uint32(d) <= 31 && 1<<(32-uint32(d)) <= int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTzero(v *Value) bool {
	// match: (SGTzero (MOVWconst [d]))
	// cond: int32(d) > 0
	// result: (MOVWconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(d) > 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTzero (MOVWconst [d]))
	// cond: int32(d) <= 0
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(d) <= 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSLL(v *Value) bool {
	// match: (SLL _ (MOVWconst [c]))
	// cond: uint32(c)>=32
	// result: (MOVWconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SLL x (MOVWconst [c]))
	// cond:
	// result: (SLLconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSLLconst(v *Value) bool {
	// match: (SLLconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(uint32(d)<<uint32(c)))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(d) << uint32(c)))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRA(v *Value) bool {
	// match: (SRA x (MOVWconst [c]))
	// cond: uint32(c)>=32
	// result: (SRAconst x [31])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	// match: (SRA x (MOVWconst [c]))
	// cond:
	// result: (SRAconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSRAconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRAconst(v *Value) bool {
	// match: (SRAconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(d)>>uint32(c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(d) >> uint32(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRL(v *Value) bool {
	// match: (SRL _ (MOVWconst [c]))
	// cond: uint32(c)>=32
	// result: (MOVWconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SRL x (MOVWconst [c]))
	// cond:
	// result: (SRLconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSRLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRLconst(v *Value) bool {
	// match: (SRLconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(uint32(d)>>uint32(c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(uint32(d) >> uint32(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSUB(v *Value) bool {
	// match: (SUB x (MOVWconst [c]))
	// cond:
	// result: (SUBconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSUBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUB x x)
	// cond:
	// result: (MOVWconst [0])
	for {
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SUB (MOVWconst [0]) x)
	// cond:
	// result: (NEG x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSUBconst(v *Value) bool {
	// match: (SUBconst [0] x)
	// cond:
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [int64(int32(d-c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(d - c))
		return true
	}
	// match: (SUBconst [c] (SUBconst [d] x))
	// cond:
	// result: (ADDconst [int64(int32(-c-d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(-c - d))
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (ADDconst [d] x))
	// cond:
	// result: (ADDconst [int64(int32(-c+d))] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(-c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSXOR(v *Value) bool {
	// match: (XOR x (MOVWconst [c]))
	// cond:
	// result: (XORconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSXORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (XOR (MOVWconst [c]) x)
	// cond:
	// result: (XORconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (XOR x x)
	// cond:
	// result: (MOVWconst [0])
	for {
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSXORconst(v *Value) bool {
	// match: (XORconst [0] x)
	// cond:
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (XORconst [-1] x)
	// cond:
	// result: (NORconst [0] x)
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
	// match: (XORconst [c] (MOVWconst [d]))
	// cond:
	// result: (MOVWconst [c^d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c ^ d
		return true
	}
	// match: (XORconst [c] (XORconst [d] x))
	// cond:
	// result: (XORconst [c^d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSXORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSXORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMod16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod16 x y)
	// cond:
	// result: (Select0 (DIV (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, MakeTuple(types.Int32, types.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod16u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod16u x y)
	// cond:
	// result: (Select0 (DIVU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, MakeTuple(types.UInt32, types.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod32 x y)
	// cond:
	// result: (Select0 (DIV x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, MakeTuple(types.Int32, types.Int32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod32u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod32u x y)
	// cond:
	// result: (Select0 (DIVU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, MakeTuple(types.UInt32, types.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod8 x y)
	// cond:
	// result: (Select0 (DIV (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, MakeTuple(types.Int32, types.Int32))
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod8u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod8u x y)
	// cond:
	// result: (Select0 (DIVU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, MakeTuple(types.UInt32, types.UInt32))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMove(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	types := &b.Func.Config.Types
	_ = types
	// match: (Move [0] _ _ mem)
	// cond:
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// cond:
	// result: (MOVBstore dst (MOVBUload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, types.UInt8)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] {t} dst src mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore dst (MOVHUload src mem) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHUload, types.UInt16)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// cond:
	// result: (MOVBstore [1] dst (MOVBUload [1] src mem) 		(MOVBstore dst (MOVBUload src mem) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 1
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, types.UInt8)
		v0.AuxInt = 1
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVBUload, types.UInt8)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, types.UInt32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore [2] dst (MOVHUload [2] src mem) 		(MOVHstore dst (MOVHUload src mem) mem))
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHUload, types.UInt16)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVHUload, types.UInt16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [4] dst src mem)
	// cond:
	// result: (MOVBstore [3] dst (MOVBUload [3] src mem) 		(MOVBstore [2] dst (MOVBUload [2] src mem) 			(MOVBstore [1] dst (MOVBUload [1] src mem) 				(MOVBstore dst (MOVBUload src mem) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, types.UInt8)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVBUload, types.UInt8)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v3.AuxInt = 1
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVBUload, types.UInt8)
		v4.AuxInt = 1
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVBUload, types.UInt8)
		v6.AddArg(src)
		v6.AddArg(mem)
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [3] dst src mem)
	// cond:
	// result: (MOVBstore [2] dst (MOVBUload [2] src mem) 		(MOVBstore [1] dst (MOVBUload [1] src mem) 			(MOVBstore dst (MOVBUload src mem) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, types.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v1.AuxInt = 1
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVBUload, types.UInt8)
		v2.AuxInt = 1
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVBUload, types.UInt8)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore [4] dst (MOVWload [4] src mem) 		(MOVWstore dst (MOVWload src mem) mem))
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, types.UInt32)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWload, types.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore [6] dst (MOVHload [6] src mem) 		(MOVHstore [4] dst (MOVHload [4] src mem) 			(MOVHstore [2] dst (MOVHload [2] src mem) 				(MOVHstore dst (MOVHload src mem) mem))))
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHload, types.Int16)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVHload, types.Int16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVHstore, TypeMem)
		v3.AuxInt = 2
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVHload, types.Int16)
		v4.AuxInt = 2
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVHstore, TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVHload, types.Int16)
		v6.AddArg(src)
		v6.AddArg(mem)
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [6] {t} dst src mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore [4] dst (MOVHload [4] src mem) 		(MOVHstore [2] dst (MOVHload [2] src mem) 			(MOVHstore dst (MOVHload src mem) mem)))
	for {
		if v.AuxInt != 6 {
			break
		}
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHload, types.Int16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVHload, types.Int16)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVHstore, TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVHload, types.Int16)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [12] {t} dst src mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore [8] dst (MOVWload [8] src mem) 		(MOVWstore [4] dst (MOVWload [4] src mem) 			(MOVWstore dst (MOVWload src mem) mem)))
	for {
		if v.AuxInt != 12 {
			break
		}
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, types.UInt32)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWload, types.UInt32)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWload, types.UInt32)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [16] {t} dst src mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore [12] dst (MOVWload [12] src mem) 		(MOVWstore [8] dst (MOVWload [8] src mem) 			(MOVWstore [4] dst (MOVWload [4] src mem) 				(MOVWstore dst (MOVWload src mem) mem))))
	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 12
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, types.UInt32)
		v0.AuxInt = 12
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v1.AuxInt = 8
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWload, types.UInt32)
		v2.AuxInt = 8
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v3.AuxInt = 4
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWload, types.UInt32)
		v4.AuxInt = 4
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVWload, types.UInt32)
		v6.AddArg(src)
		v6.AddArg(mem)
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: (s > 16 || t.(Type).Alignment()%4 != 0)
	// result: (LoweredMove [t.(Type).Alignment()] 		dst 		src 		(ADDconst <src.Type> src [s-moveSize(t.(Type).Alignment(), config)]) 		mem)
	for {
		s := v.AuxInt
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 16 || t.(Type).Alignment()%4 != 0) {
			break
		}
		v.reset(OpMIPSLoweredMove)
		v.AuxInt = t.(Type).Alignment()
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpMIPSADDconst, src.Type)
		v0.AuxInt = s - moveSize(t.(Type).Alignment(), config)
		v0.AddArg(src)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMul16(v *Value) bool {
	// match: (Mul16 x y)
	// cond:
	// result: (MUL x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpMul32(v *Value) bool {
	// match: (Mul32 x y)
	// cond:
	// result: (MUL x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpMul32F(v *Value) bool {
	// match: (Mul32F x y)
	// cond:
	// result: (MULF x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpMul32uhilo(v *Value) bool {
	// match: (Mul32uhilo x y)
	// cond:
	// result: (MULTU x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMULTU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpMul64F(v *Value) bool {
	// match: (Mul64F x y)
	// cond:
	// result: (MULD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpMul8(v *Value) bool {
	// match: (Mul8 x y)
	// cond:
	// result: (MUL x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpNeg16(v *Value) bool {
	// match: (Neg16 x)
	// cond:
	// result: (NEG x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpNeg32(v *Value) bool {
	// match: (Neg32 x)
	// cond:
	// result: (NEG x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpNeg32F(v *Value) bool {
	// match: (Neg32F x)
	// cond:
	// result: (NEGF x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSNEGF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpNeg64F(v *Value) bool {
	// match: (Neg64F x)
	// cond:
	// result: (NEGD x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSNEGD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpNeg8(v *Value) bool {
	// match: (Neg8 x)
	// cond:
	// result: (NEG x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpNeq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Neq16 x y)
	// cond:
	// result: (SGTU (XOR (ZeroExt16to32 x) (ZeroExt16to32 y)) (MOVWconst [0]))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, types.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpNeq32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Neq32 x y)
	// cond:
	// result: (SGTU (XOR x y) (MOVWconst [0]))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, types.UInt32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpNeq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Neq32F x y)
	// cond:
	// result: (FPFlagFalse (CMPEQF x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQF, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpNeq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Neq64F x y)
	// cond:
	// result: (FPFlagFalse (CMPEQD x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQD, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpNeq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Neq8 x y)
	// cond:
	// result: (SGTU (XOR (ZeroExt8to32 x) (ZeroExt8to32 y)) (MOVWconst [0]))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, types.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpNeqB(v *Value) bool {
	// match: (NeqB x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpNeqPtr(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (NeqPtr x y)
	// cond:
	// result: (SGTU (XOR x y) (MOVWconst [0]))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, types.UInt32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpNilCheck(v *Value) bool {
	// match: (NilCheck ptr mem)
	// cond:
	// result: (LoweredNilCheck ptr mem)
	for {
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSLoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpNot(v *Value) bool {
	// match: (Not x)
	// cond:
	// result: (XORconst [1] x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpOffPtr(v *Value) bool {
	// match: (OffPtr [off] ptr:(SP))
	// cond:
	// result: (MOVWaddr [off] ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpMIPSMOVWaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// cond:
	// result: (ADDconst [off] ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueMIPS_OpOr16(v *Value) bool {
	// match: (Or16 x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpOr32(v *Value) bool {
	// match: (Or32 x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpOr8(v *Value) bool {
	// match: (Or8 x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpOrB(v *Value) bool {
	// match: (OrB x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpRound32F(v *Value) bool {
	// match: (Round32F x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpRound64F(v *Value) bool {
	// match: (Round64F x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpRsh16Ux16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux16 <t> x y)
	// cond:
	// result: (CMOVZ (SRL <t> (ZeroExt16to32 x) (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh16Ux32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux32 <t> x y)
	// cond:
	// result: (CMOVZ (SRL <t> (ZeroExt16to32 x) y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh16Ux64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux64 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SRLconst (SLLconst <types.UInt32> x [16]) [c+16])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpMIPSSRLconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 _ (Const64 [c]))
	// cond: uint32(c) >= 16
	// result: (MOVWconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh16Ux8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux8 <t> x y)
	// cond:
	// result: (CMOVZ (SRL <t> (ZeroExt16to32 x) (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh16x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x16 x y)
	// cond:
	// result: (SRA (SignExt16to32 x) ( CMOVZ <types.UInt32> (ZeroExt16to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt16to32 y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, types.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v3.AuxInt = -1
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh16x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x32 x y)
	// cond:
	// result: (SRA (SignExt16to32 x) ( CMOVZ <types.UInt32> y (MOVWconst [-1]) (SGTUconst [32] y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, types.UInt32)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = -1
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh16x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SRAconst (SLLconst <types.UInt32> x [16]) [c+16])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint32(c) >= 16
	// result: (SRAconst (SLLconst <types.UInt32> x [16]) [31])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh16x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x8 x y)
	// cond:
	// result: (SRA (SignExt16to32 x) ( CMOVZ <types.UInt32> (ZeroExt8to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt8to32 y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, types.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v3.AuxInt = -1
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh32Ux16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32Ux16 <t> x y)
	// cond:
	// result: (CMOVZ (SRL <t> x (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh32Ux32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32Ux32 <t> x y)
	// cond:
	// result: (CMOVZ (SRL <t> x y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueMIPS_OpRsh32Ux64(v *Value) bool {
	// match: (Rsh32Ux64 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SRLconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpMIPSSRLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux64 _ (Const64 [c]))
	// cond: uint32(c) >= 32
	// result: (MOVWconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh32Ux8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32Ux8 <t> x y)
	// cond:
	// result: (CMOVZ (SRL <t> x (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh32x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32x16 x y)
	// cond:
	// result: (SRA x ( CMOVZ <types.UInt32> (ZeroExt16to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt16to32 y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPSCMOVZ, types.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = -1
		v0.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpRsh32x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32x32 x y)
	// cond:
	// result: (SRA x ( CMOVZ <types.UInt32> y (MOVWconst [-1]) (SGTUconst [32] y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPSCMOVZ, types.UInt32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = -1
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpRsh32x64(v *Value) bool {
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SRAconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint32(c) >= 32
	// result: (SRAconst x [31])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh32x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32x8 x y)
	// cond:
	// result: (SRA x ( CMOVZ <types.UInt32> (ZeroExt8to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt8to32 y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPSCMOVZ, types.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = -1
		v0.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpRsh8Ux16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux16 <t> x y)
	// cond:
	// result: (CMOVZ (SRL <t> (ZeroExt8to32 x) (ZeroExt16to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt16to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh8Ux32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux32 <t> x y)
	// cond:
	// result: (CMOVZ (SRL <t> (ZeroExt8to32 x) y) (MOVWconst [0]) (SGTUconst [32] y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh8Ux64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux64 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SRLconst (SLLconst <types.UInt32> x [24]) [c+24])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpMIPSSRLconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 _ (Const64 [c]))
	// cond: uint32(c) >= 8
	// result: (MOVWconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh8Ux8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux8 <t> x y)
	// cond:
	// result: (CMOVZ (SRL <t> (ZeroExt8to32 x) (ZeroExt8to32 y) ) (MOVWconst [0]) (SGTUconst [32] (ZeroExt8to32 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh8x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x16 x y)
	// cond:
	// result: (SRA (SignExt16to32 x) ( CMOVZ <types.UInt32> (ZeroExt16to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt16to32 y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, types.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v3.AuxInt = -1
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh8x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x32 x y)
	// cond:
	// result: (SRA (SignExt16to32 x) ( CMOVZ <types.UInt32> y (MOVWconst [-1]) (SGTUconst [32] y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, types.UInt32)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = -1
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh8x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SRAconst (SLLconst <types.UInt32> x [24]) [c+24])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint32(c) >= 8
	// result: (SRAconst (SLLconst <types.UInt32> x [24]) [31])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, types.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh8x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x8 x y)
	// cond:
	// result: (SRA (SignExt16to32 x) ( CMOVZ <types.UInt32> (ZeroExt8to32 y) (MOVWconst [-1]) (SGTUconst [32] (ZeroExt8to32 y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, types.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v3.AuxInt = -1
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, types.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpSelect0(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Select0 (Add32carry <t> x y))
	// cond:
	// result: (ADD <t.FieldType(0)> x y)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32carry {
			break
		}
		t := v_0.Type
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpMIPSADD)
		v.Type = t.FieldType(0)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Select0 (Sub32carry <t> x y))
	// cond:
	// result: (SUB <t.FieldType(0)> x y)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSub32carry {
			break
		}
		t := v_0.Type
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpMIPSSUB)
		v.Type = t.FieldType(0)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (Select0 (MULTU (MOVWconst [0]) _))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_0.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select0 (MULTU _ (MOVWconst [0])))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select0 (MULTU (MOVWconst [1]) _))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_0.AuxInt != 1 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select0 (MULTU _ (MOVWconst [1])))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select0 (MULTU (MOVWconst [-1]) x))
	// cond:
	// result: (CMOVZ (ADDconst <x.Type> [-1] x) (MOVWconst [0]) x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_0.AuxInt != -1 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSADDconst, x.Type)
		v0.AuxInt = -1
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v.AddArg(x)
		return true
	}
	// match: (Select0 (MULTU x (MOVWconst [-1])))
	// cond:
	// result: (CMOVZ (ADDconst <x.Type> [-1] x) (MOVWconst [0]) x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_1.AuxInt != -1 {
			break
		}
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSADDconst, x.Type)
		v0.AuxInt = -1
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v.AddArg(x)
		return true
	}
	// match: (Select0 (MULTU (MOVWconst [c]) x))
	// cond: isPowerOfTwo(int64(uint32(c)))
	// result: (SRLconst [32-log2(int64(uint32(c)))] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSRLconst)
		v.AuxInt = 32 - log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}
	// match: (Select0 (MULTU x (MOVWconst [c])))
	// cond: isPowerOfTwo(int64(uint32(c)))
	// result: (SRLconst [32-log2(int64(uint32(c)))] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSRLconst)
		v.AuxInt = 32 - log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}
	// match: (Select0 (MULTU (MOVWconst [c]) (MOVWconst [d])))
	// cond:
	// result: (MOVWconst [(c*d)>>32])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = (c * d) >> 32
		return true
	}
	// match: (Select0 (MULTU (MOVWconst [d]) (MOVWconst [c])))
	// cond:
	// result: (MOVWconst [(c*d)>>32])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = (c * d) >> 32
		return true
	}
	// match: (Select0 (DIV (MOVWconst [c]) (MOVWconst [d])))
	// cond:
	// result: (MOVWconst [int64(int32(c)%int32(d))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSDIV {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c) % int32(d))
		return true
	}
	// match: (Select0 (DIVU (MOVWconst [c]) (MOVWconst [d])))
	// cond:
	// result: (MOVWconst [int64(int32(uint32(c)%uint32(d)))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSDIVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(c) % uint32(d)))
		return true
	}
	return false
}
func rewriteValueMIPS_OpSelect1(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Select1 (Add32carry <t> x y))
	// cond:
	// result: (SGTU <types.Bool> x (ADD <t.FieldType(0)> x y))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32carry {
			break
		}
		t := v_0.Type
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpMIPSSGTU)
		v.Type = types.Bool
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPSADD, t.FieldType(0))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (Select1 (Sub32carry <t> x y))
	// cond:
	// result: (SGTU <types.Bool> (SUB <t.FieldType(0)> x y) x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSub32carry {
			break
		}
		t := v_0.Type
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpMIPSSGTU)
		v.Type = types.Bool
		v0 := b.NewValue0(v.Pos, OpMIPSSUB, t.FieldType(0))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULTU (MOVWconst [0]) _))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_0.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select1 (MULTU _ (MOVWconst [0])))
	// cond:
	// result: (MOVWconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select1 (MULTU (MOVWconst [1]) x))
	// cond:
	// result: x
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_0.AuxInt != 1 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULTU x (MOVWconst [1])))
	// cond:
	// result: x
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULTU (MOVWconst [-1]) x))
	// cond:
	// result: (NEG <x.Type> x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_0.AuxInt != -1 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpMIPSNEG)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULTU x (MOVWconst [-1])))
	// cond:
	// result: (NEG <x.Type> x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_1.AuxInt != -1 {
			break
		}
		v.reset(OpMIPSNEG)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULTU (MOVWconst [c]) x))
	// cond: isPowerOfTwo(int64(uint32(c)))
	// result: (SLLconst [log2(int64(uint32(c)))] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULTU x (MOVWconst [c])))
	// cond: isPowerOfTwo(int64(uint32(c)))
	// result: (SLLconst [log2(int64(uint32(c)))] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULTU (MOVWconst [c]) (MOVWconst [d])))
	// cond:
	// result: (MOVWconst [int64(int32(uint32(c)*uint32(d)))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(c) * uint32(d)))
		return true
	}
	// match: (Select1 (MULTU (MOVWconst [d]) (MOVWconst [c])))
	// cond:
	// result: (MOVWconst [int64(int32(uint32(c)*uint32(d)))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(c) * uint32(d)))
		return true
	}
	// match: (Select1 (DIV (MOVWconst [c]) (MOVWconst [d])))
	// cond:
	// result: (MOVWconst [int64(int32(c)/int32(d))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSDIV {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c) / int32(d))
		return true
	}
	// match: (Select1 (DIVU (MOVWconst [c]) (MOVWconst [d])))
	// cond:
	// result: (MOVWconst [int64(int32(uint32(c)/uint32(d)))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSDIVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(c) / uint32(d)))
		return true
	}
	return false
}
func rewriteValueMIPS_OpSignExt16to32(v *Value) bool {
	// match: (SignExt16to32 x)
	// cond:
	// result: (MOVHreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpSignExt8to16(v *Value) bool {
	// match: (SignExt8to16 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpSignExt8to32(v *Value) bool {
	// match: (SignExt8to32 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpSignmask(v *Value) bool {
	// match: (Signmask x)
	// cond:
	// result: (SRAconst x [31])
	for {
		x := v.Args[0]
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpSlicemask(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Slicemask <t> x)
	// cond:
	// result: (SRAconst (NEG <t> x) [31])
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpMIPSNEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpSqrt(v *Value) bool {
	// match: (Sqrt x)
	// cond:
	// result: (SQRTD x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSSQRTD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpStaticCall(v *Value) bool {
	// match: (StaticCall [argwid] {target} mem)
	// cond:
	// result: (CALLstatic [argwid] {target} mem)
	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpMIPSCALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpStore(v *Value) bool {
	// match: (Store {t} ptr val mem)
	// cond: t.(Type).Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := v.Aux
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 1) {
			break
		}
		v.reset(OpMIPSMOVBstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(Type).Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := v.Aux
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 2) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(Type).Size() == 4 && !is32BitFloat(val.Type)
	// result: (MOVWstore ptr val mem)
	for {
		t := v.Aux
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 4 && !is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(Type).Size() == 4 && is32BitFloat(val.Type)
	// result: (MOVFstore ptr val mem)
	for {
		t := v.Aux
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPSMOVFstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(Type).Size() == 8 && is64BitFloat(val.Type)
	// result: (MOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPSMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpSub16(v *Value) bool {
	// match: (Sub16 x y)
	// cond:
	// result: (SUB x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpSub32(v *Value) bool {
	// match: (Sub32 x y)
	// cond:
	// result: (SUB x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpSub32F(v *Value) bool {
	// match: (Sub32F x y)
	// cond:
	// result: (SUBF x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUBF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpSub32withcarry(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Sub32withcarry <t> x y c)
	// cond:
	// result: (SUB (SUB <t> x y) c)
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		c := v.Args[2]
		v.reset(OpMIPSSUB)
		v0 := b.NewValue0(v.Pos, OpMIPSSUB, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(c)
		return true
	}
}
func rewriteValueMIPS_OpSub64F(v *Value) bool {
	// match: (Sub64F x y)
	// cond:
	// result: (SUBD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUBD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpSub8(v *Value) bool {
	// match: (Sub8 x y)
	// cond:
	// result: (SUB x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpSubPtr(v *Value) bool {
	// match: (SubPtr x y)
	// cond:
	// result: (SUB x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpTrunc16to8(v *Value) bool {
	// match: (Trunc16to8 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpTrunc32to16(v *Value) bool {
	// match: (Trunc32to16 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpTrunc32to8(v *Value) bool {
	// match: (Trunc32to8 x)
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpXor16(v *Value) bool {
	// match: (Xor16 x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpXor32(v *Value) bool {
	// match: (Xor32 x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpXor8(v *Value) bool {
	// match: (Xor8 x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpZero(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	types := &b.Func.Config.Types
	_ = types
	// match: (Zero [0] _ mem)
	// cond:
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v.Args[1]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Zero [1] ptr mem)
	// cond:
	// result: (MOVBstore ptr (MOVWconst [0]) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSMOVBstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] {t} ptr mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore ptr (MOVWconst [0]) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] ptr mem)
	// cond:
	// result: (MOVBstore [1] ptr (MOVWconst [0]) 		(MOVBstore [0] ptr (MOVWconst [0]) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 1
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore ptr (MOVWconst [0]) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore [2] ptr (MOVWconst [0]) 		(MOVHstore [0] ptr (MOVWconst [0]) mem))
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [4] ptr mem)
	// cond:
	// result: (MOVBstore [3] ptr (MOVWconst [0]) 		(MOVBstore [2] ptr (MOVWconst [0]) 			(MOVBstore [1] ptr (MOVWconst [0]) 				(MOVBstore [0] ptr (MOVWconst [0]) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 3
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v3.AuxInt = 1
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [3] ptr mem)
	// cond:
	// result: (MOVBstore [2] ptr (MOVWconst [0]) 		(MOVBstore [1] ptr (MOVWconst [0]) 			(MOVBstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v1.AuxInt = 1
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [6] {t} ptr mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore [4] ptr (MOVWconst [0]) 		(MOVHstore [2] ptr (MOVWconst [0]) 			(MOVHstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if v.AuxInt != 6 {
			break
		}
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVHstore, TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore [4] ptr (MOVWconst [0]) 			(MOVWstore [0] ptr (MOVWconst [0]) mem))
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [12] {t} ptr mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore [8] ptr (MOVWconst [0]) 		(MOVWstore [4] ptr (MOVWconst [0]) 			(MOVWstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if v.AuxInt != 12 {
			break
		}
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 8
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v1.AuxInt = 4
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [16] {t} ptr mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore [12] ptr (MOVWconst [0]) 		(MOVWstore [8] ptr (MOVWconst [0]) 			(MOVWstore [4] ptr (MOVWconst [0]) 				(MOVWstore [0] ptr (MOVWconst [0]) mem))))
	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 12
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v1.AuxInt = 8
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v3.AuxInt = 4
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVWstore, TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: (s > 16  || t.(Type).Alignment()%4 != 0)
	// result: (LoweredZero [t.(Type).Alignment()] 		ptr 		(ADDconst <ptr.Type> ptr [s-moveSize(t.(Type).Alignment(), config)]) 		mem)
	for {
		s := v.AuxInt
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(s > 16 || t.(Type).Alignment()%4 != 0) {
			break
		}
		v.reset(OpMIPSLoweredZero)
		v.AuxInt = t.(Type).Alignment()
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSADDconst, ptr.Type)
		v0.AuxInt = s - moveSize(t.(Type).Alignment(), config)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpZeroExt16to32(v *Value) bool {
	// match: (ZeroExt16to32 x)
	// cond:
	// result: (MOVHUreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVHUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpZeroExt8to16(v *Value) bool {
	// match: (ZeroExt8to16 x)
	// cond:
	// result: (MOVBUreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpZeroExt8to32(v *Value) bool {
	// match: (ZeroExt8to32 x)
	// cond:
	// result: (MOVBUreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpZeromask(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Zeromask x)
	// cond:
	// result: (NEG (SGTU x (MOVWconst [0])))
	for {
		x := v.Args[0]
		v.reset(OpMIPSNEG)
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, types.Bool)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, types.UInt32)
		v1.AuxInt = 0
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteBlockMIPS(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	types := &config.Types
	_ = types
	switch b.Kind {
	case BlockMIPSEQ:
		// match: (EQ (FPFlagTrue cmp) yes no)
		// cond:
		// result: (FPF cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSFPFlagTrue {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSFPF
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (FPFlagFalse cmp) yes no)
		// cond:
		// result: (FPT cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSFPFlagFalse {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSFPT
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGT _ _)) yes no)
		// cond:
		// result: (NE cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGT {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTU _ _)) yes no)
		// cond:
		// result: (NE cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTU {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTconst _)) yes no)
		// cond:
		// result: (NE cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTconst {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTUconst _)) yes no)
		// cond:
		// result: (NE cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTUconst {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTzero _)) yes no)
		// cond:
		// result: (NE cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTzero {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTUzero _)) yes no)
		// cond:
		// result: (NE cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTUzero {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (SGTUconst [1] x) yes no)
		// cond:
		// result: (NE x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSSGTUconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSNE
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (SGTUzero x) yes no)
		// cond:
		// result: (EQ x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSSGTUzero {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSEQ
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (SGTconst [0] x) yes no)
		// cond:
		// result: (GEZ x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSSGTconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSGEZ
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (SGTzero x) yes no)
		// cond:
		// result: (LEZ x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSSGTzero {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSLEZ
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (MOVWconst [0]) yes no)
		// cond:
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockFirst
			b.SetControl(nil)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (MOVWconst [c]) yes no)
		// cond: c != 0
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c != 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.swapSuccessors()
			_ = no
			_ = yes
			return true
		}
	case BlockMIPSGEZ:
		// match: (GEZ (MOVWconst [c]) yes no)
		// cond: int32(c) >= 0
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(int32(c) >= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			_ = yes
			_ = no
			return true
		}
		// match: (GEZ (MOVWconst [c]) yes no)
		// cond: int32(c) <  0
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(int32(c) < 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.swapSuccessors()
			_ = no
			_ = yes
			return true
		}
	case BlockMIPSGTZ:
		// match: (GTZ (MOVWconst [c]) yes no)
		// cond: int32(c) >  0
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(int32(c) > 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			_ = yes
			_ = no
			return true
		}
		// match: (GTZ (MOVWconst [c]) yes no)
		// cond: int32(c) <= 0
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(int32(c) <= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.swapSuccessors()
			_ = no
			_ = yes
			return true
		}
	case BlockIf:
		// match: (If cond yes no)
		// cond:
		// result: (NE cond yes no)
		for {
			v := b.Control
			_ = v
			cond := b.Control
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSNE
			b.SetControl(cond)
			_ = yes
			_ = no
			return true
		}
	case BlockMIPSLEZ:
		// match: (LEZ (MOVWconst [c]) yes no)
		// cond: int32(c) <= 0
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(int32(c) <= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			_ = yes
			_ = no
			return true
		}
		// match: (LEZ (MOVWconst [c]) yes no)
		// cond: int32(c) >  0
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(int32(c) > 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.swapSuccessors()
			_ = no
			_ = yes
			return true
		}
	case BlockMIPSLTZ:
		// match: (LTZ (MOVWconst [c]) yes no)
		// cond: int32(c) <  0
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(int32(c) < 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			_ = yes
			_ = no
			return true
		}
		// match: (LTZ (MOVWconst [c]) yes no)
		// cond: int32(c) >= 0
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(int32(c) >= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.swapSuccessors()
			_ = no
			_ = yes
			return true
		}
	case BlockMIPSNE:
		// match: (NE (FPFlagTrue cmp) yes no)
		// cond:
		// result: (FPT cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSFPFlagTrue {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSFPT
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (FPFlagFalse cmp) yes no)
		// cond:
		// result: (FPF cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSFPFlagFalse {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSFPF
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGT _ _)) yes no)
		// cond:
		// result: (EQ cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGT {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTU _ _)) yes no)
		// cond:
		// result: (EQ cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTU {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTconst _)) yes no)
		// cond:
		// result: (EQ cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTconst {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTUconst _)) yes no)
		// cond:
		// result: (EQ cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTUconst {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTzero _)) yes no)
		// cond:
		// result: (EQ cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTzero {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTUzero _)) yes no)
		// cond:
		// result: (EQ cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTUzero {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (SGTUconst [1] x) yes no)
		// cond:
		// result: (EQ x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSSGTUconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSEQ
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (SGTUzero x) yes no)
		// cond:
		// result: (NE x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSSGTUzero {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSNE
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (SGTconst [0] x) yes no)
		// cond:
		// result: (LTZ x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSSGTconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSLTZ
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (SGTzero x) yes no)
		// cond:
		// result: (GTZ x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSSGTzero {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPSGTZ
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (MOVWconst [0]) yes no)
		// cond:
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.swapSuccessors()
			_ = no
			_ = yes
			return true
		}
		// match: (NE (MOVWconst [c]) yes no)
		// cond: c != 0
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c != 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			_ = yes
			_ = no
			return true
		}
	}
	return false
}
