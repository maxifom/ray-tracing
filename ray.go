package main

// Луч
type Ray struct {
	Origin, Direction Vec3
}

func (r Ray) PointAtParameter(t float64) Vec3 {
	return r.Origin.Add(r.Direction.MulN(t))
}
