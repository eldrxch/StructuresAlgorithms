package structures

import (
	"reflect"
	"testing"
)

func TestAddEdge(t *testing.T) {
	tests := []struct {
		verts     [][]int
		wantCount int
	}{
		{[][]int{{1, 2}}, 2},
		{[][]int{{1, 2}, {3, 4}}, 4},
		{[][]int{{1, 2}, {3, 4}, {5, 4}}, 6},
	}

	for _, test := range tests {
		g := &SimpleGraph{}
		for i := range test.verts {
			vert1 := test.verts[i][0]
			vert2 := test.verts[i][1]
			g.AddEdge(vert1, vert2)
		}
		if g.Count() != test.wantCount {
			t.Errorf("got %v, want %v", g.Count(), test.wantCount)
		}
	}
}

func TestNeighborsOf(t *testing.T) {
	tests := []struct {
		verts  [][]int
		lookup int
		want   []int
	}{
		{[][]int{}, 9, []int{}},                                              // empty graph
		{[][]int{{1, 2}}, 2, []int{1}},                                       // single edge - 1st vert
		{[][]int{{3, 2}}, 3, []int{2}},                                       // single edge - 2nd vert
		{[][]int{{1, 2}, {3, 2}}, 2, []int{1, 3}},                            // two edges
		{[][]int{{1, 2}, {3, 4}, {5, 4}, {5, 6}, {5, 1}}, 5, []int{4, 6, 1}}, // multiple edges
	}

	for _, test := range tests {
		g := &SimpleGraph{}
		for i := range test.verts {
			vert1 := test.verts[i][0]
			vert2 := test.verts[i][1]
			g.AddEdge(vert1, vert2)
		}
		neighs := g.NeighborsOf(test.lookup)
		if !reflect.DeepEqual(neighs, test.want) {
			t.Errorf("got %v, want %v", neighs, test.want)
		}
	}
}
