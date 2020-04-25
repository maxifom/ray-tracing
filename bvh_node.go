package main

import (
	"math/rand"
	"sort"
)

type BVHNode struct {
	Box         AABB
	Left, Right Hittable
}

func SortByX(list HittableList) HittableList {
	sort.Slice(list, func(i, j int) bool {
		box1, _ := list[i].BoundingBox(0, 0)
		box2, _ := list[j].BoundingBox(0, 0)
		return box1.Min.X > box2.Min.X
	})
	return list
}

func SortByY(list HittableList) HittableList {
	sort.Slice(list, func(i, j int) bool {
		box1, _ := list[i].BoundingBox(0, 0)
		box2, _ := list[j].BoundingBox(0, 0)
		return box1.Min.Y > box2.Min.Y
	})
	return list
}

func SortByZ(list HittableList) HittableList {
	sort.Slice(list, func(i, j int) bool {
		box1, _ := list[i].BoundingBox(0, 0)
		box2, _ := list[j].BoundingBox(0, 0)
		return box1.Min.Z > box2.Min.Z
	})
	return list
}

func NewBVHNode(list HittableList, n int64, time0, time1 float64) BVHNode {
	var node BVHNode
	axis := int64(3 * rand.Float64())
	if axis == 0 {
		list = SortByX(list)
	} else if axis == 1 {
		list = SortByY(list)
	} else {
		list = SortByZ(list)
	}
	if n == 1 {
		node.Left = list[0]
		node.Right = list[0]
	} else if n == 2 {
		node.Left = list[0]
		node.Right = list[1]
	} else {
		node.Left = NewBVHNode(list, n/2, time0, time1)
		node.Right = NewBVHNode(list[n/2:], n/2, time0, time1)
	}

	boxLeft, _ := node.Left.BoundingBox(time0, time1)
	boxRight, _ := node.Right.BoundingBox(time0, time1)

	node.Box = SurroundingBox(boxLeft, boxRight)
	return node
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

func (N BVHNode) PDFValue(o, v Vec3) float64 {
	return 0
}

func (N BVHNode) Random(o Vec3) Vec3 {
	return Vec3{1, 0, 0}
}
