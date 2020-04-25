package main

type NoiseTextureFloat struct {
	Noise PerlinFloat
	Scale float64
}

func (n NoiseTextureFloat) Value(u, v float64, p Vec3) Vec3 {
	return Vec3{1, 1, 1}.MulN(n.Noise.Noise(p.MulN(n.Scale)))
}
