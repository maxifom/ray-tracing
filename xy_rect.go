package main

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
