// Code generated by command: go run sub_asm.go -out ../../matfuncs/sub_amd64.s -stubs ../../matfuncs/sub_amd64_stubs.go -pkg matfuncs. DO NOT EDIT.

//go:build amd64 && gc && !purego

#include "textflag.h"

// func SubAVX32(x1 []float32, x2 []float32, y []float32)
// Requires: AVX, SSE
TEXT ·SubAVX32(SB), NOSPLIT, $0-72
	MOVQ x1_base+0(FP), AX
	MOVQ x2_base+24(FP), CX
	MOVQ y_base+48(FP), DX
	MOVQ x1_len+8(FP), BX

unrolledLoop:
	CMPQ    BX, $0x00000080
	JL      singleRegisterLoop
	VMOVUPS (AX), Y0
	VMOVUPS 32(AX), Y1
	VMOVUPS 64(AX), Y2
	VMOVUPS 96(AX), Y3
	VMOVUPS 128(AX), Y4
	VMOVUPS 160(AX), Y5
	VMOVUPS 192(AX), Y6
	VMOVUPS 224(AX), Y7
	VMOVUPS 256(AX), Y8
	VMOVUPS 288(AX), Y9
	VMOVUPS 320(AX), Y10
	VMOVUPS 352(AX), Y11
	VMOVUPS 384(AX), Y12
	VMOVUPS 416(AX), Y13
	VMOVUPS 448(AX), Y14
	VMOVUPS 480(AX), Y15
	VSUBPS  (CX), Y0, Y0
	VSUBPS  32(CX), Y1, Y1
	VSUBPS  64(CX), Y2, Y2
	VSUBPS  96(CX), Y3, Y3
	VSUBPS  128(CX), Y4, Y4
	VSUBPS  160(CX), Y5, Y5
	VSUBPS  192(CX), Y6, Y6
	VSUBPS  224(CX), Y7, Y7
	VSUBPS  256(CX), Y8, Y8
	VSUBPS  288(CX), Y9, Y9
	VSUBPS  320(CX), Y10, Y10
	VSUBPS  352(CX), Y11, Y11
	VSUBPS  384(CX), Y12, Y12
	VSUBPS  416(CX), Y13, Y13
	VSUBPS  448(CX), Y14, Y14
	VSUBPS  480(CX), Y15, Y15
	VMOVUPS Y0, (DX)
	VMOVUPS Y1, 32(DX)
	VMOVUPS Y2, 64(DX)
	VMOVUPS Y3, 96(DX)
	VMOVUPS Y4, 128(DX)
	VMOVUPS Y5, 160(DX)
	VMOVUPS Y6, 192(DX)
	VMOVUPS Y7, 224(DX)
	VMOVUPS Y8, 256(DX)
	VMOVUPS Y9, 288(DX)
	VMOVUPS Y10, 320(DX)
	VMOVUPS Y11, 352(DX)
	VMOVUPS Y12, 384(DX)
	VMOVUPS Y13, 416(DX)
	VMOVUPS Y14, 448(DX)
	VMOVUPS Y15, 480(DX)
	ADDQ    $0x00000200, AX
	ADDQ    $0x00000200, CX
	ADDQ    $0x00000200, DX
	SUBQ    $0x00000080, BX
	JMP     unrolledLoop

singleRegisterLoop:
	CMPQ    BX, $0x00000008
	JL      tailLoop
	VMOVUPS (AX), Y0
	VSUBPS  (CX), Y0, Y0
	VMOVUPS Y0, (DX)
	ADDQ    $0x00000020, AX
	ADDQ    $0x00000020, CX
	ADDQ    $0x00000020, DX
	SUBQ    $0x00000008, BX
	JMP     singleRegisterLoop

tailLoop:
	CMPQ  BX, $0x00000000
	JE    end
	MOVSS (AX), X0
	SUBSS (CX), X0
	MOVSS X0, (DX)
	ADDQ  $0x00000004, AX
	ADDQ  $0x00000004, CX
	ADDQ  $0x00000004, DX
	DECQ  BX
	JMP   tailLoop

