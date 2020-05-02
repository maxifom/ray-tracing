package material

import (
	"ray-tracing/pkg/common"
	. "ray-tracing/pkg/scene"
	. "ray-tracing/pkg/vec3"
)

type DiffuseLight struct {
	Emit common.Texture
}

func (d DiffuseLight) Emitted(rIn Ray, u, v float64, rec common.HitRecord, p Vec3) Vec3 {
	// We also need to flip the light so its normals point in the -y direction
	if rec.FrontFace {
		return d.Emit.Value(u, v, p)
	}

	return Vec3{0, 0, 0}
}

func (d DiffuseLight) Scatter(Ray, common.HitRecord) (common.ScatterRecord, bool) {
	return common.ScatterRecord{}, false
}

func (d DiffuseLight) ScatteringPDF(Ray, common.HitRecord, Ray) float64 {
	return 0
}
