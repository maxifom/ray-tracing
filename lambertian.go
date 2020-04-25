package main

type Lambertian struct {
	Albedo Texture
}

func (l Lambertian) Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, hasScattered bool) {
	scatterDirection := rec.Normal.Add(RandomUnitVector())
	scattered = Ray{rec.P, scatterDirection, r.Time}
	attenuation = l.Albedo.Value(rec.U, rec.V, rec.P)

	return scattered, attenuation, true
}
