package texture

import (
	"math"

	"ray-tracing/pkg/common"
	. "ray-tracing/pkg/vec3"
)

type CheckerTexture struct {
	Odd, Even common.Texture
}

func (c CheckerTexture) Value(u, v float64, p Vec3) Vec3 {
	sines := math.Sin(10*p.X) * math.Sin(10*p.Y) * math.Sin(10*p.Z)
	if sines < 0 {
		return c.Odd.Value(u, v, p)
	}

	return c.Even.Value(u, v, p)
}
