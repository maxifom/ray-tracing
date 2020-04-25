package main

import "math"

type MovingSphere struct {
	Center0, Center1 Vec3
	Radius           float64
	Material         Material

	Time0, Time1 float64
}

func (s MovingSphere) Center(time float64) Vec3 {
	return s.Center0.
		Add(
			s.Center1.
				Sub(s.Center0).
				MulN((time - s.Time0) / (s.Time1 - s.Time0)),
		)
}

func (s MovingSphere) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	oc := r.Origin.Sub(s.Center(r.Time))
	a := r.Direction.SqrLength()
	halfB := Dot(oc, r.Direction)
	c := oc.SqrLength() - s.Radius*s.Radius
	discriminant := halfB*halfB - a*c
	if discriminant > 0 {
		root := math.Sqrt(halfB*halfB - a*c)
		temp := (-halfB - root) / a
		if temp < tMax && temp > tMin {
			h := HitRecord{
				T:        temp,
				P:        r.PointAtParameter(temp),
				Normal:   r.PointAtParameter(temp).Sub(s.Center(r.Time)).DivN(s.Radius),
				Material: s.Material,
			}
			h.U, h.V = GetSphereUV(h.P.Sub(s.Center(r.Time)).DivN(s.Radius))
			outwardNormal := h.P.Sub(s.Center(r.Time)).DivN(s.Radius)
			h = h.SetFaceNormal(r, outwardNormal)
			return h, true
		}

		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			h := HitRecord{
				T:        temp,
				P:        r.PointAtParameter(temp),
				Normal:   r.PointAtParameter(temp).Sub(s.Center(r.Time)).DivN(s.Radius),
				Material: s.Material,
			}
			h.U, h.V = GetSphereUV(h.P.Sub(s.Center(r.Time)).DivN(s.Radius))
			outwardNormal := h.P.Sub(s.Center(r.Time)).DivN(s.Radius)
			h = h.SetFaceNormal(r, outwardNormal)
			return h, true
		}
	}

	return HitRecord{}, false
}

func (s MovingSphere) BoundingBox(t0, t1 float64) (AABB, bool) {
	box0 := AABB{
		s.Center0.Sub(Vec3{s.Radius, s.Radius, s.Radius}),
		s.Center0.Add(Vec3{s.Radius, s.Radius, s.Radius}),
	}

	box1 := AABB{
		s.Center1.Sub(Vec3{s.Radius, s.Radius, s.Radius}),
		s.Center1.Add(Vec3{s.Radius, s.Radius, s.Radius}),
	}

	return SurroundingBox(box0, box1), true
}

func (s MovingSphere) PDFValue(o, v Vec3) float64 {
	return 0
}

func (s MovingSphere) Random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
