package main

type Hittable interface {
	Hit(r Ray, tMin, tMax float64) (HitRecord, bool)
}
