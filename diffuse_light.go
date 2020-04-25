package main

type DiffuseLight struct {
	Emit Texture
}

func (d DiffuseLight) Scatter(r Ray, rec HitRecord) (scattered Ray, attenuation Vec3, hasScattered bool) {
	return Ray{}, Vec3{}, false
}

func (d DiffuseLight) Emitted(u, v float64, p Vec3) Vec3 {
	return d.Emit.Value(u, v, p)
}
