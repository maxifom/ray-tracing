package main

import "math"

type RotateY struct {
	H                  Hittable
	SinTheta, CosTheta float64
	HasBox             bool
	BBox               AABB
}

func (ry RotateY) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	origin := r.Origin
	direction := r.Direction
	origin.X = ry.CosTheta*r.Origin.X - ry.SinTheta*r.Origin.Z
	origin.Z = ry.SinTheta*r.Origin.X + ry.CosTheta*r.Origin.Z

	direction.X = ry.CosTheta*r.Direction.X - ry.SinTheta*r.Direction.Z
	direction.Z = ry.SinTheta*r.Direction.X + ry.CosTheta*r.Direction.Z

	rotatedR := Ray{origin, direction, r.Time}

	h, isHit := ry.H.Hit(rotatedR, tMin, tMax)
	if !isHit {
		return HitRecord{}, false
	}

	p := h.P
	normal := h.Normal

	p.X = ry.CosTheta*h.P.X + ry.SinTheta*h.P.Z
	p.Z = -ry.SinTheta*h.P.X + ry.CosTheta*h.P.Z

	normal.X = ry.CosTheta*h.Normal.X + ry.SinTheta*h.Normal.Z
	normal.Z = -ry.SinTheta*h.Normal.X + ry.CosTheta*h.Normal.Z

	h.P = p
	h = h.SetFaceNormal(rotatedR, normal)

	return h, true
}

func (ry RotateY) BoundingBox(t0, t1 float64) (AABB, bool) {
	return ry.BBox, ry.HasBox
}

func (ry RotateY) PDFValue(o, v Vec3) float64 {
	return 0
}

func (ry RotateY) Random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}

func NewRotateY(h Hittable, angle float64) RotateY {
	radians := angle * (math.Pi / 180.0)
	sinTheta := math.Sin(radians)
	cosTheta := math.Cos(radians)
	bbox, hasBox := h.BoundingBox(0, 1)
	min := Vec3{math.Inf(1), math.Inf(1), math.Inf(1)}
	max := Vec3{math.Inf(-1), math.Inf(-1), math.Inf(-1)}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				i1 := float64(i)
				j1 := float64(j)
				k1 := float64(k)
				x := i1*bbox.Max.X + (1.0-i1)*bbox.Min.X
				y := j1*bbox.Max.Y + (1.0-j1)*bbox.Min.Y
				z := k1*bbox.Max.Z + (1.0-k1)*bbox.Min.Z

				newX := cosTheta*x + sinTheta*z
				newZ := -sinTheta*x + cosTheta*z

				tester := Vec3{newX, y, newZ}

				min.X = math.Min(min.X, tester.X)
				min.Y = math.Min(min.Y, tester.Y)
				min.Z = math.Min(min.Z, tester.Z)

				max.X = math.Max(max.X, tester.X)
				max.Y = math.Max(max.Y, tester.Y)
				max.Z = math.Max(max.Z, tester.Z)
			}
		}
	}

	bbox = AABB{min, max}
	return RotateY{
		H:        h,
		SinTheta: sinTheta,
		CosTheta: cosTheta,
		HasBox:   hasBox,
		BBox:     bbox,
	}
}
