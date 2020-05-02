package scenes

import (
	"math"
	"math/rand"

	. "ray-tracing/pkg/hittable"
	. "ray-tracing/pkg/material"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/texture"
	. "ray-tracing/pkg/utils"
	. "ray-tracing/pkg/vec3"
)

func SimpleScene(width, height int) (scene Scene) {
	m := int(math.Max(float64(width), float64(height)))
	scene.Width = m
	// width / height
	aspect := 16.0 / 9.0
	scene.Height = int(float64(scene.Width) / aspect)

	scene.World = append(scene.World, Sphere{Vec3{0, -1000, 0}, 1000, Lambertian{CheckerTexture{ConstantTexture{Vec3{0, 0, 0}}, ConstantTexture{Vec3{1, 1, 1}}}}})
	light := Sphere{Vec3{0, 10, 0}, 5, DiffuseLight{ConstantTexture{Vec3{8, 8, 8}}}}
	scene.World = append(scene.World, light)
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMaterial := rand.Float64()
			chooseForm := rand.Float64()
			center := Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			if center.Sub(Vec3{4, 0.2, 0}).Length() > 0.9 {
				if chooseMaterial < 0.8 {
					albedo := ConstantTexture{Vec3{rand.Float64(), rand.Float64(), rand.Float64()}}
					if chooseForm < 0.6 {
						scene.World = append(scene.World, Sphere{center, 0.2, Lambertian{albedo}})
					} else {
						scene.World = append(scene.World, NewOctahedron(center, 0.2, Lambertian{albedo}))
					}
				} else if chooseMaterial < 0.95 {
					albedo := Vec3{RandomDouble(.5, 1), RandomDouble(.5, 1), RandomDouble(.5, 1)}
					fuzz := RandomDouble(0, .5)
					scene.World = append(scene.World, Sphere{center, .2, Metal{albedo, fuzz}})
				} else {
					scene.World = append(scene.World, Sphere{center, 0.2, Dielectric{1.5}})
				}
			}
		}
	}

	scene.World = append(scene.World, Sphere{Vec3{0, 1, 0}, 1, Dielectric{1.5}})
	scene.World = append(scene.World, Sphere{Vec3{-4, 1, 0}, 1, Lambertian{ConstantTexture{Vec3{.4, .2, .1}}}})
	scene.World = append(scene.World, Sphere{Vec3{4, 1, 0}, 1, Metal{Vec3{.7, .6, .5}, 0}})

	lookFrom := Vec3{13, 2, 3}
	lookAt := Vec3{0, 0, 0}
	up := Vec3{0, 1, 0}
	distToFocus := 10.0
	aperture := 0.1
	vFov := 20.0
	t0 := 0.0
	t1 := 1.0

	scene.Lights = NewList(
		light,
	)

	scene.Camera = NewCamera(lookFrom, lookAt, up, vFov, aspect, aperture, distToFocus, t0, t1)

	return scene
}
