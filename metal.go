package main

type Metal struct {
	Albedo Vec3
}

func (m Metal) Scatter(r Ray, rec HitRecord) (scattered ScatterRecord, hasScattered bool) {
	reflected := Reflect(r.Direction.UnitVector(), rec.Normal)
	scattered.Ray = Ray{rec.P, reflected.Add(RandomInUnitSphere()), 0.0}
	scattered.Attenuation = m.Albedo
	scattered.IsSpecular = true
	scattered.PDF = DefaultPDF{}
	return scattered, true
}

func (m Metal) ScatteringPDF(r Ray, rec HitRecord, scatteredRay Ray) float64 {
	return 0
}

func (m Metal) Emitted(rIn Ray, u, v float64, rec HitRecord, p Vec3) Vec3 {
	return Vec3{}
}
