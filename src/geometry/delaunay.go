package geometry

import "sort"

// Delaunay performs a Delaunay triangulation
// Reference:
// - Algorithm: https://en.wikipedia.org/wiki/Delaunay_triangulation#Incremental
func Delaunay(p []Pointer) []Triangle {
	sort.Sort(&PointSorter{p, SortByX})
	return (&triangulation{Points: p}).triangulate()
}

type triangulation struct {
	Points []Pointer
}

func (t *triangulation) triangulate() []Triangle {
	outer := t.calculateOuterTriangle()
	triangles := []Triangle{outer}

	for _, p := range t.Points {
		edges := []Edge{}

		k := 0
		for _, t := range triangles {
			if t.CircumcircleContains(p) {
				edges = append(edges, Edge{t.A, t.B})
				edges = append(edges, Edge{t.B, t.C})
				edges = append(edges, Edge{t.C, t.A})
			} else {
				triangles[k] = t
				k++
			}
		}
		triangles = triangles[:k]

		k = 0
		for i := 0; i < len(edges); i++ {
			var dup bool
			for j := i + 1; j < len(edges); j++ {
				if edges[i].Eq(edges[j]) {
					dup = true
				}
			}
			if !dup {
				edges[k] = edges[i]
				k++
			}
		}
		edges = edges[:k]

		for _, e := range edges {
			triangles = append(triangles, Triangle{e.A, e.B, p})
		}
	}

	k := 0
	for _, t := range triangles {
		if !(t.IsPoint(outer.A) || t.IsPoint(outer.B) || t.IsPoint(outer.C)) {
			triangles[k] = t
			k++
		}
	}

	return triangles[:k]
}

func (t *triangulation) calculateOuterTriangle() Triangle {
	min := Point{t.Points[0].Point().X, t.Points[0].Point().Y}
	max := min
	for _, p := range t.Points {
		point := p.Point()
		if point.X < min.X {
			min.X = point.X
		}
		if point.Y < min.Y {
			min.Y = point.Y
		}

		if point.X > max.X {
			max.X = point.X
		}
		if point.Y > max.Y {
			max.Y = point.Y
		}
	}

	diff := max.Diff(min)
	overallMax := diff.X
	if diff.Y > diff.X {
		overallMax = diff.Y
	}
	mid := Point{(max.X + min.X) / 2, (max.Y + min.Y) / 2}

	tri := Triangle{
		A: Point{mid.X - 2*overallMax, mid.Y - overallMax},
		B: Point{mid.X, mid.Y + 2*overallMax},
		C: Point{mid.Y + 2*overallMax, mid.Y - overallMax},
	}
	t.Points = append(t.Points, tri.A, tri.B, tri.C)
	return tri
}
