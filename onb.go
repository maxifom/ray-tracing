package main

import "math"

type ONB struct {
	U, V, W Vec3
}

func NewONB(n Vec3) ONB {
	var o ONB
	o.W = n.UnitVector()
	var a Vec3
	if math.Abs(o.W.X) > 0.9 {
		a = Vec3{0, 1, 0}
	} else {
		a = Vec3{1, 0, 0}
	}
	o.V = Cross(o.W, a)
	o.U = Cross(o.W, o.V)

	return o
}

func (o ONB) Local(v Vec3) Vec3 {
	return o.U.MulN(v.X).Add(o.V.MulN(v.Y)).Add(o.W.MulN(v.Z))
}
