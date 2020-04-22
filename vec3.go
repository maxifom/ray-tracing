package main

import "math"

// Реализация трехмерного вектора
type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Negative() Vec3 {
	return Vec3{
		X: -v.X,
		Y: -v.Y,
		Z: -v.Z,
	}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) SqrLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) UnitVector() Vec3 {
	l := v.Length()
	return Vec3{
		X: v.X / l,
		Y: v.Y / l,
		Z: v.Z / l,
	}
}

func (v Vec3) Add(v1 Vec3) Vec3 {
	return Vec3{
		X: v.X + v1.X,
		Y: v.Y + v1.Y,
		Z: v.Z + v1.Z,
	}
}

func (v Vec3) Sub(v1 Vec3) Vec3 {
	return Vec3{
		X: v.X - v1.X,
		Y: v.Y - v1.Y,
		Z: v.Z - v1.Z,
	}
}

func (v Vec3) Mul(v1 Vec3) Vec3 {
	return Vec3{
		X: v.X * v1.X,
		Y: v.Y * v1.Y,
		Z: v.Z * v1.Z,
	}
}

func (v Vec3) Div(v1 Vec3) Vec3 {
	return Vec3{
		X: v.X / v1.X,
		Y: v.Y / v1.Y,
		Z: v.Z / v1.Z,
	}
}

func (v Vec3) MulN(n float64) Vec3 {
	return Vec3{
		X: v.X * n,
		Y: v.Y * n,
		Z: v.Z * n,
	}
}

func (v Vec3) DivN(n float64) Vec3 {
	return Vec3{
		X: v.X / n,
		Y: v.Y / n,
		Z: v.Z / n,
	}
}

func Dot(v, v1 Vec3) float64 {
	return v.X*v1.X + v.Y*v1.Y + v.Z*v1.Z
}

func Cross(v, v1 Vec3) Vec3 {
	return Vec3{
		X: v.Y*v1.Z - v1.Z*v1.Y,
		Y: v.Z*v1.X - v.X*v1.Z,
		Z: v.X*v1.Y - v.Y*v1.X,
	}
}
