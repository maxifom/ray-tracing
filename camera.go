package main

import (
	"math"
)

type Camera struct {
	Origin          Vec3
	LowerLeftCorner Vec3
	Horizontal      Vec3
	Vertical        Vec3

	U          Vec3
	V          Vec3
	W          Vec3
	LensRadius float64

	Time0, Time1 float64
}

func NewCamera(lookFrom, lookAt, vUp Vec3, vFov, aspect, aperture, focusDist, t0, t1 float64) Camera {
	theta := vFov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	w := lookFrom.Sub(lookAt).UnitVector()
	u := Cross(vUp, w).UnitVector()
	v := Cross(w, u)
	return Camera{
		Origin: lookFrom,
		LowerLeftCorner: lookFrom.
			Sub(u.MulN(halfWidth * focusDist)).
			Sub(v.MulN(halfHeight * focusDist)).
			Sub(w.MulN(focusDist)),
		Horizontal: u.MulN(2 * halfWidth * focusDist),
		Vertical:   v.MulN(2 * halfHeight * focusDist),
		U:          u,
		V:          v,
		W:          w,
		LensRadius: aperture / 2,

		Time0: t0,
		Time1: t1,
	}
}

func (c Camera) Ray(s, t float64) Ray {
	rd := RandomInUnitDisk().MulN(c.LensRadius)
	offset := c.U.MulN(rd.X).Add(c.V.MulN(rd.Y))
	return Ray{
		Origin: c.Origin.Add(offset),
		Direction: c.LowerLeftCorner.
			Add(c.Horizontal.
				MulN(s)).
			Add(c.Vertical.
				MulN(t)).
			Sub(c.Origin).
			Sub(offset),
		Time: RandomDouble(c.Time0, c.Time1),
	}
}
