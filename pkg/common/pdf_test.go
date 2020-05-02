package common

import (
	"math"
	"testing"

	. "ray-tracing/pkg/hittable"
	. "ray-tracing/pkg/pdf"
	. "ray-tracing/pkg/vec3"
)

func TestCosinePDF_Value(t *testing.T) {
	c := CosinePDF{ONB: NewONB(Vec3{1.23, 5.66, 3.123})}
	v := c.Value(Vec3{1.0, 2.324, 6.324})

	if math.Abs(v-0.2424100) > 1e-6 {
		t.Error("failed to cosine pdf value")
	}
}

func TestHittablePDF_Value(t *testing.T) {
	h := XZRect{0, 100, 0, 200, 100, nil}
	c := HittablePDF{O: Vec3{1.23, 5.66, 3.123}, H: h}
	v := c.Value(Vec3{2, 3, 5})

	if math.Abs(v-3.86077) > 1e-6 {
		t.Error("failed to hittable pdf value")
	}
}
