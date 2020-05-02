package hittable

import (
	"math"
	"math/rand"

	"ray-tracing/pkg/common"
	. "ray-tracing/pkg/vec3"
)

func SurroundingBox(box, box1 common.AABB) common.AABB {
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

	return common.AABB{small, big}
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
func GetSphereUV(p Vec3) (u float64, v float64) {
	phi := math.Atan2(p.Z, p.X)
	theta := math.Asin(p.Y)
	return 1 - (phi+math.Pi)/(2*math.Pi),
		(theta + math.Pi/2) / math.Pi
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
