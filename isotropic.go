package main

type Isotropic struct {
	Albedo Texture
}

func (i Isotropic) Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, hasScattered bool) {
	return Ray{
		Origin:    rec.P,
		Direction: RandomInUnitSphere(),
		Time:      r.Time,
	}, i.Albedo.Value(rec.U, rec.V, rec.P), true
}

func (i Isotropic) Emitted(u, v float64, p Vec3) Vec3 {
	return Vec3{}
}
