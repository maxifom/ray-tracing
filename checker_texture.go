package main

import "math"

type CheckerTexture struct {
	Odd, Even Texture
}

func (c CheckerTexture) Value(u, v float64, p Vec3) Vec3 {
	sines := math.Sin(10*p.X) * math.Sin(10*p.Y) * math.Sin(10*p.Z)
	if sines < 0 {
		return c.Odd.Value(u, v, p)
	}

	return c.Even.Value(u, v, p)
}
