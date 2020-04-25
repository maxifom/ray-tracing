package main

type Material interface {
	Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, hasScattered bool)
	Emitted(u, v float64, p Vec3) Vec3
}
