package scene

import (
	"fmt"

	. "ray-tracing/pkg/vec3"
)

type Ray struct {
	Origin, Direction Vec3

	Time float64
}

func (r Ray) PointAtParameter(t float64) Vec3 {
	return r.Origin.Add(r.Direction.MulN(t))
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray {O: %s, D: %s}", r.Origin, r.Direction)
}
