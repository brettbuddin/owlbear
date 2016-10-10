package geometry

import "testing"

func TestEdgeEquality(t *testing.T) {
	var examples = []struct {
		a, b   Edge
		expect bool
	}{
		{Edge{Point{0, 0}, Point{0, 0}}, Edge{Point{0, 0}, Point{0, 0}}, true},
		{Edge{Point{1, 0}, Point{0, 0}}, Edge{Point{0, 0}, Point{0, 0}}, false},
		{Edge{Point{1, 0}, Point{0, 1}}, Edge{Point{0, 0}, Point{0, 0}}, false},
		{Edge{Point{1, 1}, Point{1, 1}}, Edge{Point{1, 1}, Point{1, 1}}, true},
		{Edge{Point{1, 0}, Point{0, 1}}, Edge{Point{0, 1}, Point{1, 0}}, true},
	}
	for i, ex := range examples {
		actual := ex.a.Eq(ex.b)
		if actual != ex.expect {
			t.Errorf("[example %d] expected expected=%v got=%v", i, ex.expect, actual)
		}
	}
}

func TestEdgeDistance(t *testing.T) {
	var examples = []struct {
		a      Edge
		expect float64
	}{
		{Edge{Point{0, 0}, Point{0, 0}}, 0},
		{Edge{Point{1, 0}, Point{0, 0}}, 1},
		{Edge{Point{1, 0}, Point{0, 1}}, 1.4142135623730951},
		{Edge{Point{1, 1}, Point{1, 1}}, 0},
		{Edge{Point{1, 0}, Point{0, 1}}, 1.4142135623730951},
		{Edge{Point{5, 0}, Point{0, 1}}, 5.0990195135927845},
	}
	for i, ex := range examples {
		actual := ex.a.Distance()
		if actual != ex.expect {
			t.Errorf("[example %d] expected expected=%v got=%v", i, ex.expect, actual)
		}
	}
}
