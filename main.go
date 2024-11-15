package main

import "fmt"

func main() {
	input := []string{
		"x := x + y",
		"y := y + 2z",
		"x := 3x + z",
		"z := y - z",
	}
	sigma := []rune{'a', 'b', 'c', 'd'}

	var sets Sets
	err := sets.Initialize(input, sigma)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(sets.String())

	sets.GenerateGraph()
}