package geometry

import "math"

type Triangle struct {
	A, B, C Pointer
}

func (t Triangle) IsPoint(p Pointer) bool {
	return PointerEq(p, t.A) || PointerEq(p, t.B) || PointerEq(p, t.C)
}

// CircumcircleContains determins whether a Pointer's Point is inside the triangle's circumcircle.
// Reference:
//	- https://en.wikipedia.org/wiki/Circumscribed_circle#Cartesian_coordinates_2
func (t Triangle) CircumcircleContains(p Pointer) bool {
	tA := t.A.Point()
	tB := t.B.Point()
	tC := t.C.Point()

	a := math.Pow(tA.X, 2) + math.Pow(tA.Y, 2)
	b := math.Pow(tB.X, 2) + math.Pow(tB.Y, 2)
	c := math.Pow(tC.X, 2) + math.Pow(tC.Y, 2)

	x := (a*(tC.Y-tB.Y) + b*(tA.Y-tC.Y) + c*(tB.Y-tA.Y)) /
		(tA.X*(tC.Y-tB.Y) + tB.X*(tA.Y-tC.Y) + tC.X*(tB.Y-tA.Y)) / 2
	y := (a*(tC.X-tB.X) + b*(tA.X-tC.X) + c*(tB.X-tA.X)) /
		(tA.Y*(tC.X-tB.X) + tB.Y*(tA.X-tC.X) + tC.Y*(tB.X-tA.X)) / 2
	radius := math.Sqrt(math.Pow(tA.X-x, 2) + math.Pow(tA.Y-y, 2))

	point := p.Point()
	dist := math.Sqrt(math.Pow(point.X-x, 2) + math.Pow(point.Y-y, 2))
	return dist <= radius
}
