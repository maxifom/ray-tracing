package hittable

import (
	"math"

	. "ray-tracing/pkg/common"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/vec3"
)

// |x-a| + |y-b| + |z-c| = r
type Octahedron struct {
	Center    Vec3
	Radius    float64
	Material  Material
	triangles HittableList
}

func NewOctahedron(center Vec3, radius float64, m Material) Octahedron {
	a := Vec3{radius + center.X, center.Y, center.Z}
	b := Vec3{-radius + center.X, center.Y, center.Z}
	c := Vec3{center.X, radius + center.Y, center.Z}
	d := Vec3{center.X, -radius + center.Y, center.Z}
	e := Vec3{center.X, center.Y, radius + center.Z}
	f := Vec3{center.X, center.Y, -radius + center.Z}
	triangles := make(HittableList, 0, 8)
	triangles = NewList(
		NewTriangle(a, e, d, m),
		NewTriangle(b, d, e, m),
		NewTriangle(a, c, e, m),
		NewTriangle(e, b, c, m),
		NewTriangle(f, d, a, m),
		NewTriangle(a, c, f, m),
		NewTriangle(f, c, b, m),
		NewTriangle(f, b, d, m),
	)

	return Octahedron{Center: center, Radius: radius, triangles: triangles, Material: m}
}

func (o Octahedron) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	return o.triangles.Hit(r, tMin, tMax)
}

func (o Octahedron) BoundingBox(t0, t1 float64) (AABB, bool) {
	return AABB{
		Min: Vec3{
			X: o.Center.X - o.Radius,
			Y: o.Center.Y - o.Radius,
			Z: o.Center.Z - o.Radius,
		},
		Max: Vec3{
			X: o.Center.X + o.Radius,
			Y: o.Center.Y + o.Radius,
			Z: o.Center.Z + o.Radius,
		},
	}, true
}

func (o Octahedron) PDFValue(ov, v Vec3) float64 {
	_, isHit := o.Hit(Ray{ov, v, 0}, 0.001, math.Inf(1))
	if !isHit {
		return 0
	}
	cosThetaMax := math.Sqrt(1.0 - o.Radius*o.Radius/(o.Center.Sub(v).SqrLength()))
	solidAngle := 2 * math.Pi * (1 - cosThetaMax)
	return 1 / solidAngle
}

func (o Octahedron) Random(origin Vec3) Vec3 {
	direction := o.Center.Sub(origin)
	distanceSquared := direction.SqrLength()
	onb := NewONB(direction)
	return onb.Local(RandomToSphere(o.Radius, distanceSquared))
}
