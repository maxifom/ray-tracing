package texture

import (
	"math/rand"

	. "ray-tracing/pkg/vec3"
)

type PerlinFloat struct {
	RanFloat []float64
	PermX    []int64
	PermY    []int64
	PermZ    []int64
}

func NewPerlinFloat() PerlinFloat {
	return PerlinFloat{
		RanFloat: PerlinGenerateFloat(),
		PermX:    PerlinGeneratePermute(),
		PermY:    PerlinGeneratePermute(),
		PermZ:    PerlinGeneratePermute(),
	}
}

func (pn PerlinFloat) Noise(p Vec3) float64 {
	i := int(4*p.X) & 255
	j := int(4*p.Y) & 255
	k := int(4*p.Z) & 255

	g := pn.RanFloat

	x := pn.PermX
	y := pn.PermY
	z := pn.PermZ

	return g[x[i]^y[j]^z[k]]
}

func PerlinGenerateFloat() []float64 {
	p := make([]float64, 256)
	for i := 0; i < 256; i++ {
		p[i] = rand.Float64()
	}

	return p
}
