package main

import (
	"fmt"

	"github.com/Mielecki/Traces/internal/graph"
	"github.com/Mielecki/Traces/internal/sets"
)

func main() {
	// Path where the files will be saved
	const path = "./output/"

	// Creating the output folder if does not exist
	err := checkIfDirExists(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parsing input file
	input, err := parseFile("./example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Creating a Sets structure to store dependency/independence sets
	sets, err := sets.New(input.Tasks, input.Sigma)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Printing dependency and independency sets
	fmt.Println(sets.String())

	// Creating a dependency graph from the set of dependent tasks
	dependencyGraph := graph.ParseSets(sets.Dependent)

	// Save the dependency graph in a .dot format for Graphviz visualization
	err = writeFile(dependencyGraph.ToDot(), path + "graph_dependency.gv")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Creating an independence graph from the set of independent tasks
	independenceGraph := graph.ParseSets(sets.Independent)

	// Save the independence graph in a .dot format for Graphviz visualization
	err = writeFile(independenceGraph.ToDot(), path + "graph_independence.gv")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Creating a DiekertGraph structure to store unminimized dependency graph for the given word 
	diekertGraph := dependencyGraph.NewDiekertGraph(input.Word)

	// Save the Diekert graph in a .dot format for Graphviz visualization
	err = writeFile(diekertGraph.ToDot(), path + "graph_diekert.gv")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Creating a HasseDiagram structure to store minimized dependency graph for the given word
	hasseDiagram := diekertGraph.NewHasseDiagram()

	// Save the Hasse Diagram in a .dot format for Graphviz visualization
	err = writeFile(hasseDiagram.ToDot(), path + "graph_hasse.gv")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Printing the Foaty Normal Form obtained from Hasse Diagram
	fmt.Println("FNF(w) = ", hasseDiagram.GetFNF())
}
