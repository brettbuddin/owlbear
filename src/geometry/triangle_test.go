package geometry

import "testing"

func TestTrianglePointMembership(t *testing.T) {
	var examples = []struct {
		t      Triangle
		p      Point
		expect bool
	}{
		{
			Triangle{Point{1, 1}, Point{0, 0}, Point{0, 1}},
			Point{0, 1},
			true,
		},
		{
			Triangle{Point{1, 1}, Point{0, 0}, Point{0, 1}},
			Point{5, 1},
			false,
		},
	}
	for i, ex := range examples {
		actual := ex.t.IsPoint(ex.p)
		if actual != ex.expect {
			t.Errorf("[example %d] expected expected=%v got=%v", i, ex.expect, actual)
		}
	}
}

func TestTriangleCircumcircle(t *testing.T) {
	var examples = []struct {
		t      Triangle
		p      Point
		expect bool
	}{
		{
			Triangle{Point{3, 3}, Point{1, 1}, Point{0, 0}},
			Point{2, 2},
			true,
		},
		{
			Triangle{Point{0, 1}, Point{-1, 0}, Point{1, 0}},
			Point{5, 5},
			false,
		},
		{
			Triangle{Point{10, 10}, Point{0, 0}, Point{10, 0}},
			Point{2, 15},
			false,
		},
	}
	for i, ex := range examples {
		actual := ex.t.CircumcircleContains(ex.p)
		if actual != ex.expect {
			t.Errorf("[example %d] expected expected=%v got=%v", i, ex.expect, actual)
		}
	}
}
