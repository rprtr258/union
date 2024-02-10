package union

// TaggedF2 is a value which is either A or B, uses closures
type TaggedF2[A, B any] func(func(A), func(B))

func NewTaggedF2A[A, B any](l A) TaggedF2[A, B] {
	return func(fl func(A), _ func(B)) {
		fl(l)
	}
}

func NewTaggedF2B[A, B any](r B) TaggedF2[A, B] {
	return func(_ func(A), fr func(B)) {
		fr(r)
	}
}

func (t TaggedF2[A, B]) A() A {
	var res A
	t(
		func(a A) { res = a },
		func(B) {
			panic("cannot get A while there is B")
		},
	)
	return res
}

func (t TaggedF2[A, B]) B() B {
	var res B
	t(
		func(A) {
			panic("cannot get B while there is A")
		},
		func(b B) { res = b },
	)
	return res
}

func SwitchF[A, B any](
	u TaggedF2[A, B],
	fa func(A),
	fb func(B),
) {
	u(fa, fb)
}

func SwitchF2[A, B, Out any](
	u TaggedF2[A, B],
	fa func(A) Out,
	fb func(B) Out,
) Out {
	var res Out
	u(
		func(l A) { res = fa(l) },
		func(r B) { res = fb(r) },
	)
	return res
}

type Result[T any] TaggedF2[T, error]

func Ok[T any](t T) Result[T] {
	return Result[T](NewTaggedF2A[T, error](t))
}

func Err[T any](err error) Result[T] {
	return Result[T](NewTaggedF2B[T, error](err))
}

func Pack[T any](t T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}

	return Ok(t)
}

func (e Result[T]) Unpack() (T, error) {
	var res T
	var err error
	e(
		func(t T) { res = t },
		func(e error) { err = e },
	)
	return res, err
}

func MapResult[T, R any](r Result[T], f func(T) R) Result[R] {
	return SwitchF2[T, error, Result[R]](
		TaggedF2[T, error](r),
		func(t T) Result[R] {
			return Ok(f(t))
		},
		func(err error) Result[R] {
			return Err[R](err)
		})
}

func FlatMapResult[T, R any](r Result[T], f func(T) Result[R]) Result[R] {
	return SwitchF2[T, error, Result[R]](
		TaggedF2[T, error](r),
		func(t T) Result[R] {
			return f(t)
		},
		func(err error) Result[R] {
			return Err[R](err)
		})
}
