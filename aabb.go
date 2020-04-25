package main

import "math"

type AABB struct {
	Min, Max Vec3
}

func (A AABB) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	x := (A.Min.X - r.Origin.X) / r.Direction.X
	x1 := (A.Max.X - r.Origin.X) / r.Direction.X
	t0 := math.Min(x, x1)
	t1 := math.Max(x, x1)

	tMin = math.Max(t0, tMin)
	tMax = math.Max(t1, tMax)
	if tMax <= tMin {
		return HitRecord{}, false
	}

	y := (A.Min.X - r.Origin.Y) / r.Direction.Y
	y1 := (A.Max.X - r.Origin.Y) / r.Direction.Y
	t0 = math.Min(y, y1)
	t1 = math.Max(y, y1)

	tMin = math.Max(t0, tMin)
	tMax = math.Max(t1, tMax)
	if tMax <= tMin {
		return HitRecord{}, false
	}

	z := (A.Min.X - r.Origin.Z) / r.Direction.Z
	z1 := (A.Max.X - r.Origin.Z) / r.Direction.Z
	t0 = math.Min(x, x1)
	t1 = math.Max(x, x1)

	tMin = math.Max(z, tMin)
	tMax = math.Max(z1, tMax)
	if tMax <= tMin {
		return HitRecord{}, false
	}

	return HitRecord{}, true
}
