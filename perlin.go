package main

import (
	"math/rand"
)

type Perlin struct {
	RanFloat []float64
	PermX    []int64
	PermY    []int64
	PermZ    []int64
}

func NewPerlin() Perlin {
	return Perlin{
		RanFloat: PerlinGenerate(),
		PermX:    PerlinGeneratePermute(),
		PermY:    PerlinGeneratePermute(),
		PermZ:    PerlinGeneratePermute(),
	}
}

func (pn Perlin) Noise(p Vec3) float64 {
	// u := p.X - math.Floor(p.X)
	// v := p.Y - math.Floor(p.Y)
	// w := p.Z - math.Floor(p.Z)
	i := int64(4*p.X) & 255
	j := int64(4*p.Y) & 255
	k := int64(4*p.Z) & 255
	generate := pn.RanFloat
	return generate[pn.PermX[i]^pn.PermY[j]^pn.PermZ[k]]
}

func PerlinGenerate() []float64 {
	p := make([]float64, 256)
	for i := 0; i < 256; i++ {
		p[i] = rand.Float64()
	}
	return p
}

func Permute(p []int64) []int64 {
	for i := range p {
		target := int64(rand.Float64() * (float64(i) + 1))
		p[i], p[target] = p[target], p[i]
	}

	return p
}

func PerlinGeneratePermute() []int64 {
	p := make([]int64, 256)
	for i := int64(0); i < 256; i++ {
		p[i] = i
	}

	p = Permute(p)
	return p
}
