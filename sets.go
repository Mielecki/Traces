package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type dataItem struct {
	symbol rune
	modified rune
	read []rune
}

type pair struct {
	first rune
	second rune
}

type Sets struct {
	sigma []rune
	data []dataItem
	dependent map[pair]bool
	independent map[pair]bool
}

func (sets *Sets) Initialize(input []string, sigma []rune) error {
	sets.sigma = make([]rune, 0)
	sets.sigma = append(sets.sigma, sigma...)

	if err := sets.parseInput(input); err != nil {
		return err
	}

	if err := sets.createSets(); err != nil {
		return err
	}

	return nil
}

func (sets *Sets) parseInput(input []string) error {
	var parsedInput []dataItem

	for i, item := range input {
		parsedItem := []rune{}

		for _, ch := range item {
			if !strings.ContainsRune("+-:=1234567890", ch) {
				parsedItem = append(parsedItem, ch)
			}
		}

		if len(parsedItem) < 2 {
			return errors.New("invalid task format")
		}

		parsedInput = append(parsedInput, dataItem {
			symbol: sets.sigma[i],
			modified: parsedItem[0],
			read: parsedItem[1:],
		})
	}

	sets.data = parsedInput
	return nil
}

func (sets *Sets) createSets() error {
	dependent := make(map[pair]bool)
	independent := make(map[pair]bool)

	for i := 0; i < len(sets.data); i++ {
		itemI := sets.data[i]

		dependent[pair{itemI.symbol, itemI.symbol}] = true

		for j := i + 1; j < len(sets.data); j++ {
			itemJ := sets.data[j]

			if contains(itemJ.read, itemI.modified) || contains(itemI.read, itemJ.modified) {
				dependent[pair{itemI.symbol, itemJ.symbol}] = true
				dependent[pair{itemJ.symbol, itemI.symbol}] = true
			} else {
				independent[pair{itemI.symbol, itemJ.symbol}] = true
				independent[pair{itemJ.symbol, itemI.symbol}] = true
			}
		}
	}

	sets.independent = independent
	sets.dependent = dependent
	return nil
}

func (sets *Sets) String() string {
	dependentStr := ""
	for p := range sets.dependent {
		dependentStr += fmt.Sprintf("(%c, %c) ", p.first, p.second)
	}

	independentStr := ""
	for p := range sets.independent {
		independentStr += fmt.Sprintf("(%c, %c) ", p.first, p.second)
	}

	return fmt.Sprintf("D = {%s}\nI = {%s}", dependentStr, independentStr)
}

func contains(slice []rune, target rune) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func removeDuplicatePairs(m map[pair]bool) map[pair]bool {
	processed := make(map[pair]bool)
	seen := make(map[pair]bool)

	for p := range m {
		reverse := pair{p.second, p.first}
		if p.second != p.first && !seen[reverse] {
			processed[p] = true
			seen[p] = true
		}
	}

	return processed
}

func (sets *Sets) GenerateGraph() error {
	depEdges := removeDuplicatePairs(sets.dependent)
	indepEdges := removeDuplicatePairs(sets.independent)

	g := graph.New(graph.StringHash)

	for _, symbol := range sets.sigma {
		g.AddVertex(string(symbol))
	}
	
	for edge, _ := range depEdges {
		if err := g.AddEdge(string(edge.first), string(edge.second), graph.EdgeAttribute("color", "red")); err != nil {
			return err
		}
	}

	for edge, _ := range indepEdges {
		if err := g.AddEdge(string(edge.first), string(edge.second), graph.EdgeAttribute("color", "green")); err != nil {
			return err
		}
	}

	file, _ := os.Create("./graph.gv")
	_ = draw.DOT(g, file)	

	return nil
}