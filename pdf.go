package main

import "math"

type PDF interface {
	Value(direction Vec3) float64
	Generate() Vec3
}

type CosinePDF struct {
	ONB ONB
}

func (p CosinePDF) Value(direction Vec3) float64 {
	cosine := Dot(direction.UnitVector(), p.ONB.W)
	if cosine <= 0 {
		return 0
	}

	return cosine / math.Pi
}

func (p CosinePDF) Generate() Vec3 {
	return p.ONB.Local(RandomCosineDirection())
}

type DefaultPDF struct {
}

func (d DefaultPDF) Value(direction Vec3) float64 {
	return 0
}

func (d DefaultPDF) Generate() Vec3 {
	return Vec3{}
}
