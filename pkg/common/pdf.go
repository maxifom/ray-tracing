package common

import (
	. "ray-tracing/pkg/vec3"
)

// Probability density function
type PDF interface {
	Value(direction Vec3) float64
	Generate() Vec3
}
