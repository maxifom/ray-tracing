package pdf

import (
	"ray-tracing/pkg/common"
	. "ray-tracing/pkg/vec3"
)

type HittablePDF struct {
	H common.Hittable
	O Vec3
}

func (p HittablePDF) Value(direction Vec3) float64 {
	return p.H.PDFValue(p.O, direction)
}

func (p HittablePDF) Generate() Vec3 {
	return p.H.Random(p.O)
}
