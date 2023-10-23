package structures

type SimpleGraph struct {
	vertices map[int]DoublyLinkedList
}

// AddEdge adds an edge to the graph.
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

// NeighborsOf returns the neighbors of a vertex.
func (g *SimpleGraph) NeighborsOf(vert int) []int {
	if g.vertices == nil {
		return []int{}
	}

	neighs := []int{}
	visited := map[int]bool{}
	var queue []int
	for k := range g.vertices {
		if _, ok := visited[k]; ok {
			continue
		}
		queue := append(queue, k) // initialize queue with first vertex
		for len(queue) > 0 {
			elm := queue[0]                         // dequeue 1st element
			queue = append(queue[:0], queue[1:]...) // remove 1st element to advance queue
			if _, ok := visited[elm]; ok {
				continue
			}
			visited[elm] = true
			root := g.vertices[elm]
			links := root.Values()
			if len(links) <= 0 {
				continue
			}
			links = append(links[:0], links[1:]...)
			switch {
			case vert == elm:
				neighs = append(neighs, links...) // current elm is at the root and has link/descendants
			case contains(links, vert):
				neighs = append(neighs, elm) // vert is a link/descendant of current elm
			}
			queue = append(queue, links...) // enqueue elements to queue
		}
	}
	return neighs
}

// contains returns true if a slice contains an element.
func contains(s []int, elm int) bool {
	for _, v := range s {
		if v == elm {
			return true
		}
	}
	return false
}

// Count returns the number of vertices in the graph.
func (g *SimpleGraph) Count() int {
	if g.vertices == nil {
		return 0
	}
	size := 0
	for _, v := range g.vertices {
		size = v.Size() + size
	}
	return size
}
