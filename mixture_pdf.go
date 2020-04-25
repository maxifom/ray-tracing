package main

import "math/rand"

type MixturePDF struct {
	P0, P1 PDF
}

func (m MixturePDF) Generate() Vec3 {
	if rand.Float64() < 0.5 {
		return m.P0.Generate()
	}
	return m.P1.Generate()
}

func (m MixturePDF) Value(direction Vec3) float64 {
	return 0.5*m.P0.Value(direction) + 0.5*m.P1.Value(direction)
}
