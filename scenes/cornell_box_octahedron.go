package scenes

import (
	"math"

	. "ray-tracing/pkg/common"
	. "ray-tracing/pkg/hittable"
	. "ray-tracing/pkg/material"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/texture"
	. "ray-tracing/pkg/vec3"
)

func CornellBoxOctahedron(width, height int) (scene Scene) {
	m := int(math.Max(float64(width), float64(height)))
	scene.Width = m
	// width / height
	aspect := 1.0
	scene.Height = int(float64(scene.Width) / aspect)
	red := Lambertian{ConstantTexture{Vec3{.65, .05, .05}}}
	white := Lambertian{ConstantTexture{Vec3{.73, .73, .73}}}
	green := Lambertian{ConstantTexture{Vec3{.12, .45, .15}}}
	light := DiffuseLight{ConstantTexture{Vec3{15, 15, 15}}}
	blue := Lambertian{ConstantTexture{Vec3{.05, .05, .9}}}

	scene.World = append(scene.World, FlipFace{YZRect{0, 600, 0, 600, 600, green}})
	scene.World = append(scene.World, YZRect{0, 600, 0, 600, 0, red})
	scene.World = append(scene.World, FlipFace{XZRect{213, 343, 227, 332, 599, light}})
	scene.World = append(scene.World, FlipFace{XZRect{0, 600, 0, 600, 600, white}})
	scene.World = append(scene.World, XZRect{0, 600, 0, 600, 0, white})
	scene.World = append(scene.World, FlipFace{XYRect{0, 600, 0, 600, 600, white}})

	colors := []Material{red, blue, green}

	for i, x := range []int{100, 300, 500} {
		var octahedron Hittable
		octahedron = NewOctahedron(Vec3{0, 0, 0}, 100, colors[i])
		octahedron = Translate{octahedron, Vec3{float64(x), 400, 295}}
		scene.World = append(scene.World, octahedron)
		var sphere Hittable

		sphere = Sphere{Vec3{float64(x), 200, 295}, 100, colors[i]}
		scene.World = append(scene.World, sphere)
	}

	lookFrom := Vec3{278, 278, -800}
	lookAt := Vec3{278, 278, 0}
	up := Vec3{0, 1, 0}
	distToFocus := 10.0
	aperture := 0.0
	vFov := 40.0
	t0 := 0.0
	t1 := 1.0

	scene.Lights = NewList(
		XZRect{213, 343, 227, 332, 599, nil},
	)

	scene.Camera = NewCamera(lookFrom, lookAt, up, vFov, aspect, aperture, distToFocus, t0, t1)

	return scene
}
