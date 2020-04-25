package main

type HittableList []Hittable

func NewList(h ...Hittable) HittableList {
	return h
}

func (hl HittableList) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	var hitAnything bool
	var outputRecord HitRecord
	closest := tMax
	for _, h := range hl {
		record, hit := h.Hit(r, tMin, closest)
		if hit {
			hitAnything = true
			closest = record.T
			outputRecord = record
		}
	}

	return outputRecord, hitAnything
}

func (hl HittableList) BoundingBox(t0, t1 float64) (AABB, bool) {
	return NewBVHNode(hl, int64(len(hl)), t0, t1).BoundingBox(t0, t1)
}

func (hl HittableList) PDFValue(o, v Vec3) float64 {
	return 0
}

func (hl HittableList) Random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
