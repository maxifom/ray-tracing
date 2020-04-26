package main

type YZRect struct {
	Y0, Y1, Z0, Z1, X float64
	Material          Material
}

func (r YZRect) Hit(ray Ray, tMin, tMax float64) (HitRecord, bool) {
	t := (r.X - ray.Origin.X) / ray.Direction.X
	if t < tMin || t > tMax {
		return HitRecord{}, false
	}
	y := ray.Origin.Y + t*ray.Direction.Y
	z := ray.Origin.Z + t*ray.Direction.Z
	if y < r.Y0 || y > r.Y1 || z < r.Z0 || z > r.Z1 {
		return HitRecord{}, false
	}

	h := HitRecord{
		U:        (y - r.Y0) / (r.Y1 - r.Y0),
		V:        (z - r.Z0) / (r.Z1 - r.Z0),
		T:        t,
		Material: r.Material,
		P:        ray.PointAtParameter(t),
	}

	outwardNormal := Vec3{1, 0, 0}
	h = h.SetFaceNormal(ray, outwardNormal)
	return h, true
}

func (r YZRect) BoundingBox(t0, t1 float64) (AABB, bool) {
	return AABB{Vec3{r.X - 0.0001, r.Y0, r.Z0}, Vec3{r.X + 0.0001, r.Y1, r.Z1}}, true
}

func (r YZRect) PDFValue(o, v Vec3) float64 {
	return 0
}

func (r YZRect) Random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
