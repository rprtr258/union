package union_test

import (
	"testing"

	"github.com/rprtr258/union"
	"github.com/stretchr/testify/assert"
)

func TestFloat32AsUint32(t *testing.T) {
	u := union.NewUnion2B[float32, uint32](0x40490fdb)
	assert.Equal(t, float32(3.14159265), *u.A())
	assert.Equal(t, uint32(0x40490fdb), *u.B())
}

func TestReinterpretReinterpret(t *testing.T) {
	pi := float32(3.14159265)
	ui := *union.NewUnion2A[float32, uint32](pi).B()
	got := *union.NewUnion2B[float32, uint32](ui).A()
	assert.Equal(t, pi, got)
}
