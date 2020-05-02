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
	// aluminium := Metal{Vec3{0.8, 0.85, 0.88}, 0}
	glass := Dielectric{1.5}

	scene.World = append(scene.World, FlipFace{YZRect{0, 600, 0, 600, 600, green}})
	scene.World = append(scene.World, YZRect{0, 600, 0, 600, 0, red})
	scene.World = append(scene.World, FlipFace{XZRect{213, 343, 227, 332, 599, light}})
	scene.World = append(scene.World, FlipFace{XZRect{0, 600, 0, 600, 600, white}})
	scene.World = append(scene.World, XZRect{0, 600, 0, 600, 0, white})
	scene.World = append(scene.World, FlipFace{XYRect{0, 600, 0, 600, 600, white}})

	var octahedron Hittable

	octahedron = NewOctahedron(Vec3{0, 0, 0}, 100, red)
	octahedron = Translate{octahedron, Vec3{100, 400, 295}}
	scene.World = append(scene.World, octahedron)

	var octahedron1 Hittable

	octahedron1 = NewOctahedron(Vec3{0, 0, 0}, 100, blue)
	octahedron1 = Translate{octahedron1, Vec3{300, 400, 295}}
	scene.World = append(scene.World, octahedron1)

	var octahedron2 Hittable

	octahedron2 = NewOctahedron(Vec3{0, 0, 0}, 100, green)
	octahedron2 = Translate{octahedron2, Vec3{500, 400, 295}}
	scene.World = append(scene.World, octahedron2)

	var sphere Hittable

	sphere = Sphere{Vec3{100, 200, 295}, 100, red}
	scene.World = append(scene.World, sphere)

	var sphere1 Hittable

	sphere1 = Sphere{Vec3{300, 200, 295}, 100, glass}
	scene.World = append(scene.World, sphere1)

	var sphere2 Hittable

	sphere2 = Sphere{Vec3{500, 200, 295}, 100, green}
	scene.World = append(scene.World, sphere2)

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
