package utils

import (
	"os"
	"log"
	"io/ioutil"
	"flag"
	"strings"
)

type PuzzlePart int
const (
	PartOne PuzzlePart = iota
	PartTwo
)

// Struct representing the puzzle input of a specific day.
type PuzzleInput struct {
	Data []byte
	Part PuzzlePart
}

// Puzzle input filename flag. Defaults to "input.txt"
var puzzleFileFlag = flag.String("file", "input.txt", "Filename of text file containing puzzle input.")
var puzzlePart = flag.Int("part", 1, "Part 1 or 2 of the puzzle")

// Return a PuzzleInput struct from a file. Panics on error
func ParseInputFromFile() *PuzzleInput {
	var input = &PuzzleInput{}
	flag.Parse()

	if *puzzlePart == 1 {
		input.Part = PartOne
	} else if *puzzlePart == 2 {
		input.Part = PartTwo
	} else {
		log.Fatal("Enter a valid puzzle part number")
	}

	file, err := os.Open(*puzzleFileFlag)
	if err != nil {
		log.Fatal(err)
	}

	input.Data, err = ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return input
}

// Parses puzzle input as array of strings, separated by a delimiter
func (input *PuzzleInput) ParseStringArrayDelimiter(delimiter string) []string {
	return strings.Split(string(input.Data), delimiter)
}

// Parses puzzle input as array of strings separated by newlines
func (input *PuzzleInput) ParseStringArrayNewlines() []string {
	return input.ParseStringArrayDelimiter("\n")
}
