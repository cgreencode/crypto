// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64,!appengine,!gccgo

// This code was translated into a form compatible with 6a from the public
// domain sources in SUPERCOP: http://bench.cr.yp.to/supercop.html

// func salsa2020XORKeyStream(out, in *byte, n uint64, nonce, key *byte)
TEXT ·salsa2020XORKeyStream(SB),0,$512-40
	MOVQ out+0(FP),DI
	MOVQ in+8(FP),SI
	MOVQ n+16(FP),DX
	MOVQ nonce+24(FP),CX
	MOVQ key+32(FP),R8

	MOVQ SP,R11
	MOVQ $31,R9
	NOTQ R9
	ANDQ R9,SP
	ADDQ $32,SP

	MOVQ R11,352(SP)
	MOVQ R12,360(SP)
	MOVQ R13,368(SP)
	MOVQ R14,376(SP)
	MOVQ R15,384(SP)
	MOVQ BX,392(SP)
	MOVQ BP,400(SP)
	MOVQ DX,R9
	MOVQ CX,DX
	MOVQ R8,R10
	CMPQ R9,$0
	JBE DONE
	START:
	MOVL 20(R10),CX
	MOVL 0(R10),R8
	MOVL 0(DX),AX
	MOVL 16(R10),R11
	MOVL CX,0(SP)
	MOVL R8, 4 (SP)
	MOVL AX, 8 (SP)
	MOVL R11, 12 (SP)
	MOVL 8(DX),CX
	MOVL 24(R10),R8
	MOVL 4(R10),AX
	MOVL 4(DX),R11
	MOVL CX,16(SP)
	MOVL R8, 20 (SP)
	MOVL AX, 24 (SP)
	MOVL R11, 28 (SP)
	MOVL 12(DX),CX
	MOVL 12(R10),DX
	MOVL 28(R10),R8
	MOVL 8(R10),AX
	MOVL DX,32(SP)
	MOVL CX, 36 (SP)
	MOVL R8, 40 (SP)
	MOVL AX, 44 (SP)
	MOVQ $1634760805,DX
	MOVQ $857760878,CX
	MOVQ $2036477234,R8
	MOVQ $1797285236,AX
	MOVL DX,48(SP)
	MOVL CX, 52 (SP)
	MOVL R8, 56 (SP)
	MOVL AX, 60 (SP)
	CMPQ R9,$256
	JB BYTESBETWEEN1AND255
	MOVOA 48(SP),X0
	PSHUFL $0X55,X0,X1
	PSHUFL $0XAA,X0,X2
	PSHUFL $0XFF,X0,X3
	PSHUFL $0X00,X0,X0
	MOVOA X1,64(SP)
	MOVOA X2,80(SP)
	MOVOA X3,96(SP)
	MOVOA X0,112(SP)
	MOVOA 0(SP),X0
	PSHUFL $0XAA,X0,X1
	PSHUFL $0XFF,X0,X2
	PSHUFL $0X00,X0,X3
	PSHUFL $0X55,X0,X0
	MOVOA X1,128(SP)
	MOVOA X2,144(SP)
	MOVOA X3,160(SP)
	MOVOA X0,176(SP)
	MOVOA 16(SP),X0
	PSHUFL $0XFF,X0,X1
	PSHUFL $0X55,X0,X2
	PSHUFL $0XAA,X0,X0
	MOVOA X1,192(SP)
	MOVOA X2,208(SP)
	MOVOA X0,224(SP)
	MOVOA 32(SP),X0
	PSHUFL $0X00,X0,X1
	PSHUFL $0XAA,X0,X2
	PSHUFL $0XFF,X0,X0
	MOVOA X1,240(SP)
	MOVOA X2,256(SP)
	MOVOA X0,272(SP)
	BYTESATLEAST256:
	MOVL 16(SP),DX
	MOVL  36 (SP),CX
	MOVL DX,288(SP)
	MOVL CX,304(SP)
	ADDQ $1,DX
	SHLQ $32,CX
	ADDQ CX,DX
	MOVQ DX,CX
	SHRQ $32,CX
	MOVL DX, 292 (SP)
	MOVL CX, 308 (SP)
	ADDQ $1,DX
	SHLQ $32,CX
	ADDQ CX,DX
	MOVQ DX,CX
	SHRQ $32,CX
	MOVL DX, 296 (SP)
	MOVL CX, 312 (SP)
	ADDQ $1,DX
	SHLQ $32,CX
	ADDQ CX,DX
	MOVQ DX,CX
	SHRQ $32,CX
	MOVL DX, 300 (SP)
	MOVL CX, 316 (SP)
	ADDQ $1,DX
	SHLQ $32,CX
	ADDQ CX,DX
	MOVQ DX,CX
	SHRQ $32,CX
	MOVL DX,16(SP)
	MOVL CX, 36 (SP)
	MOVQ R9,408(SP)
	MOVQ $20,DX
	MOVOA 64(SP),X0
	MOVOA 80(SP),X1
	MOVOA 96(SP),X2
	MOVOA 256(SP),X3
	MOVOA 272(SP),X4
	MOVOA 128(SP),X5
	MOVOA 144(SP),X6
	MOVOA 176(SP),X7
	MOVOA 192(SP),X8
	MOVOA 208(SP),X9
	MOVOA 224(SP),X10
	MOVOA 304(SP),X11
	MOVOA 112(SP),X12
	MOVOA 160(SP),X13
	MOVOA 240(SP),X14
	MOVOA 288(SP),X15
	MAINLOOP1:
	MOVOA X1,320(SP)
	MOVOA X2,336(SP)
	MOVOA X13,X1
	PADDL X12,X1
	MOVOA X1,X2
	PSLLL $7,X1
	PXOR X1,X14
	PSRLL $25,X2
	PXOR X2,X14
	MOVOA X7,X1
	PADDL X0,X1
	MOVOA X1,X2
	PSLLL $7,X1
	PXOR X1,X11
	PSRLL $25,X2
	PXOR X2,X11
	MOVOA X12,X1
	PADDL X14,X1
	MOVOA X1,X2
	PSLLL $9,X1
	PXOR X1,X15
	PSRLL $23,X2
	PXOR X2,X15
	MOVOA X0,X1
	PADDL X11,X1
	MOVOA X1,X2
	PSLLL $9,X1
	PXOR X1,X9
	PSRLL $23,X2
	PXOR X2,X9
	MOVOA X14,X1
	PADDL X15,X1
	MOVOA X1,X2
	PSLLL $13,X1
	PXOR X1,X13
	PSRLL $19,X2
	PXOR X2,X13
	MOVOA X11,X1
	PADDL X9,X1
	MOVOA X1,X2
	PSLLL $13,X1
	PXOR X1,X7
	PSRLL $19,X2
	PXOR X2,X7
	MOVOA X15,X1
	PADDL X13,X1
	MOVOA X1,X2
	PSLLL $18,X1
	PXOR X1,X12
	PSRLL $14,X2
	PXOR X2,X12
	MOVOA 320(SP),X1
	MOVOA X12,320(SP)
	MOVOA X9,X2
	PADDL X7,X2
	MOVOA X2,X12
	PSLLL $18,X2
	PXOR X2,X0
	PSRLL $14,X12
	PXOR X12,X0
	MOVOA X5,X2
	PADDL X1,X2
	MOVOA X2,X12
	PSLLL $7,X2
	PXOR X2,X3
	PSRLL $25,X12
	PXOR X12,X3
	MOVOA 336(SP),X2
	MOVOA X0,336(SP)
	MOVOA X6,X0
	PADDL X2,X0
	MOVOA X0,X12
	PSLLL $7,X0
	PXOR X0,X4
	PSRLL $25,X12
	PXOR X12,X4
	MOVOA X1,X0
	PADDL X3,X0
	MOVOA X0,X12
	PSLLL $9,X0
	PXOR X0,X10
	PSRLL $23,X12
	PXOR X12,X10
	MOVOA X2,X0
	PADDL X4,X0
	MOVOA X0,X12
	PSLLL $9,X0
	PXOR X0,X8
	PSRLL $23,X12
	PXOR X12,X8
	MOVOA X3,X0
	PADDL X10,X0
	MOVOA X0,X12
	PSLLL $13,X0
	PXOR X0,X5
	PSRLL $19,X12
	PXOR X12,X5
	MOVOA X4,X0
	PADDL X8,X0
	MOVOA X0,X12
	PSLLL $13,X0
	PXOR X0,X6
	PSRLL $19,X12
	PXOR X12,X6
	MOVOA X10,X0
	PADDL X5,X0
	MOVOA X0,X12
	PSLLL $18,X0
	PXOR X0,X1
	PSRLL $14,X12
	PXOR X12,X1
	MOVOA 320(SP),X0
	MOVOA X1,320(SP)
	MOVOA X4,X1
	PADDL X0,X1
	MOVOA X1,X12
	PSLLL $7,X1
	PXOR X1,X7
	PSRLL $25,X12
	PXOR X12,X7
	MOVOA X8,X1
	PADDL X6,X1
	MOVOA X1,X12
	PSLLL $18,X1
	PXOR X1,X2
	PSRLL $14,X12
	PXOR X12,X2
	MOVOA 336(SP),X12
	MOVOA X2,336(SP)
	MOVOA X14,X1
	PADDL X12,X1
	MOVOA X1,X2
	PSLLL $7,X1
	PXOR X1,X5
	PSRLL $25,X2
	PXOR X2,X5
	MOVOA X0,X1
	PADDL X7,X1
	MOVOA X1,X2
	PSLLL $9,X1
	PXOR X1,X10
	PSRLL $23,X2
	PXOR X2,X10
	MOVOA X12,X1
	PADDL X5,X1
	MOVOA X1,X2
	PSLLL $9,X1
	PXOR X1,X8
	PSRLL $23,X2
	PXOR X2,X8
	MOVOA X7,X1
	PADDL X10,X1
	MOVOA X1,X2
	PSLLL $13,X1
	PXOR X1,X4
	PSRLL $19,X2
	PXOR X2,X4
	MOVOA X5,X1
	PADDL X8,X1
	MOVOA X1,X2
	PSLLL $13,X1
	PXOR X1,X14
	PSRLL $19,X2
	PXOR X2,X14
	MOVOA X10,X1
	PADDL X4,X1
	MOVOA X1,X2
	PSLLL $18,X1
	PXOR X1,X0
	PSRLL $14,X2
	PXOR X2,X0
	MOVOA 320(SP),X1
	MOVOA X0,320(SP)
	MOVOA X8,X0
	PADDL X14,X0
	MOVOA X0,X2
	PSLLL $18,X0
	PXOR X0,X12
	PSRLL $14,X2
	PXOR X2,X12
	MOVOA X11,X0
	PADDL X1,X0
	MOVOA X0,X2
	PSLLL $7,X0
	PXOR X0,X6
	PSRLL $25,X2
	PXOR X2,X6
	MOVOA 336(SP),X2
	MOVOA X12,336(SP)
	MOVOA X3,X0
	PADDL X2,X0
	MOVOA X0,X12
	PSLLL $7,X0
	PXOR X0,X13
	PSRLL $25,X12
	PXOR X12,X13
	MOVOA X1,X0
	PADDL X6,X0
	MOVOA X0,X12
	PSLLL $9,X0
	PXOR X0,X15
	PSRLL $23,X12
	PXOR X12,X15
	MOVOA X2,X0
	PADDL X13,X0
	MOVOA X0,X12
	PSLLL $9,X0
	PXOR X0,X9
	PSRLL $23,X12
	PXOR X12,X9
	MOVOA X6,X0
	PADDL X15,X0
	MOVOA X0,X12
	PSLLL $13,X0
	PXOR X0,X11
	PSRLL $19,X12
	PXOR X12,X11
	MOVOA X13,X0
	PADDL X9,X0
	MOVOA X0,X12
	PSLLL $13,X0
	PXOR X0,X3
	PSRLL $19,X12
	PXOR X12,X3
	MOVOA X15,X0
	PADDL X11,X0
	MOVOA X0,X12
	PSLLL $18,X0
	PXOR X0,X1
	PSRLL $14,X12
	PXOR X12,X1
	MOVOA X9,X0
	PADDL X3,X0
	MOVOA X0,X12
	PSLLL $18,X0
	PXOR X0,X2
	PSRLL $14,X12
	PXOR X12,X2
	MOVOA 320(SP),X12
	MOVOA 336(SP),X0
	SUBQ $2,DX
	JA MAINLOOP1
	PADDL 112(SP),X12
	PADDL 176(SP),X7
	PADDL 224(SP),X10
	PADDL 272(SP),X4
	MOVD X12,DX
	MOVD X7,CX
	MOVD X10,R8
	MOVD X4,R9
	PSHUFL $0X39,X12,X12
	PSHUFL $0X39,X7,X7
	PSHUFL $0X39,X10,X10
	PSHUFL $0X39,X4,X4
	XORL 0(SI),DX
	XORL 4(SI),CX
	XORL 8(SI),R8
	XORL 12(SI),R9
	MOVL DX,0(DI)
	MOVL CX,4(DI)
	MOVL R8,8(DI)
	MOVL R9,12(DI)
	MOVD X12,DX
	MOVD X7,CX
	MOVD X10,R8
	MOVD X4,R9
	PSHUFL $0X39,X12,X12
	PSHUFL $0X39,X7,X7
	PSHUFL $0X39,X10,X10
	PSHUFL $0X39,X4,X4
	XORL 64(SI),DX
	XORL 68(SI),CX
	XORL 72(SI),R8
	XORL 76(SI),R9
	MOVL DX,64(DI)
	MOVL CX,68(DI)
	MOVL R8,72(DI)
	MOVL R9,76(DI)
	MOVD X12,DX
	MOVD X7,CX
	MOVD X10,R8
	MOVD X4,R9
	PSHUFL $0X39,X12,X12
	PSHUFL $0X39,X7,X7
	PSHUFL $0X39,X10,X10
	PSHUFL $0X39,X4,X4
	XORL 128(SI),DX
	XORL 132(SI),CX
	XORL 136(SI),R8
	XORL 140(SI),R9
	MOVL DX,128(DI)
	MOVL CX,132(DI)
	MOVL R8,136(DI)
	MOVL R9,140(DI)
	MOVD X12,DX
	MOVD X7,CX
	MOVD X10,R8
	MOVD X4,R9
	XORL 192(SI),DX
	XORL 196(SI),CX
	XORL 200(SI),R8
	XORL 204(SI),R9
	MOVL DX,192(DI)
	MOVL CX,196(DI)
	MOVL R8,200(DI)
	MOVL R9,204(DI)
	PADDL 240(SP),X14
	PADDL 64(SP),X0
	PADDL 128(SP),X5
	PADDL 192(SP),X8
	MOVD X14,DX
	MOVD X0,CX
	MOVD X5,R8
	MOVD X8,R9
	PSHUFL $0X39,X14,X14
	PSHUFL $0X39,X0,X0
	PSHUFL $0X39,X5,X5
	PSHUFL $0X39,X8,X8
	XORL 16(SI),DX
	XORL 20(SI),CX
	XORL 24(SI),R8
	XORL 28(SI),R9
	MOVL DX,16(DI)
	MOVL CX,20(DI)
	MOVL R8,24(DI)
	MOVL R9,28(DI)
	MOVD X14,DX
	MOVD X0,CX
	MOVD X5,R8
	MOVD X8,R9
	PSHUFL $0X39,X14,X14
	PSHUFL $0X39,X0,X0
	PSHUFL $0X39,X5,X5
	PSHUFL $0X39,X8,X8
	XORL 80(SI),DX
	XORL 84(SI),CX
	XORL 88(SI),R8
	XORL 92(SI),R9
	MOVL DX,80(DI)
	MOVL CX,84(DI)
	MOVL R8,88(DI)
	MOVL R9,92(DI)
	MOVD X14,DX
	MOVD X0,CX
	MOVD X5,R8
	MOVD X8,R9
	PSHUFL $0X39,X14,X14
	PSHUFL $0X39,X0,X0
	PSHUFL $0X39,X5,X5
	PSHUFL $0X39,X8,X8
	XORL 144(SI),DX
	XORL 148(SI),CX
	XORL 152(SI),R8
	XORL 156(SI),R9
	MOVL DX,144(DI)
	MOVL CX,148(DI)
	MOVL R8,152(DI)
	MOVL R9,156(DI)
	MOVD X14,DX
	MOVD X0,CX
	MOVD X5,R8
	MOVD X8,R9
	XORL 208(SI),DX
	XORL 212(SI),CX
	XORL 216(SI),R8
	XORL 220(SI),R9
	MOVL DX,208(DI)
	MOVL CX,212(DI)
	MOVL R8,216(DI)
	MOVL R9,220(DI)
	PADDL 288(SP),X15
	PADDL 304(SP),X11
	PADDL 80(SP),X1
	PADDL 144(SP),X6
	MOVD X15,DX
	MOVD X11,CX
	MOVD X1,R8
	MOVD X6,R9
	PSHUFL $0X39,X15,X15
	PSHUFL $0X39,X11,X11
	PSHUFL $0X39,X1,X1
	PSHUFL $0X39,X6,X6
	XORL 32(SI),DX
	XORL 36(SI),CX
	XORL 40(SI),R8
	XORL 44(SI),R9
	MOVL DX,32(DI)
	MOVL CX,36(DI)
	MOVL R8,40(DI)
	MOVL R9,44(DI)
	MOVD X15,DX
	MOVD X11,CX
	MOVD X1,R8
	MOVD X6,R9
	PSHUFL $0X39,X15,X15
	PSHUFL $0X39,X11,X11
	PSHUFL $0X39,X1,X1
	PSHUFL $0X39,X6,X6
	XORL 96(SI),DX
	XORL 100(SI),CX
	XORL 104(SI),R8
	XORL 108(SI),R9
	MOVL DX,96(DI)
	MOVL CX,100(DI)
	MOVL R8,104(DI)
	MOVL R9,108(DI)
	MOVD X15,DX
	MOVD X11,CX
	MOVD X1,R8
	MOVD X6,R9
	PSHUFL $0X39,X15,X15
	PSHUFL $0X39,X11,X11
	PSHUFL $0X39,X1,X1
	PSHUFL $0X39,X6,X6
	XORL 160(SI),DX
	XORL 164(SI),CX
	XORL 168(SI),R8
	XORL 172(SI),R9
	MOVL DX,160(DI)
	MOVL CX,164(DI)
	MOVL R8,168(DI)
	MOVL R9,172(DI)
	MOVD X15,DX
	MOVD X11,CX
	MOVD X1,R8
	MOVD X6,R9
	XORL 224(SI),DX
	XORL 228(SI),CX
	XORL 232(SI),R8
	XORL 236(SI),R9
	MOVL DX,224(DI)
	MOVL CX,228(DI)
	MOVL R8,232(DI)
	MOVL R9,236(DI)
	PADDL 160(SP),X13
	PADDL 208(SP),X9
	PADDL 256(SP),X3
	PADDL 96(SP),X2
	MOVD X13,DX
	MOVD X9,CX
	MOVD X3,R8
	MOVD X2,R9
	PSHUFL $0X39,X13,X13
	PSHUFL $0X39,X9,X9
	PSHUFL $0X39,X3,X3
	PSHUFL $0X39,X2,X2
	XORL 48(SI),DX
	XORL 52(SI),CX
	XORL 56(SI),R8
	XORL 60(SI),R9
	MOVL DX,48(DI)
	MOVL CX,52(DI)
	MOVL R8,56(DI)
	MOVL R9,60(DI)
	MOVD X13,DX
	MOVD X9,CX
	MOVD X3,R8
	MOVD X2,R9
	PSHUFL $0X39,X13,X13
	PSHUFL $0X39,X9,X9
	PSHUFL $0X39,X3,X3
	PSHUFL $0X39,X2,X2
	XORL 112(SI),DX
	XORL 116(SI),CX
	XORL 120(SI),R8
	XORL 124(SI),R9
	MOVL DX,112(DI)
	MOVL CX,116(DI)
	MOVL R8,120(DI)
	MOVL R9,124(DI)
	MOVD X13,DX
	MOVD X9,CX
	MOVD X3,R8
	MOVD X2,R9
	PSHUFL $0X39,X13,X13
	PSHUFL $0X39,X9,X9
	PSHUFL $0X39,X3,X3
	PSHUFL $0X39,X2,X2
	XORL 176(SI),DX
	XORL 180(SI),CX
	XORL 184(SI),R8
	XORL 188(SI),R9
	MOVL DX,176(DI)
	MOVL CX,180(DI)
	MOVL R8,184(DI)
	MOVL R9,188(DI)
	MOVD X13,DX
	MOVD X9,CX
	MOVD X3,R8
	MOVD X2,R9
	XORL 240(SI),DX
	XORL 244(SI),CX
	XORL 248(SI),R8
	XORL 252(SI),R9
	MOVL DX,240(DI)
	MOVL CX,244(DI)
	MOVL R8,248(DI)
	MOVL R9,252(DI)
	MOVQ 408(SP),R9
	SUBQ $256,R9
	ADDQ $256,SI
	ADDQ $256,DI
	CMPQ R9,$256
	JAE BYTESATLEAST256
	CMPQ R9,$0
	JBE DONE
	BYTESBETWEEN1AND255:
	CMPQ R9,$64
	JAE NOCOPY
	MOVQ DI,DX
	LEAQ 416(SP),DI
	MOVQ R9,CX
	REP; MOVSB
	LEAQ 416(SP),DI
	LEAQ 416(SP),SI
	NOCOPY:
	MOVQ R9,408(SP)
	MOVOA 48(SP),X0
	MOVOA 0(SP),X1
	MOVOA 16(SP),X2
	MOVOA 32(SP),X3
	MOVOA X1,X4
	MOVQ $20,CX
	MAINLOOP2:
	PADDL X0,X4
	MOVOA X0,X5
	MOVOA X4,X6
	PSLLL $7,X4
	PSRLL $25,X6
	PXOR X4,X3
	PXOR X6,X3
	PADDL X3,X5
	MOVOA X3,X4
	MOVOA X5,X6
	PSLLL $9,X5
	PSRLL $23,X6
	PXOR X5,X2
	PSHUFL $0X93,X3,X3
	PXOR X6,X2
	PADDL X2,X4
	MOVOA X2,X5
	MOVOA X4,X6
	PSLLL $13,X4
	PSRLL $19,X6
	PXOR X4,X1
	PSHUFL $0X4E,X2,X2
	PXOR X6,X1
	PADDL X1,X5
	MOVOA X3,X4
	MOVOA X5,X6
	PSLLL $18,X5
	PSRLL $14,X6
	PXOR X5,X0
	PSHUFL $0X39,X1,X1
	PXOR X6,X0
	PADDL X0,X4
	MOVOA X0,X5
	MOVOA X4,X6
	PSLLL $7,X4
	PSRLL $25,X6
	PXOR X4,X1
	PXOR X6,X1
	PADDL X1,X5
	MOVOA X1,X4
	MOVOA X5,X6
	PSLLL $9,X5
	PSRLL $23,X6
	PXOR X5,X2
	PSHUFL $0X93,X1,X1
	PXOR X6,X2
	PADDL X2,X4
	MOVOA X2,X5
	MOVOA X4,X6
	PSLLL $13,X4
	PSRLL $19,X6
	PXOR X4,X3
	PSHUFL $0X4E,X2,X2
	PXOR X6,X3
	PADDL X3,X5
	MOVOA X1,X4
	MOVOA X5,X6
	PSLLL $18,X5
	PSRLL $14,X6
	PXOR X5,X0
	PSHUFL $0X39,X3,X3
	PXOR X6,X0
	PADDL X0,X4
	MOVOA X0,X5
	MOVOA X4,X6
	PSLLL $7,X4
	PSRLL $25,X6
	PXOR X4,X3
	PXOR X6,X3
	PADDL X3,X5
	MOVOA X3,X4
	MOVOA X5,X6
	PSLLL $9,X5
	PSRLL $23,X6
	PXOR X5,X2
	PSHUFL $0X93,X3,X3
	PXOR X6,X2
	PADDL X2,X4
	MOVOA X2,X5
	MOVOA X4,X6
	PSLLL $13,X4
	PSRLL $19,X6
	PXOR X4,X1
	PSHUFL $0X4E,X2,X2
	PXOR X6,X1
	PADDL X1,X5
	MOVOA X3,X4
	MOVOA X5,X6
	PSLLL $18,X5
	PSRLL $14,X6
	PXOR X5,X0
	PSHUFL $0X39,X1,X1
	PXOR X6,X0
	PADDL X0,X4
	MOVOA X0,X5
	MOVOA X4,X6
	PSLLL $7,X4
	PSRLL $25,X6
	PXOR X4,X1
	PXOR X6,X1
	PADDL X1,X5
	MOVOA X1,X4
	MOVOA X5,X6
	PSLLL $9,X5
	PSRLL $23,X6
	PXOR X5,X2
	PSHUFL $0X93,X1,X1
	PXOR X6,X2
	PADDL X2,X4
	MOVOA X2,X5
	MOVOA X4,X6
	PSLLL $13,X4
	PSRLL $19,X6
	PXOR X4,X3
	PSHUFL $0X4E,X2,X2
	PXOR X6,X3
	SUBQ $4,CX
	PADDL X3,X5
	MOVOA X1,X4
	MOVOA X5,X6
	PSLLL $18,X5
	PXOR X7,X7
	PSRLL $14,X6
	PXOR X5,X0
	PSHUFL $0X39,X3,X3
	PXOR X6,X0
	JA MAINLOOP2
	PADDL 48(SP),X0
	PADDL 0(SP),X1
	PADDL 16(SP),X2
	PADDL 32(SP),X3
	MOVD X0,CX
	MOVD X1,R8
	MOVD X2,R9
	MOVD X3,AX
	PSHUFL $0X39,X0,X0
	PSHUFL $0X39,X1,X1
	PSHUFL $0X39,X2,X2
	PSHUFL $0X39,X3,X3
	XORL 0(SI),CX
	XORL 48(SI),R8
	XORL 32(SI),R9
	XORL 16(SI),AX
	MOVL CX,0(DI)
	MOVL R8,48(DI)
	MOVL R9,32(DI)
	MOVL AX,16(DI)
	MOVD X0,CX
	MOVD X1,R8
	MOVD X2,R9
	MOVD X3,AX
	PSHUFL $0X39,X0,X0
	PSHUFL $0X39,X1,X1
	PSHUFL $0X39,X2,X2
	PSHUFL $0X39,X3,X3
	XORL 20(SI),CX
	XORL 4(SI),R8
	XORL 52(SI),R9
	XORL 36(SI),AX
	MOVL CX,20(DI)
	MOVL R8,4(DI)
	MOVL R9,52(DI)
	MOVL AX,36(DI)
	MOVD X0,CX
	MOVD X1,R8
	MOVD X2,R9
	MOVD X3,AX
	PSHUFL $0X39,X0,X0
	PSHUFL $0X39,X1,X1
	PSHUFL $0X39,X2,X2
	PSHUFL $0X39,X3,X3
	XORL 40(SI),CX
	XORL 24(SI),R8
	XORL 8(SI),R9
	XORL 56(SI),AX
	MOVL CX,40(DI)
	MOVL R8,24(DI)
	MOVL R9,8(DI)
	MOVL AX,56(DI)
	MOVD X0,CX
	MOVD X1,R8
	MOVD X2,R9
	MOVD X3,AX
	XORL 60(SI),CX
	XORL 44(SI),R8
	XORL 28(SI),R9
	XORL 12(SI),AX
	MOVL CX,60(DI)
	MOVL R8,44(DI)
	MOVL R9,28(DI)
	MOVL AX,12(DI)
	MOVQ 408(SP),R9
	MOVL 16(SP),CX
	MOVL  36 (SP),R8
	ADDQ $1,CX
	SHLQ $32,R8
	ADDQ R8,CX
	MOVQ CX,R8
	SHRQ $32,R8
	MOVL CX,16(SP)
	MOVL R8, 36 (SP)
	CMPQ R9,$64
	JA BYTESATLEAST65
	JAE BYTESATLEAST64
	MOVQ DI,SI
	MOVQ DX,DI
	MOVQ R9,CX
	REP; MOVSB
	BYTESATLEAST64:
	DONE:
	MOVQ 352(SP),R11
	MOVQ 360(SP),R12
	MOVQ 368(SP),R13
	MOVQ 376(SP),R14
	MOVQ 384(SP),R15
	MOVQ 392(SP),BX
	MOVQ 400(SP),BP
	MOVQ R11,SP
	RET
	BYTESATLEAST65:
	SUBQ $64,R9
	ADDQ $64,DI
	ADDQ $64,SI
	JMP BYTESBETWEEN1AND255
