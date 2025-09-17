package addcompare

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

type T1[T Number] struct {
	a T
}

type T2[T Number] struct {
	a, b float64
}

type T3[T Number] struct {
	a, b, c float64
}

type T4[T Number] struct {
	a, b, c, d float64
}

type T5[T Number] struct {
	a, b, c, d, e float64
}

type T6[T Number] struct {
	a, b, c, d, e, f float64
}

type T7[T Number] struct {
	a, b, c, d, e, f, g float64
}

type T8[T Number] struct {
	a, b, c, d, e, f, g, h float64
}

//go:noinline
func Add(x, y float64) float64 {
	return x + y
}

//go:noinline
func Add1[T Number](x, y T1[T]) (z T1[T]) {
	z.a = x.a + y.a
	return
}

//go:noinline
func Add2[T Number](x, y T2[T]) (z T2[T]) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	return
}

//go:noinline
func Add3[T Number](x, y T3[T]) (z T3[T]) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	return
}

//go:noinline
func Add4[T Number](x, y T4[T]) (z T4[T]) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	return
}

//go:noinline
func Add5[T Number](x, y T5[T]) (z T5[T]) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	z.e = x.e + y.e
	return
}

//go:noinline
func Add6[T Number](x, y T6[T]) (z T6[T]) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	z.e = x.e + y.e
	z.f = x.f + y.f
	return
}

//go:noinline
func Add7[T Number](x, y T7[T]) (z T7[T]) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	z.e = x.e + y.e
	z.f = x.f + y.f
	z.g = x.g + y.g
	return
}

//go:noinline
func Add8[T Number](x, y T8[T]) (z T8[T]) {
	z.a = x.a + y.a
	z.b = x.b + y.b
	z.c = x.c + y.c
	z.d = x.d + y.d
	z.e = x.e + y.e
	z.f = x.f + y.f
	z.g = x.g + y.g
	z.h = x.h + y.h
	return
}
