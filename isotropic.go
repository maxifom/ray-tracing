package main

type Isotropic struct {
	Albedo Texture
}

func (i Isotropic) Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, pdf float64, hasScattered bool) {
	return Ray{
		Origin:    rec.P,
		Direction: RandomInUnitSphere(),
		Time:      r.Time,
	}, i.Albedo.Value(rec.U, rec.V, rec.P), pdf, true
}

func (i Isotropic) ScatteringPDF(r Ray, rec HitRecord, scatteredRay Ray) float64 {
	return 0
}

func (i Isotropic) Emitted(u, v float64, rec HitRecord, p Vec3) Vec3 {
	return Vec3{}
}
