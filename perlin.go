package main

import (
	"math"
	"math/rand"
)

type Perlin struct {
	RanVec []Vec3
	PermX  []int64
	PermY  []int64
	PermZ  []int64
}

func NewPerlin() Perlin {
	return Perlin{
		RanVec: PerlinGenerate(),
		PermX:  PerlinGeneratePermute(),
		PermY:  PerlinGeneratePermute(),
		PermZ:  PerlinGeneratePermute(),
	}
}

func (pn Perlin) Noise(p Vec3) float64 {
	u := p.X - math.Floor(p.X)
	v := p.Y - math.Floor(p.Y)
	w := p.Z - math.Floor(p.Z)

	i := int(math.Floor(p.X))
	j := int(math.Floor(p.Y))
	k := int(math.Floor(p.Z))

	g := pn.RanVec

	x := pn.PermX
	y := pn.PermY
	z := pn.PermZ
	var c [2][2][2]Vec3
	for di := 0; di < 2; di++ {
		for dj := 0; dj < 2; dj++ {
			for dk := 0; dk < 2; dk++ {
				idx := x[(i+di)&255] ^ y[(j+dj)&255] ^ z[(k+dk)&255]
				c[di][dj][dk] = g[idx]
			}
		}
	}

	return PerlinInterpolation(c, u, v, w)
}

func (pn Perlin) Turb(p Vec3, depth int) float64 {
	accum := 0.0
	weight := 1.0
	for i := 0; i < depth; i++ {
		accum += weight * pn.Noise(p)
		weight *= 0.5
		p = p.MulN(2)
	}

	return math.Abs(accum)
}

func PerlinGenerate() []Vec3 {
	p := make([]Vec3, 256)
	for i := 0; i < 256; i++ {
		p[i] = Vec3{-1 + 2*rand.Float64(), -1 + 2*rand.Float64(), -1 + 2*rand.Float64()}.UnitVector()
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
