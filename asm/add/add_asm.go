// Copyright 2022 The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//go:generate go run . -out ../../matfuncs/add_amd64.s -stubs ../../matfuncs/add_amd64_stubs.go -pkg matfuncs

import (
	"fmt"

	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
)

func main() {
	ConstraintExpr("amd64,gc,!purego")

	buildAVX(32)
	buildAVX(64)

	buildSSE(32)
	buildSSE(64)

	Generate()
}

const unroll = 16 // number of XMM or YMM registers

type bitsToFuncOps map[int]func(...Op)
type bitsToFunc2Ops map[int]func(Op, Op)

var (
	ADDP   = bitsToFunc2Ops{32: ADDPS, 64: ADDPD}
	ADDS   = bitsToFunc2Ops{32: ADDSS, 64: ADDSD}
	MOVS   = bitsToFunc2Ops{32: MOVSS, 64: MOVSD}
	MOVUP  = bitsToFunc2Ops{32: MOVUPS, 64: MOVUPD}
	VADDP  = bitsToFuncOps{32: VADDPS, 64: VADDPD}
	VADDS  = bitsToFuncOps{32: VADDSS, 64: VADDSD}
	VMOVS  = bitsToFuncOps{32: VMOVSS, 64: VMOVSD}
	VMOVUP = bitsToFuncOps{32: VMOVUPS, 64: VMOVUPD}
)

func buildAVX(bits int) {
	name := fmt.Sprintf("AddAVX%d", bits)
	signature := fmt.Sprintf("func(x1, x2, y []float%d)", bits)
	TEXT(name, NOSPLIT, signature)
	Pragma("noescape")
	Doc(fmt.Sprintf("%s adds x1 and x2 element-wise, storing the result in y (%d bits, AVX required).", name, bits))

	x1 := Mem{Base: Load(Param("x1").Base(), GP64())}
	x2 := Mem{Base: Load(Param("x2").Base(), GP64())}
	y := Mem{Base: Load(Param("y").Base(), GP64())}
	n := Load(Param("x1").Len(), GP64())

	regs := make([]VecVirtual, unroll)
	for i := 0; i < unroll; i++ {
		regs[i] = YMM()
	}

	bytesPerRegister := 32 // size of one YMM register
	bytesPerValue := bits / 8
	itemsPerRegister := 8 * bytesPerRegister / bits // 4 64-bit values, or 8 32-bit values

	Label("unrolledLoop")

	blockItems := itemsPerRegister * unroll
	blockBytesSize := bytesPerValue * blockItems

	CMPQ(n, U32(blockItems))
	JL(LabelRef("singleRegisterLoop"))

	for i, reg := range regs {
		VMOVUP[bits](x1.Offset(bytesPerRegister*i), reg)
	}

	for i, reg := range regs {
		VADDP[bits](x2.Offset(bytesPerRegister*i), reg, reg)
	}

	for i, reg := range regs {
		VMOVUP[bits](reg, y.Offset(bytesPerRegister*i))
	}

	ADDQ(U32(blockBytesSize), x1.Base)
	ADDQ(U32(blockBytesSize), x2.Base)
	ADDQ(U32(blockBytesSize), y.Base)
	SUBQ(U32(blockItems), n)

	JMP(LabelRef("unrolledLoop"))

	// ---

	Label("singleRegisterLoop")

	blockItems = itemsPerRegister
	blockBytesSize = (bits / 8) * blockItems

	reg := regs[0]

	CMPQ(n, U32(blockItems))
	JL(LabelRef("tailLoop"))

	VMOVUP[bits](x1, reg)
	VADDP[bits](x2, reg, reg)
	VMOVUP[bits](reg, y)

	ADDQ(U32(blockBytesSize), x1.Base)
	ADDQ(U32(blockBytesSize), x2.Base)
	ADDQ(U32(blockBytesSize), y.Base)
	SUBQ(U32(blockItems), n)

	JMP(LabelRef("singleRegisterLoop"))

	// ---

	Label("tailLoop")

	reg = XMM()

	CMPQ(n, U32(0))
	JE(LabelRef("end"))

	VMOVS[bits](x1, reg)
	VADDS[bits](x2, reg, reg)
	VMOVS[bits](reg, y)

	ADDQ(U32(bits/8), x1.Base)
	ADDQ(U32(bits/8), x2.Base)
	ADDQ(U32(bits/8), y.Base)
	DECQ(n)

	JMP(LabelRef("tailLoop"))

	Label("end")
	RET()
}

