package geometry

import "math"

// Point represents a coordinate on the map
type Point struct {
	X, Y float64
}

// Point implements Pointer interface. This allows us to us Pointer almost exclusively in other packages without caring
// if we have a concrete Point or some other type.
func (pointer Point) Point() Point {
	return pointer
}

// Eq determines whether two Points are equal
func (p Point) Eq(o Point) bool {
	return p.X == o.X && p.Y == o.Y
}

// Diff calculates slope between two Points
func (p Point) Diff(o Point) Point {
	return Point{X: p.X - o.X, Y: p.Y - o.Y}
}

// Distance calculates the distance between two Points
func (p Point) Distance(o Point) float64 {
	return math.Sqrt(math.Pow(float64(p.X-o.X), 2) + math.Pow(float64(p.Y-o.Y), 2))
}

// PointSorter sorts a list of Pointers using a specified operation
type PointSorter struct {
	Points []Pointer
	Op     func(a, b Point) bool
}

func (s *PointSorter) Len() int {
	return len(s.Points)
}

func (s *PointSorter) Less(i, j int) bool {
	return s.Op(s.Points[i].Point(), s.Points[j].Point())
}

func (s *PointSorter) Swap(i, j int) {
	s.Points[i], s.Points[j] = s.Points[j], s.Points[i]
}

func SortByX(a, b Point) bool { return a.X < b.X }

// Pointer yields a Point. This allows other types to be used for their intrinsic location on the map.
type Pointer interface {
	Point() Point
}

// PointerDistance calculates the distance between two Pointers
func PointerDistance(a, b Pointer) float64 {
	return a.Point().Distance(b.Point())
}

// PointerEq determines whether two Pointers are equal
func PointerEq(a, b Pointer) bool {
	return a.Point().Eq(b.Point())
}
