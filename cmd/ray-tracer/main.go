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
	"runtime"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/jessevdk/go-flags"

	. "ray-tracing/pkg/common"
	. "ray-tracing/pkg/pdf"
	. "ray-tracing/pkg/scene"
	"ray-tracing/pkg/utils"
	. "ray-tracing/pkg/vec3"
	"ray-tracing/scenes"
)

const MaxDepth = 50

func main() {
	var opts struct {
		Width             int    `long:"width" default:"555"`
		Height            int    `long:"height" default:"555"`
		NumberOfSamples   int    `long:"number_of_samples" default:"10"`
		OutputFileName    string `long:"output_file_name" default:"output.png"`
		ShowAfterComplete int    `long:"show_after_complete" default:"1"`
		NumberOfWorkers   int    `long:"number_of_workers" default:"-1"`
		Scene             string `long:"scene" default:"cornell_box"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalf("failed to parse flags: %s", err)
	}

	numberOfWorkers := opts.NumberOfWorkers
	if numberOfWorkers == -1 {
		numberOfWorkers = runtime.NumCPU()
	}

	rand.Seed(time.Now().UnixNano())

	width := opts.Width
	height := opts.Height
	numberOfSamples := opts.NumberOfSamples

	background := Vec3{0, 0, 0}

	var scene scenes.Scene

	switch opts.Scene {
	case "cornell_box":
		scene = scenes.CornellBox(width, height)
	default:
		scene = scenes.CornellBox(width, height)
	}
	outputImage := image.NewRGBA(image.Rect(0, 0, scene.Width, scene.Height))
	count := scene.Width * scene.Height
	progressBar := pb.Full.Start(count)

	workerChan := make(chan WorkerInput)
	var wg sync.WaitGroup
	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go worker(WorkerConfig{
			NumberOfSamples: numberOfSamples,
			Background:      background,
			Scene:           scene,
			OutputImage:     outputImage,
			InputChan:       workerChan,
			Wg:              &wg,
		})
	}

	t := time.Now()
	for j := 0; j < scene.Height; j++ {
		for i := 0; i < scene.Width; i++ {
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

	f, err := os.OpenFile(opts.OutputFileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer f.Close()

	err = png.Encode(f, outputImage)
	if err != nil {
		log.Fatalf("failed to encode image: %s", err)
	}
	if opts.ShowAfterComplete > 0 {
		display := exec.Command("display", opts.OutputFileName)
		err = display.Run()
		if err != nil {
			log.Panicf("failed to execute display command: %s", err)
		}
	}
}

type WorkerInput struct {
	X, Y int
}

type WorkerConfig struct {
	Scene           scenes.Scene
	NumberOfSamples int
	Background      Vec3
	OutputImage     *image.RGBA
	InputChan       chan WorkerInput
	Wg              *sync.WaitGroup
}

func worker(config WorkerConfig) {
	defer config.Wg.Done()
	for input := range config.InputChan {
		i := float64(input.X)
		j := float64(config.Scene.Height - input.Y)
		col := Vec3{0, 0, 0}
		for s := 0; s < config.NumberOfSamples; s++ {
			u := (i + rand.Float64()) / float64(config.Scene.Width)
			v := (j + rand.Float64()) / float64(config.Scene.Height)
			ray := config.Scene.Camera.Ray(u, v)
			col = col.Add(RayColor(ray, config.Background, config.Scene.World, config.Scene.Lights, MaxDepth))
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
