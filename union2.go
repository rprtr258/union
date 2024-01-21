package union

import "unsafe"

// Union2 is a value which can be interpreted as either A or B.
// Effectively this is pointer to C union
type Union2[A, B any] struct {
	data []byte // of constant len=max(sizeof(A), sizeof(B))
}

// NewUnion2 creates new Union2 filled with zeros
func NewUnion2[A, B any]() Union2[A, B] {
	return Union2[A, B]{make([]byte, max(
		unsafe.Sizeof(*new(A)),
		unsafe.Sizeof(*new(B)),
	))}
}

// NewUnion2A creates new Union2 filled with value of type A
func NewUnion2A[A, B any](a A) Union2[A, B] {
	u := NewUnion2[A, B]()
	*u.A() = a
	return u
}

// NewUnion2B creates new Union2 filled with value of type B
func NewUnion2B[A, B any](b B) Union2[A, B] {
	u := NewUnion2[A, B]()
	*u.B() = b
	return u
}

// A - returns pointer to value interpreted as type A
func (u Union2[A, B]) A() *A {
	return reinterpret[A](u.data)
}

// B - returns pointer to value interpreted as type B
func (u Union2[A, B]) B() *B {
	return reinterpret[B](u.data)
}