end:
	RET

// func SubAVX64(x1 []float64, x2 []float64, y []float64)
// Requires: AVX, SSE2
TEXT ·SubAVX64(SB), NOSPLIT, $0-72
	MOVQ x1_base+0(FP), AX
	MOVQ x2_base+24(FP), CX
	MOVQ y_base+48(FP), DX
	MOVQ x1_len+8(FP), BX

unrolledLoop:
	CMPQ    BX, $0x00000040
	JL      singleRegisterLoop
	VMOVUPD (AX), Y0
	VMOVUPD 32(AX), Y1
	VMOVUPD 64(AX), Y2
	VMOVUPD 96(AX), Y3
	VMOVUPD 128(AX), Y4
	VMOVUPD 160(AX), Y5
	VMOVUPD 192(AX), Y6
	VMOVUPD 224(AX), Y7
	VMOVUPD 256(AX), Y8
	VMOVUPD 288(AX), Y9
	VMOVUPD 320(AX), Y10
	VMOVUPD 352(AX), Y11
	VMOVUPD 384(AX), Y12
	VMOVUPD 416(AX), Y13
	VMOVUPD 448(AX), Y14
	VMOVUPD 480(AX), Y15
	VSUBPD  (CX), Y0, Y0
	VSUBPD  32(CX), Y1, Y1
	VSUBPD  64(CX), Y2, Y2
	VSUBPD  96(CX), Y3, Y3
	VSUBPD  128(CX), Y4, Y4
	VSUBPD  160(CX), Y5, Y5
	VSUBPD  192(CX), Y6, Y6
	VSUBPD  224(CX), Y7, Y7
	VSUBPD  256(CX), Y8, Y8
	VSUBPD  288(CX), Y9, Y9
	VSUBPD  320(CX), Y10, Y10
	VSUBPD  352(CX), Y11, Y11
	VSUBPD  384(CX), Y12, Y12
	VSUBPD  416(CX), Y13, Y13
	VSUBPD  448(CX), Y14, Y14
	VSUBPD  480(CX), Y15, Y15
	VMOVUPD Y0, (DX)
	VMOVUPD Y1, 32(DX)
	VMOVUPD Y2, 64(DX)
	VMOVUPD Y3, 96(DX)
	VMOVUPD Y4, 128(DX)
	VMOVUPD Y5, 160(DX)
	VMOVUPD Y6, 192(DX)
	VMOVUPD Y7, 224(DX)
	VMOVUPD Y8, 256(DX)
	VMOVUPD Y9, 288(DX)
	VMOVUPD Y10, 320(DX)
	VMOVUPD Y11, 352(DX)
	VMOVUPD Y12, 384(DX)
	VMOVUPD Y13, 416(DX)
	VMOVUPD Y14, 448(DX)
	VMOVUPD Y15, 480(DX)
	ADDQ    $0x00000200, AX
	ADDQ    $0x00000200, CX
	ADDQ    $0x00000200, DX
	SUBQ    $0x00000040, BX
	JMP     unrolledLoop

singleRegisterLoop:
	CMPQ    BX, $0x00000004
	JL      tailLoop
	VMOVUPD (AX), Y0
	VSUBPD  (CX), Y0, Y0
	VMOVUPD Y0, (DX)
	ADDQ    $0x00000020, AX
	ADDQ    $0x00000020, CX
	ADDQ    $0x00000020, DX
	SUBQ    $0x00000004, BX
	JMP     singleRegisterLoop

tailLoop:
	CMPQ  BX, $0x00000000
	JE    end
	MOVSD (AX), X0
	SUBSD (CX), X0
	MOVSD X0, (DX)
	ADDQ  $0x00000008, AX
	ADDQ  $0x00000008, CX
	ADDQ  $0x00000008, DX
	DECQ  BX
	JMP   tailLoop

