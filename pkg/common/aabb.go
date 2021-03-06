package common

import (
	"math"

	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/vec3"
)

type AABB struct {
	Min, Max Vec3
}

func (A AABB) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	x := (A.Min.X - r.Origin.X) / r.Direction.X
	x1 := (A.Max.X - r.Origin.X) / r.Direction.X
	t0 := math.Min(x, x1)
	t1 := math.Max(x, x1)

	tMin = math.Max(t0, tMin)
	tMax = math.Min(t1, tMax)
	if tMax <= tMin {
		return HitRecord{}, false
	}

	y := (A.Min.Y - r.Origin.Y) / r.Direction.Y
	y1 := (A.Max.Y - r.Origin.Y) / r.Direction.Y
	t0 = math.Min(y, y1)
	t1 = math.Max(y, y1)

	tMin = math.Max(t0, tMin)
	tMax = math.Min(t1, tMax)
	if tMax <= tMin {
		return HitRecord{}, false
	}

	z := (A.Min.Z - r.Origin.Z) / r.Direction.Z
	z1 := (A.Max.Z - r.Origin.Z) / r.Direction.Z
	t0 = math.Min(z, z1)
	t1 = math.Max(z, z1)

	tMin = math.Max(z, tMin)
	tMax = math.Min(z1, tMax)
	if tMax <= tMin {
		return HitRecord{}, false
	}

	return HitRecord{}, true
}

func (A AABB) BoundingBox(t0, t1 float64) (AABB, bool) {
	return A, true
}

func (A AABB) PDFValue(o, v Vec3) float64 {
	return 0
}

func (A AABB) Random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
