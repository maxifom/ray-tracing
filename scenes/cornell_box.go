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

func CornellBox(width, height int) (scene Scene) {
	m := int(math.Max(float64(width), float64(height)))
	scene.Width = m
	// width / height
	aspect := 1.0
	scene.Height = int(float64(scene.Width) / aspect)
	red := Lambertian{ConstantTexture{Vec3{.65, .05, .05}}}
	white := Lambertian{ConstantTexture{Vec3{.73, .73, .73}}}
	green := Lambertian{ConstantTexture{Vec3{.12, .45, .15}}}
	light := DiffuseLight{ConstantTexture{Vec3{15, 15, 15}}}

	scene.World = append(scene.World, FlipFace{YZRect{0, 555, 0, 555, 555, green}})
	scene.World = append(scene.World, YZRect{0, 555, 0, 555, 0, red})
	scene.World = append(scene.World, FlipFace{XZRect{213, 343, 227, 332, 554, light}})
	scene.World = append(scene.World, FlipFace{XZRect{0, 555, 0, 555, 555, white}})
	scene.World = append(scene.World, XZRect{0, 555, 0, 555, 0, white})
	scene.World = append(scene.World, FlipFace{XYRect{0, 555, 0, 555, 555, white}})

	var box1 Hittable

	box1 = NewBox(Vec3{0, 0, 0}, Vec3{165, 330, 165}, white)
	box1 = NewRotateY(box1, 15)
	box1 = Translate{box1, Vec3{265, 0, 295}}
	scene.World = append(scene.World, box1)

	glass := Dielectric{1.5}
	scene.World = append(scene.World, Sphere{Vec3{190, 90, 190}, 90, glass})

	lookFrom := Vec3{278, 278, -800}
	lookAt := Vec3{278, 278, 0}
	up := Vec3{0, 1, 0}
	distToFocus := 10.0
	aperture := 0.0
	vFov := 40.0
	t0 := 0.0
	t1 := 1.1

	scene.Lights = NewList(
		XZRect{213, 343, 227, 332, 554, nil},
		Sphere{Vec3{190, 90, 190}, 90, nil},
	)

	scene.Camera = NewCamera(lookFrom, lookAt, up, vFov, aspect, aperture, distToFocus, t0, t1)

	return scene
}
