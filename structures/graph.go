package structures

import (
	"bytes"
)

type SimpleGraph struct {
	vertices map[int]DoublyLinkedList
}

func (g *SimpleGraph) AddEdge(vert1, vert2 int) {
	if g.vertices == nil {
		g.vertices = make(map[int]DoublyLinkedList)
	}
	if _, ok := g.vertices[vert1]; !ok {
		root := DoublyLinkedList{}
		root.Add(vert1)
		root.Add(vert2)
		g.vertices[vert1] = root
		return
	}
	root := g.vertices[vert1]
	root.Add(vert2)
	g.vertices[vert1] = root
}

func (g *SimpleGraph) NeighborsOf(vert int) []int {
	if g.vertices == nil {
		return nil
	}

	neighs := []int{}
	visited := map[int]bool{}
	var queue []int
	for k := range g.vertices {
		if _, ok := visited[k]; ok {
			continue
		}
		queue := append(queue, k)
		for len(queue) > 0 {
			elm := queue[0]
			queue = append(queue[:0], queue[1:]...)
			if _, ok := visited[elm]; ok {
				continue
			}
			visited[elm] = true
			root := g.vertices[elm]
			links := root.Values()
			if links != nil {
				links = append(links[:0], links[1:]...)
			}
			if len(links) > 0 && elm == vert {
				neighs = append(neighs, links...)
			}
			if len(links) > 0 && contains(links, vert) {
				neighs = append(neighs, elm)
			}
			queue = append(queue, links...)
		}
	}
	return neighs
}

func contains(s []int, elm int) bool {
	for _, v := range s {
		if v == elm {
			return true
		}
	}
	return false
}

func (g *SimpleGraph) Count() int {
	if g.vertices == nil {
		return 0
	}
	size := 0
	for _, v := range g.vertices {
		size = v.Size() + size
		buf := new(bytes.Buffer)
		v.WriteString(buf)
	}
	return size
}
