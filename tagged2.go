package union

// Tagged2 is a value which is either A or B
type Tagged2[A, B any] struct {
	Union2[A, B]
	isB bool
}

// NewTagged2A creates new Tagged2 which is value of type A
func NewTagged2A[A, B any](a A) Tagged2[A, B] {
	return Tagged2[A, B]{
		Union2: NewUnion2A[A, B](a),
		isB:    false,
	}
}

// NewTagged2B creates new Tagged2 which is value of type B
func NewTagged2B[A, B any](b B) Tagged2[A, B] {
	return Tagged2[A, B]{
		Union2: NewUnion2B[A, B](b),
		isB:    true,
	}
}

// A - returns pointer to value interpreted as type A.
// If it is actually B, it is reinterpreted as A.
func (u Tagged2[A, B]) A() *A {
	return u.Union2.A()
}

// B - returns pointer to value interpreted as type B.
// If it is actually A, it is reinterpreted as B.
func (u Tagged2[A, B]) B() *B {
	return u.Union2.B()
}

// Switch - calls function depending on value type
func (u Tagged2[A, B]) Switch(
	fa func(*A),
	fb func(*B),
) {
	if u.isB {
		fb(u.B())
	} else {
		fa(u.A())
	}
}

// Switch2 evaluates function depending on value type
func Switch2[A, B, R any](
	u Tagged2[A, B],
	fa func(A) R,
	fb func(B) R,
) R {
	var res R
	u.Switch(
		func(a *A) { res = fa(*a) },
		func(b *B) { res = fb(*b) },
	)
	return res
}
