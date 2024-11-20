package sets

import (
	"fmt"
	"slices"
	"strings"
)

// Structure to store a single task data
type dataItem struct {
	symbol   rune	// The symbol representing the task from the sigma set
	modified rune   // The variable that the task modifies
	read     []rune // Variables that are read by the task
}

// Structure to store represent a relation between two tasks
type Pair struct {
	First  rune
	Second rune
}

// Structure to store on which other algorithms will work
type Sets struct {
	Sigma       []rune            // The sigma set
	Data        []dataItem        // A slice of task data
	Dependent   map[Pair]struct{} // The set of depenedent relations
	Independent map[Pair]struct{} // The set of independent relations
}

// Function to initalize a new Sets structure
func New(input []string, sigma []rune) (Sets, error) {
	newSets := Sets{}

	// Fill Sigma field in the structure
	newSets.Sigma = make([]rune, 0)
	newSets.Sigma = append(newSets.Sigma, sigma...)

	// Fill Data field in the structure
	if err := newSets.parseInput(input); err != nil {
		return Sets{}, err
	}

	// Fill Dependent and Independent fields in the structure
	if err := newSets.createSets(); err != nil {
		return Sets{}, err
	}

	return newSets, nil
}

// Function to parse a list of tasks in string format to a list of tasks in the dataItem structure format
func (sets *Sets) parseInput(input []string) error {
	var parsedInput []dataItem

	// Iterate over each task
	for i, item := range input {
		parsedItem := []rune{}

		// Transform each task into a list of runes containing only variables (e.g. "x := y - z" -> "xyz")
		for _, ch := range item {
			if !strings.ContainsRune("+-:=1234567890 ", ch) {
				parsedItem = append(parsedItem, ch)
			}
		}

		parsedInput = append(parsedInput, dataItem{
			symbol:   sets.Sigma[i],  // the i-th task corresponds to the i-th symbol in the sigma set
			modified: parsedItem[0],  // the first variable in parsedItem is the one that is modified
			read:     parsedItem[1:], // the remaining variables are read by the task
		})
	}

	sets.Data = parsedInput
	return nil
}

// Function to create Dependent and Independent sets from the Data field
func (sets *Sets) createSets() error {
	dependent := make(map[Pair]struct{})
	independent := make(map[Pair]struct{})

	for i := 0; i < len(sets.Data); i++ {
		itemI := sets.Data[i] // the i-th item from the Data field

		dependent[Pair{itemI.symbol, itemI.symbol}] = struct{}{}

		for j := i + 1; j < len(sets.Data); j++ {
			itemJ := sets.Data[j] // the j-th item from the Data field
			
			// If the one task modifies a variable that another reads, then they are dependent
			if slices.Contains(itemJ.read, itemI.modified) || slices.Contains(itemI.read, itemJ.modified) {
				dependent[Pair{itemI.symbol, itemJ.symbol}] = struct{}{}
				dependent[Pair{itemJ.symbol, itemI.symbol}] = struct{}{}
			} else { // In the other case, they are independent
				independent[Pair{itemI.symbol, itemJ.symbol}] = struct{}{}
				independent[Pair{itemJ.symbol, itemI.symbol}] = struct{}{}
			}
		}
	}

	sets.Independent = independent
	sets.Dependent = dependent
	return nil
}

// Function to print Dependency and Independence sets
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
