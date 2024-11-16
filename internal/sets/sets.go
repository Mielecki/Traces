package sets

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type dataItem struct {
	symbol rune
	modified rune
	read []rune
}

type Pair struct {
	First rune
	Second rune
}

type Sets struct {
	Sigma []rune
	Data []dataItem
	Dependent map[Pair]bool
	Independent map[Pair]bool
}

func New(input []string, sigma []rune) (Sets, error) {
	newSets := Sets{}
	
	newSets.Sigma = make([]rune, 0)
	newSets.Sigma = append(newSets.Sigma, sigma...)

	if err := newSets.parseInput(input); err != nil {
		return Sets{},err
	}

	if err := newSets.createSets(); err != nil {
		return Sets{},err
	}

	return newSets, nil
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
			symbol: sets.Sigma[i],
			modified: parsedItem[0],
			read: parsedItem[1:],
		})
	}

	sets.Data = parsedInput
	return nil
}

func (sets *Sets) createSets() error {
	dependent := make(map[Pair]bool)
	independent := make(map[Pair]bool)

	for i := 0; i < len(sets.Data); i++ {
		itemI := sets.Data[i]

		dependent[Pair{itemI.symbol, itemI.symbol}] = true

		for j := i + 1; j < len(sets.Data); j++ {
			itemJ := sets.Data[j]

			if slices.Contains(itemJ.read, itemI.modified) || slices.Contains(itemI.read, itemJ.modified) {
				dependent[Pair{itemI.symbol, itemJ.symbol}] = true
				dependent[Pair{itemJ.symbol, itemI.symbol}] = true
			} else {
				independent[Pair{itemI.symbol, itemJ.symbol}] = true
				independent[Pair{itemJ.symbol, itemI.symbol}] = true
			}
		}
	}

	sets.Independent = independent
	sets.Dependent = dependent
	return nil
}

func (sets *Sets) String() string {
	dependentStr := ""
	for p := range sets.Dependent {
		dependentStr += fmt.Sprintf("(%c, %c) ", p.First, p.Second)
	}

	independentStr := ""
	for p := range sets.Independent {
		independentStr += fmt.Sprintf("(%c, %c) ", p.First, p.Second)
	}

	return fmt.Sprintf("D = {%s}\nI = {%s}", dependentStr, independentStr)
}