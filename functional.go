package union

// Example of function sum type

// Type declaration

type ExprNode func(
	Number func(float64),
	Sum func([]ExprNode),
	Product func([]ExprNode),
)

// Variant constructors

func ExprNodeNumber(n float64) ExprNode {
	return func(
		f func(float64),
		_ func([]ExprNode),
		_ func([]ExprNode),
	) {
		if f != nil {
			f(n)
		}
	}
}

func ExprNodeSum(xs ...ExprNode) ExprNode {
	return func(
		_ func(float64),
		f func([]ExprNode),
		_ func([]ExprNode),
	) {
		if f != nil {
			f(xs)
		}
	}
}

func ExprNodeProduct(xs ...ExprNode) ExprNode {
	return func(
		_ func(float64),
		_ func([]ExprNode),
		f func([]ExprNode),
	) {
		if f != nil {
			f(xs)
		}
	}
}

// Type mapping

type ExprNodeMapping[T any] struct {
	Number  func(float64) T
	Sum     func([]ExprNode) T
	Product func([]ExprNode) T
}

func ExprNodeMap[T any](n ExprNode, f ExprNodeMapping[T]) (T, bool) {
	var res T
	ok := false
	// exhaustive switch is done by just calling value with handlers
	n(
		func(n float64) {
			if f.Number != nil {
				res = f.Number(n)
				ok = true
			}
		},
		func(xs []ExprNode) {
			if f.Sum != nil {
				res = f.Sum(xs)
				ok = true
			}
		},
		func(xs []ExprNode) {
			if f.Product != nil {
				res = f.Product(xs)
				ok = true
			}
		},
	)
	return res, ok
}

func ExprNodeMapDefault[T any](n ExprNode, f ExprNodeMapping[T], defaultT T) T {
	res := defaultT
	n(
		func(n float64) {
			if f.Number != nil {
				res = f.Number(n)
			}
		},
		func(xs []ExprNode) {
			if f.Sum != nil {
				res = f.Sum(xs)
			}
		},
		func(xs []ExprNode) {
			if f.Product != nil {
				res = f.Product(xs)
			}
		},
	)
	return res
}

// Variant predicates

func (n ExprNode) IsNumber() bool {
	return ExprNodeMapDefault(n, ExprNodeMapping[bool]{Number: func(float64) bool { return true }}, false)
}

func (n ExprNode) IsSum() bool {
	return ExprNodeMapDefault(n, ExprNodeMapping[bool]{Sum: func([]ExprNode) bool { return true }}, false)
}

func (n ExprNode) IsProduct() bool {
	return ExprNodeMapDefault(n, ExprNodeMapping[bool]{Product: func([]ExprNode) bool { return true }}, false)
}

// Mutation

func (n *ExprNode) Update(
	Number func(*float64),
	Sum func(*[]ExprNode),
	Product func(*[]ExprNode),
) {
	(*n)(
		func(x float64) {
			if Number != nil {
				Number(&x)
				*n = ExprNodeNumber(x)
			}
		},
		func(xs []ExprNode) {
			if Sum != nil {
				Sum(&xs)
				*n = ExprNodeSum(xs...)
			}
		},
		func(xs []ExprNode) {
			if Product != nil {
				Product(&xs)
				*n = ExprNodeProduct(xs...)
			}
		},
	)
}
