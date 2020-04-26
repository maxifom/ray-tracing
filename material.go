package main

type ScatterRecord struct {
	Ray         Ray
	IsSpecular  bool
	Attenuation Vec3
	PDF         PDF
}

type Material interface {
	Scatter(r Ray, rec HitRecord) (scattered ScatterRecord, hasScattered bool)
	ScatteringPDF(r Ray, rec HitRecord, scatteredRay Ray) float64
	Emitted(rIn Ray, u, v float64, rec HitRecord, p Vec3) Vec3
}
