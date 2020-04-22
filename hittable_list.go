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
