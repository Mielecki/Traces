package main

import (
	"fmt"

	"github.com/Mielecki/Traces/internal/graph"
	"github.com/Mielecki/Traces/internal/sets"
)

func main() {
	// Parsing input file
	input, err := parseFile("./example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Creating a Sets structure to store dependency/independency sets
	sets, err := sets.New(input.Tasks, input.Sigma)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Printing dependency and independency sets
	fmt.Println(sets.String())

	// Creating a dependency graph from the set of dependent tasks
	dependencyGraph := graph.ParseSets(sets.Dependent)

	// Printing the dependency graph in a .dot format for Graphviz visualization
	fmt.Println("Graf zaleznosci: \n", dependencyGraph.ToDot())

	// Creating an independence graph from the set of independent tasks
	independenceGraph := graph.ParseSets(sets.Independent)

	// Printing the independence graph in a .dot format for Graphviz visualization
	fmt.Println("Graf niezaleznosci: \n", independenceGraph.ToDot())

	// Creating a DiekertGraph structure to store unminimized dependency graph for the given word 
	diekertGraph := dependencyGraph.NewDiekertGraph(input.Word)

	// Printing the Diekert graph in a .dot format for Graphviz visualization
	fmt.Println("Graf Diekerta: \n", diekertGraph.ToDot())

	// Creating a HasseDiagram structure to store minimized dependency graph for the given word
	hasseDiagram := diekertGraph.NewHasseDiagram()

	// Printing the Hasse Diagram in a .dot format for Graphviz visualization
	fmt.Println("Diagram Hassego: \n", hasseDiagram.ToDot())

	// Printing the Foaty Normal Form obtained from Hasse Diagram
	fmt.Println("FNF(w) = ", hasseDiagram.GetFNF())
}
