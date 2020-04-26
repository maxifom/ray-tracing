package main

import (
	"math"
	"testing"
)

func TestAABB_Random(t *testing.T) {
	a := AABB{Vec3{1.123, 2.234, 5.123}, Vec3{5.324, 3.345, 10.123}}
	r := Vec3{34.435, 1.123, 6.234}
	v := a.Random(r)

	if v.Sub(Vec3{1, 0, 0}).SqrLength() > 1e-6 {
		t.Error("failed aabb random")
	}
}

func TestAABB_BoundingBox(t *testing.T) {
	a := AABB{Vec3{1.123, 2.234, 5.123}, Vec3{5.324, 3.345, 10.123}}
	v, hasBox := a.BoundingBox(0, 1)

	if v != a || hasBox != true {
		t.Error("failed aabb bounding box")
	}
}

func TestAABB_Hit(t *testing.T) {
	a := AABB{Vec3{1.123, 2.234, 5.123}, Vec3{5.324, 3.345, 10.123}}
	r := Ray{Vec3{9.2628, -3.9560, 6.1588}, Vec3{-7.3183, 8.9552, 2.8173}, 0}
	_, isHit := a.Hit(r, 0.0001, math.Inf(1))
	if isHit != true {
		t.Error("failed to hit aabb")
	}
}

func TestAABB_PDFValue(t *testing.T) {
	a := AABB{Vec3{1.123, 2.234, 5.123}, Vec3{5.324, 3.345, 10.123}}
	v := a.PDFValue(Vec3{}, Vec3{})
	if v != 0 {
		t.Error("failed to pdf value aabb")
	}
}
