package main

type Material interface {
	Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, hasScattered bool)
}
