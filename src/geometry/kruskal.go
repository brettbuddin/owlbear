package geometry

import "sort"

func Kruskal(pointers []Pointer, triangles []Triangle) []Edge {
	mst := []Edge{}
	edges := []Edge{}
	for _, t := range triangles {
		edges = append(edges, Edge{t.A, t.B})
		edges = append(edges, Edge{t.B, t.C})
		edges = append(edges, Edge{t.C, t.A})
	}
	sort.Sort(&EdgeSorter{edges, SortByDistance})

	union := NewDisjointPointers(pointers)
	for len(edges) > 0 && len(mst) <= len(pointers) {
		var e Edge
		e, edges = edges[0], edges[1:]
		if !union.AreJoined(e.A, e.B) {
			union.Join(e.A, e.B)
			mst = append(mst, e)
		}
	}

	return mst
}

// DisjointPointers is an implementation of a disjoint-set data structure for use as the "trees" in Kruskal's algorithm.
// Reference:
// - https://en.wikipedia.org/wiki/Disjoint-set_data_structure
type DisjointPointers struct {
	Pointers map[Pointer]int
}

// NewDisjointPointers returns a new DistjointPointers. It implicitely performs the "makeset" operation of disjoint-set
// data structure.
func NewDisjointPointers(pointers []Pointer) *DisjointPointers {
	u := &DisjointPointers{map[Pointer]int{}}
	for i, p := range pointers {
		u.Pointers[p] = i
	}
	return u
}

// IsConnected lets us know if two partitions are connected
func (d *DisjointPointers) AreJoined(a, b Pointer) bool {
	return d.Pointers[a] == d.Pointers[b]
}

// Join connects two disjointed sets
func (d *DisjointPointers) Join(a, b Pointer) {
	if d.AreJoined(a, b) {
		return
	}
	aIdx, bIdx := d.Pointers[a], d.Pointers[b]
	for p, i := range d.Pointers {
		if i == bIdx {
			d.Pointers[p] = aIdx
		}
	}
}
