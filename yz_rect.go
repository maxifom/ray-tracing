package main

type YZRect struct {
	Y0, Y1, Z0, Z1, K float64
	Material          Material
}

func (r YZRect) Hit(ray Ray, tMin, tMax float64) (HitRecord, bool) {
	t := (r.K - ray.Origin.Z) / ray.Direction.Z
	if t < tMin || t > tMax {
		return HitRecord{}, false
	}
	x := ray.Origin.X + t*ray.Direction.X
	y := ray.Origin.Y + t*ray.Direction.Y
	if x < r.Y0 || x > r.Y1 || y < r.Z0 || y > r.Z1 {
		return HitRecord{}, false
	}

	h := HitRecord{
		U:        (x - r.Y0) / (r.Y1 - r.Y0),
		V:        (y - r.Z0) / (r.Z1 - r.Z0),
		T:        t,
		Material: r.Material,
		P:        ray.PointAtParameter(t),
	}

	outwardNormal := Vec3{0, 0, 1}
	h = h.SetFaceNormal(ray, outwardNormal)
	return h, true
}

func (r YZRect) BoundingBox(t0, t1 float64) (AABB, bool) {
	return AABB{Vec3{r.Y0, r.Z0, r.K - 0.0001}, Vec3{r.Y1, r.Z1, r.K + 0.0001}}, true
}
