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
	p := Vec3{RandomDouble(-1, 1), RandomDouble(-1, 1), 0}

	for p.SqrLength() >= 1 {
		p = Vec3{RandomDouble(-1, 1), RandomDouble(-1, 1), 0}
	}

	return p
}

func Reflect(v, n Vec3) Vec3 {
	return v.Sub(n.MulN(2 * Dot(v, n)))
}

func Refract(uv, n Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(Dot(uv.Negative(), n), 1.0)
	rOutParallel := (uv.Add(n.MulN(cosTheta))).MulN(etaiOverEtat)
	rOutPerp := n.MulN(-math.Sqrt(1.0 - rOutParallel.SqrLength()))
	return rOutParallel.Add(rOutPerp)
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
		math.Max(box.Max.X, box1.Max.X),
		math.Max(box.Max.Y, box1.Max.Y),
		math.Max(box.Max.Z, box1.Max.Z),
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

func GetSphereUV(p Vec3) (u float64, v float64) {
	phi := math.Atan2(p.Z, p.X)
	theta := math.Asin(p.Y)
	return 1 - (phi+math.Pi)/(2*math.Pi),
		(theta + math.Pi/2) / math.Pi
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func RandomUnitVector() Vec3 {
	a := rand.Float64() * 2 * math.Pi // 0 -> 2PI
	z := -1 + 2*rand.Float64()        // -1 -> 1
	r := math.Sqrt(1.0 - z*z)
	return Vec3{r * math.Cos(a), r * math.Sin(a), z}
}

func RandomCosineDirection() Vec3 {
	r1 := rand.Float64()
	r2 := rand.Float64()
	z := math.Sqrt(1.0 - r2)
	phi := 2 * math.Pi * r1
	x := math.Cos(phi) * math.Sqrt(r2)
	y := math.Sin(phi) * math.Sqrt(r2)
	return Vec3{x, y, z}
}

func RandomDouble(a, b float64) float64 {
	return a + (b-a)*rand.Float64()
}

func RandomToSphere(radius, distanceSquared float64) Vec3 {
	r1 := rand.Float64()
	r2 := rand.Float64()
	z := 1 + r2*(math.Sqrt(1.0-radius*radius/distanceSquared)-1)

	phi := 2 * math.Pi * r1
	x := math.Cos(phi) * math.Sqrt(1-z*z)
	y := math.Sin(phi) * math.Sqrt(1-z*z)
	return Vec3{x, y, z}
}
