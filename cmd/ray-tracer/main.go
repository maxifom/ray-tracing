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

	"github.com/cheggaaa/pb/v3"

	"ray-tracing/pkg/common"
	. "ray-tracing/pkg/pdf"
	. "ray-tracing/pkg/scene"
	"ray-tracing/pkg/utils"
	. "ray-tracing/pkg/vec3"
	"ray-tracing/scenes"
)

const MaxDepth = 50

func main() {
	rand.Seed(time.Now().UnixNano())

	width := 555
	height := 555
	count := width * height
	progressBar := pb.Full.Start(count)

	outputImage := image.NewRGBA(image.Rect(0, 0, width, height))
	numberOfSamples := 10
	background := Vec3{0, 0, 0}

	world, camera, lights := scenes.CornellBox(1)

	workerChan := make(chan WorkerInput)
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go worker(WorkerConfig{
			Width:           width,
			Height:          height,
			NumberOfSamples: numberOfSamples,
			Background:      background,
			World:           world,
			Camera:          camera,
			OutputImage:     outputImage,
			InputChan:       workerChan,
			Wg:              &wg,
			Lights:          lights,
		})
	}

	t := time.Now()
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			workerChan <- WorkerInput{
				X: i,
				Y: j,
			}

			progressBar.Increment()
		}
	}

	close(workerChan)
	wg.Wait()
	progressBar.Finish()

	seconds := time.Since(t).Seconds()
	log.Printf("Finished in %.2f seconds: Average speed: %.2f/s", seconds, float64(count)/seconds)

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

type WorkerInput struct {
	X, Y int
}

type WorkerConfig struct {
	Width, Height, NumberOfSamples int
	Background                     Vec3
	World                          common.Hittable
	Camera                         Camera
	OutputImage                    *image.RGBA
	InputChan                      chan WorkerInput
	Wg                             *sync.WaitGroup
	Lights                         common.Hittable
}

func worker(config WorkerConfig) {
	defer config.Wg.Done()
	for input := range config.InputChan {
		i := float64(input.X)
		j := float64(config.Height - input.Y)
		col := Vec3{0, 0, 0}
		for s := 0; s < config.NumberOfSamples; s++ {
			u := (i + rand.Float64()) / float64(config.Width)
			v := (j + rand.Float64()) / float64(config.Height)
			ray := config.Camera.Ray(u, v)
			col = col.Add(RayColor(ray, config.Background, config.World, config.Lights, MaxDepth))
		}

		col = col.DivN(float64(config.NumberOfSamples))
		col = Vec3{
			X: math.Sqrt(col.X),
			Y: math.Sqrt(col.Y),
			Z: math.Sqrt(col.Z),
		}
		ir := uint8(256 * utils.Clamp(col.X, 0, 0.999))
		ig := uint8(256 * utils.Clamp(col.Y, 0, 0.999))
		ib := uint8(256 * utils.Clamp(col.Z, 0, 0.999))
		config.OutputImage.SetRGBA(input.X, input.Y, color.RGBA{
			R: ir,
			G: ig,
			B: ib,
			A: 255,
		})
	}
}

// Цвет который получает луч
func RayColor(r Ray, background Vec3, world common.Hittable, lights common.Hittable, depth int64) Vec3 {
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
