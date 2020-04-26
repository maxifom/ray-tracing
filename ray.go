package main

import "fmt"

// Луч +++
type Ray struct {
	Origin, Direction Vec3

	Time float64
}

func (r Ray) PointAtParameter(t float64) Vec3 {
	return r.Origin.Add(r.Direction.MulN(t))
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray {Origin: %#v, Direction: %#v, Time: %.1f}", r.Origin, r.Direction, r.Time)
}
