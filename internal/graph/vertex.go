package graph

import "strconv"

type vertex struct {
	name rune
	index int
}

func (v vertex) String() string {
	return string(v.name) + strconv.Itoa(v.index)
}