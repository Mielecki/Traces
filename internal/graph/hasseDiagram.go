package graph

import (
	"sort"
)


func (g *Graph) dfs(v vertex, visited map[vertex]bool) {
	visited[v] = true
	for _, neighbor := range g.adjacencyList[v] {
		if !visited[neighbor] {
			g.dfs(neighbor, visited)
		}
	}
}

func (graph *Graph) NewHasseDiagram() Graph {
	reducedGraph := Graph{
		adjacencyList: make(map[vertex][]vertex),
		directed: true,
	}

	vertices := []vertex{}
	
	for k, v := range graph.adjacencyList {
		reducedGraph.adjacencyList[k] = v
		vertices = append(vertices, k)
	}

	reachable := make(map[vertex]map[vertex]bool)

	for _, u := range vertices {
		reachable[u] = make(map[vertex]bool)

		for _, v := range graph.adjacencyList[u] {
			graph.dfs(v, reachable[u])
		}
	}

	for _, u := range vertices {
		for _, v := range reducedGraph.adjacencyList[u] {
			for w, _ := range reachable[u] {
				if w != u && w != v && reachable[u][v] && reachable[w][v] {
					reducedGraph.removeEdge(u, v)
				}
			}
		}
	}

	return reducedGraph
}

func (g *Graph) GetFNF() string {
	times := make(map[vertex]int)
	queue := []vertex{}

	for v := range g.adjacencyList {
		queue = append(queue, v)
		times[v] = 0
	}


	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range g.adjacencyList[current] {
			times[neighbor] = times[current] + 1
			queue = append(queue, neighbor)
		}
	}

	grouped := make(map[int][]string)
	for key, value := range times {
		grouped[value] = append(grouped[value], string(key.name))
	}


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