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
	if len(hl) == 0 {
		return AABB{}, false
	}
	var outputBox AABB

	box, isBounded := hl[0].BoundingBox(t0, t1)
	if !isBounded {
		return AABB{}, false
	}

	outputBox = box

	for _, h := range hl {
		box, isBounded = h.BoundingBox(t0, t1)
		if !isBounded {
			return AABB{}, false
		}

		outputBox = SurroundingBox(outputBox, box)
	}

	return outputBox, true
}

func (hl HittableList) PDFValue(o, v Vec3) float64 {
	weight := 1.0 / float64(len(hl))
	sum := 0.0
	for _, h := range hl {
		sum += weight * h.PDFValue(o, v)
	}

	return sum
}

func (hl HittableList) Random(o Vec3) Vec3 {
	if len(hl) == 0 {
		return Vec3{1, 0, 0}
	}
	return hl[int(RandomDouble(0.0, float64(len(hl)-1)))].Random(o)
}
