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

func CornellBoxSmoke() Hittable {
	red := Lambertian{ConstantTexture{Vec3{0.65, 0.05, 0.05}}}
	white := Lambertian{ConstantTexture{Vec3{0.73, 0.73, 0.73}}}
	green := Lambertian{ConstantTexture{Vec3{0.12, 0.45, 0.15}}}
	light := DiffuseLight{ConstantTexture{Vec3{7, 7, 7}}}
	var box1, box2 Hittable
	box1 = NewBox(Vec3{0, 0, 0}, Vec3{165, 330, 165}, white)
	box1 = NewRotateY(box1, 15)
	box1 = Translate{box1, Vec3{265, 0, 295}}
	box1 = ConstantMedium{box1, Isotropic{ConstantTexture{Vec3{0, 0, 0}}}, -1 / 0.01}

	box2 = NewBox(Vec3{0, 0, 0}, Vec3{165, 165, 165}, white)
	box2 = NewRotateY(box2, -18)
	box2 = Translate{box2, Vec3{130, 0, 65}}
	box2 = ConstantMedium{box2, Isotropic{ConstantTexture{Vec3{1, 1, 1}}}, -1 / 0.01}

	return NewList(
		FlipFace{YZRect{0, 555, 0, 555, 555, green}},
		YZRect{0, 555, 0, 555, 0, red},
		XZRect{113, 443, 127, 432, 554, light},
		XZRect{0, 555, 0, 555, 0, white},
		FlipFace{XZRect{0, 555, 0, 555, 555, white}},
		FlipFace{XYRect{0, 555, 0, 555, 555, white}},
		box1, box2,
	)

}

func FinalScene() Hittable {
	var boxes HittableList
	ground := Lambertian{ConstantTexture{Vec3{0.48, 0.83, 0.53}}}
	boxesPerSide := 20
	for i := 0; i < boxesPerSide; i++ {
		for j := 0; j < boxesPerSide; j++ {
			i1 := float64(i)
			j1 := float64(j)
			w := 100.0
			x0 := -1000.0 + i1*w
			z0 := -1000.0 + j1*w
			y0 := 0.0
			x1 := x0 + w
			y1 := 1 + 100*rand.Float64() // 1-> 101
			z1 := z0 + w

			boxes = append(boxes, NewBox(Vec3{x0, y0, z0}, Vec3{x1, y1, z1}, ground))
		}
	}

	var objects HittableList

	objects = append(objects, NewBVHNode(boxes, int64(len(boxes)), 0, 1))
	light := DiffuseLight{ConstantTexture{Vec3{15, 15, 15}}}
	objects = append(objects, XZRect{123, 423, 147, 412, 554, light})
	center1 := Vec3{400, 400, 200}
	center2 := center1.Add(Vec3{30, 0, 0})

	movingSphereMaterial := Lambertian{ConstantTexture{Vec3{0.7, 0.3, 0.1}}}

	objects = append(objects, MovingSphere{center1, center2, 50, movingSphereMaterial, 0, 1})
	objects = append(objects, Sphere{Vec3{260, 150, 45}, 50, Dielectric{1.5}})
	objects = append(objects, Sphere{Vec3{0, 150, 145}, 50, Metal{Vec3{0.8, 0.8, 0.9}}})

	boundary := Sphere{Vec3{360, 150, 145}, 70, Dielectric{1.5}}
	objects = append(objects, boundary)

	objects = append(objects, ConstantMedium{boundary, Isotropic{ConstantTexture{Vec3{0.2, 0.4, 0.9}}}, -1 / 0.2})
	objects = append(objects, Sphere{Vec3{0, 0, 0}, 5000, Dielectric{1.5}})
	objects = append(objects, ConstantMedium{boundary, Isotropic{ConstantTexture{Vec3{1, 1, 1}}}, -1 / 0.0001})

	f, err := os.Open("earth.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	imageData, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	emat := Lambertian{ImageTexture{
		Nx:   imageData.Bounds().Size().X,
		Ny:   imageData.Bounds().Size().Y,
		Data: imageData,
	}}

	objects = append(objects, Sphere{Vec3{400, 200, 400}, 100, emat})
	objects = append(objects, Sphere{Vec3{220, 280, 300}, 80, Lambertian{NoiseTexture{NewPerlin(), 0.1}}})

	var boxes2 HittableList

	white := Lambertian{ConstantTexture{Vec3{0.73, 0.73, 0.73}}}
	ns := 1000
	for j := 0; j < ns; j++ {
		boxes2 = append(boxes2, Sphere{Vec3{165 * rand.Float64(), 165 * rand.Float64(), 165 * rand.Float64()}, 10, white})
	}

	objects = append(objects, Translate{NewRotateY(NewBVHNode(boxes2, int64(len(boxes2)), 0, 1), 15), Vec3{-100, 270, 395}})

	return objects
}

func CornellBox() Hittable {
	red := Lambertian{ConstantTexture{Vec3{0.65, 0.05, 0.05}}}
	white := Lambertian{ConstantTexture{Vec3{0.73, 0.73, 0.73}}}
	green := Lambertian{ConstantTexture{Vec3{0.12, 0.45, 0.15}}}
	light := DiffuseLight{ConstantTexture{Vec3{15, 15, 15}}}

	var box1, box2 Hittable
	box1 = NewBox(Vec3{0, 0, 0}, Vec3{165, 330, 165}, white)
	box1 = NewRotateY(box1, 15)
	box1 = Translate{box1, Vec3{265, 0, 295}}

	box2 = NewBox(Vec3{0, 0, 0}, Vec3{165, 165, 165}, white)
	box2 = NewRotateY(box2, -18)
	box2 = Translate{box2, Vec3{130, 0, 65}}

	return NewList(
		FlipFace{YZRect{0, 555, 0, 555, 555, green}},
		YZRect{0, 555, 0, 555, 0, red},
		XZRect{213, 343, 227, 332, 554, light},
		XZRect{0, 555, 0, 555, 0, white},
		FlipFace{XZRect{0, 555, 0, 555, 555, white}},
		FlipFace{XYRect{0, 555, 0, 555, 555, white}},
		box1, box2,
	)
}

func main() {
	file, err := os.OpenFile("output.ppm", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	width := 100
	height := 50
	numberOfTimes := 20

	background := Vec3{0, 0, 0}

	fmt.Fprintf(file, "P3\n%d %d\n255\n", width, height)

	// world := RandomScene()
	// world := TwoPerlinSpheres()
	// world := TestImageTexture()
	// world := DiffuseLightScene()
	// world := CornellBox()
	// world := CornellBoxSmoke()
	world := FinalScene()
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
