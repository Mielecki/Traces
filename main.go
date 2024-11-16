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

	gt, _ := graph.ParseSets(sets)

	fmt.Println("graf zaleznosci: \n", gt.ToDot())

	dg, err := gt.NewDiekertGraph(word)
	if err != nil {
		return
	}

	fmt.Println("graf diekerta: \n", dg.ToDot())

	fmt.Println("hassego diagram: \n", dg.Minimize().ToDot())
}