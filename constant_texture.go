package main

type ConstantTexture struct {
	Color Vec3
}

func (c ConstantTexture) Value(u, v float64, p Vec3) Vec3 {
	return c.Color
}
