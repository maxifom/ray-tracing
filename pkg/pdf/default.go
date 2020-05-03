package pdf

import . "ray-tracing/pkg/vec3"

type DefaultPDF struct {
}

func (d DefaultPDF) Value(direction Vec3) float64 {
	return 0
}

func (d DefaultPDF) Generate() Vec3 {
	return Vec3{1, 0, 0}
}
