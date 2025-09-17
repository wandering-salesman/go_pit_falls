package addcompare

import "testing"

type testType byte

func BenchmarkAdd1(b *testing.B) {
	for b.Loop() {
		var x, y T1[testType]
		Add1(x, y)
	}
}

func BenchmarkAdd2(b *testing.B) {
	for b.Loop() {
		var x, y T2[testType]
		Add2(x, y)
	}
}
func BenchmarkAdd3(b *testing.B) {
	for b.Loop() {
		var x, y T3[testType]
		Add3(x, y)
	}
}

func BenchmarkAdd4(b *testing.B) {
	for b.Loop() {
		var x, y T4[testType]
		Add4(x, y)
	}
}

func BenchmarkAdd5(b *testing.B) {
	for b.Loop() {
		var x, y T5[testType]
		Add5(x, y)
	}
}

func BenchmarkAdd6(b *testing.B) {
	for b.Loop() {
		var x, y T6[testType]
		Add6(x, y)
	}
}

func BenchmarkAdd7(b *testing.B) {
	for b.Loop() {
		var x, y T7[testType]
		Add7(x, y)
	}
}

func BenchmarkAdd8(b *testing.B) {
	for b.Loop() {
		var x, y T8[testType]
		Add8(x, y)
	}
}
