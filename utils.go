package main

import (
	"math"
	"math/rand"
)

func RandomInUnitSphere() Vec3 {
	p := Vec3{rand.Float64(), rand.Float64(), rand.Float64()}.MulN(2).Sub(Vec3{1, 1, 1})

	for p.SqrLength() >= 1 {
		p = Vec3{rand.Float64(), rand.Float64(), rand.Float64()}.MulN(2).Sub(Vec3{1, 1, 1})
	}

	return p
}

func RandomInUnitDisk() Vec3 {
	p := Vec3{rand.Float64(), rand.Float64(), 0}.MulN(2).Sub(Vec3{1, 1, 0})

	for Dot(p, p) >= 1 {
		p = Vec3{rand.Float64(), rand.Float64(), 0}.MulN(2).Sub(Vec3{1, 1, 0})
	}

	return p
}

func Reflect(v, v1 Vec3) Vec3 {
	return v.Sub(v1.MulN(2 * Dot(v, v1)))
}

func Refract(v, v1 Vec3, niOverNt float64) (Vec3, bool) {
	uv := v.UnitVector()
	dt := Dot(uv, v1)
	discriminant := 1 - niOverNt*niOverNt*(1-dt*dt)

	if discriminant <= 0 {
		return Vec3{}, false

	}
	return uv.Sub(v1.MulN(dt)).MulN(niOverNt).Sub(v1.MulN(math.Sqrt(discriminant))), true
}

func Schlick(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

func SurroundingBox(box, box1 AABB) AABB {
	small := Vec3{
		math.Min(box.Min.X, box1.Min.X),
		math.Min(box.Min.Y, box1.Min.Y),
		math.Min(box.Min.Z, box1.Min.Z),
	}

	big := Vec3{
		math.Max(box.Min.X, box1.Min.X),
		math.Max(box.Min.Y, box1.Min.Y),
		math.Max(box.Min.Z, box1.Min.Z),
	}

	return AABB{small, big}
}

func PerlinInterpolation(c [2][2][2]Vec3, u, v, w float64) float64 {
	uu := u * u * (3 - 2*u)
	vv := v * v * (3 - 2*v)
	ww := w * w * (3 - 2*w)
	accum := 0.0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				i1 := float64(i)
				j1 := float64(j)
				k1 := float64(k)
				weightV := Vec3{u - i1, v - j1, w - k1}
				accum += (i1*uu + (1-i1)*(1-uu)) *
					(j1*vv + (1-j1)*(1-vv)) *
					(k1*ww + (1-k1)*(1-ww)) *
					Dot(c[i][j][k], weightV)
			}
		}
	}

	return accum
}
