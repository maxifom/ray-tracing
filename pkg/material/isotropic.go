package material

import (
	"ray-tracing/pkg/common"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/vec3"
)

type Isotropic struct {
	Albedo common.Texture
}

func (i Isotropic) Scatter(r Ray, rec common.HitRecord) (scattered common.ScatterRecord, hasScattered bool) {
	scattered.Attenuation = i.Albedo.Value(rec.U, rec.V, rec.P)
	scattered.Ray = Ray{rec.P, RandomInUnitSphere(), r.Time}
	scattered.PDF = nil
	return scattered, true
}

func (i Isotropic) ScatteringPDF(r Ray, rec common.HitRecord, scatteredRay Ray) float64 {
	return 0
}

func (i Isotropic) Emitted(rIn Ray, u, v float64, rec common.HitRecord, p Vec3) Vec3 {
	return Vec3{}
}
