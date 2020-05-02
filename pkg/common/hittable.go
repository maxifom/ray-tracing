package common

import (
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/vec3"
)

type Hittable interface {
	Hit(r Ray, tMin, tMax float64) (HitRecord, bool)
	BoundingBox(t0, t1 float64) (AABB, bool)
	PDFValue(o, v Vec3) float64
	Random(o Vec3) Vec3
}
