package main

import (
	"github.com/devosray/Advent-of-Code-2017/utils"
	"fmt"
)

type partOneSquare struct {
	Number      int
	X, Y        int
	CurrentRing int
}

func (square *partOneSquare) move(stepX, stepY int) {
	square.X += stepX
	square.Y += stepY
	square.Number++
}

func (square *partOneSquare) getStepsToCentre() int {
	fmt.Printf("%d, %d\n", square.X, square.Y)
	return utils.IntAbs(square.X) + utils.IntAbs(square.Y)
}

func main() {
	input := utils.ParseInputFromFile()
	endingSquare := input.ParseAsInt()

	var steps int = 0

	switch input.Part {
	case utils.PartOne:
		square := partOneSquare{
			1, 0, 0, 0,
		}

		// Algorithm for iterating through ring
		for square.Number < endingSquare {
			square.CurrentRing++

			/*
			TEMPORARY: not proud of this
			 */

			// Move one step to the right
			square.move(1, 0)
			if square.Number == endingSquare {
				steps = square.getStepsToCentre()
			}

			// Move (ring. no - 1) * 2 + 1 steps up
			for i := 0; i < (square.CurrentRing - 1) * 2 + 1; i++ {
				square.move(0, 1)
				if square.Number == endingSquare {
					steps = square.getStepsToCentre()
				}
			}

			// Move ring. no * 2 + 1 steps to the left
			for i := 0; i < (square.CurrentRing - 1) * 2 + 2; i++ {
				square.move(-1, 0)
				if square.Number == endingSquare {
					steps = square.getStepsToCentre()
				}
			}

			// Move ring. no * 2 + 1 steps down
			for i := 0; i < (square.CurrentRing - 1) * 2 + 2; i++ {
				square.move(0, -1)
				if square.Number == endingSquare {
					steps = square.getStepsToCentre()
				}
			}

			// Move ring. no * 2 + 1 steps to the right
			for i := 0; i < (square.CurrentRing - 1) * 2 + 2; i++ {
				square.move(1, 0)
				if square.Number == endingSquare {
					steps = square.getStepsToCentre()
				}
			}
		}

	}

	fmt.Printf("Steps needed: %d\n", steps)
}
