package main

import (
	"math"
	"math/rand"
)

type ConstantMedium struct {
	Boundary      Hittable
	PhaseFunction Material

	// -1/d в создании
	NegInvDensity float64
}

func (c ConstantMedium) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	h1, isHit1 := c.Boundary.Hit(r, math.Inf(-1), math.Inf(1))
	if !isHit1 {
		return HitRecord{}, false
	}

	h2, isHit2 := c.Boundary.Hit(r, h1.T+0.0001, math.Inf(1))
	if !isHit2 {
		return HitRecord{}, false
	}

	if h1.T < tMin {
		h1.T = tMin
	}

	if h2.T > tMax {
		h2.T = tMax
	}

	if h1.T >= h2.T {
		return HitRecord{}, false
	}

	if h1.T < 0 {
		h1.T = 0
	}

	rayLength := r.Direction.Length()
	distanceInsideBoundary := (h2.T - h1.T) * rayLength
	hitDistance := c.NegInvDensity * math.Log(rand.Float64())
	if hitDistance > distanceInsideBoundary {
		return HitRecord{}, false
	}

	var h HitRecord

	h.T = h1.T + hitDistance/rayLength
	h.P = r.PointAtParameter(h.T)
	h.Normal = Vec3{1, 0, 0}
	h.FrontFace = true
	h.Material = c.PhaseFunction

	return h, true
}

func (c ConstantMedium) BoundingBox(t0, t1 float64) (AABB, bool) {
	return c.Boundary.BoundingBox(t0, t1)
}

func (c ConstantMedium) PDFValue(o, v Vec3) float64 {
	return 0
}

func (c ConstantMedium) Random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
