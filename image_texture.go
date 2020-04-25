package main

import (
	"image"
	"log"
)

type ImageTexture struct {
	Nx, Ny int
	Data   image.Image
}

func (it ImageTexture) Value(u, v float64, p Vec3) Vec3 {
	i := int(u * float64(it.Nx))
	j := int((1-v)*float64(it.Ny) - 0.001)
	if i < 0 {
		i = 0
	}
	if j < 0 {
		j = 0
	}

	if i > it.Nx-1 {
		i = it.Nx - 1
	}
	if j > it.Ny-1 {
		j = it.Ny - 1
	}
	log.Print(i, j)
	r, g, b, _ := it.Data.At(3*i, 3*j).RGBA()
	vec3 := Vec3{float64(r / 255), float64(g / 255), float64(b / 255)}
	return vec3
}
