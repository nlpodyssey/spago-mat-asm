// Copyright 2022 The NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//go:generate go run . -out ../../matfuncs/log_amd64.s -stubs ../../matfuncs/log_amd64_stubs.go -pkg matfuncs

import (
	"fmt"
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
)

func main() {
	ConstraintExpr("amd64,gc,!purego")

	buildSSE(32)

	Generate()
}

func buildSSE(bits int) {
	globlData4 := func(name string, v U32) Mem {
		m := GLOBL(name, RODATA|NOPTR)
		DATA(0, v)
		DATA(4, v)
		DATA(8, v)
		DATA(12, v)
		return m
	}

	LCPI0_0 := globlData4("SSE_LCPI0_0", 0x00800000)   // float 1.17549435E-38
	LCPI0_1 := globlData4("SSE_LCPI0_1", 2155872255)   // 0x807fffff
	LCPI0_2 := globlData4("SSE_LCPI0_2", 1056964608)   // 0x3f000000
	LCPI0_3 := globlData4("SSE_LCPI0_3", 4294967169)   // 0xffffff81
	LCPI0_4 := globlData4("SSE_LCPI0_4", 0x3f800000)   // float 1
	LCPI0_5 := globlData4("SSE_LCPI0_5", 0x3f3504f3)   // float 0.707106769
	LCPI0_6 := globlData4("SSE_LCPI0_6", 0xbf800000)   // float -1
	LCPI0_7 := globlData4("SSE_LCPI0_7", 0x3d9021bb)   // float 0.0703768358
	LCPI0_8 := globlData4("SSE_LCPI0_8", 0xbdebd1b8)   // float -0.115146101
	LCPI0_9 := globlData4("SSE_LCPI0_9", 0x3def251a)   // float 0.116769984
	LCPI0_10 := globlData4("SSE_LCPI0_10", 0xbdfe5d4f) // float -0.12420141
	LCPI0_11 := globlData4("SSE_LCPI0_11", 0x3e11e9bf) // float 0.142493233
	LCPI0_12 := globlData4("SSE_LCPI0_12", 0xbe2aae50) // float -0.166680574
	LCPI0_13 := globlData4("SSE_LCPI0_13", 0x3e4cceac) // float 0.200007141
	LCPI0_14 := globlData4("SSE_LCPI0_14", 0xbe7ffffc) // float -0.24999994
	LCPI0_15 := globlData4("SSE_LCPI0_15", 0x3eaaaaaa) // float 0.333333313
	LCPI0_16 := globlData4("SSE_LCPI0_16", 0xb95e8083) // float -2.12194442E-4
	LCPI0_17 := globlData4("SSE_LCPI0_17", 0xbf000000) // float -0.5
	LCPI0_18 := globlData4("SSE_LCPI0_18", 0x3f318000) // float 0.693359375

	name := fmt.Sprintf("LogSSE%d", bits)
	signature := fmt.Sprintf("func(x, y []float%d)", bits)
	TEXT(name, NOSPLIT, signature)
	Pragma("noescape")
	Doc(fmt.Sprintf(
		"%s computes the natural logarithm of each element of x, storing the result in y (%d bits, SSE required).",
		name, bits,
	))

	x := Mem{Base: Load(Param("x").Base(), GP64())}
	y := Mem{Base: Load(Param("y").Base(), GP64())}

	MOVUPS(x, X0)

	// ---

	//        xorps   %xmm2, %xmm2
	XORPS(X2, X2)

	//        movaps  %xmm0, %xmm1
	MOVAPS(X0, X1)

	//        cmpleps %xmm2, %xmm1
	CMPPS(X2, X1, U8(2))

	//        maxps   .LCPI0_0(%rip), %xmm0
	MAXPS(LCPI0_0, X0)

	//        movaps  %xmm0, %xmm2
	MOVAPS(X0, X2)

	//        psrld   $23, %xmm2
	PSRLL(U8(23), X2)

	//        andps   .LCPI0_1(%rip), %xmm0
	ANDPS(LCPI0_1, X0)

	//        orps    .LCPI0_2(%rip), %xmm0
	ORPS(LCPI0_2, X0)

	//        paddd   .LCPI0_3(%rip), %xmm2
	PADDD(LCPI0_3, X2)

	//        movaps  %xmm0, %xmm4
	MOVAPS(X0, X4)

	//        cmpltps .LCPI0_5(%rip), %xmm4
	CMPPS(LCPI0_5, X4, U8(1))

	//        movaps  %xmm4, %xmm3
	MOVAPS(X4, X3)

	//        andps   %xmm0, %xmm3
	ANDPS(X0, X3)

	//        addps   .LCPI0_6(%rip), %xmm0
	ADDPS(LCPI0_6, X0)

	//        addps   %xmm3, %xmm0
	ADDPS(X3, X0)

	//        movaps  .LCPI0_7(%rip), %xmm3           # xmm3 = [7.03768358E-2,7.03768358E-2,7.03768358E-2,7.03768358E-2]
	MOVAPS(LCPI0_7, X3)

	//        mulps   %xmm0, %xmm3
	MULPS(X0, X3)

	//        addps   .LCPI0_8(%rip), %xmm3
	ADDPS(LCPI0_8, X3)

	//        cvtdq2ps        %xmm2, %xmm2
	CVTPL2PS(X2, X2)

	//        mulps   %xmm0, %xmm3
	MULPS(X0, X3)

	//        addps   .LCPI0_9(%rip), %xmm3
	ADDPS(LCPI0_9, X3)

	//        movaps  .LCPI0_4(%rip), %xmm5           # xmm5 = [1.0E+0,1.0E+0,1.0E+0,1.0E+0]
	MOVAPS(LCPI0_4, X5)

	//        mulps   %xmm0, %xmm3
	MULPS(X0, X3)

	//        addps   .LCPI0_10(%rip), %xmm3
	ADDPS(LCPI0_10, X3)

	//        addps   %xmm5, %xmm2
	ADDPS(X5, X2)

	//        mulps   %xmm0, %xmm3
	MULPS(X0, X3)

	//        addps   .LCPI0_11(%rip), %xmm3
	ADDPS(LCPI0_11, X3)

	//        andps   %xmm5, %xmm4
	ANDPS(X5, X4)

	//        mulps   %xmm0, %xmm3
	MULPS(X0, X3)

	//        addps   .LCPI0_12(%rip), %xmm3
	ADDPS(LCPI0_12, X3)

	//        subps   %xmm4, %xmm2
	SUBPS(X4, X2)

	//        movaps  %xmm0, %xmm4
	MOVAPS(X0, X4)

	//        mulps   %xmm0, %xmm3
	MULPS(X0, X3)

	//        addps   .LCPI0_13(%rip), %xmm3
	ADDPS(LCPI0_13, X3)

	//        mulps   %xmm0, %xmm4
	MULPS(X0, X4)

	//        mulps   %xmm0, %xmm3
	MULPS(X0, X3)

	//        addps   .LCPI0_14(%rip), %xmm3
	ADDPS(LCPI0_14, X3)

	//        mulps   %xmm0, %xmm3
	MULPS(X0, X3)

	//        addps   .LCPI0_15(%rip), %xmm3
	ADDPS(LCPI0_15, X3)

	//        mulps   %xmm0, %xmm3
	MULPS(X0, X3)

	//        mulps   %xmm4, %xmm3
	MULPS(X4, X3)

	//        movaps  .LCPI0_16(%rip), %xmm5          # xmm5 = [-2.12194442E-4,-2.12194442E-4,-2.12194442E-4,-2.12194442E-4]
	MOVAPS(LCPI0_16, X5)

	//        mulps   %xmm2, %xmm5
	MULPS(X2, X5)

	//        addps   %xmm3, %xmm5
	ADDPS(X3, X5)

	//        mulps   .LCPI0_17(%rip), %xmm4
	MULPS(LCPI0_17, X4)

	//        mulps   .LCPI0_18(%rip), %xmm2
	MULPS(LCPI0_18, X2)

	//        addps   %xmm5, %xmm4
	ADDPS(X5, X4)

	//        addps   %xmm4, %xmm0
	ADDPS(X4, X0)

	//        addps   %xmm2, %xmm0
	ADDPS(X2, X0)

	//        orps    %xmm1, %xmm0
	ORPS(X1, X0)

	// ---

	MOVUPS(X0, y)

	RET()
}
