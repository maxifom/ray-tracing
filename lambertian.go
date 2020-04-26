package main

import "math"

type Lambertian struct {
	Albedo Texture
}

func (l Lambertian) Scatter(r Ray, rec HitRecord) (scattered ScatterRecord, hasScattered bool) {
	scattered.IsSpecular = false
	scattered.Attenuation = l.Albedo.Value(rec.U, rec.V, rec.P)
	scattered.PDF = CosinePDF{NewONB(rec.Normal)}
	return scattered, true
	// onb := NewONB(rec.Normal)
	// direction := onb.Local(RandomCosineDirection())
	// scattered = Ray{rec.P, direction.UnitVector(), r.Time}
	// attenuation = l.Albedo.Value(rec.U, rec.V, rec.P)
	// pdf = Dot(onb.W, scattered.Direction) / math.Pi

	// return scattered, true
}

func (l Lambertian) Emitted(Ray, float64, float64, HitRecord, Vec3) Vec3 {
	return Vec3{}
}

func (l Lambertian) ScatteringPDF(r Ray, rec HitRecord, scatteredRay Ray) (result float64) {
	cosine := Dot(rec.Normal, scatteredRay.Direction.UnitVector())
	if cosine < 0 {
		return 0
	}

	return cosine / math.Pi
}
