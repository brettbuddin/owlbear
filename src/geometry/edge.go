package geometry

// Edge is a line between two Pointers
type Edge struct {
	A, B Pointer
}

// IsPoint determines whether a Pointer is one of the points on the Edge
func (e Edge) IsPoint(p Pointer) bool {
	return PointerEq(p, e.A) || PointerEq(p, e.B)
}

// Distance calculates the distance of an Edge
func (e Edge) Distance() float64 {
	return PointerDistance(e.A, e.B)
}

// Eq determines whether two Edges are equal
func (e Edge) Eq(o Edge) bool {
	return (PointerEq(e.A, o.A) && PointerEq(e.B, o.B)) || (PointerEq(e.A, o.B) && PointerEq(e.B, o.A))
}

// EdgeSorter sorts a list of Edges using a specified operation
type EdgeSorter struct {
	Edges []Edge
	Op    func(a, b Edge) bool
}

func (s *EdgeSorter) Len() int           { return len(s.Edges) }
func (s *EdgeSorter) Less(i, j int) bool { return s.Op(s.Edges[i], s.Edges[j]) }
func (s *EdgeSorter) Swap(i, j int)      { s.Edges[i], s.Edges[j] = s.Edges[j], s.Edges[i] }

func SortByDistance(a, b Edge) bool {
	return a.Distance() < b.Distance()
}
