package scenes

import (
	. "ray-tracing/pkg/hittable"
	. "ray-tracing/pkg/scene"
)

type Scene struct {
	World         HittableList
	Camera        Camera
	Lights        HittableList
	Width, Height int
}
