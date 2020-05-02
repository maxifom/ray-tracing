package hittable

import (
	"ray-tracing/pkg/common"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/vec3"
)

type Box struct {
	BoxMin, BoxMax Vec3
	Sides          common.Hittable
}

func NewBox(p0, p1 Vec3, m common.Material) Box {
	return Box{
		BoxMin: p0,
		BoxMax: p1,
		Sides: NewList(
			XYRect{p0.X, p1.X, p0.Y, p1.Y, p1.Z, m},
			FlipFace{XYRect{p0.X, p1.X, p0.Y, p1.Y, p0.Z, m}},

			XZRect{p0.X, p1.X, p0.Z, p1.Z, p1.Y, m},
			FlipFace{XZRect{p0.X, p1.X, p0.Z, p1.Z, p0.Y, m}},

			YZRect{p0.Y, p1.Y, p0.Z, p1.Z, p1.X, m},
			FlipFace{YZRect{p0.Y, p1.Y, p0.Z, p1.Z, p0.X, m}},
		),
	}
}

func (b Box) Hit(r Ray, tMin, tMax float64) (common.HitRecord, bool) {
	return b.Sides.Hit(r, tMin, tMax)
}

func (b Box) BoundingBox(t0, t1 float64) (common.AABB, bool) {
	return common.AABB{b.BoxMin, b.BoxMax}, true
}

func (b Box) PDFValue(o, v Vec3) float64 {
	return 0
}

func (b Box) Random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
