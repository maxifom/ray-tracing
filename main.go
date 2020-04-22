package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
)

// Цвет который получает луч
func Color(ray Ray, world Hittable, depth int64) Vec3 {
	record, hit := world.Hit(ray, 0.001, math.MaxFloat64)
	if hit {
		scattered, attenuation, hasScattered := record.Material.Scatter(ray, record)
		if depth < 50 && hasScattered {
			return attenuation.Mul(Color(scattered, world, depth+1))
		}
		return Vec3{0, 0, 0}
	}

	unitDirection := ray.Direction.UnitVector()
	t := 0.5*unitDirection.Y + 1.0

	return Vec3{1, 1, 1}.MulN(1 - t).Add(Vec3{0.5, 0.7, 1.0}.MulN(t))
}

func RandomScene() Hittable {
	n := 500
	hl := make(HittableList, 0, n)
	hl = append(hl, Sphere{Vec3{0, -1000, 0}, 1000, Lambertian{Vec3{0.5, 0.5, 0.5}}})
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			if center.Sub(Vec3{4, 0.2, 0}).Length() > 0.9 {
				if chooseMat < 0.8 {
					hl = append(hl, MovingSphere{center, center.Add(Vec3{0, 0.5 * rand.Float64(), 0}), 0.2, Lambertian{Vec3{rand.Float64() * rand.Float64(), rand.Float64() * rand.Float64(), rand.Float64() * rand.Float64()}}, 0, 1})
				} else if chooseMat < 0.95 {
					hl = append(hl, Sphere{center, 0.2, Metal{Vec3{0.5 * (1 + rand.Float64()), 0.5 * (1 + rand.Float64()), 0.5 * (1 + rand.Float64())}}})
				} else {
					hl = append(hl, Sphere{center, 0.2, Dielectric{1.5}})
				}
			}
		}
	}

	hl = append(hl, Sphere{Vec3{0, 1, 0}, 1, Dielectric{1.5}})
	hl = append(hl, Sphere{Vec3{-4, 1, 0}, 1, Lambertian{Vec3{0.4, 0.2, 0.1}}})
	hl = append(hl, Sphere{Vec3{4, 1, 0}, 1, Metal{Vec3{0.7, 0.6, 0.5}}})
	return hl
}

func main() {
	file, err := os.OpenFile("output.ppm", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	width := 16 * 100
	height := 9 * 100
	numberOfTimes := 2

	fmt.Fprintf(file, "P3\n%d %d\n255\n", width, height)

	world := RandomScene()

	lookFrom := Vec3{13, 2, 5}
	lookAt := Vec3{0, 0, 0}
	focusDist := 10.0
	aperture := 0.1
	cam := NewCamera(
		lookFrom,
		lookAt,
		Vec3{0, 1, 0},
		20,
		float64(width)/float64(height),
		aperture,
		focusDist,
		0,
		1,
	)

	for j := height - 1; j >= 0; j-- {
		log.Println(j)
		for i := 0; i < width; i++ {
			col := Vec3{0, 0, 0}
			for s := 0; s < numberOfTimes; s++ {
				u := (float64(i) + rand.Float64()) / float64(width)
				v := (float64(j) + rand.Float64()) / float64(height)
				ray := cam.Ray(u, v)
				col = col.Add(Color(ray, world, 0))
			}

			col = col.DivN(float64(numberOfTimes))
			col = Vec3{
				X: math.Sqrt(col.X),
				Y: math.Sqrt(col.Y),
				Z: math.Sqrt(col.Z),
			}
			ir := int64(255.99 * col.X)
			ig := int64(255.99 * col.Y)
			ib := int64(255.99 * col.Z)

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
