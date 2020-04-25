package main

type Metal struct {
	Albedo Vec3
}

func (m Metal) Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, pdf float64, hasScattered bool) {
	reflected := Reflect(r.Direction.UnitVector(), rec.Normal)
	scattered = Ray{rec.P, reflected, 0}

	return scattered, m.Albedo, pdf, Dot(scattered.Direction, rec.Normal) > 0
}

func (m Metal) ScatteringPDF(r Ray, rec HitRecord, scatteredRay Ray) float64 {
	return 0
}

func (m Metal) Emitted(u, v float64, rec HitRecord, p Vec3) Vec3 {
	return Vec3{}
}