end:
	RET

// func SubSSE32(x1 []float32, x2 []float32, y []float32)
// Requires: SSE
TEXT ·SubSSE32(SB), NOSPLIT, $0-72
	MOVQ x1_base+0(FP), AX
	MOVQ x2_base+24(FP), CX
	MOVQ y_base+48(FP), DX
	MOVQ x1_len+8(FP), BX
	CMPQ BX, $0x00000000
	JE   end
	MOVQ CX, SI
	ANDQ $0x0000000f, SI
	JZ   unrolledLoop
	XORQ $0x0000000f, SI
	INCQ SI
	SHRQ $0x02, SI

alignmentLoop:
	MOVSS (AX), X0
	SUBSS (CX), X0
	MOVSS X0, (DX)
	ADDQ  $0x00000004, AX
	ADDQ  $0x00000004, CX
	ADDQ  $0x00000004, DX
	DECQ  BX
	JZ    end
	DECQ  SI
	JNZ   alignmentLoop

unrolledLoop:
	CMPQ   BX, $0x00000040
	JL     singleRegisterLoop
	MOVUPS (AX), X0
	MOVUPS 16(AX), X1
	MOVUPS 32(AX), X2
	MOVUPS 48(AX), X3
	MOVUPS 64(AX), X4
	MOVUPS 80(AX), X5
	MOVUPS 96(AX), X6
	MOVUPS 112(AX), X7
	MOVUPS 128(AX), X8
	MOVUPS 144(AX), X9
	MOVUPS 160(AX), X10
	MOVUPS 176(AX), X11
	MOVUPS 192(AX), X12
	MOVUPS 208(AX), X13
	MOVUPS 224(AX), X14
	MOVUPS 240(AX), X15
	SUBPS  (CX), X0
	SUBPS  16(CX), X1
	SUBPS  32(CX), X2
	SUBPS  48(CX), X3
	SUBPS  64(CX), X4
	SUBPS  80(CX), X5
	SUBPS  96(CX), X6
	SUBPS  112(CX), X7
	SUBPS  128(CX), X8
	SUBPS  144(CX), X9
	SUBPS  160(CX), X10
	SUBPS  176(CX), X11
	SUBPS  192(CX), X12
	SUBPS  208(CX), X13
	SUBPS  224(CX), X14
	SUBPS  240(CX), X15
	MOVUPS X0, (DX)
	MOVUPS X1, 16(DX)
	MOVUPS X2, 32(DX)
	MOVUPS X3, 48(DX)
	MOVUPS X4, 64(DX)
	MOVUPS X5, 80(DX)
	MOVUPS X6, 96(DX)
	MOVUPS X7, 112(DX)
	MOVUPS X8, 128(DX)
	MOVUPS X9, 144(DX)
	MOVUPS X10, 160(DX)
	MOVUPS X11, 176(DX)
	MOVUPS X12, 192(DX)
	MOVUPS X13, 208(DX)
	MOVUPS X14, 224(DX)
	MOVUPS X15, 240(DX)
	ADDQ   $0x00000100, AX
	ADDQ   $0x00000100, CX
	ADDQ   $0x00000100, DX
	SUBQ   $0x00000040, BX
	JMP    unrolledLoop

singleRegisterLoop:
	CMPQ   BX, $0x00000004
	JL     tailLoop
	MOVUPS (AX), X0
	SUBPS  (CX), X0
	MOVUPS X0, (DX)
	ADDQ   $0x00000010, AX
	ADDQ   $0x00000010, CX
	ADDQ   $0x00000010, DX
	SUBQ   $0x00000004, BX
	JMP    singleRegisterLoop

tailLoop:
	CMPQ  BX, $0x00000000
	JE    end
	MOVSS (AX), X0
	SUBSS (CX), X0
	MOVSS X0, (DX)
	ADDQ  $0x00000004, AX
	ADDQ  $0x00000004, CX
	ADDQ  $0x00000004, DX
	DECQ  BX
	JMP   tailLoop

