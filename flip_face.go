package main

type FlipFace struct {
	H Hittable
}

func (f FlipFace) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	h, isHit := f.H.Hit(r, tMin, tMax)
	if !isHit {
		return HitRecord{}, false
	}

	h.FrontFace = !h.FrontFace
	return h, true
}

func (f FlipFace) BoundingBox(t0, t1 float64) (AABB, bool) {
	return f.H.BoundingBox(t0, t1)
}

func (f FlipFace) PDFValue(o, v Vec3) float64 {
	return f.H.PDFValue(o, v)
}

func (f FlipFace) Random(o Vec3) Vec3 {
	return f.H.Random(o)
}
