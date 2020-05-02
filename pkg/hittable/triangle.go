package hittable

import (
	. "ray-tracing/pkg/common"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/vec3"
)

type Triangle struct {
	A, B, C  Vec3
	Material Material
	Normal   Vec3
}

func NewTriangle(A, B, C Vec3, material Material) Triangle {
	ab := B.Sub(A)
	ac := C.Sub(A)
	normal := Cross(ab, ac).UnitVector()

	return Triangle{
		A, B, C, material, normal,
	}

}

func (t Triangle) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	res := MollerTrumbor(r.Origin, r.Direction, t.A, t.B, t.C)
	if res > tMin && res < tMax {
		return HitRecord{
			T:        res,
			P:        r.PointAtParameter(res),
			Normal:   t.Normal,
			Material: t.Material,
		}, true
	}

	return HitRecord{}, false
}

func (t Triangle) BoundingBox(t0, t1 float64) (AABB, bool) {
	return AABB{}, false
}

func (t Triangle) PDFValue(o, v Vec3) float64 {
	return 0
}

func (t Triangle) Random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
