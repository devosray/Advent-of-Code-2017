package utils

import (
	"strings"
	"strconv"
	"log"
)

// Converts a string to array of integers separated by a delimiter.
// No delimiter = no space between numbers. "1234" -> [1, 2, 3, 4]
func StringToIntegerArray(input, delimiter string) []int {
	intArray := []int{}

	split := strings.Split(input, delimiter)
	for _, str := range split {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("Error: Could not convert `%s` to int. Ignoring value.", str)
			continue
		}
		intArray = append(intArray, num)
	}

	return intArray
}
