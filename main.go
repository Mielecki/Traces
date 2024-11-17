package main

import (
	"fmt"

	"github.com/Mielecki/Traces/internal/graph"
	"github.com/Mielecki/Traces/internal/sets"
)

func main() {
	input, err := parseFile("./example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}


	sets, err := sets.New(input.Tasks, input.Sigma)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sets.String())

	dependentGraph, _ := graph.ParseSets(sets, true)

	fmt.Println("Graf zaleznosci: \n", dependentGraph.ToDot())

	independentGraph, _ := graph.ParseSets(sets, false)

	fmt.Println("Graf niezaleznosci: \n", independentGraph.ToDot())

	diekertGraph, err := dependentGraph.NewDiekertGraph(input.Word)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Graf Diekerta: \n", diekertGraph.ToDot())

	hasseDiagram := diekertGraph.NewHasseDiagram()

	fmt.Println("Diagram Hassego: \n", hasseDiagram.ToDot())

	fmt.Println("FNF(w) = ", hasseDiagram.GetFNF())
}
