package graph


func (g *Graph) dfs(v vertex, visited map[vertex]bool) {
	visited[v] = true
	for _, neighbor := range g.adjacencyList[v] {
		if !visited[neighbor] {
			g.dfs(neighbor, visited)
		}
	}
}

func (graph *Graph) Minimize() Graph {
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