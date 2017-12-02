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
		largest := numArray[0]
		smallest := numArray[0]
		for _, num := range numArray {
			if num > largest {
				largest = num
			}
			if num < smallest {
				smallest = num
			}
		}
		checksum += largest - smallest
	}

	fmt.Printf("Checksum: %d\n", checksum)
}
