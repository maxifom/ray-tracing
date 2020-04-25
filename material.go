package main

type Material interface {
	Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, pdf float64, hasScattered bool)
	ScatteringPDF(r Ray, rec HitRecord, scatteredRay Ray) float64
	Emitted(u, v float64, rec HitRecord, p Vec3) Vec3
}
