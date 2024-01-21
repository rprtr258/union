package union

import "unsafe"

// Union3 is a value which can be interpreted as either A or B, or C.
// Effectively this is pointer to C union
type Union3[A, B, C any] struct {
	data []byte // of constant len=max(sizeof(A), sizeof(B), sizeof(C))
}

// NewUnion3 creates new Union2 filled with zeros
func NewUnion3[A, B, C any]() Union3[A, B, C] {
	return Union3[A, B, C]{make([]byte, max(
		unsafe.Sizeof(*new(A)),
		unsafe.Sizeof(*new(B)),
		unsafe.Sizeof(*new(C)),
	))}
}

// NewUnion2A creates new Union3 filled with value of type A
func NewUnion3A[A, B, C any](a A) Union3[A, B, C] {
	u := NewUnion3[A, B, C]()
	*u.A() = a
	return u
}

// NewUnion2B creates new Union3 filled with value of type B
func NewUnion3B[A, B, C any](b B) Union3[A, B, C] {
	u := NewUnion3[A, B, C]()
	*u.B() = b
	return u
}

// NewUnion2C creates new Union3 filled with value of type C
func NewUnion3C[A, B, C any](c C) Union3[A, B, C] {
	u := NewUnion3[A, B, C]()
	*u.C() = c
	return u
}

// C - returns pointer to value interpreted as type C
func (u Union3[A, B, C]) A() *A {
	return reinterpret[A](u.data)
}

// B - returns pointer to value interpreted as type B
func (u Union3[A, B, C]) B() *B {
	return reinterpret[B](u.data)
}

// C - returns pointer to value interpreted as type C
func (u Union3[A, B, C]) C() *C {
	return reinterpret[C](u.data)
}