func buildSSE(bits int) {
	name := fmt.Sprintf("AddSSE%d", bits)
	signature := fmt.Sprintf("func(x1, x2, y []float%d)", bits)
	TEXT(name, NOSPLIT, signature)
	Pragma("noescape")
	Doc(fmt.Sprintf("%s adds x1 and x2 element-wise, storing the result in y (%d bits, SSE required).", name, bits))

	x1 := Mem{Base: Load(Param("x1").Base(), GP64())}
	x2 := Mem{Base: Load(Param("x2").Base(), GP64())}
	y := Mem{Base: Load(Param("y").Base(), GP64())}
	n := Load(Param("x1").Len(), GP64())

	// x2 memory alignment

	CMPQ(n, U32(0))
	JE(LabelRef("end"))

	x2StartByte := GP64()
	MOVQ(x2.Base, x2StartByte)
	ANDQ(U32(15), x2StartByte)
	JZ(LabelRef("unrolledLoop"))

	switch bits {
	case 32:
		shifts := x2StartByte
		// 4 - floor(x2StartByte % 16 / 4)
		XORQ(U32(15), shifts)
		INCQ(shifts)
		SHRQ(U8(2), shifts)

		Label("alignmentLoop")

		reg := XMM()

		MOVS[bits](x1, reg)
		ADDS[bits](x2, reg)
		MOVS[bits](reg, y)

		ADDQ(U32(bits/8), x1.Base)
		ADDQ(U32(bits/8), x2.Base)
		ADDQ(U32(bits/8), y.Base)
		DECQ(n)
		JZ(LabelRef("end"))

		DECQ(shifts)
		JNZ(LabelRef("alignmentLoop"))
	case 64:
		reg := XMM()

		MOVS[bits](x1, reg)
		ADDS[bits](x2, reg)
		MOVS[bits](reg, y)

		ADDQ(U32(bits/8), x1.Base)
		ADDQ(U32(bits/8), x2.Base)
		ADDQ(U32(bits/8), y.Base)
		DECQ(n)
	default:
		panic(fmt.Errorf("unexpected bits %d", bits))
	}

	regs := make([]VecVirtual, unroll)
	for i := 0; i < unroll; i++ {
		regs[i] = XMM()
	}

	bytesPerRegister := 16 // size of one XMM register
	bytesPerValue := bits / 8
	itemsPerRegister := 8 * bytesPerRegister / bits // 2 64-bit values, or 4 32-bit values

	Label("unrolledLoop")

	blockItems := itemsPerRegister * unroll
	blockBytesSize := bytesPerValue * blockItems

	CMPQ(n, U32(blockItems))
	JL(LabelRef("singleRegisterLoop"))

	for i, reg := range regs {
		MOVUP[bits](x1.Offset(bytesPerRegister*i), reg)
	}

	for i, reg := range regs {
		ADDP[bits](x2.Offset(bytesPerRegister*i), reg)
	}

	for i, reg := range regs {
		MOVUP[bits](reg, y.Offset(bytesPerRegister*i))
	}

	ADDQ(U32(blockBytesSize), x1.Base)
	ADDQ(U32(blockBytesSize), x2.Base)
	ADDQ(U32(blockBytesSize), y.Base)
	SUBQ(U32(blockItems), n)

	JMP(LabelRef("unrolledLoop"))

	// ---

	Label("singleRegisterLoop")

	blockItems = itemsPerRegister
	blockBytesSize = (bits / 8) * blockItems

	reg := regs[0]

	CMPQ(n, U32(blockItems))
	JL(LabelRef("tailLoop"))

	MOVUP[bits](x1, reg)
	ADDP[bits](x2, reg)
	MOVUP[bits](reg, y)

	ADDQ(U32(blockBytesSize), x1.Base)
	ADDQ(U32(blockBytesSize), x2.Base)
	ADDQ(U32(blockBytesSize), y.Base)
	SUBQ(U32(blockItems), n)

	JMP(LabelRef("singleRegisterLoop"))

	// ---

	Label("tailLoop")

	reg = XMM()

	CMPQ(n, U32(0))
	JE(LabelRef("end"))

	MOVS[bits](x1, reg)
	ADDS[bits](x2, reg)
	MOVS[bits](reg, y)

	ADDQ(U32(bits/8), x1.Base)
	ADDQ(U32(bits/8), x2.Base)
	ADDQ(U32(bits/8), y.Base)
	DECQ(n)

	JMP(LabelRef("tailLoop"))

	Label("end")
	RET()
}
