package graph

import "slices"

// Structure to store Diekert graph, created to avoid using functions meant for Diekert graph instead of a normal graph
type DiekertGraph struct {
	Graph
}

// Function to initialize a Diekert Graph for the given word
func (graph Graph) NewDiekertGraph(wordInput string) DiekertGraph {
	diekertGraph := DiekertGraph{
		Graph{
			adjacencyList: make(map[vertex][]vertex),
			directed:      true,
		},
	}

	word := []rune(wordInput)

	// Map storing current number of occurrences of a symbol
	indices := make(map[rune]int)
	for v := range graph.adjacencyList {
		indices[v.name] = -1
	}

	// Iterate over characters in the word and create a list of vertices with their indices 
	vertices := []vertex{}
	for _, c := range word {
		indices[c] += 1
		vertices = append(vertices, vertex{
			name:  c,
			index: indices[c],
		})
	}

	// Iterate over all possible edges, and if the edge exists in the dependency graph, append it to the Diekert graph
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

	return diekertGraph
}
