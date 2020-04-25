package main

import "math"

type Sphere struct {
	Center   Vec3
	Radius   float64
	Material Material
}

func (s Sphere) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.SqrLength()
	halfB := Dot(oc, r.Direction)
	c := oc.SqrLength() - s.Radius*s.Radius
	discriminant := halfB*halfB - a*c
	if discriminant > 0 {
		root := math.Sqrt(discriminant)
		temp := (-halfB - root) / a
		if temp < tMax && temp > tMin {
			h := HitRecord{
				T:        temp,
				P:        r.PointAtParameter(temp),
				Normal:   r.PointAtParameter(temp).Sub(s.Center).DivN(s.Radius),
				Material: s.Material,
			}
			h.U, h.V = GetSphereUV(h.P.Sub(s.Center).DivN(s.Radius))
			outwardNormal := (h.P.Sub(s.Center)).DivN(s.Radius)
			h = h.SetFaceNormal(r, outwardNormal)
			return h, true
		}

		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			h := HitRecord{
				T:        temp,
				P:        r.PointAtParameter(temp),
				Normal:   r.PointAtParameter(temp).Sub(s.Center).DivN(s.Radius),
				Material: s.Material,
			}

			h.U, h.V = GetSphereUV(h.P.Sub(s.Center).DivN(s.Radius))

			outwardNormal := (h.P.Sub(s.Center)).DivN(s.Radius)
			h = h.SetFaceNormal(r, outwardNormal)
			return h, true
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
