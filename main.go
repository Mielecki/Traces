package main

import (
	"fmt"

	"github.com/Mielecki/Traces/internal/graph"
	"github.com/Mielecki/Traces/internal/sets"
)

func main() {
	input := []string{
		"x := x + 1",
		"y := y + 2z",
		"x := 3x + z",
		"w := w + v",
		"z := y − z",
		"v := x + v",
	}
	sigma := []rune{'a', 'b', 'c', 'd', 'e', 'f'}

	word := "acdcfbbe"

	sets, err := sets.New(input, sigma)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(sets.String())

	dependentGraph, _ := graph.ParseSets(sets, true)

	fmt.Println("Graf zaleznosci: \n", dependentGraph.ToDot())

	independentGraph, _ := graph.ParseSets(sets, false)

	fmt.Println("Graf niezaleznosci: \n", independentGraph.ToDot())

	diekertGraph, err := dependentGraph.NewDiekertGraph(word)
	if err != nil {
		return
	}

	fmt.Println("Graf Diekerta: \n", diekertGraph.ToDot())

	hasseDiagram := diekertGraph.NewHasseDiagram()

	fmt.Println("Diagram Hassego: \n", hasseDiagram.ToDot())

	fmt.Println("FNF(w) = ", hasseDiagram.GetFNF())
}
