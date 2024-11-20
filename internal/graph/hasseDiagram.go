package graph

import (
	"container/list"
	"sort"
)

// Structure to store Hasse diagram, created to avoid using functions meant for Hasse diagram instead of a normal graph/Diekert graph
type HasseDiagram struct {
	DiekertGraph
}

// Standard DFS
func (g *DiekertGraph) dfs(v vertex, visited map[vertex]bool) {
	visited[v] = true
	for _, neighbor := range g.adjacencyList[v] {
		if !visited[neighbor] {
			g.dfs(neighbor, visited)
		}
	}
}

// Function to initialize a Hasse Diagram
func (graph *DiekertGraph) NewHasseDiagram() HasseDiagram {
	// A Hasse Diagram is just a minimized (reduced) Diekert graph
	reducedGraph := HasseDiagram{
		DiekertGraph{
			Graph{
				adjacencyList: make(map[vertex][]vertex),
				directed:      true,
			},
		},
	}

	// Copy edges from the Diekert graph to a new graph
	vertices := []vertex{}
	for k, v := range graph.adjacencyList {
		reducedGraph.adjacencyList[k] = v
		vertices = append(vertices, k)
	}

	// Map storing the set of vertices that are reachable from each vertex
	reachable := make(map[vertex]map[vertex]bool)

	// For each vertex, use DFS to determine which vertices are reachable
	for _, u := range vertices {
		reachable[u] = make(map[vertex]bool)

		for _, v := range graph.adjacencyList[u] {
			graph.dfs(v, reachable[u])
		}
	}

	// For each edge u -> v, check if the path u -> w -> v exists. If so, remove the u -> v edge
	for _, u := range vertices {
		for _, v := range reducedGraph.adjacencyList[u] {
			for w := range reachable[u] {
				if w != u && w != v && reachable[u][w] && reachable[w][v] {
					reducedGraph.removeEdge(u, v)
				}
			}
		}
	}

	return reducedGraph
}

// Function to extract Foaty Normal Form from the Hasse diagram using BFS
func (g *HasseDiagram) GetFNF() string {
	
	// BFS counting visit time, but without checking if the vertex has been visited
	times := make(map[vertex]int)
	queue := list.New()

	for v := range g.adjacencyList {
		queue.PushBack(v)
		times[v] = 0
	}

	for queue.Len() > 0 {
		currentElement := queue.Front()
		queue.Remove(currentElement)
		current := currentElement.Value.(vertex)

		for _, neighbor := range g.adjacencyList[current] {
			times[neighbor] = times[current] + 1
			queue.PushBack(neighbor)
		}
	}
	// End BFS

	// Grouping vertices by their visit time
	grouped := make(map[int][]string)
	for key, value := range times {
		grouped[value] = append(grouped[value], string(key.name))
	}

	// Creating a readable output
	result := ""
	for i := range len(grouped) {
		sort.Strings(grouped[i])

		result = result + "("

		for _, c := range grouped[i] {
			result = result + c
		}

		result = result + ")"
	}

	return result
}
