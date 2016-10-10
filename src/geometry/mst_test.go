package geometry

import "testing"

func TestMST(t *testing.T) {
	pointers := []Pointer{
		Point{10, 10},
		Point{10, 0},
		Point{0, 10},
		Point{5, 5},
		Point{-5, -5},
	}

	graph := Delaunay(pointers)
	mst := Kruskal(pointers, graph)

	expected := []Edge{
		{Point{10, 10}, Point{5, 5}},
		{Point{0, 10}, Point{5, 5}},
		{Point{5, 5}, Point{10, 0}},
		{Point{-5, -5}, Point{5, 5}},
	}

	for i, e := range expected {
		if !e.Eq(mst[i]) {
			t.Errorf("[example %d] expected expected=%v got=%v", i, e, mst[i])
		}
	}
}
