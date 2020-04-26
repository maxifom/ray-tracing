package main

import (
	"testing"
)

func TestNewONB(t *testing.T) {
	onb := NewONB(Vec3{1.234, 3.435, 5.234})

	u := Vec3{-0.981122, 0.106107, 0.161679}
	v := Vec3{0, 0.836034, -0.548677}
	w := Vec3{0.193388, 0.53832, 0.820252}

	l1 := onb.U.Sub(u).SqrLength()
	l2 := onb.V.Sub(v).SqrLength()
	l3 := onb.W.Sub(w).SqrLength()

	if l1 > 1e-6 || l2 > 1e-6 || l3 > 1e-6 {
		t.Error("failed new onb")
	}
}

func TestONB_Local(t *testing.T) {
	onb := NewONB(Vec3{1.234, 3.435, 5.234})
	v := onb.Local(Vec3{150.053345, 3408.2348, 1231.34589})

	v1 := Vec3{90.9064, 3528.18, -835.747}

	if v1.Sub(v).SqrLength() > 1e-6 {
		t.Error("failed onb local")
	}
}
