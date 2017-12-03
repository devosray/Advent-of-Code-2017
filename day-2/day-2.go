package main

import (
	"github.com/devosray/Advent-of-Code-2017/utils"
	"fmt"
)

func main() {
	input := utils.ParseInputFromFile()
	inputLines := input.ParseStringArrayNewlines()
	var checksum int = 0

	for _, line := range inputLines {
		numArray := utils.StringToIntegerArray(line, "\t")

		// Needed for part 1
		largest := numArray[0]
		smallest := numArray[0]

		// Needed for part 2
		var toAdd int = 0

		for i, numOne := range numArray {
			switch input.Part {
			case utils.PartOne:
				if numOne > largest {
					largest = numOne
				}
				if numOne < smallest {
					smallest = numOne
				}

			case utils.PartTwo:
				for j, numTwo := range numArray {
					if i != j && numOne > numTwo && numOne%numTwo == 0 {
						toAdd = numOne / numTwo
					}
				}
			}
		}

		// Add result to checksum
		switch input.Part {
		case utils.PartOne:
			checksum += largest - smallest
		case utils.PartTwo:
			checksum += toAdd
		}
	}

	fmt.Printf("Checksum: %d\n", checksum)
}
