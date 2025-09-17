package main

import "testing"


func return1(input float64) (float64) {
	return input + 0
}

func BenchmarkReturn1(b *testing.B) {
	for b.Loop(){
		return1(0)
	}
}

func return2(input float64) (float64, float64) {
	return input + 0, input + 1
}

func BenchmarkReturn2(b *testing.B) {
	for b.Loop(){
		return2(0)
	}
}

func return3(input float64) (float64, float64, float64) {
	return input + 0, input + 1, input + 2
}

func BenchmarkReturn3(b *testing.B) {
	for b.Loop(){
		return3(0)
	}
}

func return4(input float64) (float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3
}

func BenchmarkReturn4(b *testing.B) {
	for b.Loop(){
		return4(0)
	}
}

func return5(input float64) (float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4
}

func BenchmarkReturn5(b *testing.B) {
	for b.Loop(){
		return5(0)
	}
}

func return6(input float64) (float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5
}

func BenchmarkReturn6(b *testing.B) {
	for b.Loop(){
		return6(0)
	}
}

func return7(input float64) (float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6
}

func BenchmarkReturn7(b *testing.B) {
	for b.Loop(){
		return7(0)
	}
}

func return8(input float64) (float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7
}

func BenchmarkReturn8(b *testing.B) {
	for b.Loop(){
		return8(0)
	}
}

func return9(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8
}

func BenchmarkReturn9(b *testing.B) {
	for b.Loop(){
		return9(0)
	}
}

func return10(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9
}

func BenchmarkReturn10(b *testing.B) {
	for b.Loop(){
		return10(0)
	}
}

func return11(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10
}

func BenchmarkReturn11(b *testing.B) {
	for b.Loop(){
		return11(0)
	}
}

func return12(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11
}

func BenchmarkReturn12(b *testing.B) {
	for b.Loop(){
		return12(0)
	}
}

func return13(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12
}

func BenchmarkReturn13(b *testing.B) {
	for b.Loop(){
		return13(0)
	}
}

func return14(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13
}

func BenchmarkReturn14(b *testing.B) {
	for b.Loop(){
		return14(0)
	}
}

func return15(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14
}

func BenchmarkReturn15(b *testing.B) {
	for b.Loop(){
		return15(0)
	}
}

func return16(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15
}

func BenchmarkReturn16(b *testing.B) {
	for b.Loop(){
		return16(0)
	}
}

func return17(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16
}

func BenchmarkReturn17(b *testing.B) {
	for b.Loop(){
		return17(0)
	}
}

func return18(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17
}

func BenchmarkReturn18(b *testing.B) {
	for b.Loop(){
		return18(0)
	}
}

func return19(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18
}

func BenchmarkReturn19(b *testing.B) {
	for b.Loop(){
		return19(0)
	}
}

func return20(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19
}

func BenchmarkReturn20(b *testing.B) {
	for b.Loop(){
		return20(0)
	}
}

func return21(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20
}

func BenchmarkReturn21(b *testing.B) {
	for b.Loop(){
		return21(0)
	}
}

func return22(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21
}

func BenchmarkReturn22(b *testing.B) {
	for b.Loop(){
		return22(0)
	}
}

func return23(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22
}

func BenchmarkReturn23(b *testing.B) {
	for b.Loop(){
		return23(0)
	}
}

func return24(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23
}

func BenchmarkReturn24(b *testing.B) {
	for b.Loop(){
		return24(0)
	}
}

func return25(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24
}

func BenchmarkReturn25(b *testing.B) {
	for b.Loop(){
		return25(0)
	}
}

func return26(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25
}

func BenchmarkReturn26(b *testing.B) {
	for b.Loop(){
		return26(0)
	}
}

func return27(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26
}

func BenchmarkReturn27(b *testing.B) {
	for b.Loop(){
		return27(0)
	}
}

func return28(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27
}

func BenchmarkReturn28(b *testing.B) {
	for b.Loop(){
		return28(0)
	}
}

func return29(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28
}

func BenchmarkReturn29(b *testing.B) {
	for b.Loop(){
		return29(0)
	}
}

func return30(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29
}

func BenchmarkReturn30(b *testing.B) {
	for b.Loop(){
		return30(0)
	}
}

func return31(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30
}

func BenchmarkReturn31(b *testing.B) {
	for b.Loop(){
		return31(0)
	}
}

func return32(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31
}

func BenchmarkReturn32(b *testing.B) {
	for b.Loop(){
		return32(0)
	}
}

func return33(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32
}

func BenchmarkReturn33(b *testing.B) {
	for b.Loop(){
		return33(0)
	}
}

func return34(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33
}

func BenchmarkReturn34(b *testing.B) {
	for b.Loop(){
		return34(0)
	}
}

func return35(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34
}

func BenchmarkReturn35(b *testing.B) {
	for b.Loop(){
		return35(0)
	}
}

func return36(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35
}

func BenchmarkReturn36(b *testing.B) {
	for b.Loop(){
		return36(0)
	}
}

func return37(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36
}

func BenchmarkReturn37(b *testing.B) {
	for b.Loop(){
		return37(0)
	}
}

