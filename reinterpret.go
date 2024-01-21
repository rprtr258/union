package union

import "unsafe"

func reinterpret[T any](b []byte) *T {
	return (*T)(unsafe.Pointer(unsafe.SliceData(b)))
}
