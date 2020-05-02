package texture

import . "ray-tracing/pkg/vec3"

func PerlinInterpolation(c [2][2][2]Vec3, u, v, w float64) float64 {
	uu := u * u * (3 - 2*u)
	vv := v * v * (3 - 2*v)
	ww := w * w * (3 - 2*w)
	accum := 0.0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				i1 := float64(i)
				j1 := float64(j)
				k1 := float64(k)
				weightV := Vec3{u - i1, v - j1, w - k1}
				accum += (i1*uu + (1-i1)*(1-uu)) *
					(j1*vv + (1-j1)*(1-vv)) *
					(k1*ww + (1-k1)*(1-ww)) *
					Dot(c[i][j][k], weightV)
			}
		}
	}

	return accum
}
