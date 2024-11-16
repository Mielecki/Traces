package graph

import "slices"

type DiekertGraph struct {
	Graph
}

func (graph Graph) NewDiekertGraph(wordInput string) (DiekertGraph, error) {
	diekertGraph := DiekertGraph{
		Graph{
			adjacencyList: make(map[vertex][]vertex),
			directed:      true,
		},
	}

	word := []rune(wordInput)
	vertices := []vertex{}

	indices := make(map[rune]int)
	for v := range graph.adjacencyList {
		indices[v.name] = -1
	}

	for _, c := range word {
		indices[c] += 1
		vertices = append(vertices, vertex{
			name:  c,
			index: indices[c],
		})
	}

	for i := 0; i < len(vertices); i++ {
		v1 := vertices[i]
		diekertGraph.adjacencyList[v1] = make([]vertex, 0)
		for j := i + 1; j < len(vertices); j++ {
			v2 := vertices[j]

			if slices.Contains(graph.adjacencyList[vertex{
				name:  v1.name,
				index: 0,
			}], vertex{
				name:  v2.name,
				index: 0,
			}) {
				diekertGraph.adjacencyList[v1] = append(diekertGraph.adjacencyList[v1], v2)
			}
		}
	}

	return diekertGraph, nil
}
