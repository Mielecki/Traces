package graph

import (
	"github.com/Mielecki/Traces/internal/sets"
)

// Structure to store graph using adjacency lists
type Graph struct {
	adjacencyList map[vertex][]vertex
	directed      bool
}

// Function to initialize a new graph by parsing dependency/independence set
func ParseSets(set map[sets.Pair]struct{}) Graph {
	newGraph := Graph{
		adjacencyList: make(map[vertex][]vertex),
		directed:      false, // The dependency/independence graph is undirected
	}

	// Adding edges to the adjacency list based on the given set
	for pair := range set {
		v1 := vertex{
			name:  pair.First,
			index: 0,
		}
		v2 := vertex{
			name:  pair.Second,
			index: 0,
		}
		newGraph.adjacencyList[v1] = append(newGraph.adjacencyList[v1], v2)
	}

	return newGraph
}

// Function to parse Graph structure to .dot format for Graphviz visualization
func (graph Graph) ToDot() string {
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

	return result
}

// Function to remove an edge from the graph
func (graph *Graph) removeEdge(start vertex, end vertex) {
	for i, v := range graph.adjacencyList[start] {
		if v == end {
			graph.adjacencyList[start] = append(graph.adjacencyList[start][:i], graph.adjacencyList[start][i+1:]...)
			return
		}
	}
}
