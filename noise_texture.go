package main

type NoiseTexture struct {
	Noise Perlin
	Scale float64
}

func (n NoiseTexture) Value(u, v float64, p Vec3) Vec3 {
	// TODO: not working
	return Vec3{1, 1, 1}.MulN(n.Noise.Noise(p.MulN(n.Scale)))
	// return Vec3{1, 1, 1}.MulN(0.5 * (1 + n.Noise.Turb(p.MulN(n.Scale), 7)))
	// return Vec3{1, 1, 1}.MulN(n.Noise.Turb(p.MulN(n.Scale), 7)) // Черные линии стр 27 верхняя картинка
	// return Vec3{1, 1, 1}.MulN(0.5*(1+math.Sin(n.Scale*p.Z + 10 * n.Noise.Turb(p, 7)))) // 27 стр нижняя картинка
}
