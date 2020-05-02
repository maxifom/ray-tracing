package scenes

import (
	"ray-tracing/pkg/common"
	. "ray-tracing/pkg/hittable"
	. "ray-tracing/pkg/material"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/texture"
	. "ray-tracing/pkg/vec3"
)

func CornellBox(aspect float64) (world HittableList, camera Camera, lights HittableList) {
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

	var box1 common.Hittable

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

	lights = NewList(
		XZRect{213, 343, 227, 332, 554, nil},
		Sphere{Vec3{190, 90, 190}, 90, nil},
	)

	camera = NewCamera(lookFrom, lookAt, up, vFov, aspect, aperture, distToFocus, t0, t1)
	return world, camera, lights
}
