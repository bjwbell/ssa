// Code generated from gen/PPC64.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "math"
import "cmd/internal/obj"

var _ = math.MinInt8 // in case not otherwise used
var _ = obj.ANOP     // in case not otherwise used
func rewriteValuePPC64(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValuePPC64_OpAdd16(v)
	case OpAdd32:
		return rewriteValuePPC64_OpAdd32(v)
	case OpAdd32F:
		return rewriteValuePPC64_OpAdd32F(v)
	case OpAdd64:
		return rewriteValuePPC64_OpAdd64(v)
	case OpAdd64F:
		return rewriteValuePPC64_OpAdd64F(v)
	case OpAdd8:
		return rewriteValuePPC64_OpAdd8(v)
	case OpAddPtr:
		return rewriteValuePPC64_OpAddPtr(v)
	case OpAddr:
		return rewriteValuePPC64_OpAddr(v)
	case OpAnd16:
		return rewriteValuePPC64_OpAnd16(v)
	case OpAnd32:
		return rewriteValuePPC64_OpAnd32(v)
	case OpAnd64:
		return rewriteValuePPC64_OpAnd64(v)
	case OpAnd8:
		return rewriteValuePPC64_OpAnd8(v)
	case OpAndB:
		return rewriteValuePPC64_OpAndB(v)
	case OpAtomicAdd32:
		return rewriteValuePPC64_OpAtomicAdd32(v)
	case OpAtomicAdd64:
		return rewriteValuePPC64_OpAtomicAdd64(v)
	case OpAtomicAnd8:
		return rewriteValuePPC64_OpAtomicAnd8(v)
	case OpAtomicCompareAndSwap32:
		return rewriteValuePPC64_OpAtomicCompareAndSwap32(v)
	case OpAtomicCompareAndSwap64:
		return rewriteValuePPC64_OpAtomicCompareAndSwap64(v)
	case OpAtomicExchange32:
		return rewriteValuePPC64_OpAtomicExchange32(v)
	case OpAtomicExchange64:
		return rewriteValuePPC64_OpAtomicExchange64(v)
	case OpAtomicLoad32:
		return rewriteValuePPC64_OpAtomicLoad32(v)
	case OpAtomicLoad64:
		return rewriteValuePPC64_OpAtomicLoad64(v)
	case OpAtomicLoadPtr:
		return rewriteValuePPC64_OpAtomicLoadPtr(v)
	case OpAtomicOr8:
		return rewriteValuePPC64_OpAtomicOr8(v)
	case OpAtomicStore32:
		return rewriteValuePPC64_OpAtomicStore32(v)
	case OpAtomicStore64:
		return rewriteValuePPC64_OpAtomicStore64(v)
	case OpAvg64u:
		return rewriteValuePPC64_OpAvg64u(v)
	case OpClosureCall:
		return rewriteValuePPC64_OpClosureCall(v)
	case OpCom16:
		return rewriteValuePPC64_OpCom16(v)
	case OpCom32:
		return rewriteValuePPC64_OpCom32(v)
	case OpCom64:
		return rewriteValuePPC64_OpCom64(v)
	case OpCom8:
		return rewriteValuePPC64_OpCom8(v)
	case OpConst16:
		return rewriteValuePPC64_OpConst16(v)
	case OpConst32:
		return rewriteValuePPC64_OpConst32(v)
	case OpConst32F:
		return rewriteValuePPC64_OpConst32F(v)
	case OpConst64:
		return rewriteValuePPC64_OpConst64(v)
	case OpConst64F:
		return rewriteValuePPC64_OpConst64F(v)
	case OpConst8:
		return rewriteValuePPC64_OpConst8(v)
	case OpConstBool:
		return rewriteValuePPC64_OpConstBool(v)
	case OpConstNil:
		return rewriteValuePPC64_OpConstNil(v)
	case OpConvert:
		return rewriteValuePPC64_OpConvert(v)
	case OpCvt32Fto32:
		return rewriteValuePPC64_OpCvt32Fto32(v)
	case OpCvt32Fto64:
		return rewriteValuePPC64_OpCvt32Fto64(v)
	case OpCvt32Fto64F:
		return rewriteValuePPC64_OpCvt32Fto64F(v)
	case OpCvt32to32F:
		return rewriteValuePPC64_OpCvt32to32F(v)
	case OpCvt32to64F:
		return rewriteValuePPC64_OpCvt32to64F(v)
	case OpCvt64Fto32:
		return rewriteValuePPC64_OpCvt64Fto32(v)
	case OpCvt64Fto32F:
		return rewriteValuePPC64_OpCvt64Fto32F(v)
	case OpCvt64Fto64:
		return rewriteValuePPC64_OpCvt64Fto64(v)
	case OpCvt64to32F:
		return rewriteValuePPC64_OpCvt64to32F(v)
	case OpCvt64to64F:
		return rewriteValuePPC64_OpCvt64to64F(v)
	case OpDiv16:
		return rewriteValuePPC64_OpDiv16(v)
	case OpDiv16u:
		return rewriteValuePPC64_OpDiv16u(v)
	case OpDiv32:
		return rewriteValuePPC64_OpDiv32(v)
	case OpDiv32F:
		return rewriteValuePPC64_OpDiv32F(v)
	case OpDiv32u:
		return rewriteValuePPC64_OpDiv32u(v)
	case OpDiv64:
		return rewriteValuePPC64_OpDiv64(v)
	case OpDiv64F:
		return rewriteValuePPC64_OpDiv64F(v)
	case OpDiv64u:
		return rewriteValuePPC64_OpDiv64u(v)
	case OpDiv8:
		return rewriteValuePPC64_OpDiv8(v)
	case OpDiv8u:
		return rewriteValuePPC64_OpDiv8u(v)
	case OpEq16:
		return rewriteValuePPC64_OpEq16(v)
	case OpEq32:
		return rewriteValuePPC64_OpEq32(v)
	case OpEq32F:
		return rewriteValuePPC64_OpEq32F(v)
	case OpEq64:
		return rewriteValuePPC64_OpEq64(v)
	case OpEq64F:
		return rewriteValuePPC64_OpEq64F(v)
	case OpEq8:
		return rewriteValuePPC64_OpEq8(v)
	case OpEqB:
		return rewriteValuePPC64_OpEqB(v)
	case OpEqPtr:
		return rewriteValuePPC64_OpEqPtr(v)
	case OpGeq16:
		return rewriteValuePPC64_OpGeq16(v)
	case OpGeq16U:
		return rewriteValuePPC64_OpGeq16U(v)
	case OpGeq32:
		return rewriteValuePPC64_OpGeq32(v)
	case OpGeq32F:
		return rewriteValuePPC64_OpGeq32F(v)
	case OpGeq32U:
		return rewriteValuePPC64_OpGeq32U(v)
	case OpGeq64:
		return rewriteValuePPC64_OpGeq64(v)
	case OpGeq64F:
		return rewriteValuePPC64_OpGeq64F(v)
	case OpGeq64U:
		return rewriteValuePPC64_OpGeq64U(v)
	case OpGeq8:
		return rewriteValuePPC64_OpGeq8(v)
	case OpGeq8U:
		return rewriteValuePPC64_OpGeq8U(v)
	case OpGetClosurePtr:
		return rewriteValuePPC64_OpGetClosurePtr(v)
	case OpGreater16:
		return rewriteValuePPC64_OpGreater16(v)
	case OpGreater16U:
		return rewriteValuePPC64_OpGreater16U(v)
	case OpGreater32:
		return rewriteValuePPC64_OpGreater32(v)
	case OpGreater32F:
		return rewriteValuePPC64_OpGreater32F(v)
	case OpGreater32U:
		return rewriteValuePPC64_OpGreater32U(v)
	case OpGreater64:
		return rewriteValuePPC64_OpGreater64(v)
	case OpGreater64F:
		return rewriteValuePPC64_OpGreater64F(v)
	case OpGreater64U:
		return rewriteValuePPC64_OpGreater64U(v)
	case OpGreater8:
		return rewriteValuePPC64_OpGreater8(v)
	case OpGreater8U:
		return rewriteValuePPC64_OpGreater8U(v)
	case OpHmul32:
		return rewriteValuePPC64_OpHmul32(v)
	case OpHmul32u:
		return rewriteValuePPC64_OpHmul32u(v)
	case OpHmul64:
		return rewriteValuePPC64_OpHmul64(v)
	case OpHmul64u:
		return rewriteValuePPC64_OpHmul64u(v)
	case OpInterCall:
		return rewriteValuePPC64_OpInterCall(v)
	case OpIsInBounds:
		return rewriteValuePPC64_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValuePPC64_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValuePPC64_OpIsSliceInBounds(v)
	case OpLeq16:
		return rewriteValuePPC64_OpLeq16(v)
	case OpLeq16U:
		return rewriteValuePPC64_OpLeq16U(v)
	case OpLeq32:
		return rewriteValuePPC64_OpLeq32(v)
	case OpLeq32F:
		return rewriteValuePPC64_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValuePPC64_OpLeq32U(v)
	case OpLeq64:
		return rewriteValuePPC64_OpLeq64(v)
	case OpLeq64F:
		return rewriteValuePPC64_OpLeq64F(v)
	case OpLeq64U:
		return rewriteValuePPC64_OpLeq64U(v)
	case OpLeq8:
		return rewriteValuePPC64_OpLeq8(v)
	case OpLeq8U:
		return rewriteValuePPC64_OpLeq8U(v)
	case OpLess16:
		return rewriteValuePPC64_OpLess16(v)
	case OpLess16U:
		return rewriteValuePPC64_OpLess16U(v)
	case OpLess32:
		return rewriteValuePPC64_OpLess32(v)
	case OpLess32F:
		return rewriteValuePPC64_OpLess32F(v)
	case OpLess32U:
		return rewriteValuePPC64_OpLess32U(v)
	case OpLess64:
		return rewriteValuePPC64_OpLess64(v)
	case OpLess64F:
		return rewriteValuePPC64_OpLess64F(v)
	case OpLess64U:
		return rewriteValuePPC64_OpLess64U(v)
	case OpLess8:
		return rewriteValuePPC64_OpLess8(v)
	case OpLess8U:
		return rewriteValuePPC64_OpLess8U(v)
	case OpLoad:
		return rewriteValuePPC64_OpLoad(v)
	case OpLsh16x16:
		return rewriteValuePPC64_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValuePPC64_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValuePPC64_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValuePPC64_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValuePPC64_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValuePPC64_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValuePPC64_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValuePPC64_OpLsh32x8(v)
	case OpLsh64x16:
		return rewriteValuePPC64_OpLsh64x16(v)
	case OpLsh64x32:
		return rewriteValuePPC64_OpLsh64x32(v)
	case OpLsh64x64:
		return rewriteValuePPC64_OpLsh64x64(v)
	case OpLsh64x8:
		return rewriteValuePPC64_OpLsh64x8(v)
	case OpLsh8x16:
		return rewriteValuePPC64_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValuePPC64_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValuePPC64_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValuePPC64_OpLsh8x8(v)
	case OpMod16:
		return rewriteValuePPC64_OpMod16(v)
	case OpMod16u:
		return rewriteValuePPC64_OpMod16u(v)
	case OpMod32:
		return rewriteValuePPC64_OpMod32(v)
	case OpMod32u:
		return rewriteValuePPC64_OpMod32u(v)
	case OpMod64:
		return rewriteValuePPC64_OpMod64(v)
	case OpMod64u:
		return rewriteValuePPC64_OpMod64u(v)
	case OpMod8:
		return rewriteValuePPC64_OpMod8(v)
	case OpMod8u:
		return rewriteValuePPC64_OpMod8u(v)
	case OpMove:
		return rewriteValuePPC64_OpMove(v)
	case OpMul16:
		return rewriteValuePPC64_OpMul16(v)
	case OpMul32:
		return rewriteValuePPC64_OpMul32(v)
	case OpMul32F:
		return rewriteValuePPC64_OpMul32F(v)
	case OpMul64:
		return rewriteValuePPC64_OpMul64(v)
	case OpMul64F:
		return rewriteValuePPC64_OpMul64F(v)
	case OpMul8:
		return rewriteValuePPC64_OpMul8(v)
	case OpNeg16:
		return rewriteValuePPC64_OpNeg16(v)
	case OpNeg32:
		return rewriteValuePPC64_OpNeg32(v)
	case OpNeg32F:
		return rewriteValuePPC64_OpNeg32F(v)
	case OpNeg64:
		return rewriteValuePPC64_OpNeg64(v)
	case OpNeg64F:
		return rewriteValuePPC64_OpNeg64F(v)
	case OpNeg8:
		return rewriteValuePPC64_OpNeg8(v)
	case OpNeq16:
		return rewriteValuePPC64_OpNeq16(v)
	case OpNeq32:
		return rewriteValuePPC64_OpNeq32(v)
	case OpNeq32F:
		return rewriteValuePPC64_OpNeq32F(v)
	case OpNeq64:
		return rewriteValuePPC64_OpNeq64(v)
	case OpNeq64F:
		return rewriteValuePPC64_OpNeq64F(v)
	case OpNeq8:
		return rewriteValuePPC64_OpNeq8(v)
	case OpNeqB:
		return rewriteValuePPC64_OpNeqB(v)
	case OpNeqPtr:
		return rewriteValuePPC64_OpNeqPtr(v)
	case OpNilCheck:
		return rewriteValuePPC64_OpNilCheck(v)
	case OpNot:
		return rewriteValuePPC64_OpNot(v)
	case OpOffPtr:
		return rewriteValuePPC64_OpOffPtr(v)
	case OpOr16:
		return rewriteValuePPC64_OpOr16(v)
	case OpOr32:
		return rewriteValuePPC64_OpOr32(v)
	case OpOr64:
		return rewriteValuePPC64_OpOr64(v)
	case OpOr8:
		return rewriteValuePPC64_OpOr8(v)
	case OpOrB:
		return rewriteValuePPC64_OpOrB(v)
	case OpPPC64ADD:
		return rewriteValuePPC64_OpPPC64ADD(v)
	case OpPPC64ADDconst:
		return rewriteValuePPC64_OpPPC64ADDconst(v)
	case OpPPC64AND:
		return rewriteValuePPC64_OpPPC64AND(v)
	case OpPPC64ANDconst:
		return rewriteValuePPC64_OpPPC64ANDconst(v)
	case OpPPC64CMP:
		return rewriteValuePPC64_OpPPC64CMP(v)
	case OpPPC64CMPU:
		return rewriteValuePPC64_OpPPC64CMPU(v)
	case OpPPC64CMPUconst:
		return rewriteValuePPC64_OpPPC64CMPUconst(v)
	case OpPPC64CMPW:
		return rewriteValuePPC64_OpPPC64CMPW(v)
	case OpPPC64CMPWU:
		return rewriteValuePPC64_OpPPC64CMPWU(v)
	case OpPPC64CMPWUconst:
		return rewriteValuePPC64_OpPPC64CMPWUconst(v)
	case OpPPC64CMPWconst:
		return rewriteValuePPC64_OpPPC64CMPWconst(v)
	case OpPPC64CMPconst:
		return rewriteValuePPC64_OpPPC64CMPconst(v)
	case OpPPC64Equal:
		return rewriteValuePPC64_OpPPC64Equal(v)
	case OpPPC64FADD:
		return rewriteValuePPC64_OpPPC64FADD(v)
	case OpPPC64FADDS:
		return rewriteValuePPC64_OpPPC64FADDS(v)
	case OpPPC64FMOVDload:
		return rewriteValuePPC64_OpPPC64FMOVDload(v)
	case OpPPC64FMOVDstore:
		return rewriteValuePPC64_OpPPC64FMOVDstore(v)
	case OpPPC64FMOVSload:
		return rewriteValuePPC64_OpPPC64FMOVSload(v)
	case OpPPC64FMOVSstore:
		return rewriteValuePPC64_OpPPC64FMOVSstore(v)
	case OpPPC64FSUB:
		return rewriteValuePPC64_OpPPC64FSUB(v)
	case OpPPC64FSUBS:
		return rewriteValuePPC64_OpPPC64FSUBS(v)
	case OpPPC64GreaterEqual:
		return rewriteValuePPC64_OpPPC64GreaterEqual(v)
	case OpPPC64GreaterThan:
		return rewriteValuePPC64_OpPPC64GreaterThan(v)
	case OpPPC64LessEqual:
		return rewriteValuePPC64_OpPPC64LessEqual(v)
	case OpPPC64LessThan:
		return rewriteValuePPC64_OpPPC64LessThan(v)
	case OpPPC64MOVBZload:
		return rewriteValuePPC64_OpPPC64MOVBZload(v)
	case OpPPC64MOVBZreg:
		return rewriteValuePPC64_OpPPC64MOVBZreg(v)
	case OpPPC64MOVBreg:
		return rewriteValuePPC64_OpPPC64MOVBreg(v)
	case OpPPC64MOVBstore:
		return rewriteValuePPC64_OpPPC64MOVBstore(v)
	case OpPPC64MOVBstorezero:
		return rewriteValuePPC64_OpPPC64MOVBstorezero(v)
	case OpPPC64MOVDload:
		return rewriteValuePPC64_OpPPC64MOVDload(v)
	case OpPPC64MOVDstore:
		return rewriteValuePPC64_OpPPC64MOVDstore(v)
	case OpPPC64MOVDstorezero:
		return rewriteValuePPC64_OpPPC64MOVDstorezero(v)
	case OpPPC64MOVHZload:
		return rewriteValuePPC64_OpPPC64MOVHZload(v)
	case OpPPC64MOVHZreg:
		return rewriteValuePPC64_OpPPC64MOVHZreg(v)
	case OpPPC64MOVHload:
		return rewriteValuePPC64_OpPPC64MOVHload(v)
	case OpPPC64MOVHreg:
		return rewriteValuePPC64_OpPPC64MOVHreg(v)
	case OpPPC64MOVHstore:
		return rewriteValuePPC64_OpPPC64MOVHstore(v)
	case OpPPC64MOVHstorezero:
		return rewriteValuePPC64_OpPPC64MOVHstorezero(v)
	case OpPPC64MOVWZload:
		return rewriteValuePPC64_OpPPC64MOVWZload(v)
	case OpPPC64MOVWZreg:
		return rewriteValuePPC64_OpPPC64MOVWZreg(v)
	case OpPPC64MOVWload:
		return rewriteValuePPC64_OpPPC64MOVWload(v)
	case OpPPC64MOVWreg:
		return rewriteValuePPC64_OpPPC64MOVWreg(v)
	case OpPPC64MOVWstore:
		return rewriteValuePPC64_OpPPC64MOVWstore(v)
	case OpPPC64MOVWstorezero:
		return rewriteValuePPC64_OpPPC64MOVWstorezero(v)
	case OpPPC64MaskIfNotCarry:
		return rewriteValuePPC64_OpPPC64MaskIfNotCarry(v)
	case OpPPC64NotEqual:
		return rewriteValuePPC64_OpPPC64NotEqual(v)
	case OpPPC64OR:
		return rewriteValuePPC64_OpPPC64OR(v)
	case OpPPC64ORN:
		return rewriteValuePPC64_OpPPC64ORN(v)
	case OpPPC64ORconst:
		return rewriteValuePPC64_OpPPC64ORconst(v)
	case OpPPC64SUB:
		return rewriteValuePPC64_OpPPC64SUB(v)
	case OpPPC64XOR:
		return rewriteValuePPC64_OpPPC64XOR(v)
	case OpPPC64XORconst:
		return rewriteValuePPC64_OpPPC64XORconst(v)
	case OpRound32F:
		return rewriteValuePPC64_OpRound32F(v)
	case OpRound64F:
		return rewriteValuePPC64_OpRound64F(v)
	case OpRsh16Ux16:
		return rewriteValuePPC64_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValuePPC64_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValuePPC64_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValuePPC64_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValuePPC64_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValuePPC64_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValuePPC64_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValuePPC64_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValuePPC64_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValuePPC64_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValuePPC64_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValuePPC64_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValuePPC64_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValuePPC64_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValuePPC64_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValuePPC64_OpRsh32x8(v)
	case OpRsh64Ux16:
		return rewriteValuePPC64_OpRsh64Ux16(v)
	case OpRsh64Ux32:
		return rewriteValuePPC64_OpRsh64Ux32(v)
	case OpRsh64Ux64:
		return rewriteValuePPC64_OpRsh64Ux64(v)
	case OpRsh64Ux8:
		return rewriteValuePPC64_OpRsh64Ux8(v)
	case OpRsh64x16:
		return rewriteValuePPC64_OpRsh64x16(v)
	case OpRsh64x32:
		return rewriteValuePPC64_OpRsh64x32(v)
	case OpRsh64x64:
		return rewriteValuePPC64_OpRsh64x64(v)
	case OpRsh64x8:
		return rewriteValuePPC64_OpRsh64x8(v)
	case OpRsh8Ux16:
		return rewriteValuePPC64_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValuePPC64_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValuePPC64_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValuePPC64_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValuePPC64_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValuePPC64_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValuePPC64_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValuePPC64_OpRsh8x8(v)
	case OpSignExt16to32:
		return rewriteValuePPC64_OpSignExt16to32(v)
	case OpSignExt16to64:
		return rewriteValuePPC64_OpSignExt16to64(v)
	case OpSignExt32to64:
		return rewriteValuePPC64_OpSignExt32to64(v)
	case OpSignExt8to16:
		return rewriteValuePPC64_OpSignExt8to16(v)
	case OpSignExt8to32:
		return rewriteValuePPC64_OpSignExt8to32(v)
	case OpSignExt8to64:
		return rewriteValuePPC64_OpSignExt8to64(v)
	case OpSlicemask:
		return rewriteValuePPC64_OpSlicemask(v)
	case OpSqrt:
		return rewriteValuePPC64_OpSqrt(v)
	case OpStaticCall:
		return rewriteValuePPC64_OpStaticCall(v)
	case OpStore:
		return rewriteValuePPC64_OpStore(v)
	case OpSub16:
		return rewriteValuePPC64_OpSub16(v)
	case OpSub32:
		return rewriteValuePPC64_OpSub32(v)
	case OpSub32F:
		return rewriteValuePPC64_OpSub32F(v)
	case OpSub64:
		return rewriteValuePPC64_OpSub64(v)
	case OpSub64F:
		return rewriteValuePPC64_OpSub64F(v)
	case OpSub8:
		return rewriteValuePPC64_OpSub8(v)
	case OpSubPtr:
		return rewriteValuePPC64_OpSubPtr(v)
	case OpTrunc16to8:
		return rewriteValuePPC64_OpTrunc16to8(v)
	case OpTrunc32to16:
		return rewriteValuePPC64_OpTrunc32to16(v)
	case OpTrunc32to8:
		return rewriteValuePPC64_OpTrunc32to8(v)
	case OpTrunc64to16:
		return rewriteValuePPC64_OpTrunc64to16(v)
	case OpTrunc64to32:
		return rewriteValuePPC64_OpTrunc64to32(v)
	case OpTrunc64to8:
		return rewriteValuePPC64_OpTrunc64to8(v)
	case OpXor16:
		return rewriteValuePPC64_OpXor16(v)
	case OpXor32:
		return rewriteValuePPC64_OpXor32(v)
	case OpXor64:
		return rewriteValuePPC64_OpXor64(v)
	case OpXor8:
		return rewriteValuePPC64_OpXor8(v)
	case OpZero:
		return rewriteValuePPC64_OpZero(v)
	case OpZeroExt16to32:
		return rewriteValuePPC64_OpZeroExt16to32(v)
	case OpZeroExt16to64:
		return rewriteValuePPC64_OpZeroExt16to64(v)
	case OpZeroExt32to64:
		return rewriteValuePPC64_OpZeroExt32to64(v)
	case OpZeroExt8to16:
		return rewriteValuePPC64_OpZeroExt8to16(v)
	case OpZeroExt8to32:
		return rewriteValuePPC64_OpZeroExt8to32(v)
	case OpZeroExt8to64:
		return rewriteValuePPC64_OpZeroExt8to64(v)
	}
	return false
}
func rewriteValuePPC64_OpAdd16(v *Value) bool {
	// match: (Add16 x y)
	// cond:
	// result: (ADD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAdd32(v *Value) bool {
	// match: (Add32 x y)
	// cond:
	// result: (ADD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAdd32F(v *Value) bool {
	// match: (Add32F x y)
	// cond:
	// result: (FADDS x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FADDS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAdd64(v *Value) bool {
	// match: (Add64 x y)
	// cond:
	// result: (ADD  x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAdd64F(v *Value) bool {
	// match: (Add64F x y)
	// cond:
	// result: (FADD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAdd8(v *Value) bool {
	// match: (Add8 x y)
	// cond:
	// result: (ADD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAddPtr(v *Value) bool {
	// match: (AddPtr x y)
	// cond:
	// result: (ADD  x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAddr(v *Value) bool {
	// match: (Addr {sym} base)
	// cond:
	// result: (MOVDaddr {sym} base)
	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpPPC64MOVDaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValuePPC64_OpAnd16(v *Value) bool {
	// match: (And16 x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAnd32(v *Value) bool {
	// match: (And32 x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAnd64(v *Value) bool {
	// match: (And64 x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAnd8(v *Value) bool {
	// match: (And8 x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAndB(v *Value) bool {
	// match: (AndB x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAtomicAdd32(v *Value) bool {
	// match: (AtomicAdd32 ptr val mem)
	// cond:
	// result: (LoweredAtomicAdd32 ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicAdd32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicAdd64(v *Value) bool {
	// match: (AtomicAdd64 ptr val mem)
	// cond:
	// result: (LoweredAtomicAdd64 ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicAdd64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicAnd8(v *Value) bool {
	// match: (AtomicAnd8 ptr val mem)
	// cond:
	// result: (LoweredAtomicAnd8 ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicAnd8)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicCompareAndSwap32(v *Value) bool {
	// match: (AtomicCompareAndSwap32 ptr old new_ mem)
	// cond:
	// result: (LoweredAtomicCas32 ptr old new_ mem)
	for {
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpPPC64LoweredAtomicCas32)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicCompareAndSwap64(v *Value) bool {
	// match: (AtomicCompareAndSwap64 ptr old new_ mem)
	// cond:
	// result: (LoweredAtomicCas64 ptr old new_ mem)
	for {
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpPPC64LoweredAtomicCas64)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicExchange32(v *Value) bool {
	// match: (AtomicExchange32 ptr val mem)
	// cond:
	// result: (LoweredAtomicExchange32 ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicExchange32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicExchange64(v *Value) bool {
	// match: (AtomicExchange64 ptr val mem)
	// cond:
	// result: (LoweredAtomicExchange64 ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicExchange64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicLoad32(v *Value) bool {
	// match: (AtomicLoad32 ptr mem)
	// cond:
	// result: (LoweredAtomicLoad32 ptr mem)
	for {
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64LoweredAtomicLoad32)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicLoad64(v *Value) bool {
	// match: (AtomicLoad64 ptr mem)
	// cond:
	// result: (LoweredAtomicLoad64 ptr mem)
	for {
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64LoweredAtomicLoad64)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicLoadPtr(v *Value) bool {
	// match: (AtomicLoadPtr ptr mem)
	// cond:
	// result: (LoweredAtomicLoadPtr ptr mem)
	for {
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64LoweredAtomicLoadPtr)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicOr8(v *Value) bool {
	// match: (AtomicOr8 ptr val mem)
	// cond:
	// result: (LoweredAtomicOr8  ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicOr8)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicStore32(v *Value) bool {
	// match: (AtomicStore32 ptr val mem)
	// cond:
	// result: (LoweredAtomicStore32 ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicStore32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicStore64(v *Value) bool {
	// match: (AtomicStore64 ptr val mem)
	// cond:
	// result: (LoweredAtomicStore64 ptr val mem)
	for {
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicStore64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAvg64u(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Avg64u <t> x y)
	// cond:
	// result: (ADD (SRDconst <t> (SUB <t> x y) [1]) y)
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v0 := b.NewValue0(v.Pos, OpPPC64SRDconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpPPC64SUB, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpClosureCall(v *Value) bool {
	// match: (ClosureCall [argwid] entry closure mem)
	// cond:
	// result: (CALLclosure [argwid] entry closure mem)
	for {
		argwid := v.AuxInt
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64CALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpCom16(v *Value) bool {
	// match: (Com16 x)
	// cond:
	// result: (NOR x x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCom32(v *Value) bool {
	// match: (Com32 x)
	// cond:
	// result: (NOR x x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCom64(v *Value) bool {
	// match: (Com64 x)
	// cond:
	// result: (NOR x x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCom8(v *Value) bool {
	// match: (Com8 x)
	// cond:
	// result: (NOR x x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpConst16(v *Value) bool {
	// match: (Const16 [val])
	// cond:
	// result: (MOVDconst [val])
	for {
		val := v.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConst32(v *Value) bool {
	// match: (Const32 [val])
	// cond:
	// result: (MOVDconst [val])
	for {
		val := v.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConst32F(v *Value) bool {
	// match: (Const32F [val])
	// cond:
	// result: (FMOVSconst [val])
	for {
		val := v.AuxInt
		v.reset(OpPPC64FMOVSconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConst64(v *Value) bool {
	// match: (Const64 [val])
	// cond:
	// result: (MOVDconst [val])
	for {
		val := v.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConst64F(v *Value) bool {
	// match: (Const64F [val])
	// cond:
	// result: (FMOVDconst [val])
	for {
		val := v.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConst8(v *Value) bool {
	// match: (Const8 [val])
	// cond:
	// result: (MOVDconst [val])
	for {
		val := v.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConstBool(v *Value) bool {
	// match: (ConstBool [b])
	// cond:
	// result: (MOVDconst [b])
	for {
		b := v.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValuePPC64_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// cond:
	// result: (MOVDconst [0])
	for {
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValuePPC64_OpConvert(v *Value) bool {
	// match: (Convert <t> x mem)
	// cond:
	// result: (MOVDconvert <t> x mem)
	for {
		t := v.Type
		x := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVDconvert)
		v.Type = t
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpCvt32Fto32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Cvt32Fto32 x)
	// cond:
	// result: (Xf2i64 (FCTIWZ x))
	for {
		x := v.Args[0]
		v.reset(OpPPC64Xf2i64)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIWZ, types.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt32Fto64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Cvt32Fto64 x)
	// cond:
	// result: (Xf2i64 (FCTIDZ x))
	for {
		x := v.Args[0]
		v.reset(OpPPC64Xf2i64)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIDZ, types.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt32Fto64F(v *Value) bool {
	// match: (Cvt32Fto64F x)
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
func rewriteValuePPC64_OpCvt32to32F(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Cvt32to32F x)
	// cond:
	// result: (FRSP (FCFID (Xi2f64 (SignExt32to64 x))))
	for {
		x := v.Args[0]
		v.reset(OpPPC64FRSP)
		v0 := b.NewValue0(v.Pos, OpPPC64FCFID, types.Float64)
		v1 := b.NewValue0(v.Pos, OpPPC64Xi2f64, types.Float64)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt32to64F(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Cvt32to64F x)
	// cond:
	// result: (FCFID (Xi2f64 (SignExt32to64 x)))
	for {
		x := v.Args[0]
		v.reset(OpPPC64FCFID)
		v0 := b.NewValue0(v.Pos, OpPPC64Xi2f64, types.Float64)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64Fto32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Cvt64Fto32 x)
	// cond:
	// result: (Xf2i64 (FCTIWZ x))
	for {
		x := v.Args[0]
		v.reset(OpPPC64Xf2i64)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIWZ, types.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64Fto32F(v *Value) bool {
	// match: (Cvt64Fto32F x)
	// cond:
	// result: (FRSP x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64FRSP)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCvt64Fto64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Cvt64Fto64 x)
	// cond:
	// result: (Xf2i64 (FCTIDZ x))
	for {
		x := v.Args[0]
		v.reset(OpPPC64Xf2i64)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIDZ, types.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64to32F(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Cvt64to32F x)
	// cond:
	// result: (FRSP (FCFID (Xi2f64 x)))
	for {
		x := v.Args[0]
		v.reset(OpPPC64FRSP)
		v0 := b.NewValue0(v.Pos, OpPPC64FCFID, types.Float64)
		v1 := b.NewValue0(v.Pos, OpPPC64Xi2f64, types.Float64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64to64F(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Cvt64to64F x)
	// cond:
	// result: (FCFID (Xi2f64 x))
	for {
		x := v.Args[0]
		v.reset(OpPPC64FCFID)
		v0 := b.NewValue0(v.Pos, OpPPC64Xi2f64, types.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpDiv16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div16 x y)
	// cond:
	// result: (DIVW  (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpDiv16u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div16u x y)
	// cond:
	// result: (DIVWU (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVWU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpDiv32(v *Value) bool {
	// match: (Div32 x y)
	// cond:
	// result: (DIVW  x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv32F(v *Value) bool {
	// match: (Div32F x y)
	// cond:
	// result: (FDIVS x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FDIVS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv32u(v *Value) bool {
	// match: (Div32u x y)
	// cond:
	// result: (DIVWU x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv64(v *Value) bool {
	// match: (Div64 x y)
	// cond:
	// result: (DIVD  x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv64F(v *Value) bool {
	// match: (Div64F x y)
	// cond:
	// result: (FDIV x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FDIV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv64u(v *Value) bool {
	// match: (Div64u x y)
	// cond:
	// result: (DIVDU x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVDU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div8 x y)
	// cond:
	// result: (DIVW  (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpDiv8u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div8u x y)
	// cond:
	// result: (DIVWU (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVWU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpEq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Eq16 x y)
	// cond: isSigned(x.Type) && isSigned(y.Type)
	// result: (Equal (CMPW (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		if !(isSigned(x.Type) && isSigned(y.Type)) {
			break
		}
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
	// match: (Eq16 x y)
	// cond:
	// result: (Equal (CMPW (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpEq32(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Eq32 x y)
	// cond:
	// result: (Equal (CMPW x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Eq32F x y)
	// cond:
	// result: (Equal (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEq64(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Eq64 x y)
	// cond:
	// result: (Equal (CMP x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Eq64F x y)
	// cond:
	// result: (Equal (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Eq8 x y)
	// cond: isSigned(x.Type) && isSigned(y.Type)
	// result: (Equal (CMPW (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		if !(isSigned(x.Type) && isSigned(y.Type)) {
			break
		}
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
	// match: (Eq8 x y)
	// cond:
	// result: (Equal (CMPW (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpEqB(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (EqB x y)
	// cond:
	// result: (ANDconst [1] (EQV x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpPPC64EQV, types.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEqPtr(v *Value) bool {
	b := v.Block
	_ = b
	// match: (EqPtr x y)
	// cond:
	// result: (Equal (CMP x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGeq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq16 x y)
	// cond:
	// result: (GreaterEqual (CMPW (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpGeq16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq16U x y)
	// cond:
	// result: (GreaterEqual (CMPWU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
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
func rewriteValuePPC64_OpGeq32(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Geq32 x y)
	// cond:
	// result: (GreaterEqual (CMPW x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGeq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Geq32F x y)
	// cond:
	// result: (FGreaterEqual (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGeq32U(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Geq32U x y)
	// cond:
	// result: (GreaterEqual (CMPWU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGeq64(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Geq64 x y)
	// cond:
	// result: (GreaterEqual (CMP x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGeq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Geq64F x y)
	// cond:
	// result: (FGreaterEqual (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGeq64U(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Geq64U x y)
	// cond:
	// result: (GreaterEqual (CMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGeq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq8 x y)
	// cond:
	// result: (GreaterEqual (CMPW (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpGeq8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq8U x y)
	// cond:
	// result: (GreaterEqual (CMPWU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
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
func rewriteValuePPC64_OpGetClosurePtr(v *Value) bool {
	// match: (GetClosurePtr)
	// cond:
	// result: (LoweredGetClosurePtr)
	for {
		v.reset(OpPPC64LoweredGetClosurePtr)
		return true
	}
}
func rewriteValuePPC64_OpGreater16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater16 x y)
	// cond:
	// result: (GreaterThan (CMPW (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpGreater16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater16U x y)
	// cond:
	// result: (GreaterThan (CMPWU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
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
func rewriteValuePPC64_OpGreater32(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Greater32 x y)
	// cond:
	// result: (GreaterThan (CMPW x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGreater32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Greater32F x y)
	// cond:
	// result: (FGreaterThan (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FGreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGreater32U(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Greater32U x y)
	// cond:
	// result: (GreaterThan (CMPWU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGreater64(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Greater64 x y)
	// cond:
	// result: (GreaterThan (CMP x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGreater64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Greater64F x y)
	// cond:
	// result: (FGreaterThan (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FGreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGreater64U(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Greater64U x y)
	// cond:
	// result: (GreaterThan (CMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGreater8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater8 x y)
	// cond:
	// result: (GreaterThan (CMPW (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpGreater8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater8U x y)
	// cond:
	// result: (GreaterThan (CMPWU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
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
func rewriteValuePPC64_OpHmul32(v *Value) bool {
	// match: (Hmul32 x y)
	// cond:
	// result: (MULHW  x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULHW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpHmul32u(v *Value) bool {
	// match: (Hmul32u x y)
	// cond:
	// result: (MULHWU x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULHWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpHmul64(v *Value) bool {
	// match: (Hmul64 x y)
	// cond:
	// result: (MULHD  x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULHD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpHmul64u(v *Value) bool {
	// match: (Hmul64u x y)
	// cond:
	// result: (MULHDU x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULHDU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpInterCall(v *Value) bool {
	// match: (InterCall [argwid] entry mem)
	// cond:
	// result: (CALLinter [argwid] entry mem)
	for {
		argwid := v.AuxInt
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64CALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpIsInBounds(v *Value) bool {
	b := v.Block
	_ = b
	// match: (IsInBounds idx len)
	// cond:
	// result: (LessThan (CMPU idx len))
	for {
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpIsNonNil(v *Value) bool {
	b := v.Block
	_ = b
	// match: (IsNonNil ptr)
	// cond:
	// result: (NotEqual (CMPconst [0] ptr))
	for {
		ptr := v.Args[0]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPconst, TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(ptr)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpIsSliceInBounds(v *Value) bool {
	b := v.Block
	_ = b
	// match: (IsSliceInBounds idx len)
	// cond:
	// result: (LessEqual (CMPU idx len))
	for {
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq16 x y)
	// cond:
	// result: (LessEqual (CMPW (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpLeq16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq16U x y)
	// cond:
	// result: (LessEqual (CMPWU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
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
func rewriteValuePPC64_OpLeq32(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Leq32 x y)
	// cond:
	// result: (LessEqual (CMPW x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Leq32F x y)
	// cond:
	// result: (FLessEqual (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FLessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq32U(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Leq32U x y)
	// cond:
	// result: (LessEqual (CMPWU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq64(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Leq64 x y)
	// cond:
	// result: (LessEqual (CMP x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Leq64F x y)
	// cond:
	// result: (FLessEqual (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FLessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq64U(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Leq64U x y)
	// cond:
	// result: (LessEqual (CMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLeq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq8 x y)
	// cond:
	// result: (LessEqual (CMPW (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpLeq8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq8U x y)
	// cond:
	// result: (LessEqual (CMPWU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
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
func rewriteValuePPC64_OpLess16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less16 x y)
	// cond:
	// result: (LessThan (CMPW (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpLess16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less16U x y)
	// cond:
	// result: (LessThan (CMPWU (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
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
func rewriteValuePPC64_OpLess32(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Less32 x y)
	// cond:
	// result: (LessThan (CMPW x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Less32F x y)
	// cond:
	// result: (FLessThan (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FLessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess32U(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Less32U x y)
	// cond:
	// result: (LessThan (CMPWU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess64(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Less64 x y)
	// cond:
	// result: (LessThan (CMP x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Less64F x y)
	// cond:
	// result: (FLessThan (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FLessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess64U(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Less64U x y)
	// cond:
	// result: (LessThan (CMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLess8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less8 x y)
	// cond:
	// result: (LessThan (CMPW (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpLess8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less8U x y)
	// cond:
	// result: (LessThan (CMPWU (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, TypeFlags)
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
func rewriteValuePPC64_OpLoad(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Load <t> ptr mem)
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpPPC64MOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitInt(t) && isSigned(t)
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is32BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitInt(t) && !isSigned(t)
	// result: (MOVWZload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is32BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVWZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is16BitInt(t) && isSigned(t)
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVHload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is16BitInt(t) && !isSigned(t)
	// result: (MOVHZload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVHZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVBZload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is8BitInt(t) && isSigned(t)
	// result: (MOVBreg (MOVBZload ptr mem))
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVBreg)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, types.UInt8)
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is8BitInt(t) && !isSigned(t)
	// result: (MOVBZload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (FMOVSload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpPPC64FMOVSload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (FMOVDload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpPPC64FMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpLsh16x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh16x16 x y)
	// cond:
	// result: (SLW  x                 (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -16
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh16x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh16x32 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SLWconst x [c])
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
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x32 x (MOVDconst [c]))
	// cond: uint32(c) < 16
	// result: (SLWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x32 x y)
	// cond:
	// result: (SLW  x                 (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -16
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh16x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SLWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (MOVDconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh16x64 x (MOVDconst [c]))
	// cond: uint64(c) < 16
	// result: (SLWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 x y)
	// cond:
	// result: (SLW  x                 (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -16
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh16x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh16x8 x y)
	// cond:
	// result: (SLW  x                 (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -16
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh32x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh32x16 x y)
	// cond:
	// result: (SLW x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh32x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh32x32 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SLWconst x [c])
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
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x32 x (MOVDconst [c]))
	// cond: uint32(c) < 32
	// result: (SLWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x32 x y)
	// cond:
	// result: (SLW x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh32x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SLWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (MOVDconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh32x64 x (MOVDconst [c]))
	// cond: uint64(c) < 32
	// result: (SLWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 x y)
	// cond:
	// result: (SLW  x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh32x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh32x8 x y)
	// cond:
	// result: (SLW x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh64x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh64x16 x y)
	// cond:
	// result: (SLD x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh64x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh64x32 x (Const64 [c]))
	// cond: uint32(c) < 64
	// result: (SLDconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x32 x (MOVDconst [c]))
	// cond: uint32(c) < 64
	// result: (SLDconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x32 x y)
	// cond:
	// result: (SLD x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh64x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh64x64 x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SLDconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh64x64 x (MOVDconst [c]))
	// cond: uint64(c) < 64
	// result: (SLDconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x64 x y)
	// cond:
	// result: (SLD  x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh64x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh64x8 x y)
	// cond:
	// result: (SLD x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh8x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh8x16 x y)
	// cond:
	// result: (SLW  x                (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -8
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh8x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh8x32 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SLWconst x [c])
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
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x32 x (MOVDconst [c]))
	// cond: uint32(c) < 8
	// result: (SLWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x32 x y)
	// cond:
	// result: (SLW  x                (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -8
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh8x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SLWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (MOVDconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh8x64 x (MOVDconst [c]))
	// cond: uint64(c) < 8
	// result: (SLWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 x y)
	// cond:
	// result: (SLW  x                (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -8
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpLsh8x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh8x8 x y)
	// cond:
	// result: (SLW  x                (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -8
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod16 x y)
	// cond:
	// result: (Mod32 (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpMod16u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod16u x y)
	// cond:
	// result: (Mod32u (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpMod32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod32 x y)
	// cond:
	// result: (SUB x (MULLW y (DIVW x y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLW, types.Int32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVW, types.Int32)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod32u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod32u x y)
	// cond:
	// result: (SUB x (MULLW y (DIVWU x y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLW, types.Int32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVWU, types.Int32)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod64 x y)
	// cond:
	// result: (SUB x (MULLD y (DIVD x y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLD, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVD, types.Int64)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod64u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod64u x y)
	// cond:
	// result: (SUB x (MULLD y (DIVDU x y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLD, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVDU, types.Int64)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod8 x y)
	// cond:
	// result: (Mod32 (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpMod8u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod8u x y)
	// cond:
	// result: (Mod32u (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpMove(v *Value) bool {
	b := v.Block
	_ = b
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
	// result: (MOVBstore dst (MOVBZload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, types.UInt8)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// cond:
	// result: (MOVHstore dst (MOVHZload src mem) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, types.UInt16)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
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
		v.reset(OpPPC64MOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWload, types.Int32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore [2] dst (MOVHZload [2] src mem) 		(MOVHstore dst (MOVHZload src mem) mem))
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
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, types.UInt16)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVHZload, types.UInt16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [4] dst src mem)
	// cond:
	// result: (MOVBstore [3] dst (MOVBZload [3] src mem) 		(MOVBstore [2] dst (MOVBZload [2] src mem) 			(MOVBstore [1] dst (MOVBZload [1] src mem) 				(MOVBstore dst (MOVBZload src mem) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, types.UInt8)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVBstore, TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVBZload, types.UInt8)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpPPC64MOVBstore, TypeMem)
		v3.AuxInt = 1
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpPPC64MOVBZload, types.UInt8)
		v4.AuxInt = 1
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpPPC64MOVBstore, TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpPPC64MOVBZload, types.UInt8)
		v6.AddArg(src)
		v6.AddArg(mem)
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(Type).Alignment()%8 == 0
	// result: (MOVDstore dst (MOVDload src mem) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, types.Int64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore [4] dst (MOVWZload [4] src mem) 		(MOVWstore dst (MOVWZload src mem) mem))
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
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, types.UInt32)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVWZload, types.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore [6] dst (MOVHZload [6] src mem) 		(MOVHstore [4] dst (MOVHZload [4] src mem) 			(MOVHstore [2] dst (MOVHZload [2] src mem) 				(MOVHstore dst (MOVHZload src mem) mem))))
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
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, types.UInt16)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHstore, TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVHZload, types.UInt16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpPPC64MOVHstore, TypeMem)
		v3.AuxInt = 2
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpPPC64MOVHZload, types.UInt16)
		v4.AuxInt = 2
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpPPC64MOVHstore, TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpPPC64MOVHZload, types.UInt16)
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
	// result: (MOVBstore [2] dst (MOVBZload [2] src mem)                 (MOVHstore dst (MOVHload src mem) mem))
	for {
		if v.AuxInt != 3 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, types.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVHload, types.Int16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [4] dst src mem)
	// cond:
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWload, types.Int32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [5] dst src mem)
	// cond:
	// result: (MOVBstore [4] dst (MOVBZload [4] src mem)                 (MOVWstore dst (MOVWload src mem) mem))
	for {
		if v.AuxInt != 5 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, types.UInt8)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVWload, types.Int32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [6] dst src mem)
	// cond:
	// result: (MOVHstore [4] dst (MOVHZload [4] src mem)                 (MOVWstore dst (MOVWload src mem) mem))
	for {
		if v.AuxInt != 6 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, types.UInt16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVWload, types.Int32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [7] dst src mem)
	// cond:
	// result: (MOVBstore [6] dst (MOVBZload [6] src mem)                 (MOVHstore [4] dst (MOVHZload [4] src mem)                         (MOVWstore dst (MOVWload src mem) mem)))
	for {
		if v.AuxInt != 7 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, types.UInt8)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHstore, TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVHZload, types.UInt16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpPPC64MOVWstore, TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpPPC64MOVWload, types.Int32)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [8] dst src mem)
	// cond:
	// result: (MOVDstore dst (MOVDload src mem) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVDstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, types.Int64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 8
	// result: (LoweredMove [s] dst src mem)
	for {
		s := v.AuxInt
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 8) {
			break
		}
		v.reset(OpPPC64LoweredMove)
		v.AuxInt = s
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpMul16(v *Value) bool {
	// match: (Mul16 x y)
	// cond:
	// result: (MULLW x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpMul32(v *Value) bool {
	// match: (Mul32 x y)
	// cond:
	// result: (MULLW  x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpMul32F(v *Value) bool {
	// match: (Mul32F x y)
	// cond:
	// result: (FMULS x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FMULS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpMul64(v *Value) bool {
	// match: (Mul64 x y)
	// cond:
	// result: (MULLD  x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpMul64F(v *Value) bool {
	// match: (Mul64F x y)
	// cond:
	// result: (FMUL x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpMul8(v *Value) bool {
	// match: (Mul8 x y)
	// cond:
	// result: (MULLW x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpNeg16(v *Value) bool {
	// match: (Neg16 x)
	// cond:
	// result: (NEG x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpNeg32(v *Value) bool {
	// match: (Neg32 x)
	// cond:
	// result: (NEG x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpNeg32F(v *Value) bool {
	// match: (Neg32F x)
	// cond:
	// result: (FNEG x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64FNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpNeg64(v *Value) bool {
	// match: (Neg64 x)
	// cond:
	// result: (NEG x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpNeg64F(v *Value) bool {
	// match: (Neg64F x)
	// cond:
	// result: (FNEG x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64FNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpNeg8(v *Value) bool {
	// match: (Neg8 x)
	// cond:
	// result: (NEG x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpNeq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Neq16 x y)
	// cond: isSigned(x.Type) && isSigned(y.Type)
	// result: (NotEqual (CMPW (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		if !(isSigned(x.Type) && isSigned(y.Type)) {
			break
		}
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
	// match: (Neq16 x y)
	// cond:
	// result: (NotEqual (CMPW (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpNeq32(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Neq32 x y)
	// cond:
	// result: (NotEqual (CMPW x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNeq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Neq32F x y)
	// cond:
	// result: (NotEqual (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNeq64(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Neq64 x y)
	// cond:
	// result: (NotEqual (CMP x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNeq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Neq64F x y)
	// cond:
	// result: (NotEqual (FCMPU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNeq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Neq8 x y)
	// cond: isSigned(x.Type) && isSigned(y.Type)
	// result: (NotEqual (CMPW (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		if !(isSigned(x.Type) && isSigned(y.Type)) {
			break
		}
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
	// match: (Neq8 x y)
	// cond:
	// result: (NotEqual (CMPW (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, TypeFlags)
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
func rewriteValuePPC64_OpNeqB(v *Value) bool {
	// match: (NeqB x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpNeqPtr(v *Value) bool {
	b := v.Block
	_ = b
	// match: (NeqPtr x y)
	// cond:
	// result: (NotEqual (CMP x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNilCheck(v *Value) bool {
	// match: (NilCheck ptr mem)
	// cond:
	// result: (LoweredNilCheck ptr mem)
	for {
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64LoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpNot(v *Value) bool {
	// match: (Not x)
	// cond:
	// result: (XORconst [1] x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64XORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpOffPtr(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (OffPtr [off] ptr)
	// cond:
	// result: (ADD (MOVDconst <types.Int64> [off]) ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpPPC64ADD)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, types.Int64)
		v0.AuxInt = off
		v.AddArg(v0)
		v.AddArg(ptr)
		return true
	}
}
func rewriteValuePPC64_OpOr16(v *Value) bool {
	// match: (Or16 x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpOr32(v *Value) bool {
	// match: (Or32 x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpOr64(v *Value) bool {
	// match: (Or64 x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpOr8(v *Value) bool {
	// match: (Or8 x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpOrB(v *Value) bool {
	// match: (OrB x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpPPC64ADD(v *Value) bool {
	// match: (ADD x (MOVDconst [c]))
	// cond: is32Bit(c)
	// result: (ADDconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpPPC64ADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (ADD (MOVDconst [c]) x)
	// cond: is32Bit(c)
	// result: (ADDconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpPPC64ADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64ADDconst(v *Value) bool {
	// match: (ADDconst [c] (ADDconst [d] x))
	// cond: is32Bit(c+d)
	// result: (ADDconst [c+d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpPPC64ADDconst)
		v.AuxInt = c + d
		v.AddArg(x)
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
	// match: (ADDconst [c] (MOVDaddr [d] {sym} x))
	// cond:
	// result: (MOVDaddr [c+d] {sym} x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		d := v_0.AuxInt
		sym := v_0.Aux
		x := v_0.Args[0]
		v.reset(OpPPC64MOVDaddr)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64AND(v *Value) bool {
	// match: (AND x (NOR y y))
	// cond:
	// result: (ANDN x y)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64NOR {
			break
		}
		y := v_1.Args[0]
		if y != v_1.Args[1] {
			break
		}
		v.reset(OpPPC64ANDN)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND (NOR y y) x)
	// cond:
	// result: (ANDN x y)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64NOR {
			break
		}
		y := v_0.Args[0]
		if y != v_0.Args[1] {
			break
		}
		x := v.Args[1]
		v.reset(OpPPC64ANDN)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (AND (MOVDconst [c]) (MOVDconst [d]))
	// cond:
	// result: (MOVDconst [c&d])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c & d
		return true
	}
	// match: (AND (MOVDconst [d]) (MOVDconst [c]))
	// cond:
	// result: (MOVDconst [c&d])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c & d
		return true
	}
	// match: (AND x (MOVDconst [c]))
	// cond: isU16Bit(c)
	// result: (ANDconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (AND (MOVDconst [c]) x)
	// cond: isU16Bit(c)
	// result: (ANDconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (AND (MOVDconst [c]) x:(MOVBZload _ _))
	// cond:
	// result: (ANDconst [c&0xFF] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if x.Op != OpPPC64MOVBZload {
			break
		}
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}
	// match: (AND x:(MOVBZload _ _) (MOVDconst [c]))
	// cond:
	// result: (ANDconst [c&0xFF] x)
	for {
		x := v.Args[0]
		if x.Op != OpPPC64MOVBZload {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}
	// match: (AND x:(MOVBZload _ _) (MOVDconst [c]))
	// cond:
	// result: (ANDconst [c&0xFF] x)
	for {
		x := v.Args[0]
		if x.Op != OpPPC64MOVBZload {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}
	// match: (AND (MOVDconst [c]) x:(MOVBZload _ _))
	// cond:
	// result: (ANDconst [c&0xFF] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if x.Op != OpPPC64MOVBZload {
			break
		}
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64ANDconst(v *Value) bool {
	// match: (ANDconst [c] (ANDconst [d] x))
	// cond:
	// result: (ANDconst [c&d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
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
	// match: (ANDconst [0] _)
	// cond:
	// result: (MOVDconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (ANDconst [c] y:(MOVBZreg _))
	// cond: c&0xFF == 0xFF
	// result: y
	for {
		c := v.AuxInt
		y := v.Args[0]
		if y.Op != OpPPC64MOVBZreg {
			break
		}
		if !(c&0xFF == 0xFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ANDconst [c] y:(MOVHZreg _))
	// cond: c&0xFFFF == 0xFFFF
	// result: y
	for {
		c := v.AuxInt
		y := v.Args[0]
		if y.Op != OpPPC64MOVHZreg {
			break
		}
		if !(c&0xFFFF == 0xFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ANDconst [c] y:(MOVWZreg _))
	// cond: c&0xFFFFFFFF == 0xFFFFFFFF
	// result: y
	for {
		c := v.AuxInt
		y := v.Args[0]
		if y.Op != OpPPC64MOVWZreg {
			break
		}
		if !(c&0xFFFFFFFF == 0xFFFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (ANDconst [c] (MOVBZreg x))
	// cond:
	// result: (ANDconst [c&0xFF] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVHZreg x))
	// cond:
	// result: (ANDconst [c&0xFFFF] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFFFF
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVWZreg x))
	// cond:
	// result: (ANDconst [c&0xFFFFFFFF] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFFFFFFFF
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMP(v *Value) bool {
	b := v.Block
	_ = b
	// match: (CMP x (MOVDconst [c]))
	// cond: is16Bit(c)
	// result: (CMPconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMP (MOVDconst [c]) y)
	// cond: is16Bit(c)
	// result: (InvertFlags (CMPconst y [c]))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v.Args[1]
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPconst, TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPU(v *Value) bool {
	b := v.Block
	_ = b
	// match: (CMPU x (MOVDconst [c]))
	// cond: isU16Bit(c)
	// result: (CMPUconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPU (MOVDconst [c]) y)
	// cond: isU16Bit(c)
	// result: (InvertFlags (CMPUconst y [c]))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v.Args[1]
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPUconst, TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPUconst(v *Value) bool {
	// match: (CMPUconst (MOVDconst [x]) [y])
	// cond: int64(x)==int64(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x) == int64(y)) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}
	// match: (CMPUconst (MOVDconst [x]) [y])
	// cond: uint64(x)<uint64(y)
	// result: (FlagLT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) < uint64(y)) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}
	// match: (CMPUconst (MOVDconst [x]) [y])
	// cond: uint64(x)>uint64(y)
	// result: (FlagGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) > uint64(y)) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPW(v *Value) bool {
	b := v.Block
	_ = b
	// match: (CMPW x (MOVWreg y))
	// cond:
	// result: (CMPW x y)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64CMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPW (MOVWreg x) y)
	// cond:
	// result: (CMPW x y)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVWreg {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64CMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPW x (MOVDconst [c]))
	// cond: is16Bit(c)
	// result: (CMPWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPW (MOVDconst [c]) y)
	// cond: is16Bit(c)
	// result: (InvertFlags (CMPWconst y [c]))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v.Args[1]
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWconst, TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPWU(v *Value) bool {
	b := v.Block
	_ = b
	// match: (CMPWU x (MOVWZreg y))
	// cond:
	// result: (CMPWU x y)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64CMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPWU (MOVWZreg x) y)
	// cond:
	// result: (CMPWU x y)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64CMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMPWU x (MOVDconst [c]))
	// cond: isU16Bit(c)
	// result: (CMPWUconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMPWU (MOVDconst [c]) y)
	// cond: isU16Bit(c)
	// result: (InvertFlags (CMPWUconst y [c]))
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v.Args[1]
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWUconst, TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPWUconst(v *Value) bool {
	// match: (CMPWUconst (MOVDconst [x]) [y])
	// cond: int32(x)==int32(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}
	// match: (CMPWUconst (MOVDconst [x]) [y])
	// cond: uint32(x)<uint32(y)
	// result: (FlagLT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}
	// match: (CMPWUconst (MOVDconst [x]) [y])
	// cond: uint32(x)>uint32(y)
	// result: (FlagGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPWconst(v *Value) bool {
	// match: (CMPWconst (MOVDconst [x]) [y])
	// cond: int32(x)==int32(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}
	// match: (CMPWconst (MOVDconst [x]) [y])
	// cond: int32(x)<int32(y)
	// result: (FlagLT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y)) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}
	// match: (CMPWconst (MOVDconst [x]) [y])
	// cond: int32(x)>int32(y)
	// result: (FlagGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y)) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPconst(v *Value) bool {
	// match: (CMPconst (MOVDconst [x]) [y])
	// cond: int64(x)==int64(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x) == int64(y)) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}
	// match: (CMPconst (MOVDconst [x]) [y])
	// cond: int64(x)<int64(y)
	// result: (FlagLT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x) < int64(y)) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}
	// match: (CMPconst (MOVDconst [x]) [y])
	// cond: int64(x)>int64(y)
	// result: (FlagGT)
	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x) > int64(y)) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64Equal(v *Value) bool {
	// match: (Equal (FlagEQ))
	// cond:
	// result: (MOVDconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (Equal (FlagLT))
	// cond:
	// result: (MOVDconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (FlagGT))
	// cond:
	// result: (MOVDconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (InvertFlags x))
	// cond:
	// result: (Equal x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64Equal)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FADD(v *Value) bool {
	// match: (FADD (FMUL x y) z)
	// cond:
	// result: (FMADD x y z)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMUL {
			break
		}
		x := v_0.Args[0]
		y := v_0.Args[1]
		z := v.Args[1]
		v.reset(OpPPC64FMADD)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (FADD z (FMUL x y))
	// cond:
	// result: (FMADD x y z)
	for {
		z := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64FMUL {
			break
		}
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpPPC64FMADD)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FADDS(v *Value) bool {
	// match: (FADDS (FMULS x y) z)
	// cond:
	// result: (FMADDS x y z)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMULS {
			break
		}
		x := v_0.Args[0]
		y := v_0.Args[1]
		z := v.Args[1]
		v.reset(OpPPC64FMADDS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (FADDS z (FMULS x y))
	// cond:
	// result: (FMADDS x y z)
	for {
		z := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64FMULS {
			break
		}
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpPPC64FMADDS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FMOVDload(v *Value) bool {
	// match: (FMOVDload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (FMOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64FMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (FMOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FMOVDstore(v *Value) bool {
	// match: (FMOVDstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: is16Bit(off1+off2)
	// result: (FMOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVDstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (FMOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
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
		v.reset(OpPPC64FMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FMOVSload(v *Value) bool {
	// match: (FMOVSload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (FMOVSload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64FMOVSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: is16Bit(off1+off2)
	// result: (FMOVSload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FMOVSstore(v *Value) bool {
	// match: (FMOVSstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: is16Bit(off1+off2)
	// result: (FMOVSstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVSstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (FMOVSstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (FMOVSstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
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
		v.reset(OpPPC64FMOVSstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FSUB(v *Value) bool {
	// match: (FSUB (FMUL x y) z)
	// cond:
	// result: (FMSUB x y z)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMUL {
			break
		}
		x := v_0.Args[0]
		y := v_0.Args[1]
		z := v.Args[1]
		v.reset(OpPPC64FMSUB)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FSUBS(v *Value) bool {
	// match: (FSUBS (FMULS x y) z)
	// cond:
	// result: (FMSUBS x y z)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMULS {
			break
		}
		x := v_0.Args[0]
		y := v_0.Args[1]
		z := v.Args[1]
		v.reset(OpPPC64FMSUBS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64GreaterEqual(v *Value) bool {
	// match: (GreaterEqual (FlagEQ))
	// cond:
	// result: (MOVDconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqual (FlagLT))
	// cond:
	// result: (MOVDconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterEqual (FlagGT))
	// cond:
	// result: (MOVDconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqual (InvertFlags x))
	// cond:
	// result: (LessEqual x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64LessEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64GreaterThan(v *Value) bool {
	// match: (GreaterThan (FlagEQ))
	// cond:
	// result: (MOVDconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThan (FlagLT))
	// cond:
	// result: (MOVDconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThan (FlagGT))
	// cond:
	// result: (MOVDconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterThan (InvertFlags x))
	// cond:
	// result: (LessThan x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64LessThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64LessEqual(v *Value) bool {
	// match: (LessEqual (FlagEQ))
	// cond:
	// result: (MOVDconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqual (FlagLT))
	// cond:
	// result: (MOVDconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqual (FlagGT))
	// cond:
	// result: (MOVDconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessEqual (InvertFlags x))
	// cond:
	// result: (GreaterEqual x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64GreaterEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64LessThan(v *Value) bool {
	// match: (LessThan (FlagEQ))
	// cond:
	// result: (MOVDconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThan (FlagLT))
	// cond:
	// result: (MOVDconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessThan (FlagGT))
	// cond:
	// result: (MOVDconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThan (InvertFlags x))
	// cond:
	// result: (GreaterThan x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64GreaterThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBZload(v *Value) bool {
	// match: (MOVBZload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBZload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBZload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVBZload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBZreg(v *Value) bool {
	// match: (MOVBZreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0xFF
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVBZreg y:(MOVBZreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVBZreg (MOVBreg x))
	// cond:
	// result: (MOVBZreg x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVBreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg x:(MOVBZload _ _))
	// cond:
	// result: x
	for {
		x := v.Args[0]
		if x.Op != OpPPC64MOVBZload {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVBZreg (MOVDconst [c]))
	// cond:
	// result: (MOVDconst [int64(uint8(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBreg(v *Value) bool {
	// match: (MOVBreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0x7F
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0x7F) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVBreg y:(MOVBreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVBreg (MOVBZreg x))
	// cond:
	// result: (MOVBreg x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVDconst [c]))
	// cond:
	// result: (MOVDconst [int64(int8(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBstore(v *Value) bool {
	// match: (MOVBstore [off1] {sym} (ADDconst [off2] x) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {sym} x val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
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
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVDconst [c]) mem)
	// cond: c == 0
	// result: (MOVBstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(c == 0) {
			break
		}
		v.reset(OpPPC64MOVBstorezero)
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
		if v_1.Op != OpPPC64MOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBZreg x) mem)
	// cond:
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBstorezero(v *Value) bool {
	// match: (MOVBstorezero [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVBstorezero [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstorezero [off1] {sym1} (MOVDaddr [off2] {sym2} x) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBstorezero [off1+off2] {mergeSym(sym1,sym2)} x mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVDload(v *Value) bool {
	// match: (MOVDload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVDload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVDstore(v *Value) bool {
	// match: (MOVDstore [off1] {sym} (ADDconst [off2] x) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVDstore [off1+off2] {sym} x val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
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
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off] {sym} ptr (MOVDconst [c]) mem)
	// cond: c == 0
	// result: (MOVDstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(c == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVDstorezero(v *Value) bool {
	// match: (MOVDstorezero [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVDstorezero [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstorezero [off1] {sym1} (MOVDaddr [off2] {sym2} x) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDstorezero [off1+off2] {mergeSym(sym1,sym2)} x mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHZload(v *Value) bool {
	// match: (MOVHZload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHZload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64MOVHZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHZload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHZload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHZreg(v *Value) bool {
	// match: (MOVHZreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0xFFFF
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHZreg y:(MOVHZreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHZreg y:(MOVBZreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHZreg y:(MOVHreg x))
	// cond:
	// result: (MOVHZreg x)
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVHZreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg x:(MOVHZload _ _))
	// cond:
	// result: x
	for {
		x := v.Args[0]
		if x.Op != OpPPC64MOVHZload {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHZreg (MOVDconst [c]))
	// cond:
	// result: (MOVDconst [int64(uint16(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHload(v *Value) bool {
	// match: (MOVHload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHreg(v *Value) bool {
	// match: (MOVHreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0x7FFF
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0x7FFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHreg y:(MOVHreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHreg y:(MOVBreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVHreg y:(MOVHZreg x))
	// cond:
	// result: (MOVHreg x)
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHZreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHload _ _))
	// cond:
	// result: x
	for {
		x := v.Args[0]
		if x.Op != OpPPC64MOVHload {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVDconst [c]))
	// cond:
	// result: (MOVDconst [int64(int16(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHstore(v *Value) bool {
	// match: (MOVHstore [off1] {sym} (ADDconst [off2] x) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {sym} x val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
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
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVDconst [c]) mem)
	// cond: c == 0
	// result: (MOVHstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(c == 0) {
			break
		}
		v.reset(OpPPC64MOVHstorezero)
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
		if v_1.Op != OpPPC64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHZreg x) mem)
	// cond:
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHstorezero(v *Value) bool {
	// match: (MOVHstorezero [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVHstorezero [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstorezero [off1] {sym1} (MOVDaddr [off2] {sym2} x) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHstorezero [off1+off2] {mergeSym(sym1,sym2)} x mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWZload(v *Value) bool {
	// match: (MOVWZload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWZload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64MOVWZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWZload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVWZload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVWZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWZreg(v *Value) bool {
	// match: (MOVWZreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0xFFFFFFFF
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFFFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(AND (MOVDconst [c]) _))
	// cond: uint64(c) <= 0xFFFFFFFF
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64AND {
			break
		}
		y_0 := y.Args[0]
		if y_0.Op != OpPPC64MOVDconst {
			break
		}
		c := y_0.AuxInt
		if !(uint64(c) <= 0xFFFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(AND _ (MOVDconst [c])))
	// cond: uint64(c) <= 0xFFFFFFFF
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64AND {
			break
		}
		y_1 := y.Args[1]
		if y_1.Op != OpPPC64MOVDconst {
			break
		}
		c := y_1.AuxInt
		if !(uint64(c) <= 0xFFFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(MOVWZreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVWZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(MOVHZreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(MOVBZreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWZreg y:(MOVWreg x))
	// cond:
	// result: (MOVWZreg x)
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVWreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVWZreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWload(v *Value) bool {
	// match: (MOVWload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVWload [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWreg(v *Value) bool {
	// match: (MOVWreg y:(ANDconst [c] _))
	// cond: uint64(c) <= 0xFFFF
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWreg y:(AND (MOVDconst [c]) _))
	// cond: uint64(c) <= 0x7FFFFFFF
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64AND {
			break
		}
		y_0 := y.Args[0]
		if y_0.Op != OpPPC64MOVDconst {
			break
		}
		c := y_0.AuxInt
		if !(uint64(c) <= 0x7FFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWreg y:(AND _ (MOVDconst [c])))
	// cond: uint64(c) <= 0x7FFFFFFF
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64AND {
			break
		}
		y_1 := y.Args[1]
		if y_1.Op != OpPPC64MOVDconst {
			break
		}
		c := y_1.AuxInt
		if !(uint64(c) <= 0x7FFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWreg y:(MOVWreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVWreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWreg y:(MOVHreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWreg y:(MOVBreg _))
	// cond:
	// result: y
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	// match: (MOVWreg y:(MOVWZreg x))
	// cond:
	// result: (MOVWreg x)
	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVWZreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVWreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWstore(v *Value) bool {
	// match: (MOVWstore [off1] {sym} (ADDconst [off2] x) val mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {sym} x val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
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
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVDconst [c]) mem)
	// cond: c == 0
	// result: (MOVWstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(c == 0) {
			break
		}
		v.reset(OpPPC64MOVWstorezero)
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
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWZreg x) mem)
	// cond:
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWstorezero(v *Value) bool {
	// match: (MOVWstorezero [off1] {sym} (ADDconst [off2] x) mem)
	// cond: is16Bit(off1+off2)
	// result: (MOVWstorezero [off1+off2] {sym} x mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstorezero [off1] {sym1} (MOVDaddr [off2] {sym2} x) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWstorezero [off1+off2] {mergeSym(sym1,sym2)} x mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MaskIfNotCarry(v *Value) bool {
	// match: (MaskIfNotCarry (ADDconstForCarry [c] (ANDconst [d] _)))
	// cond: c < 0 && d > 0 && c + d < 0
	// result: (MOVDconst [-1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconstForCarry {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64ANDconst {
			break
		}
		d := v_0_0.AuxInt
		if !(c < 0 && d > 0 && c+d < 0) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = -1
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64NotEqual(v *Value) bool {
	// match: (NotEqual (FlagEQ))
	// cond:
	// result: (MOVDconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (NotEqual (FlagLT))
	// cond:
	// result: (MOVDconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (FlagGT))
	// cond:
	// result: (MOVDconst [1])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (InvertFlags x))
	// cond:
	// result: (NotEqual x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64NotEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR(v *Value) bool {
	// match: (OR (MOVDconst [c]) (MOVDconst [d]))
	// cond:
	// result: (MOVDconst [c|d])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c | d
		return true
	}
	// match: (OR (MOVDconst [d]) (MOVDconst [c]))
	// cond:
	// result: (MOVDconst [c|d])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c | d
		return true
	}
	// match: (OR x (MOVDconst [c]))
	// cond: isU32Bit(c)
	// result: (ORconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpPPC64ORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (OR (MOVDconst [c]) x)
	// cond: isU32Bit(c)
	// result: (ORconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpPPC64ORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64ORN(v *Value) bool {
	// match: (ORN x (MOVDconst [-1]))
	// cond:
	// result: x
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64ORconst(v *Value) bool {
	// match: (ORconst [c] (ORconst [d] x))
	// cond:
	// result: (ORconst [c|d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpPPC64ORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	// match: (ORconst [-1] _)
	// cond:
	// result: (MOVDconst [-1])
	for {
		if v.AuxInt != -1 {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = -1
		return true
	}
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
	return false
}
func rewriteValuePPC64_OpPPC64SUB(v *Value) bool {
	// match: (SUB x (MOVDconst [c]))
	// cond: is32Bit(-c)
	// result: (ADDconst [-c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(-c)) {
			break
		}
		v.reset(OpPPC64ADDconst)
		v.AuxInt = -c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64XOR(v *Value) bool {
	// match: (XOR (MOVDconst [c]) (MOVDconst [d]))
	// cond:
	// result: (MOVDconst [c^d])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c ^ d
		return true
	}
	// match: (XOR (MOVDconst [d]) (MOVDconst [c]))
	// cond:
	// result: (MOVDconst [c^d])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c ^ d
		return true
	}
	// match: (XOR x (MOVDconst [c]))
	// cond: isU32Bit(c)
	// result: (XORconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpPPC64XORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (XOR (MOVDconst [c]) x)
	// cond: isU32Bit(c)
	// result: (XORconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpPPC64XORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64XORconst(v *Value) bool {
	// match: (XORconst [c] (XORconst [d] x))
	// cond:
	// result: (XORconst [c^d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64XORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpPPC64XORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
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
	return false
}
func rewriteValuePPC64_OpRound32F(v *Value) bool {
	// match: (Round32F x)
	// cond:
	// result: (LoweredRound32F x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64LoweredRound32F)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpRound64F(v *Value) bool {
	// match: (Round64F x)
	// cond:
	// result: (LoweredRound64F x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64LoweredRound64F)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpRsh16Ux16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux16 x y)
	// cond:
	// result: (SRW  (ZeroExt16to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16Ux32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux32 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SRWconst (ZeroExt16to32 x) [c])
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
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux32 x (MOVDconst [c]))
	// cond: uint32(c) < 16
	// result: (SRWconst (ZeroExt16to32 x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux32 x y)
	// cond:
	// result: (SRW  (ZeroExt16to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16Ux64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRWconst (ZeroExt16to32 x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (MOVDconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh16Ux64 x (MOVDconst [c]))
	// cond: uint64(c) < 16
	// result: (SRWconst (ZeroExt16to32 x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 x y)
	// cond:
	// result: (SRW  (ZeroExt16to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16Ux8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux8 x y)
	// cond:
	// result: (SRW  (ZeroExt16to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x16 x y)
	// cond:
	// result: (SRAW (SignExt16to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x32 x (Const64 [c]))
	// cond: uint32(c) < 16
	// result: (SRAWconst (SignExt16to32 x) [c])
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
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x32 x (MOVDconst [c]))
	// cond: uint32(c) < 16
	// result: (SRAWconst (SignExt16to32 x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x32 x y)
	// cond:
	// result: (SRAW (SignExt16to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRAWconst (SignExt16to32 x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (SRAWconst (SignExt16to32 x) [63])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (MOVDconst [c]))
	// cond: uint64(c) < 16
	// result: (SRAWconst (SignExt16to32 x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x y)
	// cond:
	// result: (SRAW (SignExt16to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh16x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x8 x y)
	// cond:
	// result: (SRAW (SignExt16to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-16] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh32Ux16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32Ux16 x y)
	// cond:
	// result: (SRW x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32Ux32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32Ux32 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SRWconst x [c])
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
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux32 x (MOVDconst [c]))
	// cond: uint32(c) < 32
	// result: (SRWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux32 x y)
	// cond:
	// result: (SRW x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32Ux64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32Ux64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (MOVDconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh32Ux64 x (MOVDconst [c]))
	// cond: uint64(c) < 32
	// result: (SRWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux64 x y)
	// cond:
	// result: (SRW  x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32Ux8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32Ux8 x y)
	// cond:
	// result: (SRW x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32x16 x y)
	// cond:
	// result: (SRAW x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32x32 x (Const64 [c]))
	// cond: uint32(c) < 32
	// result: (SRAWconst x [c])
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
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x32 x (MOVDconst [c]))
	// cond: uint32(c) < 32
	// result: (SRAWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x32 x y)
	// cond:
	// result: (SRAW x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRAWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (SRAWconst x [63])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = 63
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x (MOVDconst [c]))
	// cond: uint64(c) < 32
	// result: (SRAWconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x y)
	// cond:
	// result: (SRAW x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh32x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32x8 x y)
	// cond:
	// result: (SRAW x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-32] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64Ux16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64Ux16 x y)
	// cond:
	// result: (SRD x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64Ux32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64Ux32 x (Const64 [c]))
	// cond: uint32(c) < 64
	// result: (SRDconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux32 x (MOVDconst [c]))
	// cond: uint32(c) < 64
	// result: (SRDconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux32 x y)
	// cond:
	// result: (SRD x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64Ux64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64Ux64 x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRDconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh64Ux64 x (MOVDconst [c]))
	// cond: uint64(c) < 64
	// result: (SRDconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux64 x y)
	// cond:
	// result: (SRD  x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64Ux8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64Ux8 x y)
	// cond:
	// result: (SRD x  (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64x16 x y)
	// cond:
	// result: (SRAD x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64x32 x (Const64 [c]))
	// cond: uint32(c) < 64
	// result: (SRADconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x32 x (MOVDconst [c]))
	// cond: uint32(c) < 64
	// result: (SRADconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x32 x y)
	// cond:
	// result: (SRAD x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64x64 x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRADconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 x (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (SRADconst x [63])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = 63
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 x (MOVDconst [c]))
	// cond: uint64(c) < 64
	// result: (SRADconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 x y)
	// cond:
	// result: (SRAD x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh64x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64x8 x y)
	// cond:
	// result: (SRAD x (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-64] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRsh8Ux16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux16 x y)
	// cond:
	// result: (SRW  (ZeroExt8to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8Ux32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux32 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SRWconst (ZeroExt8to32  x) [c])
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
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux32 x (MOVDconst [c]))
	// cond: uint32(c) < 8
	// result: (SRWconst (ZeroExt8to32  x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux32 x y)
	// cond:
	// result: (SRW  (ZeroExt8to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8Ux64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRWconst (ZeroExt8to32  x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (MOVDconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh8Ux64 x (MOVDconst [c]))
	// cond: uint64(c) < 8
	// result: (SRWconst (ZeroExt8to32  x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 x y)
	// cond:
	// result: (SRW  (ZeroExt8to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8Ux8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux8 x y)
	// cond:
	// result: (SRW  (ZeroExt8to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, types.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x16 x y)
	// cond:
	// result: (SRAW (SignExt8to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt16to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x32 x (Const64 [c]))
	// cond: uint32(c) < 8
	// result: (SRAWconst (SignExt8to32  x) [c])
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
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x32 x (MOVDconst [c]))
	// cond: uint32(c) < 8
	// result: (SRAWconst (SignExt8to32  x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x32 x y)
	// cond:
	// result: (SRAW (SignExt8to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt32to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRAWconst (SignExt8to32  x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (SRAWconst (SignExt8to32  x) [63])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (MOVDconst [c]))
	// cond: uint64(c) < 8
	// result: (SRAWconst (SignExt8to32  x) [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x y)
	// cond:
	// result: (SRAW (SignExt8to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] y))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpRsh8x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x8 x y)
	// cond:
	// result: (SRAW (SignExt8to32 x) (ORN y <types.Int64> (MaskIfNotCarry (ADDconstForCarry [-8] (ZeroExt8to64 y)))))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, types.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, types.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, types.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpSignExt16to32(v *Value) bool {
	// match: (SignExt16to32 x)
	// cond:
	// result: (MOVHreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSignExt16to64(v *Value) bool {
	// match: (SignExt16to64 x)
	// cond:
	// result: (MOVHreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSignExt32to64(v *Value) bool {
	// match: (SignExt32to64 x)
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVWreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSignExt8to16(v *Value) bool {
	// match: (SignExt8to16 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSignExt8to32(v *Value) bool {
	// match: (SignExt8to32 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSignExt8to64(v *Value) bool {
	// match: (SignExt8to64 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSlicemask(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Slicemask <t> x)
	// cond:
	// result: (SRADconst (NEG <t> x) [63])
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpPPC64SRADconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpPPC64NEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpSqrt(v *Value) bool {
	// match: (Sqrt x)
	// cond:
	// result: (FSQRT x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64FSQRT)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpStaticCall(v *Value) bool {
	// match: (StaticCall [argwid] {target} mem)
	// cond:
	// result: (CALLstatic [argwid] {target} mem)
	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpPPC64CALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpStore(v *Value) bool {
	// match: (Store {t} ptr val mem)
	// cond: t.(Type).Size() == 8 && is64BitFloat(val.Type)
	// result: (FMOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpPPC64FMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(Type).Size() == 8 && is32BitFloat(val.Type)
	// result: (FMOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 8 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpPPC64FMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(Type).Size() == 4 && is32BitFloat(val.Type)
	// result: (FMOVSstore ptr val mem)
	for {
		t := v.Aux
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpPPC64FMOVSstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(Type).Size() == 8 && (is64BitInt(val.Type) || isPtr(val.Type))
	// result: (MOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 8 && (is64BitInt(val.Type) || isPtr(val.Type))) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(Type).Size() == 4 && is32BitInt(val.Type)
	// result: (MOVWstore ptr val mem)
	for {
		t := v.Aux
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 4 && is32BitInt(val.Type)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
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
		v.reset(OpPPC64MOVHstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
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
		v.reset(OpPPC64MOVBstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpSub16(v *Value) bool {
	// match: (Sub16 x y)
	// cond:
	// result: (SUB x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSub32(v *Value) bool {
	// match: (Sub32 x y)
	// cond:
	// result: (SUB x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSub32F(v *Value) bool {
	// match: (Sub32F x y)
	// cond:
	// result: (FSUBS x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FSUBS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSub64(v *Value) bool {
	// match: (Sub64 x y)
	// cond:
	// result: (SUB  x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSub64F(v *Value) bool {
	// match: (Sub64F x y)
	// cond:
	// result: (FSUB x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSub8(v *Value) bool {
	// match: (Sub8 x y)
	// cond:
	// result: (SUB x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSubPtr(v *Value) bool {
	// match: (SubPtr x y)
	// cond:
	// result: (SUB  x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpTrunc16to8(v *Value) bool {
	// match: (Trunc16to8 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc32to16(v *Value) bool {
	// match: (Trunc32to16 x)
	// cond:
	// result: (MOVHreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc32to8(v *Value) bool {
	// match: (Trunc32to8 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc64to16(v *Value) bool {
	// match: (Trunc64to16 x)
	// cond:
	// result: (MOVHreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc64to32(v *Value) bool {
	// match: (Trunc64to32 x)
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVWreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc64to8(v *Value) bool {
	// match: (Trunc64to8 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpXor16(v *Value) bool {
	// match: (Xor16 x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpXor32(v *Value) bool {
	// match: (Xor32 x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpXor64(v *Value) bool {
	// match: (Xor64 x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpXor8(v *Value) bool {
	// match: (Xor8 x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpZero(v *Value) bool {
	b := v.Block
	_ = b
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
	// match: (Zero [1] destptr mem)
	// cond:
	// result: (MOVBstorezero destptr mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVBstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] destptr mem)
	// cond:
	// result: (MOVHstorezero destptr mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVHstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [3] destptr mem)
	// cond:
	// result: (MOVBstorezero [2] destptr 		(MOVHstorezero destptr mem))
	for {
		if v.AuxInt != 3 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = 2
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHstorezero, TypeMem)
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [4] destptr mem)
	// cond:
	// result: (MOVWstorezero destptr mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVWstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [5] destptr mem)
	// cond:
	// result: (MOVBstorezero [4] destptr         	(MOVWstorezero destptr mem))
	for {
		if v.AuxInt != 5 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = 4
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWstorezero, TypeMem)
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [6] destptr mem)
	// cond:
	// result: (MOVHstorezero [4] destptr 		(MOVWstorezero destptr mem))
	for {
		if v.AuxInt != 6 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVHstorezero)
		v.AuxInt = 4
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWstorezero, TypeMem)
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [7] destptr mem)
	// cond:
	// result: (MOVBstorezero [6] destptr 		(MOVHstorezero [4] destptr 			(MOVWstorezero destptr mem)))
	for {
		if v.AuxInt != 7 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = 6
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHstorezero, TypeMem)
		v0.AuxInt = 4
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstorezero, TypeMem)
		v1.AddArg(destptr)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [8] destptr mem)
	// cond:
	// result: (MOVDstorezero destptr mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVDstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [12] destptr mem)
	// cond:
	// result: (MOVWstorezero [8] destptr                 (MOVDstorezero [0] destptr mem))
	for {
		if v.AuxInt != 12 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = 8
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [16] destptr mem)
	// cond:
	// result: (MOVDstorezero [8] destptr                 (MOVDstorezero [0] destptr mem))
	for {
		if v.AuxInt != 16 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 8
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [24] destptr mem)
	// cond:
	// result: (MOVDstorezero [16] destptr 		(MOVDstorezero [8] destptr 			(MOVDstorezero [0] destptr mem)))
	for {
		if v.AuxInt != 24 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 16
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v0.AuxInt = 8
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v1.AuxInt = 0
		v1.AddArg(destptr)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [32] destptr mem)
	// cond:
	// result: (MOVDstorezero [24] destptr 		(MOVDstorezero [16] destptr 			(MOVDstorezero [8] destptr 				(MOVDstorezero [0] destptr mem))))
	for {
		if v.AuxInt != 32 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 24
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v0.AuxInt = 16
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v1.AuxInt = 8
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v2.AuxInt = 0
		v2.AddArg(destptr)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [40] destptr mem)
	// cond:
	// result: (MOVDstorezero [32] destptr 		(MOVDstorezero [24] destptr 			(MOVDstorezero [16] destptr 				(MOVDstorezero [8] destptr 					(MOVDstorezero [0] destptr mem)))))
	for {
		if v.AuxInt != 40 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 32
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v0.AuxInt = 24
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v1.AuxInt = 16
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v2.AuxInt = 8
		v2.AddArg(destptr)
		v3 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v3.AuxInt = 0
		v3.AddArg(destptr)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [48] destptr mem)
	// cond:
	// result: (MOVDstorezero [40] destptr 		(MOVDstorezero [32] destptr 			(MOVDstorezero [24] destptr 				(MOVDstorezero [16] destptr 					(MOVDstorezero [8] destptr 						(MOVDstorezero [0] destptr mem))))))
	for {
		if v.AuxInt != 48 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 40
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v0.AuxInt = 32
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v1.AuxInt = 24
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v2.AuxInt = 16
		v2.AddArg(destptr)
		v3 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v3.AuxInt = 8
		v3.AddArg(destptr)
		v4 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v4.AuxInt = 0
		v4.AddArg(destptr)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [56] destptr mem)
	// cond:
	// result: (MOVDstorezero [48] destptr 		(MOVDstorezero [40] destptr 			(MOVDstorezero [32] destptr 				(MOVDstorezero [24] destptr 					(MOVDstorezero [16] destptr 						(MOVDstorezero [8] destptr 							(MOVDstorezero [0] destptr mem)))))))
	for {
		if v.AuxInt != 56 {
			break
		}
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 48
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v0.AuxInt = 40
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v1.AuxInt = 32
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v2.AuxInt = 24
		v2.AddArg(destptr)
		v3 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v3.AuxInt = 16
		v3.AddArg(destptr)
		v4 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v4.AuxInt = 8
		v4.AddArg(destptr)
		v5 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, TypeMem)
		v5.AuxInt = 0
		v5.AddArg(destptr)
		v5.AddArg(mem)
		v4.AddArg(v5)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Zero [s] ptr mem)
	// cond:
	// result: (LoweredZero [s] ptr mem)
	for {
		s := v.AuxInt
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64LoweredZero)
		v.AuxInt = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt16to32(v *Value) bool {
	// match: (ZeroExt16to32 x)
	// cond:
	// result: (MOVHZreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt16to64(v *Value) bool {
	// match: (ZeroExt16to64 x)
	// cond:
	// result: (MOVHZreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt32to64(v *Value) bool {
	// match: (ZeroExt32to64 x)
	// cond:
	// result: (MOVWZreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVWZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt8to16(v *Value) bool {
	// match: (ZeroExt8to16 x)
	// cond:
	// result: (MOVBZreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt8to32(v *Value) bool {
	// match: (ZeroExt8to32 x)
	// cond:
	// result: (MOVBZreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt8to64(v *Value) bool {
	// match: (ZeroExt8to64 x)
	// cond:
	// result: (MOVBZreg x)
	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteBlockPPC64(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	types := &config.Types
	_ = types
	switch b.Kind {
	case BlockPPC64EQ:
		// match: (EQ (CMPconst [0] (ANDconst [c] x)) yes no)
		// cond:
		// result: (EQ (ANDCCconst [c] x) yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64EQ
			v0 := b.NewValue0(v.Pos, OpPPC64ANDCCconst, TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (CMPWconst [0] (ANDconst [c] x)) yes no)
		// cond:
		// result: (EQ (ANDCCconst [c] x) yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64EQ
			v0 := b.NewValue0(v.Pos, OpPPC64ANDCCconst, TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (FlagEQ) yes no)
		// cond:
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
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
		// match: (EQ (FlagLT) yes no)
		// cond:
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
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
		// match: (EQ (FlagGT) yes no)
		// cond:
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
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
		// match: (EQ (InvertFlags cmp) yes no)
		// cond:
		// result: (EQ cmp yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64EQ
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
	case BlockPPC64GE:
		// match: (GE (FlagEQ) yes no)
		// cond:
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
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
		// match: (GE (FlagLT) yes no)
		// cond:
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
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
		// match: (GE (FlagGT) yes no)
		// cond:
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
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
		// match: (GE (InvertFlags cmp) yes no)
		// cond:
		// result: (LE cmp yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64LE
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
	case BlockPPC64GT:
		// match: (GT (FlagEQ) yes no)
		// cond:
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
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
		// match: (GT (FlagLT) yes no)
		// cond:
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
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
		// match: (GT (FlagGT) yes no)
		// cond:
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
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
		// match: (GT (InvertFlags cmp) yes no)
		// cond:
		// result: (LT cmp yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64LT
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
	case BlockIf:
		// match: (If (Equal cc) yes no)
		// cond:
		// result: (EQ cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64Equal {
				break
			}
			cc := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64EQ
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (If (NotEqual cc) yes no)
		// cond:
		// result: (NE cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64NotEqual {
				break
			}
			cc := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64NE
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (If (LessThan cc) yes no)
		// cond:
		// result: (LT cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64LessThan {
				break
			}
			cc := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64LT
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (If (LessEqual cc) yes no)
		// cond:
		// result: (LE cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64LessEqual {
				break
			}
			cc := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64LE
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (If (GreaterThan cc) yes no)
		// cond:
		// result: (GT cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64GreaterThan {
				break
			}
			cc := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64GT
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (If (GreaterEqual cc) yes no)
		// cond:
		// result: (GE cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64GreaterEqual {
				break
			}
			cc := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64GE
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (If (FLessThan cc) yes no)
		// cond:
		// result: (FLT cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FLessThan {
				break
			}
			cc := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64FLT
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (If (FLessEqual cc) yes no)
		// cond:
		// result: (FLE cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FLessEqual {
				break
			}
			cc := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64FLE
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (If (FGreaterThan cc) yes no)
		// cond:
		// result: (FGT cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FGreaterThan {
				break
			}
			cc := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64FGT
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (If (FGreaterEqual cc) yes no)
		// cond:
		// result: (FGE cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FGreaterEqual {
				break
			}
			cc := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64FGE
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (If cond yes no)
		// cond:
		// result: (NE (CMPWconst [0] cond) yes no)
		for {
			v := b.Control
			_ = v
			cond := b.Control
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64NE
			v0 := b.NewValue0(v.Pos, OpPPC64CMPWconst, TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(cond)
			b.SetControl(v0)
			_ = yes
			_ = no
			return true
		}
	case BlockPPC64LE:
		// match: (LE (FlagEQ) yes no)
		// cond:
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
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
		// match: (LE (FlagLT) yes no)
		// cond:
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
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
		// match: (LE (FlagGT) yes no)
		// cond:
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
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
		// match: (LE (InvertFlags cmp) yes no)
		// cond:
		// result: (GE cmp yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64GE
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
	case BlockPPC64LT:
		// match: (LT (FlagEQ) yes no)
		// cond:
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
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
		// match: (LT (FlagLT) yes no)
		// cond:
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
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
		// match: (LT (FlagGT) yes no)
		// cond:
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
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
		// match: (LT (InvertFlags cmp) yes no)
		// cond:
		// result: (GT cmp yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64GT
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
	case BlockPPC64NE:
		// match: (NE (CMPWconst [0] (Equal cc)) yes no)
		// cond:
		// result: (EQ cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64Equal {
				break
			}
			cc := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64EQ
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (CMPWconst [0] (NotEqual cc)) yes no)
		// cond:
		// result: (NE cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64NotEqual {
				break
			}
			cc := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64NE
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (CMPWconst [0] (LessThan cc)) yes no)
		// cond:
		// result: (LT cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64LessThan {
				break
			}
			cc := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64LT
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (CMPWconst [0] (LessEqual cc)) yes no)
		// cond:
		// result: (LE cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64LessEqual {
				break
			}
			cc := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64LE
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (CMPWconst [0] (GreaterThan cc)) yes no)
		// cond:
		// result: (GT cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64GreaterThan {
				break
			}
			cc := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64GT
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (CMPWconst [0] (GreaterEqual cc)) yes no)
		// cond:
		// result: (GE cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64GreaterEqual {
				break
			}
			cc := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64GE
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (CMPWconst [0] (FLessThan cc)) yes no)
		// cond:
		// result: (FLT cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64FLessThan {
				break
			}
			cc := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64FLT
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (CMPWconst [0] (FLessEqual cc)) yes no)
		// cond:
		// result: (FLE cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64FLessEqual {
				break
			}
			cc := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64FLE
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (CMPWconst [0] (FGreaterThan cc)) yes no)
		// cond:
		// result: (FGT cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64FGreaterThan {
				break
			}
			cc := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64FGT
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (CMPWconst [0] (FGreaterEqual cc)) yes no)
		// cond:
		// result: (FGE cc yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64FGreaterEqual {
				break
			}
			cc := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64FGE
			b.SetControl(cc)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (CMPconst [0] (ANDconst [c] x)) yes no)
		// cond:
		// result: (NE (ANDCCconst [c] x) yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64NE
			v0 := b.NewValue0(v.Pos, OpPPC64ANDCCconst, TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (CMPWconst [0] (ANDconst [c] x)) yes no)
		// cond:
		// result: (NE (ANDCCconst [c] x) yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64NE
			v0 := b.NewValue0(v.Pos, OpPPC64ANDCCconst, TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (FlagEQ) yes no)
		// cond:
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
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
		// match: (NE (FlagLT) yes no)
		// cond:
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
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
		// match: (NE (FlagGT) yes no)
		// cond:
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
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
		// match: (NE (InvertFlags cmp) yes no)
		// cond:
		// result: (NE cmp yes no)
		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockPPC64NE
			b.SetControl(cmp)
			_ = yes
			_ = no
			return true
		}
	}
	return false
}
