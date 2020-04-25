package main

type XZRect struct {
	X0, X1, Z0, Z1, Y float64
	Material          Material
}

func (r XZRect) Hit(ray Ray, tMin, tMax float64) (HitRecord, bool) {
	t := (r.Y - ray.Origin.Y) / ray.Direction.Y
	if t < tMin || t > tMax {
		return HitRecord{}, false
	}
	x := ray.Origin.X + t*ray.Direction.X
	z := ray.Origin.Z + t*ray.Direction.Z
	if x < r.X0 || x > r.X1 || z < r.Z0 || z > r.Z1 {
		return HitRecord{}, false
	}

	h := HitRecord{
		U:        (x - r.X0) / (r.X1 - r.X0),
		V:        (z - r.Z0) / (r.Z1 - r.Z0),
		T:        t,
		Material: r.Material,
		P:        ray.PointAtParameter(t),
	}

	outwardNormal := Vec3{0, 1, 0}
	h = h.SetFaceNormal(ray, outwardNormal)
	return h, true
}

func (r XZRect) BoundingBox(t0, t1 float64) (AABB, bool) {
	return AABB{Vec3{r.X0, r.Y - 0.0001, r.Z0}, Vec3{r.X1, r.Y + 0.0001, r.Z1}}, true
}
