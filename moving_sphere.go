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
	a := Dot(r.Direction, r.Direction)
	b := Dot(oc, r.Direction)
	c := Dot(oc, oc) - s.Radius*s.Radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			return HitRecord{
				T:        temp,
				P:        r.PointAtParameter(temp),
				Normal:   r.PointAtParameter(temp).Sub(s.Center(r.Time)).DivN(s.Radius),
				Material: s.Material,
			}, true
		}

		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			return HitRecord{
				T:        temp,
				P:        r.PointAtParameter(temp),
				Normal:   r.PointAtParameter(temp).Sub(s.Center(r.Time)).DivN(s.Radius),
				Material: s.Material,
			}, true
		}
	}

	return HitRecord{}, false
}

func (s MovingSphere) SurroundingBox(box, box1 AABB) AABB {
	small := Vec3{
		math.Min(box.Min.X, box1.Min.X),
		math.Min(box.Min.Y, box1.Min.Y),
		math.Min(box.Min.Z, box1.Min.Z),
	}

	big := Vec3{
		math.Max(box.Min.X, box1.Min.X),
		math.Max(box.Min.Y, box1.Min.Y),
		math.Max(box.Min.Z, box1.Min.Z),
	}

	return AABB{small, big}
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

	return s.SurroundingBox(box0, box1), true
}
