package main

import (
	"fmt"

	"github.com/Mielecki/Traces/internal/graph"
	"github.com/Mielecki/Traces/internal/sets"
)

func main() {
	input := []string{
		"x := x + y",
		"y := y + 2z",
		"x := 3x + z",
		"z := y - z",
	}
	sigma := []rune{'a', 'b', 'c', 'd'}

	word := "baadcb"

	sets, err := sets.New(input, sigma)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(sets.String())

	g, _ := graph.ParseSets(sets)

	g.ToDot()

	dg, err := g.NewDiekertGraph(word)
	if err != nil {
		return
	}

	fmt.Println(dg)

	dg.ToDot()

	fmt.Println(dg.Minimize())
	dg.Minimize().ToDot()
}