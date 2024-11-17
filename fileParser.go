package main

import (
	"bufio"
	"errors"
	"os"
)

type Input struct {
	Tasks []string
	Sigma []rune
	Word string
}

func parseFile(path string) (Input, error) {
	file, err := os.Open(path)
	if err != nil {
		return Input{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	state := 0
	input := Input{
		Tasks: make([]string, 0),
		Sigma: make([]rune, 0),
		Word: "",
	}

	for scanner.Scan() {
		line := scanner.Text()

		switch state{
		case 0:
			if line == "" {
				state = 1
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
				state = -1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return Input{}, err
	}

	if len(input.Sigma) != len(input.Tasks) {
		return Input{}, errors.New("invalid input format: the number of tasks must be equal to the sigma size")
	}

	sigmaSet := make(map[rune]struct{})
	for _, char := range input.Sigma {
		sigmaSet[char] = struct{}{}
	}

	for _, char := range input.Word {
		if _, exists := sigmaSet[char]; !exists {
			return Input{}, errors.New("invalid input format: the word must contain only characters from sigma set")
		}
	}

	return input, nil
}

func checkLetter(c byte) bool {
	return rune(c) >= 97 && rune(c) <= 122
}