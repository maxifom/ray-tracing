package scenes

import (
	"math/rand"

	"ray-tracing/pkg/common"
	. "ray-tracing/pkg/hittable"
	. "ray-tracing/pkg/material"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/texture"
	. "ray-tracing/pkg/utils"
	. "ray-tracing/pkg/vec3"
)

func RandomHittable(center Vec3) common.Hittable {
	f := rand.Float64()
	if f < 0.8 {
		return Sphere{center, RandomDouble(.2, .35), RandomMaterial()}
	} else if f < 0.9 {
		return MovingSphere{center, center.Add(RandomVector(0, .1)), RandomDouble(.2, .35), RandomMaterial(), 0, RandomDouble(0.33, 1)}
	} else if f < 0.95 {
		return NewOctahedron(center, RandomDouble(.2, .35), Lambertian{RandomTexture()})
	} else {
		return NewBox(center.Sub(RandomVector(.2, .35)), center.Add(RandomVector(.2, .35)), RandomMaterial())
	}
}

func RandomTexture() common.Texture {
	f := rand.Float64()
	if f < 0.6 {
		return ConstantTexture{RandomVector(0, 1)}
	} else if f < 0.7 {
		return CheckerTexture{ConstantTexture{Vec3{0, 0, 0}}, ConstantTexture{Vec3{1, 1, 1}}}
	} else if f < 0.85 {
		return NoiseTextureFloat{NewPerlinFloat(), 2}
	} else {
		return NoiseTexture{NewPerlin(), 2}
	}
}

func RandomMaterial() common.Material {
	f := rand.Float64()
	if f < 0.8 {
		return Lambertian{RandomTexture()}
	} else if f < 0.95 {
		return Metal{RandomVector(0.5, 1), RandomDouble(0, .5)}
	} else {
		return Dielectric{1.5}
	}
}

func SimpleRandomScene(width int) (scene Scene) {
	scene.Width = width
	// width / height
	aspect := 16.0 / 9.0
	scene.Height = int(float64(scene.Width) / aspect)

	scene.World = append(scene.World, Sphere{Vec3{0, -1000, 0}, 1000, Lambertian{CheckerTexture{ConstantTexture{Vec3{0, 0, 0}}, ConstantTexture{Vec3{1, 1, 1}}}}})
	light := Sphere{Vec3{0, 10, 0}, 5, DiffuseLight{ConstantTexture{Vec3{8, 8, 8}}}}
	scene.World = append(scene.World, light)
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			center := Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}
			if center.Sub(Vec3{4, 0.2, 0}).Length() > 0.9 {
				scene.World = append(scene.World, RandomHittable(center))
			}
		}
	}

	scene.World = append(scene.World, Sphere{Vec3{0, 1, 0}, 1, Dielectric{1.5}})
	scene.World = append(scene.World, Sphere{Vec3{-4, 1, 0}, 1, Lambertian{ConstantTexture{RandomVector(0, 1)}}})
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
