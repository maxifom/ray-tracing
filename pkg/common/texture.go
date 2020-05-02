package common

import (
	. "ray-tracing/pkg/vec3"
)

type Texture interface {
	Value(u, v float64, p Vec3) Vec3
}
