package main

type HitRecord struct {
	T         float64
	P, Normal Vec3
	Material  Material
	U, V      float64
	FrontFace bool
}

func (h HitRecord) SetFaceNormal(ray Ray, outwardNormal Vec3) HitRecord {
	h.FrontFace = Dot(ray.Direction, outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Negative()
	}

	return h
}
