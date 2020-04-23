package main

type NoiseTexture struct {
	Noise Perlin
}

func (n NoiseTexture) Value(u, v float64, p Vec3) Vec3 {
	return Vec3{1, 1, 1}.MulN(n.Noise.Noise(p))
}
