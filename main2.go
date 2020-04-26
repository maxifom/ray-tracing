package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
)

func main() {
	aspect := 1.0 / 1.0
	width := 600
	height := int64(float64(width) / aspect)
	samplesPerPixel := 10
	maxDepth := 50
	file, err := os.OpenFile("1.ppm", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "P3\n%d %d\n255\n", width, height)
	background := Vec3{0, 0, 0}
	world, cam := CornellBox(aspect)

	lights := NewList(
		XZRect{213, 343, 227, 332, 554, nil},
		Sphere{Vec3{190, 90, 190}, 90, nil},
	)

	for j := height - 1; j >= 0; j-- {
		log.Println(j)
		for i := 0; i < width; i++ {
			pixelColor := Vec3{0, 0, 0}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + RandomDouble(0, 1)) / float64(width-1)
				v := (float64(j) + RandomDouble(0, 1)) / float64(height-1)
				r := cam.Ray(u, v)
				pixelColor = pixelColor.Add(RayColor(r, background, world, lights, maxDepth))
			}
			writeColor(file, samplesPerPixel, pixelColor)
		}
	}

	display := exec.Command("display", "1.ppm")
	display.Stdin = os.Stdin
	display.Stdout = os.Stdout
	err = display.Run()
	if err != nil {
		log.Panic(err)
	}

}

func RayColor(r Ray, background Vec3, world Hittable, lights HittableList, depth int) Vec3 {
	if depth <= 0 {
		return Vec3{0, 0, 0}
	}

	rec, isHit := world.Hit(r, 0.001, math.Inf(1))
	if !isHit {
		return background
	}

	emitted := rec.Material.Emitted(r, rec.U, rec.V, rec, rec.P)
	srec, isScattered := rec.Material.Scatter(r, rec)
	if !isScattered {
		return emitted
	}

	if srec.IsSpecular {
		return srec.Attenuation.Mul(RayColor(srec.Ray, background, world, lights, depth-1))
	}

	lightPtr := HittablePDF{lights, rec.P}
	p := MixturePDF{lightPtr, srec.PDF}

	scattered := Ray{rec.P, p.Generate(), r.Time}

	pdfVal := p.Value(scattered.Direction)

	return emitted.
		Add(
			srec.Attenuation.
				MulN(rec.Material.ScatteringPDF(r, rec, scattered)).
				Mul(RayColor(scattered, background, world, lights, depth-1)).
				DivN(pdfVal),
		)

}

func writeColor(file *os.File, n int, color Vec3) {
	color = color.DivN(float64(n))
	color = Vec3{
		X: math.Sqrt(color.X),
		Y: math.Sqrt(color.Y),
		Z: math.Sqrt(color.Z),
	}
	ir := uint8(256 * Clamp(color.X, 0, 0.999))
	ig := uint8(256 * Clamp(color.Y, 0, 0.999))
	ib := uint8(256 * Clamp(color.Z, 0, 0.999))

	fmt.Fprintf(file, "%d %d %d\n", ir, ig, ib)
}

func CornellBox(aspect float64) (Hittable, Camera) {
	var world HittableList

	red := Lambertian{ConstantTexture{Vec3{.65, .05, .05}}}
	white := Lambertian{ConstantTexture{Vec3{.73, .73, .73}}}
	green := Lambertian{ConstantTexture{Vec3{.12, .45, .15}}}
	light := DiffuseLight{ConstantTexture{Vec3{15, 15, 15}}}

	world = append(world, FlipFace{YZRect{0, 555, 0, 555, 555, green}})
	world = append(world, YZRect{0, 555, 0, 555, 0, red})
	world = append(world, FlipFace{XZRect{213, 343, 227, 332, 554, light}})
	world = append(world, FlipFace{XZRect{0, 555, 0, 555, 555, white}})
	world = append(world, XZRect{0, 555, 0, 555, 0, white})
	world = append(world, FlipFace{XYRect{0, 555, 0, 555, 555, white}})

	var box1 Hittable

	box1 = NewBox(Vec3{0, 0, 0}, Vec3{165, 330, 165}, white)
	box1 = NewRotateY(box1, 15)
	box1 = Translate{box1, Vec3{265, 0, 295}}
	world = append(world, box1)

	glass := Dielectric{1.5}
	world = append(world, Sphere{Vec3{190, 90, 190}, 90, glass})

	lookFrom := Vec3{278, 278, -800}
	lookAt := Vec3{278, 278, 0}
	up := Vec3{0, 1, 0}
	distToFocus := 10.0
	aperture := 0.0
	vFov := 40.0
	t0 := 0.0
	t1 := 1.1

	cam := NewCamera(lookFrom, lookAt, up, vFov, aspect, aperture, distToFocus, t0, t1)
	return world, cam
}
