package common

import (
	. "ray-tracing/pkg/vec3"
)

type PDF interface {
	Value(direction Vec3) float64
	Generate() Vec3
}
