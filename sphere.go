package main

import "math"

type Sphere struct {
	Center   Vec3
	Radius   float64
	Material Material
}

func (s Sphere) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	oc := r.Origin.Sub(s.Center)
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
				Normal:   r.PointAtParameter(temp).Sub(s.Center).DivN(s.Radius),
				Material: s.Material,
			}, true
		}

		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			return HitRecord{
				T:        temp,
				P:        r.PointAtParameter(temp),
				Normal:   r.PointAtParameter(temp).Sub(s.Center).DivN(s.Radius),
				Material: s.Material,
			}, true
		}
	}

	return HitRecord{}, false
}

func (s Sphere) BoundingBox(t0, t1 float64) (AABB, bool) {
	return AABB{
		s.Center.Sub(Vec3{s.Radius, s.Radius, s.Radius}),
		s.Center.Add(Vec3{s.Radius, s.Radius, s.Radius}),
	}, true
}
