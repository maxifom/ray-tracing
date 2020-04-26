package main

type HittablePDF struct {
	H Hittable
	O Vec3
}

func (p HittablePDF) Value(direction Vec3) float64 {
	return p.H.PDFValue(p.O, direction)
}

func (p HittablePDF) Generate() Vec3 {
	return p.H.Random(p.O)
}
