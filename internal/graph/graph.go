package graph

import (
	"fmt"
	"slices"

	"github.com/Mielecki/Traces/internal/sets"
)

type Graph struct {
	adjacencyList map[vertex][]vertex
	directed bool
}

func ParseSets(sets sets.Sets) (Graph, error) {
	newGraph := Graph{
		adjacencyList: make(map[vertex][]vertex),
		directed: false,
	}

	for pair := range sets.Dependent {
		v1 := vertex{
			name: pair.First,
			index: 0,
		}
		v2 := vertex{
			name: pair.Second,
			index: 0,
		}
		newGraph.adjacencyList[v1] = append(newGraph.adjacencyList[v1], v2)
	}

	return newGraph, nil
}

func (graph Graph) ToDot() error {
	relation := "--"
	keyword := "graph"
	
	if graph.directed {
		relation = "->"
		keyword = "digraph"
	}

	result := keyword + " graphname {\n"

	for vertex := range graph.adjacencyList {
		result = result + vertex.String() + " [label=\"" + string(vertex.name) + "\"]\n"
	}

	for vertex1, list := range graph.adjacencyList {
		for _, vertex2 := range list {
			if !graph.directed && vertex2.name > vertex1.name {
				result = result + vertex1.String() + " " + relation + " " + vertex2.String() + ";\n"
			}
			if graph.directed {
				result = result + vertex1.String() + " " + relation + " " + vertex2.String() + ";\n"
			}
		}
	}

	result = result + "}"

	fmt.Println(result)

	return nil
}

func (graph Graph) NewDiekertGraph(wordInput string) (Graph, error) {
	diekertGraph := Graph{
		adjacencyList: make(map[vertex][]vertex),
		directed: true,
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
			name: c,
			index: indices[c],
		})
	}

	for i := 0; i < len(vertices); i++ {
		v1 := vertices[i]
		diekertGraph.adjacencyList[v1] = make([]vertex, 0)
		for j := i + 1; j < len(vertices); j++ {
			v2 := vertices[j]

			if slices.Contains(graph.adjacencyList[
				vertex{
				name: v1.name,
				index: 0,
			}], vertex{
				name: v2.name,
				index: 0,
			}) {
				diekertGraph.adjacencyList[v1] = append(diekertGraph.adjacencyList[v1], v2)
			}
		}
	}

	return diekertGraph, nil
}