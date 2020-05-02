package material

import (
	. "ray-tracing/pkg/common"
	. "ray-tracing/pkg/hittable"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/vec3"
)

type Metal struct {
	Albedo Vec3
	Fuzz   float64
}

func (m Metal) Scatter(r Ray, rec HitRecord) (scattered ScatterRecord, hasScattered bool) {
	reflected := Reflect(r.Direction.UnitVector(), rec.Normal)
	scattered.Ray = Ray{rec.P, reflected.Add(RandomInUnitSphere().MulN(m.Fuzz)), r.Time}
	scattered.Attenuation = m.Albedo
	scattered.IsSpecular = true
	scattered.PDF = nil
	return scattered, true
}

func (m Metal) ScatteringPDF(r Ray, rec HitRecord, scatteredRay Ray) float64 {
	return 0
}

func (m Metal) Emitted(rIn Ray, u, v float64, rec HitRecord, p Vec3) Vec3 {
	return Vec3{}
}
