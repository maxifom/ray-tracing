package texture

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	. "ray-tracing/pkg/vec3"
)

type ImageTexture struct {
	Nx, Ny int
	Data   image.Image
}

func NewImageTexture(filename string) ImageTexture {
	f, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}

	defer f.Close()

	imageData, _, err := image.Decode(f)
	if err != nil {
		log.Panic(err)
	}

	return ImageTexture{
		Nx:   imageData.Bounds().Size().X,
		Ny:   imageData.Bounds().Size().Y,
		Data: imageData,
	}

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

	r, g, b, a := it.Data.At(i, j).RGBA()
	vec3 := Vec3{float64(r) / float64(a), float64(g) / float64(a), float64(b) / float64(a)}
	return vec3
}
