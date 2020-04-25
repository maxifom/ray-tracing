package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
)

// Цвет который получает луч
func RayColor(ray Ray, background Vec3, world Hittable, depth int64) Vec3 {
	if depth <= 0 {
		return Vec3{0, 0, 0}
	}

	record, hit := world.Hit(ray, 0.001, math.Inf(1))
	if !hit {
		return background
	}

	emitted := record.Material.Emitted(record.U, record.V, record.P)
	scattered, attenuation, hasScattered := record.Material.Scatter(ray, record)
	if !hasScattered {
		return emitted
	}

	return emitted.Add(attenuation.Mul(RayColor(scattered, background, world, depth-1)))
}

func RandomScene() Hittable {
	n := 500
	hl := make(HittableList, 0, n)
	hl = append(hl, Sphere{Vec3{0, -1000, 0}, 1000, Lambertian{CheckerTexture{ConstantTexture{Vec3{0.2, 0.3, 0.1}}, ConstantTexture{Vec3{0.9, 0.9, 0.9}}}}})
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			if center.Sub(Vec3{4, 0.2, 0}).Length() > 0.9 {
				if chooseMat < 0.8 {
					hl = append(hl, MovingSphere{center, center.Add(Vec3{0, 0.5 * rand.Float64(), 0}), 0.2, Lambertian{ConstantTexture{Vec3{rand.Float64() * rand.Float64(), rand.Float64() * rand.Float64(), rand.Float64() * rand.Float64()}}}, 0, 1})
				} else if chooseMat < 0.95 {
					hl = append(hl, Sphere{center, 0.2, Metal{Vec3{0.5 * (1 + rand.Float64()), 0.5 * (1 + rand.Float64()), 0.5 * (1 + rand.Float64())}}})
				} else {
					hl = append(hl, Sphere{center, 0.2, Dielectric{1.5}})
				}
			}
		}
	}

	hl = append(hl, Sphere{Vec3{0, 1, 0}, 1, Dielectric{1.5}})
	hl = append(hl, Sphere{Vec3{-4, 1, 0}, 1, Lambertian{ConstantTexture{Vec3{0.4, 0.2, 0.1}}}})
	hl = append(hl, Sphere{Vec3{4, 1, 0}, 1, Metal{Vec3{0.7, 0.6, 0.5}}})
	return hl
}

func TwoPerlinSpheres() Hittable {
	return NewList(
		Sphere{Vec3{0, -1000, 0}, 1000, Lambertian{NoiseTexture{NewPerlin(), 2}}},
		Sphere{Vec3{0, 1, 0}, 1, Lambertian{NoiseTextureFloat{NewPerlinFloat(), 2}}},
	)
}

func TestImageTexture() Hittable {
	f, err := os.Open("earth.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	imageData, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	return NewList(
		Sphere{Vec3{0, 0, 0}, 2, Lambertian{ImageTexture{
			Nx:   imageData.Bounds().Size().X,
			Ny:   imageData.Bounds().Size().Y,
			Data: imageData,
		}}},
	)
}

func DiffuseLightScene() Hittable {
	return NewList(
		Sphere{Vec3{0, -1000, 0}, 1000, Lambertian{NoiseTexture{Noise: NewPerlin(), Scale: 4}}},
		Sphere{Vec3{0, 2, 0}, 2, Lambertian{NoiseTexture{Noise: NewPerlin(), Scale: 4}}},
		Sphere{Vec3{0, 7, 0}, 2, DiffuseLight{ConstantTexture{Vec3{4, 4, 4}}}},
		XYRect{3, 5, 1, 3, -2, DiffuseLight{ConstantTexture{Vec3{4, 4, 4}}}},
	)
}

func CornellBox() Hittable {
	red := Lambertian{ConstantTexture{Vec3{0.65, 0.05, 0.05}}}
	// yellow := Lambertian{ConstantTexture{Vec3{0.9, 0.9, 0.05}}}
	white := Lambertian{ConstantTexture{Vec3{0.73, 0.73, 0.73}}}
	// blue := Lambertian{ConstantTexture{Vec3{0.05, 0.05, 0.9}}}
	green := Lambertian{ConstantTexture{Vec3{0.12, 0.45, 0.15}}}
	light := DiffuseLight{ConstantTexture{Vec3{15, 15, 15}}}

	return NewList(
		FlipFace{YZRect{0, 555, 0, 555, 555, green}},
		YZRect{0, 555, 0, 555, 0, red},
		XZRect{213, 343, 227, 332, 554, light},
		XZRect{0, 555, 0, 555, 0, white},
		FlipFace{XZRect{0, 555, 0, 555, 555, white}},
		FlipFace{XYRect{0, 555, 0, 555, 555, white}},

		NewBox(Vec3{130, 0, 65}, Vec3{295, 165, 230}, white),
		NewBox(Vec3{265, 0, 295}, Vec3{430, 330, 460}, white),
	)
}

func main() {
	file, err := os.OpenFile("output.ppm", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	width := 16 * 100
	height := 9 * 100
	numberOfTimes := 10

	background := Vec3{0, 0, 0}

	fmt.Fprintf(file, "P3\n%d %d\n255\n", width, height)

	// world := RandomScene()
	// world := TwoPerlinSpheres()
	// world := TestImageTexture()
	// world := DiffuseLightScene()
	world := CornellBox()
	lookFrom := Vec3{278, 278, -800}
	lookAt := Vec3{278, 278, 0}
	focusDist := 10.0
	aperture := 0.0
	vUp := Vec3{0, 1, 0}
	cam := NewCamera(
		lookFrom,
		lookAt,
		vUp,
		40,
		float64(width)/float64(height),
		aperture,
		focusDist,
		0,
		1,
	)
	const MaxDepth = 50

	for j := height - 1; j >= 0; j-- {
		log.Println(j)
		for i := 0; i < width; i++ {
			col := Vec3{0, 0, 0}
			for s := 0; s < numberOfTimes; s++ {
				u := (float64(i) + rand.Float64()) / float64(width)
				v := (float64(j) + rand.Float64()) / float64(height)
				ray := cam.Ray(u, v)
				col = col.Add(RayColor(ray, background, world, MaxDepth))
			}

			col = col.DivN(float64(numberOfTimes))
			col = Vec3{
				X: math.Sqrt(col.X),
				Y: math.Sqrt(col.Y),
				Z: math.Sqrt(col.Z),
			}
			ir := int64(256 * Clamp(col.X, 0, 0.999))
			ig := int64(256 * Clamp(col.Y, 0, 0.999))
			ib := int64(256 * Clamp(col.Z, 0, 0.999))

			fmt.Fprintf(file, "%d %d %d\n", ir, ig, ib)
		}
	}

	display := exec.Command("display", "output.ppm")
	display.Stdin = os.Stdin
	display.Stdout = os.Stdout
	err = display.Run()
	if err != nil {
		log.Panic(err)
	}
}
