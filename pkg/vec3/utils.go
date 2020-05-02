package vec3

import (
	"math"
	"math/rand"

	. "ray-tracing/pkg/utils"
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