func return38(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37
}

func BenchmarkReturn38(b *testing.B) {
	for b.Loop(){
		return38(0)
	}
}

func return39(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38
}

func BenchmarkReturn39(b *testing.B) {
	for b.Loop(){
		return39(0)
	}
}

func return40(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39
}

func BenchmarkReturn40(b *testing.B) {
	for b.Loop(){
		return40(0)
	}
}

func return41(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40
}

func BenchmarkReturn41(b *testing.B) {
	for b.Loop(){
		return41(0)
	}
}

func return42(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41
}

func BenchmarkReturn42(b *testing.B) {
	for b.Loop(){
		return42(0)
	}
}

func return43(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42
}

func BenchmarkReturn43(b *testing.B) {
	for b.Loop(){
		return43(0)
	}
}

func return44(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43
}

func BenchmarkReturn44(b *testing.B) {
	for b.Loop(){
		return44(0)
	}
}

func return45(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44
}

func BenchmarkReturn45(b *testing.B) {
	for b.Loop(){
		return45(0)
	}
}

func return46(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45
}

func BenchmarkReturn46(b *testing.B) {
	for b.Loop(){
		return46(0)
	}
}

func return47(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46
}

func BenchmarkReturn47(b *testing.B) {
	for b.Loop(){
		return47(0)
	}
}

func return48(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47
}

func BenchmarkReturn48(b *testing.B) {
	for b.Loop(){
		return48(0)
	}
}

func return49(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48
}

func BenchmarkReturn49(b *testing.B) {
	for b.Loop(){
		return49(0)
	}
}

func return50(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49
}

func BenchmarkReturn50(b *testing.B) {
	for b.Loop(){
		return50(0)
	}
}

func return51(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50
}

func BenchmarkReturn51(b *testing.B) {
	for b.Loop(){
		return51(0)
	}
}

func return52(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51
}

func BenchmarkReturn52(b *testing.B) {
	for b.Loop(){
		return52(0)
	}
}

func return53(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52
}

func BenchmarkReturn53(b *testing.B) {
	for b.Loop(){
		return53(0)
	}
}

func return54(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52, input + 53
}

func BenchmarkReturn54(b *testing.B) {
	for b.Loop(){
		return54(0)
	}
}

func return55(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52, input + 53, input + 54
}

func BenchmarkReturn55(b *testing.B) {
	for b.Loop(){
		return55(0)
	}
}

func return56(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52, input + 53, input + 54, input + 55
}

func BenchmarkReturn56(b *testing.B) {
	for b.Loop(){
		return56(0)
	}
}

func return57(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52, input + 53, input + 54, input + 55, input + 56
}

func BenchmarkReturn57(b *testing.B) {
	for b.Loop(){
		return57(0)
	}
}

func return58(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52, input + 53, input + 54, input + 55, input + 56, input + 57
}

func BenchmarkReturn58(b *testing.B) {
	for b.Loop(){
		return58(0)
	}
}

func return59(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52, input + 53, input + 54, input + 55, input + 56, input + 57, input + 58
}

func BenchmarkReturn59(b *testing.B) {
	for b.Loop(){
		return59(0)
	}
}

func return60(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52, input + 53, input + 54, input + 55, input + 56, input + 57, input + 58, input + 59
}

func BenchmarkReturn60(b *testing.B) {
	for b.Loop(){
		return60(0)
	}
}

func return61(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52, input + 53, input + 54, input + 55, input + 56, input + 57, input + 58, input + 59, input + 60
}

func BenchmarkReturn61(b *testing.B) {
	for b.Loop(){
		return61(0)
	}
}

func return62(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52, input + 53, input + 54, input + 55, input + 56, input + 57, input + 58, input + 59, input + 60, input + 61
}

func BenchmarkReturn62(b *testing.B) {
	for b.Loop(){
		return62(0)
	}
}

func return63(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52, input + 53, input + 54, input + 55, input + 56, input + 57, input + 58, input + 59, input + 60, input + 61, input + 62
}

func BenchmarkReturn63(b *testing.B) {
	for b.Loop(){
		return63(0)
	}
}

func return64(input float64) (float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64, float64) {
	return input + 0, input + 1, input + 2, input + 3, input + 4, input + 5, input + 6, input + 7, input + 8, input + 9, input + 10, input + 11, input + 12, input + 13, input + 14, input + 15, input + 16, input + 17, input + 18, input + 19, input + 20, input + 21, input + 22, input + 23, input + 24, input + 25, input + 26, input + 27, input + 28, input + 29, input + 30, input + 31, input + 32, input + 33, input + 34, input + 35, input + 36, input + 37, input + 38, input + 39, input + 40, input + 41, input + 42, input + 43, input + 44, input + 45, input + 46, input + 47, input + 48, input + 49, input + 50, input + 51, input + 52, input + 53, input + 54, input + 55, input + 56, input + 57, input + 58, input + 59, input + 60, input + 61, input + 62, input + 63
}

func BenchmarkReturn64(b *testing.B) {
	for b.Loop(){
		return64(0)
	}
}
