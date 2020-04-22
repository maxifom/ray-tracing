package main

type Lambertian struct {
	Albedo Vec3
}

func (l Lambertian) Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, hasScattered bool) {
	target := rec.P.Add(rec.Normal).Add(RandomInUnitSphere())
	return Ray{rec.P, target.Sub(rec.P), 0}, l.Albedo, true
}
