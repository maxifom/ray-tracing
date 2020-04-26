package main

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"
)

// Цвет который получает луч
func RayColor(r Ray, background Vec3, world Hittable, lights Hittable, depth int64) Vec3 {
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
		return srec.Attenuation.
			Mul(RayColor(srec.Ray, background, world, lights, depth-1))
	}

	light := HittablePDF{H: lights, O: rec.P}
	p := MixturePDF{light, srec.PDF}
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

func CornellBoxNew(aspect float64) (Hittable, Camera) {
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

type Input struct {
	X, Y int
}

func worker(width, height int, numberOfSamples int, background Vec3, world Hittable, camera Camera, image *image.RGBA, inputChan chan Input, wg *sync.WaitGroup, lights Hittable) {
	defer wg.Done()
	for input := range inputChan {
		i := float64(input.X)
		j := float64(height - input.Y)
		col := Vec3{0, 0, 0}
		for s := 0; s < numberOfSamples; s++ {
			u := (i + rand.Float64()) / float64(width)
			v := (j + rand.Float64()) / float64(height)
			ray := camera.Ray(u, v)
			col = col.Add(RayColor(ray, background, world, lights, MaxDepth))
		}

		col = col.DivN(float64(numberOfSamples))
		col = Vec3{
			X: math.Sqrt(col.X),
			Y: math.Sqrt(col.Y),
			Z: math.Sqrt(col.Z),
		}
		ir := uint8(256 * Clamp(col.X, 0, 0.999))
		ig := uint8(256 * Clamp(col.Y, 0, 0.999))
		ib := uint8(256 * Clamp(col.Z, 0, 0.999))
		image.SetRGBA(input.X, input.Y, color.RGBA{
			R: ir,
			G: ig,
			B: ib,
			A: 255,
		})
	}
}

const MaxDepth = 50

func main() {
	rand.Seed(time.Now().UnixNano())
	width := 555
	height := 555
	outputImage := image.NewRGBA(image.Rect(0, 0, width, height))
	numberOfSamples := 1000
	background := Vec3{0, 0, 0}

	world, cam := CornellBoxNew(1)

	lights := NewList(
		XZRect{213, 343, 227, 332, 554, nil},
		Sphere{Vec3{190, 90, 190}, 90, nil},
	)
	workerChan := make(chan Input)
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go worker(width, height, numberOfSamples, background, world, cam, outputImage, workerChan, &wg, lights)
	}

	for j := 0; j < height; j++ {
		log.Println(height - j)
		for i := 0; i < width; i++ {
			workerChan <- Input{
				X: i,
				Y: j,
			}
		}
	}

	close(workerChan)
	wg.Wait()
	f, err := os.OpenFile("output.png", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = png.Encode(f, outputImage)
	if err != nil {
		log.Fatal(err)
	}

	display := exec.Command("display", "output.png")
	display.Stdin = os.Stdin
	display.Stdout = os.Stdout
	err = display.Run()
	if err != nil {
		log.Panic(err)
	}
}
