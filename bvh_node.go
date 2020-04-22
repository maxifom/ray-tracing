package main

import "math/rand"

type BVHNode struct {
	Box         AABB
	Left, Right Hittable
}

func NewBVHNode(list HittableList, n int64, time0, time1 float64) BVHNode {
	axis := int64(3 * rand.Float64())
	if axis == 0 {

	}
}

func (N BVHNode) Hit(r Ray, tMin, tMax float64) (HitRecord, bool) {
	_, hit := N.Box.Hit(r, tMin, tMax)
	if hit {
		leftRec, hitLeft := N.Left.Hit(r, tMin, tMax)
		rightRec, hitRight := N.Right.Hit(r, tMin, tMax)
		if hitLeft && hitRight {
			if leftRec.T < rightRec.T {
				return leftRec, true
			}
			return rightRec, true
		} else if hitLeft {
			return leftRec, true
		} else if hitRight {
			return rightRec, true
		}
		return HitRecord{}, false
	}

	return HitRecord{}, false
}

func (N BVHNode) BoundingBox(t0, t1 float64) (AABB, bool) {
	return N.Box, true
}
