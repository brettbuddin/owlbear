package geometry

import "testing"

func TestPointEquality(t *testing.T) {
	var examples = []struct {
		a, b   Point
		expect bool
	}{
		{Point{0, 0}, Point{0, 0}, true},
		{Point{1, 0}, Point{1, 0}, true},
		{Point{1, 1}, Point{1, 1}, true},
		{Point{0, 1}, Point{0, 1}, true},
		{Point{1, 1}, Point{0, 0}, false},
		{Point{1, 0}, Point{0, 1}, false},
	}
	for i, ex := range examples {
		actual := ex.a.Eq(ex.b)
		if actual != ex.expect {
			t.Errorf("[example %d] expected expected=%v got=%v", i, ex.expect, actual)
		}
	}
}

func TestPointDiff(t *testing.T) {
	var examples = []struct {
		a, b   Point
		expect Point
	}{
		{Point{0, 0}, Point{0, 0}, Point{0, 0}},
		{Point{1, 0}, Point{1, 0}, Point{0, 0}},
		{Point{1, 1}, Point{1, 1}, Point{0, 0}},
		{Point{0, 1}, Point{0, 1}, Point{0, 0}},
		{Point{1, 1}, Point{0, 0}, Point{1, 1}},
		{Point{1, 0}, Point{0, 1}, Point{1, -1}},
	}
	for i, ex := range examples {
		actual := ex.a.Diff(ex.b)
		if actual != ex.expect {
			t.Errorf("[example %d] expected expected=%v got=%v", i, ex.expect, actual)
		}
	}
}

func TestPointDistance(t *testing.T) {
	var examples = []struct {
		a, b   Point
		expect float64
	}{
		{Point{0, 0}, Point{0, 0}, 0},
		{Point{0, 1}, Point{0, 1}, 0},
		{Point{3, 1}, Point{2, 3}, 2.23606797749979},
		{Point{1, 0}, Point{0, 1}, 1.4142135623730951},
	}
	for i, ex := range examples {
		actual := ex.a.Distance(ex.b)
		if actual != ex.expect {
			t.Errorf("[example %d] expected expected=%v got=%v", i, ex.expect, actual)
		}
	}
}
