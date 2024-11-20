package main

import (
	"bufio"
	"errors"
	"os"
)

// Structure to store parsed informations from the input file
type Input struct {
	Tasks []string  // List of tasks
	Sigma []rune    // List of task symbols
	Word  string    // The word (sequence of task symbols) on which algorithms will operate
}

// Function to parse input file
func parseFile(path string) (Input, error) {
	// Open file at the given path
	file, err := os.Open(path)
	if err != nil {
		return Input{}, err
	}
	defer file.Close() // defer closing file

	scanner := bufio.NewScanner(file)

	input := Input{
		Tasks: make([]string, 0),
		Sigma: make([]rune, 0),
		Word:  "",
	}

	// iterate over every line in the given input file and do something based on the state variable
	// state 0 -> adding tasks
	// state 1 -> adding the sigma set
	// state 2 -> adding the word
	state := 0
	for scanner.Scan() {
		line := scanner.Text()

		switch state {
		case 0:
			if line == "" {
				if len(input.Tasks) != 0{
					state = 1
				}
			} else if checkLetter(line[0]) {
				input.Tasks = append(input.Tasks, line)
			}
		case 1:
			if line != "" && line[0] == 'A' {
				for i := range line {
					if checkLetter(line[i]) {
						input.Sigma = append(input.Sigma, rune(line[i]))
					}
				}
				state = 2
			}
		case 2:
			if line != "" && line[0] == 'w' {
				input.Word = line[4:]
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return Input{}, err
	}

	if len(input.Sigma) != len(input.Tasks) {
		return Input{}, errors.New("invalid input format: the number of tasks must be equal to the sigma size")
	}

	// Creating a set of symbols for fast lookup. Go doesn't have build-in set datastructure, so map with struct{} is used to emulate a set
	sigmaSet := make(map[rune]struct{})
	for _, char := range input.Sigma {
		sigmaSet[char] = struct{}{}
	}

	// Checking if the word contains only characters from the sigma set
	for _, char := range input.Word {
		if _, exists := sigmaSet[char]; !exists {
			return Input{}, errors.New("invalid input format: the word must contain only characters from sigma set")
		}
	}

	return input, nil
}

// Function to check if the character is a lowercase letter
func checkLetter(c byte) bool {
	return rune(c) >= 97 && rune(c) <= 122
}
