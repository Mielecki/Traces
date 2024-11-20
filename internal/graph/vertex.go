package graph

import "strconv"

// Structure to store vertex information
type vertex struct {
	name  rune // The label of the vertex
	index int  // The index of the vertex (necessary if multiple vertices have the same label)
}

// Function to print vertex in the format "labelindex"
func (v vertex) String() string {
	return string(v.name) + strconv.Itoa(v.index)
}
