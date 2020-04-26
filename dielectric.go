package main

import (
	"math"
	"math/rand"
)

type Dielectric struct {
	RefIdx float64
}

func (d Dielectric) Scatter(r Ray, rec HitRecord) (scattered ScatterRecord, hasScattered bool) {
	scattered.IsSpecular = true
	scattered.PDF = DefaultPDF{}

	// Attenuation is always 1 — the glass surface absorbs nothing
	scattered.Attenuation = Vec3{1, 1, 1}
	var etaiOverEtat float64
	if rec.FrontFace {
		etaiOverEtat = 1.0 / d.RefIdx
	} else {
		etaiOverEtat = d.RefIdx
	}

	unitDirection := r.Direction.UnitVector()
	cosTheta := math.Min(Dot(unitDirection.Negative(), rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	if etaiOverEtat*sinTheta > 1.0 {
		reflected := Reflect(unitDirection, rec.Normal)
		scattered.Ray = Ray{rec.P, reflected, r.Time}
		return scattered, true
	}

	reflectProb := Schlick(cosTheta, etaiOverEtat)
	if rand.Float64() < reflectProb {
		reflected := Reflect(unitDirection, rec.Normal)
		scattered.Ray = Ray{rec.P, reflected, r.Time}
		return scattered, true
	}

	refracted := Refract(unitDirection, rec.Normal, etaiOverEtat)
	scattered.Ray = Ray{rec.P, refracted, r.Time}

	return scattered, true
}

func (d Dielectric) ScatteringPDF(r Ray, rec HitRecord, scatteredRay Ray) float64 {
	return 0
}

func (d Dielectric) Emitted(Ray, float64, float64, HitRecord, Vec3) Vec3 {
	return Vec3{}
}
