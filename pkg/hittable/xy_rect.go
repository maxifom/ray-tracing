package hittable

import (
	"math"

	. "ray-tracing/pkg/common"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/utils"
	. "ray-tracing/pkg/vec3"
)

type XYRect struct {
	X0, X1, Y0, Y1, Z float64
	Material          Material
}

func (r XYRect) Hit(ray Ray, tMin, tMax float64) (HitRecord, bool) {
	t := (r.Z - ray.Origin.Z) / ray.Direction.Z
	if t < tMin || t > tMax {
		return HitRecord{}, false
	}
	x := ray.Origin.X + t*ray.Direction.X
	y := ray.Origin.Y + t*ray.Direction.Y
	if x < r.X0 || x > r.X1 || y < r.Y0 || y > r.Y1 {
		return HitRecord{}, false
	}

	h := HitRecord{
		U:        (x - r.X0) / (r.X1 - r.X0),
		V:        (y - r.Y0) / (r.Y1 - r.Y0),
		T:        t,
		Material: r.Material,
		P:        ray.PointAtParameter(t),
	}

	outwardNormal := Vec3{0, 0, 1}
	h = h.SetFaceNormal(ray, outwardNormal)
	return h, true
}

func (r XYRect) BoundingBox(t0, t1 float64) (AABB, bool) {
	return AABB{Vec3{r.X0, r.Y0, r.Z - 0.0001}, Vec3{r.X1, r.Y1, r.Z + 0.0001}}, true
}

func (r XYRect) PDFValue(o, v Vec3) float64 {
	rec, isHit := r.Hit(Ray{o, v, 0}, 0.001, math.Inf(1))
	if !isHit {
		return 0
	}

	area := (r.Y1 - r.Y0) * (r.X1 - r.X0)
	distanceSquared := rec.T * rec.T * v.SqrLength()
	cosine := math.Abs(Dot(v, rec.Normal)) / v.Length()
	return distanceSquared / (cosine * area)
}

func (r XYRect) Random(o Vec3) Vec3 {
	randomPoint := Vec3{RandomDouble(r.X0, r.X1), RandomDouble(r.Y0, r.Y1), r.Z}
	return randomPoint.Sub(o).UnitVector()
}
