// Code generated from gen/MIPS64.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "math"
import "cmd/internal/obj"

var _ = math.MinInt8 // in case not otherwise used
var _ = obj.ANOP     // in case not otherwise used
func rewriteValueMIPS64(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValueMIPS64_OpAdd16(v)
	case OpAdd32:
		return rewriteValueMIPS64_OpAdd32(v)
	case OpAdd32F:
		return rewriteValueMIPS64_OpAdd32F(v)
	case OpAdd64:
		return rewriteValueMIPS64_OpAdd64(v)
	case OpAdd64F:
		return rewriteValueMIPS64_OpAdd64F(v)
	case OpAdd8:
		return rewriteValueMIPS64_OpAdd8(v)
	case OpAddPtr:
		return rewriteValueMIPS64_OpAddPtr(v)
	case OpAddr:
		return rewriteValueMIPS64_OpAddr(v)
	case OpAnd16:
		return rewriteValueMIPS64_OpAnd16(v)
	case OpAnd32:
		return rewriteValueMIPS64_OpAnd32(v)
	case OpAnd64:
		return rewriteValueMIPS64_OpAnd64(v)
	case OpAnd8:
		return rewriteValueMIPS64_OpAnd8(v)
	case OpAndB:
		return rewriteValueMIPS64_OpAndB(v)
	case OpAvg64u:
		return rewriteValueMIPS64_OpAvg64u(v)
	case OpClosureCall:
		return rewriteValueMIPS64_OpClosureCall(v)
	case OpCom16:
		return rewriteValueMIPS64_OpCom16(v)
	case OpCom32:
		return rewriteValueMIPS64_OpCom32(v)
	case OpCom64:
		return rewriteValueMIPS64_OpCom64(v)
	case OpCom8:
		return rewriteValueMIPS64_OpCom8(v)
	case OpConst16:
		return rewriteValueMIPS64_OpConst16(v)
	case OpConst32:
		return rewriteValueMIPS64_OpConst32(v)
	case OpConst32F:
		return rewriteValueMIPS64_OpConst32F(v)
	case OpConst64:
		return rewriteValueMIPS64_OpConst64(v)
	case OpConst64F:
		return rewriteValueMIPS64_OpConst64F(v)
	case OpConst8:
		return rewriteValueMIPS64_OpConst8(v)
	case OpConstBool:
		return rewriteValueMIPS64_OpConstBool(v)
	case OpConstNil:
		return rewriteValueMIPS64_OpConstNil(v)
	case OpConvert:
		return rewriteValueMIPS64_OpConvert(v)
	case OpCvt32Fto32:
		return rewriteValueMIPS64_OpCvt32Fto32(v)
	case OpCvt32Fto64:
		return rewriteValueMIPS64_OpCvt32Fto64(v)
	case OpCvt32Fto64F:
		return rewriteValueMIPS64_OpCvt32Fto64F(v)
	case OpCvt32to32F:
		return rewriteValueMIPS64_OpCvt32to32F(v)
	case OpCvt32to64F:
		return rewriteValueMIPS64_OpCvt32to64F(v)
	case OpCvt64Fto32:
		return rewriteValueMIPS64_OpCvt64Fto32(v)
	case OpCvt64Fto32F:
		return rewriteValueMIPS64_OpCvt64Fto32F(v)
	case OpCvt64Fto64:
		return rewriteValueMIPS64_OpCvt64Fto64(v)
	case OpCvt64to32F:
		return rewriteValueMIPS64_OpCvt64to32F(v)
	case OpCvt64to64F:
		return rewriteValueMIPS64_OpCvt64to64F(v)
	case OpDiv16:
		return rewriteValueMIPS64_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueMIPS64_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueMIPS64_OpDiv32(v)
	case OpDiv32F:
		return rewriteValueMIPS64_OpDiv32F(v)
	case OpDiv32u:
		return rewriteValueMIPS64_OpDiv32u(v)
	case OpDiv64:
		return rewriteValueMIPS64_OpDiv64(v)
	case OpDiv64F:
		return rewriteValueMIPS64_OpDiv64F(v)
	case OpDiv64u:
		return rewriteValueMIPS64_OpDiv64u(v)
	case OpDiv8:
		return rewriteValueMIPS64_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueMIPS64_OpDiv8u(v)
	case OpEq16:
		return rewriteValueMIPS64_OpEq16(v)
	case OpEq32:
		return rewriteValueMIPS64_OpEq32(v)
	case OpEq32F:
		return rewriteValueMIPS64_OpEq32F(v)
	case OpEq64:
		return rewriteValueMIPS64_OpEq64(v)
	case OpEq64F:
		return rewriteValueMIPS64_OpEq64F(v)
	case OpEq8:
		return rewriteValueMIPS64_OpEq8(v)
	case OpEqB:
		return rewriteValueMIPS64_OpEqB(v)
	case OpEqPtr:
		return rewriteValueMIPS64_OpEqPtr(v)
	case OpGeq16:
		return rewriteValueMIPS64_OpGeq16(v)
	case OpGeq16U:
		return rewriteValueMIPS64_OpGeq16U(v)
	case OpGeq32:
		return rewriteValueMIPS64_OpGeq32(v)
	case OpGeq32F:
		return rewriteValueMIPS64_OpGeq32F(v)
	case OpGeq32U:
		return rewriteValueMIPS64_OpGeq32U(v)
	case OpGeq64:
		return rewriteValueMIPS64_OpGeq64(v)
	case OpGeq64F:
		return rewriteValueMIPS64_OpGeq64F(v)
	case OpGeq64U:
		return rewriteValueMIPS64_OpGeq64U(v)
	case OpGeq8:
		return rewriteValueMIPS64_OpGeq8(v)
	case OpGeq8U:
		return rewriteValueMIPS64_OpGeq8U(v)
	case OpGetClosurePtr:
		return rewriteValueMIPS64_OpGetClosurePtr(v)
	case OpGreater16:
		return rewriteValueMIPS64_OpGreater16(v)
	case OpGreater16U:
		return rewriteValueMIPS64_OpGreater16U(v)
	case OpGreater32:
		return rewriteValueMIPS64_OpGreater32(v)
	case OpGreater32F:
		return rewriteValueMIPS64_OpGreater32F(v)
	case OpGreater32U:
		return rewriteValueMIPS64_OpGreater32U(v)
	case OpGreater64:
		return rewriteValueMIPS64_OpGreater64(v)
	case OpGreater64F:
		return rewriteValueMIPS64_OpGreater64F(v)
	case OpGreater64U:
		return rewriteValueMIPS64_OpGreater64U(v)
	case OpGreater8:
		return rewriteValueMIPS64_OpGreater8(v)
	case OpGreater8U:
		return rewriteValueMIPS64_OpGreater8U(v)
	case OpHmul32:
		return rewriteValueMIPS64_OpHmul32(v)
	case OpHmul32u:
		return rewriteValueMIPS64_OpHmul32u(v)
	case OpHmul64:
		return rewriteValueMIPS64_OpHmul64(v)
	case OpHmul64u:
		return rewriteValueMIPS64_OpHmul64u(v)
	case OpInterCall:
		return rewriteValueMIPS64_OpInterCall(v)
	case OpIsInBounds:
		return rewriteValueMIPS64_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValueMIPS64_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValueMIPS64_OpIsSliceInBounds(v)
	case OpLeq16:
		return rewriteValueMIPS64_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueMIPS64_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueMIPS64_OpLeq32(v)
	case OpLeq32F:
		return rewriteValueMIPS64_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValueMIPS64_OpLeq32U(v)
	case OpLeq64:
		return rewriteValueMIPS64_OpLeq64(v)
	case OpLeq64F:
		return rewriteValueMIPS64_OpLeq64F(v)
	case OpLeq64U:
		return rewriteValueMIPS64_OpLeq64U(v)
	case OpLeq8:
		return rewriteValueMIPS64_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueMIPS64_OpLeq8U(v)
	case OpLess16:
		return rewriteValueMIPS64_OpLess16(v)
	case OpLess16U:
		return rewriteValueMIPS64_OpLess16U(v)
	case OpLess32:
		return rewriteValueMIPS64_OpLess32(v)
	case OpLess32F:
		return rewriteValueMIPS64_OpLess32F(v)
	case OpLess32U:
		return rewriteValueMIPS64_OpLess32U(v)
	case OpLess64:
		return rewriteValueMIPS64_OpLess64(v)
	case OpLess64F:
		return rewriteValueMIPS64_OpLess64F(v)
	case OpLess64U:
		return rewriteValueMIPS64_OpLess64U(v)
	case OpLess8:
		return rewriteValueMIPS64_OpLess8(v)
	case OpLess8U:
		return rewriteValueMIPS64_OpLess8U(v)
	case OpLoad:
		return rewriteValueMIPS64_OpLoad(v)
	case OpLsh16x16:
		return rewriteValueMIPS64_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueMIPS64_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValueMIPS64_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValueMIPS64_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueMIPS64_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueMIPS64_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValueMIPS64_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValueMIPS64_OpLsh32x8(v)
	case OpLsh64x16:
		return rewriteValueMIPS64_OpLsh64x16(v)
	case OpLsh64x32:
		return rewriteValueMIPS64_OpLsh64x32(v)
	case OpLsh64x64:
		return rewriteValueMIPS64_OpLsh64x64(v)
	case OpLsh64x8:
		return rewriteValueMIPS64_OpLsh64x8(v)
	case OpLsh8x16:
		return rewriteValueMIPS64_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueMIPS64_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValueMIPS64_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValueMIPS64_OpLsh8x8(v)
	case OpMIPS64ADDV:
		return rewriteValueMIPS64_OpMIPS64ADDV(v)
	case OpMIPS64ADDVconst:
		return rewriteValueMIPS64_OpMIPS64ADDVconst(v)
	case OpMIPS64AND:
		return rewriteValueMIPS64_OpMIPS64AND(v)
	case OpMIPS64ANDconst:
		return rewriteValueMIPS64_OpMIPS64ANDconst(v)
	case OpMIPS64MOVBUload:
		return rewriteValueMIPS64_OpMIPS64MOVBUload(v)
	case OpMIPS64MOVBUreg:
		return rewriteValueMIPS64_OpMIPS64MOVBUreg(v)
	case OpMIPS64MOVBload:
		return rewriteValueMIPS64_OpMIPS64MOVBload(v)
	case OpMIPS64MOVBreg:
		return rewriteValueMIPS64_OpMIPS64MOVBreg(v)
	case OpMIPS64MOVBstore:
		return rewriteValueMIPS64_OpMIPS64MOVBstore(v)
	case OpMIPS64MOVBstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVBstorezero(v)
	case OpMIPS64MOVDload:
		return rewriteValueMIPS64_OpMIPS64MOVDload(v)
	case OpMIPS64MOVDstore:
		return rewriteValueMIPS64_OpMIPS64MOVDstore(v)
	case OpMIPS64MOVFload:
		return rewriteValueMIPS64_OpMIPS64MOVFload(v)
	case OpMIPS64MOVFstore:
		return rewriteValueMIPS64_OpMIPS64MOVFstore(v)
	case OpMIPS64MOVHUload:
		return rewriteValueMIPS64_OpMIPS64MOVHUload(v)
	case OpMIPS64MOVHUreg:
		return rewriteValueMIPS64_OpMIPS64MOVHUreg(v)
	case OpMIPS64MOVHload:
		return rewriteValueMIPS64_OpMIPS64MOVHload(v)
	case OpMIPS64MOVHreg:
		return rewriteValueMIPS64_OpMIPS64MOVHreg(v)
	case OpMIPS64MOVHstore:
		return rewriteValueMIPS64_OpMIPS64MOVHstore(v)
	case OpMIPS64MOVHstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVHstorezero(v)
	case OpMIPS64MOVVload:
		return rewriteValueMIPS64_OpMIPS64MOVVload(v)
	case OpMIPS64MOVVreg:
		return rewriteValueMIPS64_OpMIPS64MOVVreg(v)
	case OpMIPS64MOVVstore:
		return rewriteValueMIPS64_OpMIPS64MOVVstore(v)
	case OpMIPS64MOVVstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVVstorezero(v)
	case OpMIPS64MOVWUload:
		return rewriteValueMIPS64_OpMIPS64MOVWUload(v)
	case OpMIPS64MOVWUreg:
		return rewriteValueMIPS64_OpMIPS64MOVWUreg(v)
	case OpMIPS64MOVWload:
		return rewriteValueMIPS64_OpMIPS64MOVWload(v)
	case OpMIPS64MOVWreg:
		return rewriteValueMIPS64_OpMIPS64MOVWreg(v)
	case OpMIPS64MOVWstore:
		return rewriteValueMIPS64_OpMIPS64MOVWstore(v)
	case OpMIPS64MOVWstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVWstorezero(v)
	case OpMIPS64NEGV:
		return rewriteValueMIPS64_OpMIPS64NEGV(v)
	case OpMIPS64NOR:
		return rewriteValueMIPS64_OpMIPS64NOR(v)
	case OpMIPS64NORconst:
		return rewriteValueMIPS64_OpMIPS64NORconst(v)
	case OpMIPS64OR:
		return rewriteValueMIPS64_OpMIPS64OR(v)
	case OpMIPS64ORconst:
		return rewriteValueMIPS64_OpMIPS64ORconst(v)
	case OpMIPS64SGT:
		return rewriteValueMIPS64_OpMIPS64SGT(v)
	case OpMIPS64SGTU:
		return rewriteValueMIPS64_OpMIPS64SGTU(v)
	case OpMIPS64SGTUconst:
		return rewriteValueMIPS64_OpMIPS64SGTUconst(v)
	case OpMIPS64SGTconst:
		return rewriteValueMIPS64_OpMIPS64SGTconst(v)
	case OpMIPS64SLLV:
		return rewriteValueMIPS64_OpMIPS64SLLV(v)
	case OpMIPS64SLLVconst:
		return rewriteValueMIPS64_OpMIPS64SLLVconst(v)
	case OpMIPS64SRAV:
		return rewriteValueMIPS64_OpMIPS64SRAV(v)
	case OpMIPS64SRAVconst:
		return rewriteValueMIPS64_OpMIPS64SRAVconst(v)
	case OpMIPS64SRLV:
		return rewriteValueMIPS64_OpMIPS64SRLV(v)
	case OpMIPS64SRLVconst:
		return rewriteValueMIPS64_OpMIPS64SRLVconst(v)
	case OpMIPS64SUBV:
		return rewriteValueMIPS64_OpMIPS64SUBV(v)
	case OpMIPS64SUBVconst:
		return rewriteValueMIPS64_OpMIPS64SUBVconst(v)
	case OpMIPS64XOR:
		return rewriteValueMIPS64_OpMIPS64XOR(v)
	case OpMIPS64XORconst:
		return rewriteValueMIPS64_OpMIPS64XORconst(v)
	case OpMod16:
		return rewriteValueMIPS64_OpMod16(v)
	case OpMod16u:
		return rewriteValueMIPS64_OpMod16u(v)
	case OpMod32:
		return rewriteValueMIPS64_OpMod32(v)
	case OpMod32u:
		return rewriteValueMIPS64_OpMod32u(v)
	case OpMod64:
		return rewriteValueMIPS64_OpMod64(v)
	case OpMod64u:
		return rewriteValueMIPS64_OpMod64u(v)
	case OpMod8:
		return rewriteValueMIPS64_OpMod8(v)
	case OpMod8u:
		return rewriteValueMIPS64_OpMod8u(v)
	case OpMove:
		return rewriteValueMIPS64_OpMove(v)
	case OpMul16:
		return rewriteValueMIPS64_OpMul16(v)
	case OpMul32:
		return rewriteValueMIPS64_OpMul32(v)
	case OpMul32F:
		return rewriteValueMIPS64_OpMul32F(v)
	case OpMul64:
		return rewriteValueMIPS64_OpMul64(v)
	case OpMul64F:
		return rewriteValueMIPS64_OpMul64F(v)
	case OpMul8:
		return rewriteValueMIPS64_OpMul8(v)
	case OpNeg16:
		return rewriteValueMIPS64_OpNeg16(v)
	case OpNeg32:
		return rewriteValueMIPS64_OpNeg32(v)
	case OpNeg32F:
		return rewriteValueMIPS64_OpNeg32F(v)
	case OpNeg64:
		return rewriteValueMIPS64_OpNeg64(v)
	case OpNeg64F:
		return rewriteValueMIPS64_OpNeg64F(v)
	case OpNeg8:
		return rewriteValueMIPS64_OpNeg8(v)
	case OpNeq16:
		return rewriteValueMIPS64_OpNeq16(v)
	case OpNeq32:
		return rewriteValueMIPS64_OpNeq32(v)
	case OpNeq32F:
		return rewriteValueMIPS64_OpNeq32F(v)
	case OpNeq64:
		return rewriteValueMIPS64_OpNeq64(v)
	case OpNeq64F:
		return rewriteValueMIPS64_OpNeq64F(v)
	case OpNeq8:
		return rewriteValueMIPS64_OpNeq8(v)
	case OpNeqB:
		return rewriteValueMIPS64_OpNeqB(v)
	case OpNeqPtr:
		return rewriteValueMIPS64_OpNeqPtr(v)
	case OpNilCheck:
		return rewriteValueMIPS64_OpNilCheck(v)
	case OpNot:
		return rewriteValueMIPS64_OpNot(v)
	case OpOffPtr:
		return rewriteValueMIPS64_OpOffPtr(v)
	case OpOr16:
		return rewriteValueMIPS64_OpOr16(v)
	case OpOr32:
		return rewriteValueMIPS64_OpOr32(v)
	case OpOr64:
		return rewriteValueMIPS64_OpOr64(v)
	case OpOr8:
		return rewriteValueMIPS64_OpOr8(v)
	case OpOrB:
		return rewriteValueMIPS64_OpOrB(v)
	case OpRound32F:
		return rewriteValueMIPS64_OpRound32F(v)
	case OpRound64F:
		return rewriteValueMIPS64_OpRound64F(v)
	case OpRsh16Ux16:
		return rewriteValueMIPS64_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueMIPS64_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueMIPS64_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueMIPS64_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueMIPS64_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueMIPS64_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueMIPS64_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueMIPS64_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueMIPS64_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueMIPS64_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueMIPS64_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueMIPS64_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueMIPS64_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueMIPS64_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueMIPS64_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueMIPS64_OpRsh32x8(v)
	case OpRsh64Ux16:
		return rewriteValueMIPS64_OpRsh64Ux16(v)
	case OpRsh64Ux32:
		return rewriteValueMIPS64_OpRsh64Ux32(v)
	case OpRsh64Ux64:
		return rewriteValueMIPS64_OpRsh64Ux64(v)
	case OpRsh64Ux8:
		return rewriteValueMIPS64_OpRsh64Ux8(v)
	case OpRsh64x16:
		return rewriteValueMIPS64_OpRsh64x16(v)
	case OpRsh64x32:
		return rewriteValueMIPS64_OpRsh64x32(v)
	case OpRsh64x64:
		return rewriteValueMIPS64_OpRsh64x64(v)
	case OpRsh64x8:
		return rewriteValueMIPS64_OpRsh64x8(v)
	case OpRsh8Ux16:
		return rewriteValueMIPS64_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueMIPS64_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueMIPS64_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueMIPS64_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueMIPS64_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueMIPS64_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueMIPS64_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueMIPS64_OpRsh8x8(v)
	case OpSelect0:
		return rewriteValueMIPS64_OpSelect0(v)
	case OpSelect1:
		return rewriteValueMIPS64_OpSelect1(v)
	case OpSignExt16to32:
		return rewriteValueMIPS64_OpSignExt16to32(v)
	case OpSignExt16to64:
		return rewriteValueMIPS64_OpSignExt16to64(v)
	case OpSignExt32to64:
		return rewriteValueMIPS64_OpSignExt32to64(v)
	case OpSignExt8to16:
		return rewriteValueMIPS64_OpSignExt8to16(v)
	case OpSignExt8to32:
		return rewriteValueMIPS64_OpSignExt8to32(v)
	case OpSignExt8to64:
		return rewriteValueMIPS64_OpSignExt8to64(v)
	case OpSlicemask:
		return rewriteValueMIPS64_OpSlicemask(v)
	case OpStaticCall:
		return rewriteValueMIPS64_OpStaticCall(v)
	case OpStore:
		return rewriteValueMIPS64_OpStore(v)
	case OpSub16:
		return rewriteValueMIPS64_OpSub16(v)
	case OpSub32:
		return rewriteValueMIPS64_OpSub32(v)
	case OpSub32F:
		return rewriteValueMIPS64_OpSub32F(v)
	case OpSub64:
		return rewriteValueMIPS64_OpSub64(v)
	case OpSub64F:
		return rewriteValueMIPS64_OpSub64F(v)
	case OpSub8:
		return rewriteValueMIPS64_OpSub8(v)
	case OpSubPtr:
		return rewriteValueMIPS64_OpSubPtr(v)
	case OpTrunc16to8:
		return rewriteValueMIPS64_OpTrunc16to8(v)
	case OpTrunc32to16:
		return rewriteValueMIPS64_OpTrunc32to16(v)
	case OpTrunc32to8:
		return rewriteValueMIPS64_OpTrunc32to8(v)
	case OpTrunc64to16:
		return rewriteValueMIPS64_OpTrunc64to16(v)
	case OpTrunc64to32:
		return rewriteValueMIPS64_OpTrunc64to32(v)
	case OpTrunc64to8:
		return rewriteValueMIPS64_OpTrunc64to8(v)
	case OpXor16:
		return rewriteValueMIPS64_OpXor16(v)
	case OpXor32:
		return rewriteValueMIPS64_OpXor32(v)
	case OpXor64:
		return rewriteValueMIPS64_OpXor64(v)
	case OpXor8:
		return rewriteValueMIPS64_OpXor8(v)
	case OpZero:
		return rewriteValueMIPS64_OpZero(v)
	case OpZeroExt16to32:
		return rewriteValueMIPS64_OpZeroExt16to32(v)
	case OpZeroExt16to64:
		return rewriteValueMIPS64_OpZeroExt16to64(v)
	case OpZeroExt32to64:
		return rewriteValueMIPS64_OpZeroExt32to64(v)
	case OpZeroExt8to16:
		return rewriteValueMIPS64_OpZeroExt8to16(v)
	case OpZeroExt8to32:
		return rewriteValueMIPS64_OpZeroExt8to32(v)
	case OpZeroExt8to64:
		return rewriteValueMIPS64_OpZeroExt8to64(v)
	}
	return false
}
func rewriteValueMIPS64_OpAdd16(v *Value) bool {
	// match: (Add16 x y)
	// cond:
	// result: (ADDV x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAdd32(v *Value) bool {
	// match: (Add32 x y)
	// cond:
	// result: (ADDV x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAdd32F(v *Value) bool {
	// match: (Add32F x y)
	// cond:
	// result: (ADDF x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAdd64(v *Value) bool {
	// match: (Add64 x y)
	// cond:
	// result: (ADDV x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAdd64F(v *Value) bool {
	// match: (Add64F x y)
	// cond:
	// result: (ADDD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAdd8(v *Value) bool {
	// match: (Add8 x y)
	// cond:
	// result: (ADDV x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAddPtr(v *Value) bool {
	// match: (AddPtr x y)
	// cond:
	// result: (ADDV x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAddr(v *Value) bool {
	// match: (Addr {sym} base)
	// cond:
	// result: (MOVVaddr {sym} base)
	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpMIPS64MOVVaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueMIPS64_OpAnd16(v *Value) bool {
	// match: (And16 x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAnd32(v *Value) bool {
	// match: (And32 x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAnd64(v *Value) bool {
	// match: (And64 x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAnd8(v *Value) bool {
	// match: (And8 x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAndB(v *Value) bool {
	// match: (AndB x y)
	// cond:
	// result: (AND x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAvg64u(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Avg64u <t> x y)
	// cond:
	// result: (ADDV (SRLVconst <t> (SUBV <t> x y) [1]) y)
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v0 := b.NewValue0(v.Pos, OpMIPS64SRLVconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpMIPS64SUBV, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpClosureCall(v *Value) bool {
	// match: (ClosureCall [argwid] entry closure mem)
	// cond:
	// result: (CALLclosure [argwid] entry closure mem)
	for {
		argwid := v.AuxInt
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64CALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpCom16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Com16 x)
	// cond:
	// result: (NOR (MOVVconst [0]) x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCom32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Com32 x)
	// cond:
	// result: (NOR (MOVVconst [0]) x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCom64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Com64 x)
	// cond:
	// result: (NOR (MOVVconst [0]) x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCom8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Com8 x)
	// cond:
	// result: (NOR (MOVVconst [0]) x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpConst16(v *Value) bool {
	// match: (Const16 [val])
	// cond:
	// result: (MOVVconst [val])
	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConst32(v *Value) bool {
	// match: (Const32 [val])
	// cond:
	// result: (MOVVconst [val])
	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConst32F(v *Value) bool {
	// match: (Const32F [val])
	// cond:
	// result: (MOVFconst [val])
	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVFconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConst64(v *Value) bool {
	// match: (Const64 [val])
	// cond:
	// result: (MOVVconst [val])
	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConst64F(v *Value) bool {
	// match: (Const64F [val])
	// cond:
	// result: (MOVDconst [val])
	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConst8(v *Value) bool {
	// match: (Const8 [val])
	// cond:
	// result: (MOVVconst [val])
	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConstBool(v *Value) bool {
	// match: (ConstBool [b])
	// cond:
	// result: (MOVVconst [b])
	for {
		b := v.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValueMIPS64_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// cond:
	// result: (MOVVconst [0])
	for {
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueMIPS64_OpConvert(v *Value) bool {
	// match: (Convert x mem)
	// cond:
	// result: (MOVVconvert x mem)
	for {
		x := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64MOVVconvert)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpCvt32Fto32(v *Value) bool {
	// match: (Cvt32Fto32 x)
	// cond:
	// result: (TRUNCFW x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64TRUNCFW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt32Fto64(v *Value) bool {
	// match: (Cvt32Fto64 x)
	// cond:
	// result: (TRUNCFV x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64TRUNCFV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt32Fto64F(v *Value) bool {
	// match: (Cvt32Fto64F x)
	// cond:
	// result: (MOVFD x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVFD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt32to32F(v *Value) bool {
	// match: (Cvt32to32F x)
	// cond:
	// result: (MOVWF x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVWF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt32to64F(v *Value) bool {
	// match: (Cvt32to64F x)
	// cond:
	// result: (MOVWD x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVWD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt64Fto32(v *Value) bool {
	// match: (Cvt64Fto32 x)
	// cond:
	// result: (TRUNCDW x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64TRUNCDW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt64Fto32F(v *Value) bool {
	// match: (Cvt64Fto32F x)
	// cond:
	// result: (MOVDF x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVDF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt64Fto64(v *Value) bool {
	// match: (Cvt64Fto64 x)
	// cond:
	// result: (TRUNCDV x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64TRUNCDV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt64to32F(v *Value) bool {
	// match: (Cvt64to32F x)
	// cond:
	// result: (MOVVF x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVVF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt64to64F(v *Value) bool {
	// match: (Cvt64to64F x)
	// cond:
	// result: (MOVVD x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVVD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpDiv16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div16 x y)
	// cond:
	// result: (Select1 (DIVV (SignExt16to64 x) (SignExt16to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, MakeTuple(types.Int64, types.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv16u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div16u x y)
	// cond:
	// result: (Select1 (DIVVU (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, MakeTuple(types.UInt64, types.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div32 x y)
	// cond:
	// result: (Select1 (DIVV (SignExt32to64 x) (SignExt32to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, MakeTuple(types.Int64, types.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv32F(v *Value) bool {
	// match: (Div32F x y)
	// cond:
	// result: (DIVF x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64DIVF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpDiv32u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div32u x y)
	// cond:
	// result: (Select1 (DIVVU (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, MakeTuple(types.UInt64, types.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div64 x y)
	// cond:
	// result: (Select1 (DIVV x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, MakeTuple(types.Int64, types.Int64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv64F(v *Value) bool {
	// match: (Div64F x y)
	// cond:
	// result: (DIVD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64DIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpDiv64u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div64u x y)
	// cond:
	// result: (Select1 (DIVVU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, MakeTuple(types.UInt64, types.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div8 x y)
	// cond:
	// result: (Select1 (DIVV (SignExt8to64 x) (SignExt8to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, MakeTuple(types.Int64, types.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv8u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Div8u x y)
	// cond:
	// result: (Select1 (DIVVU (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, MakeTuple(types.UInt64, types.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpEq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Eq16 x y)
	// cond:
	// result: (SGTU (MOVVconst [1]) (XOR (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, types.UInt64)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEq32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Eq32 x y)
	// cond:
	// result: (SGTU (MOVVconst [1]) (XOR (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, types.UInt64)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Eq32F x y)
	// cond:
	// result: (FPFlagTrue (CMPEQF x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQF, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpEq64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Eq64 x y)
	// cond:
	// result: (SGTU (MOVVconst [1]) (XOR x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, types.UInt64)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Eq64F x y)
	// cond:
	// result: (FPFlagTrue (CMPEQD x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQD, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpEq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Eq8 x y)
	// cond:
	// result: (SGTU (MOVVconst [1]) (XOR (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, types.UInt64)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEqB(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (EqB x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (XOR <types.Bool> x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, types.Bool)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEqPtr(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (EqPtr x y)
	// cond:
	// result: (SGTU (MOVVconst [1]) (XOR x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, types.UInt64)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq16 x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGT (SignExt16to64 y) (SignExt16to64 x)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, types.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq16U x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGTU (ZeroExt16to64 y) (ZeroExt16to64 x)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq32 x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGT (SignExt32to64 y) (SignExt32to64 x)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, types.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Geq32F x y)
	// cond:
	// result: (FPFlagTrue (CMPGEF x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGEF, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpGeq32U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq32U x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGTU (ZeroExt32to64 y) (ZeroExt32to64 x)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq64 x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGT y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, types.Bool)
		v1.AddArg(y)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Geq64F x y)
	// cond:
	// result: (FPFlagTrue (CMPGED x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGED, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpGeq64U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq64U x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGTU y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v1.AddArg(y)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq8 x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGT (SignExt8to64 y) (SignExt8to64 x)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, types.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Geq8U x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGTU (ZeroExt8to64 y) (ZeroExt8to64 x)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGetClosurePtr(v *Value) bool {
	// match: (GetClosurePtr)
	// cond:
	// result: (LoweredGetClosurePtr)
	for {
		v.reset(OpMIPS64LoweredGetClosurePtr)
		return true
	}
}
func rewriteValueMIPS64_OpGreater16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater16 x y)
	// cond:
	// result: (SGT (SignExt16to64 x) (SignExt16to64 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGreater16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater16U x y)
	// cond:
	// result: (SGTU (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGreater32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater32 x y)
	// cond:
	// result: (SGT (SignExt32to64 x) (SignExt32to64 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGreater32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Greater32F x y)
	// cond:
	// result: (FPFlagTrue (CMPGTF x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTF, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpGreater32U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater32U x y)
	// cond:
	// result: (SGTU (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGreater64(v *Value) bool {
	// match: (Greater64 x y)
	// cond:
	// result: (SGT x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpGreater64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Greater64F x y)
	// cond:
	// result: (FPFlagTrue (CMPGTD x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTD, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpGreater64U(v *Value) bool {
	// match: (Greater64U x y)
	// cond:
	// result: (SGTU x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpGreater8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater8 x y)
	// cond:
	// result: (SGT (SignExt8to64 x) (SignExt8to64 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGreater8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Greater8U x y)
	// cond:
	// result: (SGTU (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpHmul32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Hmul32 x y)
	// cond:
	// result: (SRAVconst (Select1 <types.Int64> (MULV (SignExt32to64 x) (SignExt32to64 y))) [32])
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpSelect1, types.Int64)
		v1 := b.NewValue0(v.Pos, OpMIPS64MULV, MakeTuple(types.Int64, types.Int64))
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpHmul32u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Hmul32u x y)
	// cond:
	// result: (SRLVconst (Select1 <types.UInt64> (MULVU (ZeroExt32to64 x) (ZeroExt32to64 y))) [32])
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRLVconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpSelect1, types.UInt64)
		v1 := b.NewValue0(v.Pos, OpMIPS64MULVU, MakeTuple(types.UInt64, types.UInt64))
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpHmul64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Hmul64 x y)
	// cond:
	// result: (Select0 (MULV x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULV, MakeTuple(types.Int64, types.Int64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpHmul64u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Hmul64u x y)
	// cond:
	// result: (Select0 (MULVU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, MakeTuple(types.UInt64, types.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpInterCall(v *Value) bool {
	// match: (InterCall [argwid] entry mem)
	// cond:
	// result: (CALLinter [argwid] entry mem)
	for {
		argwid := v.AuxInt
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64CALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpIsInBounds(v *Value) bool {
	// match: (IsInBounds idx len)
	// cond:
	// result: (SGTU len idx)
	for {
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v.AddArg(len)
		v.AddArg(idx)
		return true
	}
}
func rewriteValueMIPS64_OpIsNonNil(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (IsNonNil ptr)
	// cond:
	// result: (SGTU ptr (MOVVconst [0]))
	for {
		ptr := v.Args[0]
		v.reset(OpMIPS64SGTU)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpIsSliceInBounds(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (IsSliceInBounds idx len)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGTU idx len))
	for {
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v1.AddArg(idx)
		v1.AddArg(len)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq16 x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGT (SignExt16to64 x) (SignExt16to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, types.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq16U x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGTU (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq32 x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGT (SignExt32to64 x) (SignExt32to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, types.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Leq32F x y)
	// cond:
	// result: (FPFlagTrue (CMPGEF y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGEF, TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLeq32U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq32U x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGTU (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq64 x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGT x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, types.Bool)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Leq64F x y)
	// cond:
	// result: (FPFlagTrue (CMPGED y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGED, TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLeq64U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq64U x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGTU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq8 x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGT (SignExt8to64 x) (SignExt8to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, types.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Leq8U x y)
	// cond:
	// result: (XOR (MOVVconst [1]) (SGTU (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less16 x y)
	// cond:
	// result: (SGT (SignExt16to64 y) (SignExt16to64 x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess16U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less16U x y)
	// cond:
	// result: (SGTU (ZeroExt16to64 y) (ZeroExt16to64 x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less32 x y)
	// cond:
	// result: (SGT (SignExt32to64 y) (SignExt32to64 x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Less32F x y)
	// cond:
	// result: (FPFlagTrue (CMPGTF y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTF, TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLess32U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less32U x y)
	// cond:
	// result: (SGTU (ZeroExt32to64 y) (ZeroExt32to64 x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess64(v *Value) bool {
	// match: (Less64 x y)
	// cond:
	// result: (SGT y x)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpLess64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Less64F x y)
	// cond:
	// result: (FPFlagTrue (CMPGTD y x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTD, TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLess64U(v *Value) bool {
	// match: (Less64U x y)
	// cond:
	// result: (SGTU y x)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpLess8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less8 x y)
	// cond:
	// result: (SGT (SignExt8to64 y) (SignExt8to64 x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess8U(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Less8U x y)
	// cond:
	// result: (SGTU (ZeroExt8to64 y) (ZeroExt8to64 x))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLoad(v *Value) bool {
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
		v.reset(OpMIPS64MOVBUload)
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
		v.reset(OpMIPS64MOVBload)
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
		v.reset(OpMIPS64MOVBUload)
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
		v.reset(OpMIPS64MOVHload)
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
		v.reset(OpMIPS64MOVHUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) && isSigned(t))
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is32BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpMIPS64MOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) && !isSigned(t))
	// result: (MOVWUload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is32BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpMIPS64MOVWUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (MOVVload ptr mem)
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpMIPS64MOVVload)
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
		v.reset(OpMIPS64MOVFload)
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
		v.reset(OpMIPS64MOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpLsh16x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh16x16 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt16to64 y))) (SLLV <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh16x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh16x32 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt32to64 y))) (SLLV <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh16x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh16x64 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) y)) (SLLV <t> x y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh16x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh16x8 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt8to64  y))) (SLLV <t> x (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh32x16 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt16to64 y))) (SLLV <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh32x32 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt32to64 y))) (SLLV <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh32x64 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) y)) (SLLV <t> x y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh32x8 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt8to64  y))) (SLLV <t> x (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh64x16 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt16to64 y))) (SLLV <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh64x32 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt32to64 y))) (SLLV <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh64x64 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) y)) (SLLV <t> x y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh64x8 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt8to64  y))) (SLLV <t> x (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh8x16 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt16to64 y))) (SLLV <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh8x32 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt32to64 y))) (SLLV <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh8x64 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) y)) (SLLV <t> x y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Lsh8x8 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt8to64  y))) (SLLV <t> x (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpMIPS64ADDV(v *Value) bool {
	// match: (ADDV x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (ADDVconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (ADDV (MOVVconst [c]) x)
	// cond: is32Bit(c)
	// result: (ADDVconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (ADDV x (NEGV y))
	// cond:
	// result: (SUBV x y)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64NEGV {
			break
		}
		y := v_1.Args[0]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (ADDV (NEGV y) x)
	// cond:
	// result: (SUBV x y)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64NEGV {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64ADDVconst(v *Value) bool {
	// match: (ADDVconst [off1] (MOVVaddr [off2] {sym} ptr))
	// cond:
	// result: (MOVVaddr [off1+off2] {sym} ptr)
	for {
		off1 := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpMIPS64MOVVaddr)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		return true
	}
	// match: (ADDVconst [0] x)
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
	// match: (ADDVconst [c] (MOVVconst [d]))
	// cond:
	// result: (MOVVconst [c+d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c + d
		return true
	}
	// match: (ADDVconst [c] (ADDVconst [d] x))
	// cond: is32Bit(c+d)
	// result: (ADDVconst [c+d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}
	// match: (ADDVconst [c] (SUBVconst [d] x))
	// cond: is32Bit(c-d)
	// result: (ADDVconst [c-d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64SUBVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c - d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = c - d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64AND(v *Value) bool {
	// match: (AND x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (ANDconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (AND (MOVVconst [c]) x)
	// cond: is32Bit(c)
	// result: (ANDconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ANDconst)
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
	return false
}
func rewriteValueMIPS64_OpMIPS64ANDconst(v *Value) bool {
	// match: (ANDconst [0] _)
	// cond:
	// result: (MOVVconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
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
	// match: (ANDconst [c] (MOVVconst [d]))
	// cond:
	// result: (MOVVconst [c&d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c & d
		return true
	}
	// match: (ANDconst [c] (ANDconst [d] x))
	// cond:
	// result: (ANDconst [c&d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPS64ANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBUload(v *Value) bool {
	// match: (MOVBUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBUreg(v *Value) bool {
	// match: (MOVBUreg x:(MOVBUload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg x:(MOVBUreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (MOVVconst [c]))
	// cond:
	// result: (MOVVconst [int64(uint8(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBload(v *Value) bool {
	// match: (MOVBload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBload  [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBreg(v *Value) bool {
	// match: (MOVBreg x:(MOVBload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVVconst [c]))
	// cond:
	// result: (MOVVconst [int64(int8(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBstore(v *Value) bool {
	// match: (MOVBstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVVconst [0]) mem)
	// cond:
	// result: (MOVBstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstorezero)
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
		if v_1.Op != OpMIPS64MOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
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
		if v_1.Op != OpMIPS64MOVBUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
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
		if v_1.Op != OpMIPS64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
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
		if v_1.Op != OpMIPS64MOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
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
		if v_1.Op != OpMIPS64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVWUreg x) mem)
	// cond:
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBstorezero(v *Value) bool {
	// match: (MOVBstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstorezero [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVDload(v *Value) bool {
	// match: (MOVDload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVDload  [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVDstore(v *Value) bool {
	// match: (MOVDstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVFload(v *Value) bool {
	// match: (MOVFload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVFload  [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVFload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVFload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVFload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVFstore(v *Value) bool {
	// match: (MOVFstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVFstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVFstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHUload(v *Value) bool {
	// match: (MOVHUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHUreg(v *Value) bool {
	// match: (MOVHUreg x:(MOVBUload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHUload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVBUreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (MOVVconst [c]))
	// cond:
	// result: (MOVVconst [int64(uint16(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHload(v *Value) bool {
	// match: (MOVHload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHload  [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHreg(v *Value) bool {
	// match: (MOVHreg x:(MOVBload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVVconst [c]))
	// cond:
	// result: (MOVVconst [int64(int16(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHstore(v *Value) bool {
	// match: (MOVHstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVVconst [0]) mem)
	// cond:
	// result: (MOVHstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPS64MOVHstorezero)
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
		if v_1.Op != OpMIPS64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVHstore)
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
		if v_1.Op != OpMIPS64MOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVHstore)
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
		if v_1.Op != OpMIPS64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVWUreg x) mem)
	// cond:
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHstorezero(v *Value) bool {
	// match: (MOVHstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstorezero [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVload(v *Value) bool {
	// match: (MOVVload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVVload  [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVVload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVVload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVVload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVVload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVreg(v *Value) bool {
	// match: (MOVVreg x)
	// cond: x.Uses == 1
	// result: (MOVVnop x)
	for {
		x := v.Args[0]
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpMIPS64MOVVnop)
		v.AddArg(x)
		return true
	}
	// match: (MOVVreg (MOVVconst [c]))
	// cond:
	// result: (MOVVconst [c])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVstore(v *Value) bool {
	// match: (MOVVstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVVstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVVstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVVstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVVstore [off] {sym} ptr (MOVVconst [0]) mem)
	// cond:
	// result: (MOVVstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPS64MOVVstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVstorezero(v *Value) bool {
	// match: (MOVVstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVVstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVVstorezero [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVVstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWUload(v *Value) bool {
	// match: (MOVWUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWUload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWUreg(v *Value) bool {
	// match: (MOVWUreg x:(MOVBUload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg x:(MOVHUload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHUload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg x:(MOVWUload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVWUload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg x:(MOVBUreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg x:(MOVHUreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg x:(MOVWUreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVWUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg (MOVVconst [c]))
	// cond:
	// result: (MOVVconst [int64(uint32(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint32(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWload(v *Value) bool {
	// match: (MOVWload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWload  [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWreg(v *Value) bool {
	// match: (MOVWreg x:(MOVBload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVBUload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHUload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHUload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVWload _ _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVWload {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVBreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVBUreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVWreg _))
	// cond:
	// result: (MOVVreg x)
	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVWreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (MOVVconst [c]))
	// cond:
	// result: (MOVVconst [int64(int32(c))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(int32(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWstore(v *Value) bool {
	// match: (MOVWstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVVconst [0]) mem)
	// cond:
	// result: (MOVWstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPS64MOVWstorezero)
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
		if v_1.Op != OpMIPS64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWUreg x) mem)
	// cond:
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWstorezero(v *Value) bool {
	// match: (MOVWstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstorezero [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64NEGV(v *Value) bool {
	// match: (NEGV (MOVVconst [c]))
	// cond:
	// result: (MOVVconst [-c])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = -c
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64NOR(v *Value) bool {
	// match: (NOR x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (NORconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64NORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (NOR (MOVVconst [c]) x)
	// cond: is32Bit(c)
	// result: (NORconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64NORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64NORconst(v *Value) bool {
	// match: (NORconst [c] (MOVVconst [d]))
	// cond:
	// result: (MOVVconst [^(c|d)])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = ^(c | d)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64OR(v *Value) bool {
	// match: (OR x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (ORconst  [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (OR (MOVVconst [c]) x)
	// cond: is32Bit(c)
	// result: (ORconst  [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ORconst)
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
	return false
}
func rewriteValueMIPS64_OpMIPS64ORconst(v *Value) bool {
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
	// result: (MOVVconst [-1])
	for {
		if v.AuxInt != -1 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = -1
		return true
	}
	// match: (ORconst [c] (MOVVconst [d]))
	// cond:
	// result: (MOVVconst [c|d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c | d
		return true
	}
	// match: (ORconst [c] (ORconst [d] x))
	// cond: is32Bit(c|d)
	// result: (ORconst [c|d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c | d)) {
			break
		}
		v.reset(OpMIPS64ORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGT(v *Value) bool {
	// match: (SGT (MOVVconst [c]) x)
	// cond: is32Bit(c)
	// result: (SGTconst  [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64SGTconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGTU(v *Value) bool {
	// match: (SGTU (MOVVconst [c]) x)
	// cond: is32Bit(c)
	// result: (SGTUconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64SGTUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGTUconst(v *Value) bool {
	// match: (SGTUconst [c] (MOVVconst [d]))
	// cond: uint64(c)>uint64(d)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(uint64(c) > uint64(d)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (MOVVconst [d]))
	// cond: uint64(c)<=uint64(d)
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(uint64(c) <= uint64(d)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTUconst [c] (MOVBUreg _))
	// cond: 0xff < uint64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVBUreg {
			break
		}
		if !(0xff < uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (MOVHUreg _))
	// cond: 0xffff < uint64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVHUreg {
			break
		}
		if !(0xffff < uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (ANDconst [m] _))
	// cond: uint64(m) < uint64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ANDconst {
			break
		}
		m := v_0.AuxInt
		if !(uint64(m) < uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (SRLVconst _ [d]))
	// cond: 0 < d && d <= 63 && 1<<uint64(64-d) <= uint64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64SRLVconst {
			break
		}
		d := v_0.AuxInt
		if !(0 < d && d <= 63 && 1<<uint64(64-d) <= uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGTconst(v *Value) bool {
	// match: (SGTconst [c] (MOVVconst [d]))
	// cond: int64(c)>int64(d)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(int64(c) > int64(d)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVVconst [d]))
	// cond: int64(c)<=int64(d)
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(int64(c) <= int64(d)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVBreg _))
	// cond: 0x7f < int64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVBreg {
			break
		}
		if !(0x7f < int64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVBreg _))
	// cond: int64(c) <= -0x80
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVBreg {
			break
		}
		if !(int64(c) <= -0x80) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVBUreg _))
	// cond: 0xff < int64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVBUreg {
			break
		}
		if !(0xff < int64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVBUreg _))
	// cond: int64(c) < 0
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVBUreg {
			break
		}
		if !(int64(c) < 0) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVHreg _))
	// cond: 0x7fff < int64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVHreg {
			break
		}
		if !(0x7fff < int64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVHreg _))
	// cond: int64(c) <= -0x8000
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVHreg {
			break
		}
		if !(int64(c) <= -0x8000) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVHUreg _))
	// cond: 0xffff < int64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVHUreg {
			break
		}
		if !(0xffff < int64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVHUreg _))
	// cond: int64(c) < 0
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVHUreg {
			break
		}
		if !(int64(c) < 0) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVWUreg _))
	// cond: int64(c) < 0
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVWUreg {
			break
		}
		if !(int64(c) < 0) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (ANDconst [m] _))
	// cond: 0 <= m && m < c
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ANDconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= m && m < c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (SRLVconst _ [d]))
	// cond: 0 <= c && 0 < d && d <= 63 && 1<<uint64(64-d) <= c
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64SRLVconst {
			break
		}
		d := v_0.AuxInt
		if !(0 <= c && 0 < d && d <= 63 && 1<<uint64(64-d) <= c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SLLV(v *Value) bool {
	// match: (SLLV _ (MOVVconst [c]))
	// cond: uint64(c)>=64
	// result: (MOVVconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SLLV x (MOVVconst [c]))
	// cond:
	// result: (SLLVconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPS64SLLVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SLLVconst(v *Value) bool {
	// match: (SLLVconst [c] (MOVVconst [d]))
	// cond:
	// result: (MOVVconst [int64(d)<<uint64(c)])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(d) << uint64(c)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRAV(v *Value) bool {
	// match: (SRAV x (MOVVconst [c]))
	// cond: uint64(c)>=64
	// result: (SRAVconst x [63])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = 63
		v.AddArg(x)
		return true
	}
	// match: (SRAV x (MOVVconst [c]))
	// cond:
	// result: (SRAVconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRAVconst(v *Value) bool {
	// match: (SRAVconst [c] (MOVVconst [d]))
	// cond:
	// result: (MOVVconst [int64(d)>>uint64(c)])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(d) >> uint64(c)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRLV(v *Value) bool {
	// match: (SRLV _ (MOVVconst [c]))
	// cond: uint64(c)>=64
	// result: (MOVVconst [0])
	for {
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SRLV x (MOVVconst [c]))
	// cond:
	// result: (SRLVconst x [c])
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPS64SRLVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRLVconst(v *Value) bool {
	// match: (SRLVconst [c] (MOVVconst [d]))
	// cond:
	// result: (MOVVconst [int64(uint64(d)>>uint64(c))])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint64(d) >> uint64(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SUBV(v *Value) bool {
	// match: (SUBV x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (SUBVconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64SUBVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUBV x x)
	// cond:
	// result: (MOVVconst [0])
	for {
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SUBV (MOVVconst [0]) x)
	// cond:
	// result: (NEGV x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SUBVconst(v *Value) bool {
	// match: (SUBVconst [0] x)
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
	// match: (SUBVconst [c] (MOVVconst [d]))
	// cond:
	// result: (MOVVconst [d-c])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = d - c
		return true
	}
	// match: (SUBVconst [c] (SUBVconst [d] x))
	// cond: is32Bit(-c-d)
	// result: (ADDVconst [-c-d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64SUBVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(-c - d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = -c - d
		v.AddArg(x)
		return true
	}
	// match: (SUBVconst [c] (ADDVconst [d] x))
	// cond: is32Bit(-c+d)
	// result: (ADDVconst [-c+d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(-c + d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = -c + d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64XOR(v *Value) bool {
	// match: (XOR x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (XORconst [c] x)
	for {
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64XORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (XOR (MOVVconst [c]) x)
	// cond: is32Bit(c)
	// result: (XORconst [c] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64XORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (XOR x x)
	// cond:
	// result: (MOVVconst [0])
	for {
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64XORconst(v *Value) bool {
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
		v.reset(OpMIPS64NORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
	// match: (XORconst [c] (MOVVconst [d]))
	// cond:
	// result: (MOVVconst [c^d])
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c ^ d
		return true
	}
	// match: (XORconst [c] (XORconst [d] x))
	// cond: is32Bit(c^d)
	// result: (XORconst [c^d] x)
	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64XORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c ^ d)) {
			break
		}
		v.reset(OpMIPS64XORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMod16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod16 x y)
	// cond:
	// result: (Select0 (DIVV (SignExt16to64 x) (SignExt16to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, MakeTuple(types.Int64, types.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod16u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod16u x y)
	// cond:
	// result: (Select0 (DIVVU (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, MakeTuple(types.UInt64, types.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod32 x y)
	// cond:
	// result: (Select0 (DIVV (SignExt32to64 x) (SignExt32to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, MakeTuple(types.Int64, types.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod32u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod32u x y)
	// cond:
	// result: (Select0 (DIVVU (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, MakeTuple(types.UInt64, types.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod64 x y)
	// cond:
	// result: (Select0 (DIVV x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, MakeTuple(types.Int64, types.Int64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod64u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod64u x y)
	// cond:
	// result: (Select0 (DIVVU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, MakeTuple(types.UInt64, types.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod8 x y)
	// cond:
	// result: (Select0 (DIVV (SignExt8to64 x) (SignExt8to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, MakeTuple(types.Int64, types.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod8u(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mod8u x y)
	// cond:
	// result: (Select0 (DIVVU (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, MakeTuple(types.UInt64, types.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMove(v *Value) bool {
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
	// result: (MOVBstore dst (MOVBload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, types.Int8)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] {t} dst src mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore dst (MOVHload src mem) mem)
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
		v.reset(OpMIPS64MOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, types.Int16)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// cond:
	// result: (MOVBstore [1] dst (MOVBload [1] src mem) 		(MOVBstore dst (MOVBload src mem) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 1
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, types.Int8)
		v0.AuxInt = 1
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVBload, types.Int8)
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
		v.reset(OpMIPS64MOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVWload, types.Int32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore [2] dst (MOVHload [2] src mem) 		(MOVHstore dst (MOVHload src mem) mem))
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
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, types.Int16)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVHload, types.Int16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [4] dst src mem)
	// cond:
	// result: (MOVBstore [3] dst (MOVBload [3] src mem) 		(MOVBstore [2] dst (MOVBload [2] src mem) 			(MOVBstore [1] dst (MOVBload [1] src mem) 				(MOVBstore dst (MOVBload src mem) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, types.Int8)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVBload, types.Int8)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v3.AuxInt = 1
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVBload, types.Int8)
		v4.AuxInt = 1
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVBload, types.Int8)
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
	// result: (MOVVstore dst (MOVVload src mem) mem)
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
		v.reset(OpMIPS64MOVVstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVload, types.UInt64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
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
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVWload, types.Int32)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVWload, types.Int32)
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
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, types.Int16)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVHload, types.Int16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v3.AuxInt = 2
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVHload, types.Int16)
		v4.AuxInt = 2
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVHload, types.Int16)
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
	// result: (MOVBstore [2] dst (MOVBload [2] src mem) 		(MOVBstore [1] dst (MOVBload [1] src mem) 			(MOVBstore dst (MOVBload src mem) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, types.Int8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v1.AuxInt = 1
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVBload, types.Int8)
		v2.AuxInt = 1
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVBload, types.Int8)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
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
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, types.Int16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVHload, types.Int16)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVHload, types.Int16)
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
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVWload, types.Int32)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVWload, types.Int32)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVWload, types.Int32)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [16] {t} dst src mem)
	// cond: t.(Type).Alignment()%8 == 0
	// result: (MOVVstore [8] dst (MOVVload [8] src mem) 		(MOVVstore dst (MOVVload src mem) mem))
	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVload, types.UInt64)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVload, types.UInt64)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [24] {t} dst src mem)
	// cond: t.(Type).Alignment()%8 == 0
	// result: (MOVVstore [16] dst (MOVVload [16] src mem) 		(MOVVstore [8] dst (MOVVload [8] src mem) 			(MOVVstore dst (MOVVload src mem) mem)))
	for {
		if v.AuxInt != 24 {
			break
		}
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 16
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVload, types.UInt64)
		v0.AuxInt = 16
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, TypeMem)
		v1.AuxInt = 8
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVload, types.UInt64)
		v2.AuxInt = 8
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVload, types.UInt64)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: s > 24 || t.(Type).Alignment()%8 != 0
	// result: (LoweredMove [t.(Type).Alignment()] 		dst 		src 		(ADDVconst <src.Type> src [s-moveSize(t.(Type).Alignment(), config)]) 		mem)
	for {
		s := v.AuxInt
		t := v.Aux
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 24 || t.(Type).Alignment()%8 != 0) {
			break
		}
		v.reset(OpMIPS64LoweredMove)
		v.AuxInt = t.(Type).Alignment()
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpMIPS64ADDVconst, src.Type)
		v0.AuxInt = s - moveSize(t.(Type).Alignment(), config)
		v0.AddArg(src)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMul16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mul16 x y)
	// cond:
	// result: (Select1 (MULVU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, MakeTuple(types.UInt64, types.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMul32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mul32 x y)
	// cond:
	// result: (Select1 (MULVU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, MakeTuple(types.UInt64, types.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMul32F(v *Value) bool {
	// match: (Mul32F x y)
	// cond:
	// result: (MULF x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64MULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpMul64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mul64 x y)
	// cond:
	// result: (Select1 (MULVU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, MakeTuple(types.UInt64, types.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMul64F(v *Value) bool {
	// match: (Mul64F x y)
	// cond:
	// result: (MULD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64MULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpMul8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Mul8 x y)
	// cond:
	// result: (Select1 (MULVU x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, MakeTuple(types.UInt64, types.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpNeg16(v *Value) bool {
	// match: (Neg16 x)
	// cond:
	// result: (NEGV x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeg32(v *Value) bool {
	// match: (Neg32 x)
	// cond:
	// result: (NEGV x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeg32F(v *Value) bool {
	// match: (Neg32F x)
	// cond:
	// result: (NEGF x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeg64(v *Value) bool {
	// match: (Neg64 x)
	// cond:
	// result: (NEGV x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeg64F(v *Value) bool {
	// match: (Neg64F x)
	// cond:
	// result: (NEGD x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeg8(v *Value) bool {
	// match: (Neg8 x)
	// cond:
	// result: (NEGV x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeq16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Neq16 x y)
	// cond:
	// result: (SGTU (XOR (ZeroExt16to32 x) (ZeroExt16to64 y)) (MOVVconst [0]))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, types.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, types.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpNeq32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Neq32 x y)
	// cond:
	// result: (SGTU (XOR (ZeroExt32to64 x) (ZeroExt32to64 y)) (MOVVconst [0]))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, types.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpNeq32F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Neq32F x y)
	// cond:
	// result: (FPFlagFalse (CMPEQF x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQF, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpNeq64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Neq64 x y)
	// cond:
	// result: (SGTU (XOR x y) (MOVVconst [0]))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, types.UInt64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpNeq64F(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Neq64F x y)
	// cond:
	// result: (FPFlagFalse (CMPEQD x y))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQD, TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpNeq8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Neq8 x y)
	// cond:
	// result: (SGTU (XOR (ZeroExt8to64 x) (ZeroExt8to64 y)) (MOVVconst [0]))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, types.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpNeqB(v *Value) bool {
	// match: (NeqB x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpNeqPtr(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (NeqPtr x y)
	// cond:
	// result: (SGTU (XOR x y) (MOVVconst [0]))
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, types.UInt64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpNilCheck(v *Value) bool {
	// match: (NilCheck ptr mem)
	// cond:
	// result: (LoweredNilCheck ptr mem)
	for {
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64LoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpNot(v *Value) bool {
	// match: (Not x)
	// cond:
	// result: (XORconst [1] x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64XORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpOffPtr(v *Value) bool {
	// match: (OffPtr [off] ptr:(SP))
	// cond:
	// result: (MOVVaddr [off] ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpMIPS64MOVVaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// cond:
	// result: (ADDVconst [off] ptr)
	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueMIPS64_OpOr16(v *Value) bool {
	// match: (Or16 x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpOr32(v *Value) bool {
	// match: (Or32 x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpOr64(v *Value) bool {
	// match: (Or64 x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpOr8(v *Value) bool {
	// match: (Or8 x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpOrB(v *Value) bool {
	// match: (OrB x y)
	// cond:
	// result: (OR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpRound32F(v *Value) bool {
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
func rewriteValueMIPS64_OpRound64F(v *Value) bool {
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
func rewriteValueMIPS64_OpRsh16Ux16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux16 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt16to64 y))) (SRLV <t> (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16Ux32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux32 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt32to64 y))) (SRLV <t> (ZeroExt16to64 x) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16Ux64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux64 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) y)) (SRLV <t> (ZeroExt16to64 x) y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v4.AddArg(x)
		v3.AddArg(v4)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16Ux8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16Ux8 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt8to64  y))) (SRLV <t> (ZeroExt16to64 x) (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x16 <t> x y)
	// cond:
	// result: (SRAV (SignExt16to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt16to64 y) (Const64 <types.UInt64> [63]))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x32 <t> x y)
	// cond:
	// result: (SRAV (SignExt16to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt32to64 y) (Const64 <types.UInt64> [63]))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x64 <t> x y)
	// cond:
	// result: (SRAV (SignExt16to64 x) (OR <t> (NEGV <t> (SGTU y (Const64 <types.UInt64> [63]))) y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v3.AddArg(y)
		v4 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v4.AuxInt = 63
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh16x8 <t> x y)
	// cond:
	// result: (SRAV (SignExt16to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt8to64  y) (Const64 <types.UInt64> [63]))) (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32Ux16 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt16to64 y))) (SRLV <t> (ZeroExt32to64 x) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32Ux32 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt32to64 y))) (SRLV <t> (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32Ux64 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) y)) (SRLV <t> (ZeroExt32to64 x) y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v4.AddArg(x)
		v3.AddArg(v4)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32Ux8 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt8to64  y))) (SRLV <t> (ZeroExt32to64 x) (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32x16 <t> x y)
	// cond:
	// result: (SRAV (SignExt32to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt16to64 y) (Const64 <types.UInt64> [63]))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32x32 <t> x y)
	// cond:
	// result: (SRAV (SignExt32to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt32to64 y) (Const64 <types.UInt64> [63]))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32x64 <t> x y)
	// cond:
	// result: (SRAV (SignExt32to64 x) (OR <t> (NEGV <t> (SGTU y (Const64 <types.UInt64> [63]))) y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v3.AddArg(y)
		v4 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v4.AuxInt = 63
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh32x8 <t> x y)
	// cond:
	// result: (SRAV (SignExt32to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt8to64  y) (Const64 <types.UInt64> [63]))) (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64Ux16 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt16to64 y))) (SRLV <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64Ux32 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt32to64 y))) (SRLV <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64Ux64 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) y)) (SRLV <t> x y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64Ux8 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt8to64  y))) (SRLV <t> x (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64x16 <t> x y)
	// cond:
	// result: (SRAV x (OR <t> (NEGV <t> (SGTU (ZeroExt16to64 y) (Const64 <types.UInt64> [63]))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v4.AuxInt = 63
		v2.AddArg(v4)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v5.AddArg(y)
		v0.AddArg(v5)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64x32 <t> x y)
	// cond:
	// result: (SRAV x (OR <t> (NEGV <t> (SGTU (ZeroExt32to64 y) (Const64 <types.UInt64> [63]))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v4.AuxInt = 63
		v2.AddArg(v4)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v5.AddArg(y)
		v0.AddArg(v5)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64x64 <t> x y)
	// cond:
	// result: (SRAV x (OR <t> (NEGV <t> (SGTU y (Const64 <types.UInt64> [63]))) y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v3.AuxInt = 63
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh64x8 <t> x y)
	// cond:
	// result: (SRAV x (OR <t> (NEGV <t> (SGTU (ZeroExt8to64  y) (Const64 <types.UInt64> [63]))) (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v4.AuxInt = 63
		v2.AddArg(v4)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v5.AddArg(y)
		v0.AddArg(v5)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux16 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt16to64 y))) (SRLV <t> (ZeroExt8to64 x) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux32 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt32to64 y))) (SRLV <t> (ZeroExt8to64 x) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux64 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) y)) (SRLV <t> (ZeroExt8to64 x) y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v4.AddArg(x)
		v3.AddArg(v4)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8Ux8 <t> x y)
	// cond:
	// result: (AND (NEGV <t> (SGTU (Const64 <types.UInt64> [64]) (ZeroExt8to64  y))) (SRLV <t> (ZeroExt8to64 x) (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v2 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x16(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x16 <t> x y)
	// cond:
	// result: (SRAV (SignExt8to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt16to64 y) (Const64 <types.UInt64> [63]))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, types.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x32(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x32 <t> x y)
	// cond:
	// result: (SRAV (SignExt8to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt32to64 y) (Const64 <types.UInt64> [63]))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, types.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x64(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x64 <t> x y)
	// cond:
	// result: (SRAV (SignExt8to64 x) (OR <t> (NEGV <t> (SGTU y (Const64 <types.UInt64> [63]))) y))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v3.AddArg(y)
		v4 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v4.AuxInt = 63
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x8(v *Value) bool {
	b := v.Block
	_ = b
	types := &b.Func.Config.Types
	_ = types
	// match: (Rsh8x8 <t> x y)
	// cond:
	// result: (SRAV (SignExt8to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt8to64  y) (Const64 <types.UInt64> [63]))) (ZeroExt8to64  y)))
	for {
		t := v.Type
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, types.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, types.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpConst64, types.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, types.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpSelect0(v *Value) bool {
	// match: (Select0 (DIVVU _ (MOVVconst [1])))
	// cond:
	// result: (MOVVconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select0 (DIVVU x (MOVVconst [c])))
	// cond: isPowerOfTwo(c)
	// result: (ANDconst [c-1] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64ANDconst)
		v.AuxInt = c - 1
		v.AddArg(x)
		return true
	}
	// match: (Select0 (DIVV (MOVVconst [c]) (MOVVconst [d])))
	// cond:
	// result: (MOVVconst [int64(c)%int64(d)])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVV {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(c) % int64(d)
		return true
	}
	// match: (Select0 (DIVVU (MOVVconst [c]) (MOVVconst [d])))
	// cond:
	// result: (MOVVconst [int64(uint64(c)%uint64(d))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint64(c) % uint64(d))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpSelect1(v *Value) bool {
	// match: (Select1 (MULVU x (MOVVconst [-1])))
	// cond:
	// result: (NEGV x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != -1 {
			break
		}
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULVU (MOVVconst [-1]) x))
	// cond:
	// result: (NEGV x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_0.AuxInt != -1 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULVU _ (MOVVconst [0])))
	// cond:
	// result: (MOVVconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select1 (MULVU (MOVVconst [0]) _))
	// cond:
	// result: (MOVVconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_0.AuxInt != 0 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select1 (MULVU x (MOVVconst [1])))
	// cond:
	// result: x
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
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
	// match: (Select1 (MULVU (MOVVconst [1]) x))
	// cond:
	// result: x
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
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
	// match: (Select1 (MULVU x (MOVVconst [c])))
	// cond: isPowerOfTwo(c)
	// result: (SLLVconst [log2(c)] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64SLLVconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULVU (MOVVconst [c]) x))
	// cond: isPowerOfTwo(c)
	// result: (SLLVconst [log2(c)] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64SLLVconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULVU (MOVVconst [-1]) x))
	// cond:
	// result: (NEGV x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_0.AuxInt != -1 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULVU x (MOVVconst [-1])))
	// cond:
	// result: (NEGV x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != -1 {
			break
		}
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULVU (MOVVconst [0]) _))
	// cond:
	// result: (MOVVconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_0.AuxInt != 0 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select1 (MULVU _ (MOVVconst [0])))
	// cond:
	// result: (MOVVconst [0])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select1 (MULVU (MOVVconst [1]) x))
	// cond:
	// result: x
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
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
	// match: (Select1 (MULVU x (MOVVconst [1])))
	// cond:
	// result: x
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
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
	// match: (Select1 (MULVU (MOVVconst [c]) x))
	// cond: isPowerOfTwo(c)
	// result: (SLLVconst [log2(c)] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64SLLVconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULVU x (MOVVconst [c])))
	// cond: isPowerOfTwo(c)
	// result: (SLLVconst [log2(c)] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64SLLVconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (Select1 (DIVVU x (MOVVconst [1])))
	// cond:
	// result: x
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
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
	// match: (Select1 (DIVVU x (MOVVconst [c])))
	// cond: isPowerOfTwo(c)
	// result: (SRLVconst [log2(c)] x)
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64SRLVconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULVU (MOVVconst [c]) (MOVVconst [d])))
	// cond:
	// result: (MOVVconst [c*d])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c * d
		return true
	}
	// match: (Select1 (MULVU (MOVVconst [d]) (MOVVconst [c])))
	// cond:
	// result: (MOVVconst [c*d])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c * d
		return true
	}
	// match: (Select1 (DIVV (MOVVconst [c]) (MOVVconst [d])))
	// cond:
	// result: (MOVVconst [int64(c)/int64(d)])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVV {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(c) / int64(d)
		return true
	}
	// match: (Select1 (DIVVU (MOVVconst [c]) (MOVVconst [d])))
	// cond:
	// result: (MOVVconst [int64(uint64(c)/uint64(d))])
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint64(c) / uint64(d))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpSignExt16to32(v *Value) bool {
	// match: (SignExt16to32 x)
	// cond:
	// result: (MOVHreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSignExt16to64(v *Value) bool {
	// match: (SignExt16to64 x)
	// cond:
	// result: (MOVHreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSignExt32to64(v *Value) bool {
	// match: (SignExt32to64 x)
	// cond:
	// result: (MOVWreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVWreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSignExt8to16(v *Value) bool {
	// match: (SignExt8to16 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSignExt8to32(v *Value) bool {
	// match: (SignExt8to32 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSignExt8to64(v *Value) bool {
	// match: (SignExt8to64 x)
	// cond:
	// result: (MOVBreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSlicemask(v *Value) bool {
	b := v.Block
	_ = b
	// match: (Slicemask <t> x)
	// cond:
	// result: (SRAVconst (NEGV <t> x) [63])
	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpStaticCall(v *Value) bool {
	// match: (StaticCall [argwid] {target} mem)
	// cond:
	// result: (CALLstatic [argwid] {target} mem)
	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpMIPS64CALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpStore(v *Value) bool {
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
		v.reset(OpMIPS64MOVBstore)
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
		v.reset(OpMIPS64MOVHstore)
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
		v.reset(OpMIPS64MOVWstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(Type).Size() == 8 && !is64BitFloat(val.Type)
	// result: (MOVVstore ptr val mem)
	for {
		t := v.Aux
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 8 && !is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
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
		v.reset(OpMIPS64MOVFstore)
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
		v.reset(OpMIPS64MOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpSub16(v *Value) bool {
	// match: (Sub16 x y)
	// cond:
	// result: (SUBV x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSub32(v *Value) bool {
	// match: (Sub32 x y)
	// cond:
	// result: (SUBV x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSub32F(v *Value) bool {
	// match: (Sub32F x y)
	// cond:
	// result: (SUBF x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSub64(v *Value) bool {
	// match: (Sub64 x y)
	// cond:
	// result: (SUBV x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSub64F(v *Value) bool {
	// match: (Sub64F x y)
	// cond:
	// result: (SUBD x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSub8(v *Value) bool {
	// match: (Sub8 x y)
	// cond:
	// result: (SUBV x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSubPtr(v *Value) bool {
	// match: (SubPtr x y)
	// cond:
	// result: (SUBV x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpTrunc16to8(v *Value) bool {
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
func rewriteValueMIPS64_OpTrunc32to16(v *Value) bool {
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
func rewriteValueMIPS64_OpTrunc32to8(v *Value) bool {
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
func rewriteValueMIPS64_OpTrunc64to16(v *Value) bool {
	// match: (Trunc64to16 x)
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
func rewriteValueMIPS64_OpTrunc64to32(v *Value) bool {
	// match: (Trunc64to32 x)
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
func rewriteValueMIPS64_OpTrunc64to8(v *Value) bool {
	// match: (Trunc64to8 x)
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
func rewriteValueMIPS64_OpXor16(v *Value) bool {
	// match: (Xor16 x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpXor32(v *Value) bool {
	// match: (Xor32 x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpXor64(v *Value) bool {
	// match: (Xor64 x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpXor8(v *Value) bool {
	// match: (Xor8 x y)
	// cond:
	// result: (XOR x y)
	for {
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpZero(v *Value) bool {
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
	// result: (MOVBstore ptr (MOVVconst [0]) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64MOVBstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] {t} ptr mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore ptr (MOVVconst [0]) mem)
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
		v.reset(OpMIPS64MOVHstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] ptr mem)
	// cond:
	// result: (MOVBstore [1] ptr (MOVVconst [0]) 		(MOVBstore [0] ptr (MOVVconst [0]) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 1
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore ptr (MOVVconst [0]) mem)
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
		v.reset(OpMIPS64MOVWstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore [2] ptr (MOVVconst [0]) 		(MOVHstore [0] ptr (MOVVconst [0]) mem))
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
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [4] ptr mem)
	// cond:
	// result: (MOVBstore [3] ptr (MOVVconst [0]) 		(MOVBstore [2] ptr (MOVVconst [0]) 			(MOVBstore [1] ptr (MOVVconst [0]) 				(MOVBstore [0] ptr (MOVVconst [0]) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 3
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v3.AuxInt = 1
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.(Type).Alignment()%8 == 0
	// result: (MOVVstore ptr (MOVVconst [0]) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore [4] ptr (MOVVconst [0]) 		(MOVWstore [0] ptr (MOVVconst [0]) mem))
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
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore [6] ptr (MOVVconst [0]) 		(MOVHstore [4] ptr (MOVVconst [0]) 			(MOVHstore [2] ptr (MOVVconst [0]) 				(MOVHstore [0] ptr (MOVVconst [0]) mem))))
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 6
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v1.AuxInt = 4
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v3.AuxInt = 2
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
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
	// result: (MOVBstore [2] ptr (MOVVconst [0]) 		(MOVBstore [1] ptr (MOVVconst [0]) 			(MOVBstore [0] ptr (MOVVconst [0]) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v1.AuxInt = 1
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [6] {t} ptr mem)
	// cond: t.(Type).Alignment()%2 == 0
	// result: (MOVHstore [4] ptr (MOVVconst [0]) 		(MOVHstore [2] ptr (MOVVconst [0]) 			(MOVHstore [0] ptr (MOVVconst [0]) mem)))
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
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [12] {t} ptr mem)
	// cond: t.(Type).Alignment()%4 == 0
	// result: (MOVWstore [8] ptr (MOVVconst [0]) 		(MOVWstore [4] ptr (MOVVconst [0]) 			(MOVWstore [0] ptr (MOVVconst [0]) mem)))
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
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 8
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, TypeMem)
		v1.AuxInt = 4
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [16] {t} ptr mem)
	// cond: t.(Type).Alignment()%8 == 0
	// result: (MOVVstore [8] ptr (MOVVconst [0]) 		(MOVVstore [0] ptr (MOVVconst [0]) mem))
	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 8
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [24] {t} ptr mem)
	// cond: t.(Type).Alignment()%8 == 0
	// result: (MOVVstore [16] ptr (MOVVconst [0]) 		(MOVVstore [8] ptr (MOVVconst [0]) 			(MOVVstore [0] ptr (MOVVconst [0]) mem)))
	for {
		if v.AuxInt != 24 {
			break
		}
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 16
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, TypeMem)
		v1.AuxInt = 8
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, types.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: s%8 == 0 && s > 24 && s <= 8*128 	&& t.(Type).Alignment()%8 == 0 && !config.noDuffDevice
	// result: (DUFFZERO [8 * (128 - int64(s/8))] ptr mem)
	for {
		s := v.AuxInt
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(s%8 == 0 && s > 24 && s <= 8*128 && t.(Type).Alignment()%8 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpMIPS64DUFFZERO)
		v.AuxInt = 8 * (128 - int64(s/8))
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: (s > 8*128 || config.noDuffDevice) || t.(Type).Alignment()%8 != 0
	// result: (LoweredZero [t.(Type).Alignment()] 		ptr 		(ADDVconst <ptr.Type> ptr [s-moveSize(t.(Type).Alignment(), config)]) 		mem)
	for {
		s := v.AuxInt
		t := v.Aux
		ptr := v.Args[0]
		mem := v.Args[1]
		if !((s > 8*128 || config.noDuffDevice) || t.(Type).Alignment()%8 != 0) {
			break
		}
		v.reset(OpMIPS64LoweredZero)
		v.AuxInt = t.(Type).Alignment()
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64ADDVconst, ptr.Type)
		v0.AuxInt = s - moveSize(t.(Type).Alignment(), config)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpZeroExt16to32(v *Value) bool {
	// match: (ZeroExt16to32 x)
	// cond:
	// result: (MOVHUreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVHUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpZeroExt16to64(v *Value) bool {
	// match: (ZeroExt16to64 x)
	// cond:
	// result: (MOVHUreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVHUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpZeroExt32to64(v *Value) bool {
	// match: (ZeroExt32to64 x)
	// cond:
	// result: (MOVWUreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVWUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpZeroExt8to16(v *Value) bool {
	// match: (ZeroExt8to16 x)
	// cond:
	// result: (MOVBUreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpZeroExt8to32(v *Value) bool {
	// match: (ZeroExt8to32 x)
	// cond:
	// result: (MOVBUreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpZeroExt8to64(v *Value) bool {
	// match: (ZeroExt8to64 x)
	// cond:
	// result: (MOVBUreg x)
	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteBlockMIPS64(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	types := &config.Types
	_ = types
	switch b.Kind {
	case BlockMIPS64EQ:
		// match: (EQ (FPFlagTrue cmp) yes no)
		// cond:
		// result: (FPF cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64FPFlagTrue {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64FPF
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
			if v.Op != OpMIPS64FPFlagFalse {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64FPT
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
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGT {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64NE
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
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTU {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64NE
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
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTconst {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64NE
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
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTUconst {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64NE
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
			if v.Op != OpMIPS64SGTUconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64NE
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (SGTU x (MOVVconst [0])) yes no)
		// cond:
		// result: (EQ x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64SGTU {
				break
			}
			x := v.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != OpMIPS64MOVVconst {
				break
			}
			if v_1.AuxInt != 0 {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64EQ
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
			if v.Op != OpMIPS64SGTconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64GEZ
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (SGT x (MOVVconst [0])) yes no)
		// cond:
		// result: (LEZ x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64SGT {
				break
			}
			x := v.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != OpMIPS64MOVVconst {
				break
			}
			if v_1.AuxInt != 0 {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64LEZ
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (EQ (MOVVconst [0]) yes no)
		// cond:
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
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
		// match: (EQ (MOVVconst [c]) yes no)
		// cond: c != 0
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
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
	case BlockMIPS64GEZ:
		// match: (GEZ (MOVVconst [c]) yes no)
		// cond: c >= 0
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c >= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			_ = yes
			_ = no
			return true
		}
		// match: (GEZ (MOVVconst [c]) yes no)
		// cond: c <  0
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c < 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.swapSuccessors()
			_ = no
			_ = yes
			return true
		}
	case BlockMIPS64GTZ:
		// match: (GTZ (MOVVconst [c]) yes no)
		// cond: c >  0
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c > 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			_ = yes
			_ = no
			return true
		}
		// match: (GTZ (MOVVconst [c]) yes no)
		// cond: c <= 0
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c <= 0) {
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
			b.Kind = BlockMIPS64NE
			b.SetControl(cond)
			_ = yes
			_ = no
			return true
		}
	case BlockMIPS64LEZ:
		// match: (LEZ (MOVVconst [c]) yes no)
		// cond: c <= 0
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c <= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			_ = yes
			_ = no
			return true
		}
		// match: (LEZ (MOVVconst [c]) yes no)
		// cond: c >  0
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c > 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.swapSuccessors()
			_ = no
			_ = yes
			return true
		}
	case BlockMIPS64LTZ:
		// match: (LTZ (MOVVconst [c]) yes no)
		// cond: c <  0
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c < 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			_ = yes
			_ = no
			return true
		}
		// match: (LTZ (MOVVconst [c]) yes no)
		// cond: c >= 0
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c >= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.swapSuccessors()
			_ = no
			_ = yes
			return true
		}
	case BlockMIPS64NE:
		// match: (NE (FPFlagTrue cmp) yes no)
		// cond:
		// result: (FPT cmp yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64FPFlagTrue {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64FPT
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
			if v.Op != OpMIPS64FPFlagFalse {
				break
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64FPF
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
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGT {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64EQ
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
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTU {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64EQ
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
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTconst {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64EQ
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
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTUconst {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64EQ
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
			if v.Op != OpMIPS64SGTUconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64EQ
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (SGTU x (MOVVconst [0])) yes no)
		// cond:
		// result: (NE x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64SGTU {
				break
			}
			x := v.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != OpMIPS64MOVVconst {
				break
			}
			if v_1.AuxInt != 0 {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64NE
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
			if v.Op != OpMIPS64SGTconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64LTZ
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (SGT x (MOVVconst [0])) yes no)
		// cond:
		// result: (GTZ x yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64SGT {
				break
			}
			x := v.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != OpMIPS64MOVVconst {
				break
			}
			if v_1.AuxInt != 0 {
				break
			}
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockMIPS64GTZ
			b.SetControl(x)
			_ = yes
			_ = no
			return true
		}
		// match: (NE (MOVVconst [0]) yes no)
		// cond:
		// result: (First nil no yes)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
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
		// match: (NE (MOVVconst [c]) yes no)
		// cond: c != 0
		// result: (First nil yes no)
		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
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
