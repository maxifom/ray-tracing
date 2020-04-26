package main

type Isotropic struct {
	Albedo Texture
}

func (i Isotropic) Scatter(r Ray, rec HitRecord) (scattered ScatterRecord, hasScattered bool) {
	scattered.Attenuation = i.Albedo.Value(rec.U, rec.V, rec.P)
	scattered.Ray = Ray{rec.P, RandomInUnitSphere(), r.Time}
	scattered.PDF = DefaultPDF{}
	return scattered, true
}

func (i Isotropic) ScatteringPDF(r Ray, rec HitRecord, scatteredRay Ray) float64 {
	return 0
}

func (i Isotropic) Emitted(rIn Ray, u, v float64, rec HitRecord, p Vec3) Vec3 {
	return Vec3{}
}
