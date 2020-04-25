package main

type HittablePDF struct {
	O Vec3
	H Hittable
}

func (p HittablePDF) Value(direction Vec3) float64 {
	return p.H.PDFValue(p.O, direction)
}

func (p HittablePDF) Generate() Vec3 {
	return p.H.Random(p.O)
}
