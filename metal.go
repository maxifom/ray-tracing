package main

type Metal struct {
	Albedo Vec3
}

func (m Metal) Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, hasScattered bool) {
	reflected := Reflect(r.Direction.UnitVector(), rec.Normal)
	scattered = Ray{rec.P, reflected}

	return scattered, m.Albedo, Dot(scattered.Direction, rec.Normal) > 0
}
