package main

import (
	"math"
	"math/rand"
)

type Dielectric struct {
	RefIdx float64
}

func (d Dielectric) Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, hasScattered bool) {
	// Attenuation is always 1 — the glass surface absorbs nothing
	attenuation = Vec3{1, 1, 1}
	var niOverNt float64
	if rec.FrontFace {
		niOverNt = 1.0 / d.RefIdx
	} else {
		niOverNt = d.RefIdx
	}

	unitDirection := r.Direction.UnitVector()

	cosTheta := math.Min(Dot(unitDirection.Negative(), rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	if niOverNt*sinTheta > 1.0 {
		reflected := Reflect(unitDirection, rec.Normal)
		scattered = Ray{rec.P, reflected, 0}
		return scattered, attenuation, true
	}

	reflectProb := Schlick(cosTheta, niOverNt)
	if rand.Float64() < reflectProb {
		reflected := Reflect(unitDirection, rec.Normal)
		scattered = Ray{rec.P, reflected, 0}
		return scattered, attenuation, true
	}

	refracted := Refract(unitDirection, rec.Normal, niOverNt)
	scattered = Ray{rec.P, refracted, 0}
	return scattered, attenuation, true
}

func (d Dielectric) Emitted(u, v float64, p Vec3) Vec3 {
	return Vec3{}
}
