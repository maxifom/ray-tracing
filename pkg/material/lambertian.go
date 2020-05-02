package material

import (
	"math"

	. "ray-tracing/pkg/common"
	. "ray-tracing/pkg/hittable"
	. "ray-tracing/pkg/pdf"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/vec3"
)

type Lambertian struct {
	Albedo Texture
}

func (l Lambertian) Scatter(r Ray, rec HitRecord) (scattered ScatterRecord, hasScattered bool) {
	scattered.IsSpecular = false
	scattered.Attenuation = l.Albedo.Value(rec.U, rec.V, rec.P)
	scattered.PDF = CosinePDF{NewONB(rec.Normal)}

	return scattered, true
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
