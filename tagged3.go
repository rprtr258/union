package union

// Tagged3 is a value which is either A or B, or C
type Tagged3[A, B, C any] struct {
	Union3[A, B, C]
	isA, isB bool
}

// NewTagged3A creates new Tagged3 which is value of type A
func NewTagged3A[A, B, C any](a A) Tagged3[A, B, C] {
	return Tagged3[A, B, C]{
		Union3: NewUnion3A[A, B, C](a),
		isA:    true,
		isB:    false,
	}
}

// NewTagged3B creates new Tagged3 which is value of type B
func NewTagged3B[A, B, C any](b B) Tagged3[A, B, C] {
	return Tagged3[A, B, C]{
		Union3: NewUnion3B[A, B, C](b),
		isA:    false,
		isB:    true,
	}
}

// NewTagged3C creates new Tagged3 which is value of type C
func NewTagged3C[A, B, C any](c C) Tagged3[A, B, C] {
	return Tagged3[A, B, C]{
		Union3: NewUnion3C[A, B, C](c),
		isA:    false,
		isB:    false,
	}
}

// A - returns pointer to value interpreted as type A.
// If it is actually B or C, it is reinterpreted as A.
func (u Tagged3[A, B, C]) A() *A {
	return u.Union3.A()
}

// B - returns pointer to value interpreted as type B.
// If it is actually A or C, it is reinterpreted as B.
func (u Tagged3[A, B, C]) B() *B {
	return u.Union3.B()
}

// C - returns pointer to value interpreted as type C.
// If it is actually A or B, it is reinterpreted as C.
func (u Tagged3[A, B, C]) C() *C {
	return u.Union3.C()
}

// Switch - calls function depending on value type
func (u Tagged3[A, B, C]) Switch(
	fa func(*A),
	fb func(*B),
	fc func(*C),
) {
	switch {
	case u.isA:
		fa(u.A())
	case u.isB:
		fb(u.B())
	default:
		fc(u.C())
	}
}

// Switch3 evaluates function depending on value type
func Switch3[A, B, C, R any](
	u Tagged3[A, B, C],
	fa func(A) R,
	fb func(B) R,
	fc func(C) R,
) R {
	var res R
	u.Switch(
		func(a *A) { res = fa(*a) },
		func(b *B) { res = fb(*b) },
		func(c *C) { res = fc(*c) },
	)
	return res
}