end:
	RET

// func SubSSE64(x1 []float64, x2 []float64, y []float64)
// Requires: SSE2
TEXT ·SubSSE64(SB), NOSPLIT, $0-72
	MOVQ  x1_base+0(FP), AX
	MOVQ  x2_base+24(FP), CX
	MOVQ  y_base+48(FP), DX
	MOVQ  x1_len+8(FP), BX
	CMPQ  BX, $0x00000000
	JE    end
	MOVQ  CX, SI
	ANDQ  $0x0000000f, SI
	JZ    unrolledLoop
	MOVSD (AX), X0
	SUBSD (CX), X0
	MOVSD X0, (DX)
	ADDQ  $0x00000008, AX
	ADDQ  $0x00000008, CX
	ADDQ  $0x00000008, DX
	DECQ  BX

unrolledLoop:
	CMPQ   BX, $0x00000020
	JL     singleRegisterLoop
	MOVUPD (AX), X0
	MOVUPD 16(AX), X1
	MOVUPD 32(AX), X2
	MOVUPD 48(AX), X3
	MOVUPD 64(AX), X4
	MOVUPD 80(AX), X5
	MOVUPD 96(AX), X6
	MOVUPD 112(AX), X7
	MOVUPD 128(AX), X8
	MOVUPD 144(AX), X9
	MOVUPD 160(AX), X10
	MOVUPD 176(AX), X11
	MOVUPD 192(AX), X12
	MOVUPD 208(AX), X13
	MOVUPD 224(AX), X14
	MOVUPD 240(AX), X15
	SUBPD  (CX), X0
	SUBPD  16(CX), X1
	SUBPD  32(CX), X2
	SUBPD  48(CX), X3
	SUBPD  64(CX), X4
	SUBPD  80(CX), X5
	SUBPD  96(CX), X6
	SUBPD  112(CX), X7
	SUBPD  128(CX), X8
	SUBPD  144(CX), X9
	SUBPD  160(CX), X10
	SUBPD  176(CX), X11
	SUBPD  192(CX), X12
	SUBPD  208(CX), X13
	SUBPD  224(CX), X14
	SUBPD  240(CX), X15
	MOVUPD X0, (DX)
	MOVUPD X1, 16(DX)
	MOVUPD X2, 32(DX)
	MOVUPD X3, 48(DX)
	MOVUPD X4, 64(DX)
	MOVUPD X5, 80(DX)
	MOVUPD X6, 96(DX)
	MOVUPD X7, 112(DX)
	MOVUPD X8, 128(DX)
	MOVUPD X9, 144(DX)
	MOVUPD X10, 160(DX)
	MOVUPD X11, 176(DX)
	MOVUPD X12, 192(DX)
	MOVUPD X13, 208(DX)
	MOVUPD X14, 224(DX)
	MOVUPD X15, 240(DX)
	ADDQ   $0x00000100, AX
	ADDQ   $0x00000100, CX
	ADDQ   $0x00000100, DX
	SUBQ   $0x00000020, BX
	JMP    unrolledLoop

singleRegisterLoop:
	CMPQ   BX, $0x00000002
	JL     tailLoop
	MOVUPD (AX), X0
	SUBPD  (CX), X0
	MOVUPD X0, (DX)
	ADDQ   $0x00000010, AX
	ADDQ   $0x00000010, CX
	ADDQ   $0x00000010, DX
	SUBQ   $0x00000002, BX
	JMP    singleRegisterLoop

tailLoop:
	CMPQ  BX, $0x00000000
	JE    end
	MOVSD (AX), X0
	SUBSD (CX), X0
	MOVSD X0, (DX)
	ADDQ  $0x00000008, AX
	ADDQ  $0x00000008, CX
	ADDQ  $0x00000008, DX
	DECQ  BX
	JMP   tailLoop

end:
	RET
