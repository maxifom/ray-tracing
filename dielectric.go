package main

import "math/rand"

type Dielectric struct {
	RefIdx float64
}

func (d Dielectric) Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, hasScattered bool) {
	var outwardNormal Vec3
	reflected := Reflect(r.Direction, rec.Normal)
	attenuation = Vec3{1, 1, 1}
	var niOverNt float64
	var cosine float64
	if Dot(r.Direction, rec.Normal) > 0 {
		outwardNormal = rec.Normal.Negative()
		niOverNt = d.RefIdx
		cosine = d.RefIdx * Dot(r.Direction, rec.Normal) / r.Direction.Length()
	} else {
		outwardNormal = rec.Normal
		niOverNt = 1 / d.RefIdx
		cosine = -Dot(r.Direction, rec.Normal) / r.Direction.Length()
	}

	var reflectProb float64
	refracted, hasRefracted := Refract(r.Direction, outwardNormal, niOverNt)
	if hasRefracted {
		reflectProb = Schlick(cosine, d.RefIdx)
		scattered = Ray{rec.P, refracted}
	} else {
		scattered = Ray{rec.P, reflected}
		reflectProb = 1
	}

	if rand.Float64() < reflectProb {
		scattered = Ray{rec.P, reflected}
	} else {
		scattered = Ray{rec.P, refracted}
	}

	return scattered, attenuation, true
}
