package main

import (
	"flag"
	"log"
	"fmt"
)

var puzzleInputFlag = flag.String("input", "", "Puzzle Input")
var puzzlePart = flag.Int("part", 1, "Part 1 or 2 of the puzzle")

func main() {
	flag.Parse()
	if len(*puzzleInputFlag) == 0 {
		log.Fatal("No puzzle input received. Pass puzzle input using the `--input` program flag.")
	}

	var input = *puzzleInputFlag
	var checksum = 0

	for i, character := range input {
		currentNum := int(character - '0')
		var comparable int = 0

		if *puzzlePart == 1 {
			// First part of puzzle -> compare current number with previous number.
			// Input's length is added and then modded to ensure that position -1 wraps around to the end.
			comparable = int(input[(i-1+len(input))%len(input)] - '0')
		} else {
			// Second part of puzzle -> compare current number to number 1/2 of input length away.
			comparable = int(input[(i+len(input)/2)%len(input)] - '0')
		}

		if currentNum == comparable {
			checksum += currentNum
		}
	}

	fmt.Printf("Checksum: %d\n", checksum)
}
