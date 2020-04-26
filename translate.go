package main

type Translate struct {
	H      Hittable
	Offset Vec3
}

func (t Translate) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	movedR := Ray{r.Origin.Sub(t.Offset), r.Direction, r.Time}
	h, isHit := t.H.Hit(movedR, tMin, tMax)
	if !isHit {
		return HitRecord{}, false
	}

	h.P = h.P.Add(t.Offset)
	h = h.SetFaceNormal(movedR, h.Normal)

	return h, true
}

func (t Translate) BoundingBox(t0, t1 float64) (AABB, bool) {
	box, isBounded := t.H.BoundingBox(t0, t1)
	if !isBounded {
		return AABB{}, false
	}

	return AABB{
		Min: box.Min.Add(t.Offset),
		Max: box.Max.Add(t.Offset),
	}, true
}

func (t Translate) PDFValue(o, v Vec3) float64 {
	return 0
}

func (t Translate) Random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
