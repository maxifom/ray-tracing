package main

type DiffuseLight struct {
	Emit Texture
}

func (d DiffuseLight) Emitted(rIn Ray, u, v float64, rec HitRecord, p Vec3) Vec3 {
	// We also need to flip the light so its normals point in the -y direction
	if rec.FrontFace {
		return d.Emit.Value(u, v, p)
	}

	return Vec3{0, 0, 0}
}

func (d DiffuseLight) Scatter(Ray, HitRecord) (ScatterRecord, bool) {
	return ScatterRecord{}, false
}

func (d DiffuseLight) ScatteringPDF(Ray, HitRecord, Ray) float64 {
	return 0
}
