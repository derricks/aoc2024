package main

import (
	"fmt"
	"strings"
)

type GuardDirection int

const (
	GuardFacingUp    GuardDirection = 1
	GuardFacingDown  GuardDirection = 2
	GuardFacingLeft  GuardDirection = 3
	GuardFacingRight GuardDirection = 4
)

type GuardAdvanceResult int

const (
	SuccessfulAdvance GuardAdvanceResult = 1
	AdvanceBlocked    GuardAdvanceResult = 2
	OffGrid           GuardAdvanceResult = 3
)

type AoC6Sub1Processor struct {
	layout               [][]string
	currentCoordinates   []int // row, column
	currentDirection     GuardDirection
	currentProcessedLine int // so we can track the row of the guard
}

func (p *AoC6Sub1Processor) ProcessLine(line string) error {
	p.currentProcessedLine++
	// if line contains guard, note coordinates (regex will tell you position)
	row := strings.Split(line, "")

	// once we find the guard, we don't have to look any more
	if len(p.currentCoordinates) == 0 {
		for idx, occupant := range row {
			if occupant == "." || occupant == "#" {
				continue
			}

			// found something
			if occupant == "^" {
				p.currentDirection = GuardFacingUp
			}

			if occupant == ">" {
				p.currentDirection = GuardFacingRight
			}

			if occupant == "<" {
				p.currentDirection = GuardFacingLeft
			}

			if occupant == "v" {
				p.currentDirection = GuardFacingDown
			}
			p.currentCoordinates = []int{p.currentProcessedLine, idx}

			// now that we've recorded the direction, set that spot to a .
			// so down below looks fine
			row[idx] = "."
			break
		}
	}
	// done here so it has the edit in place
	p.layout = append(p.layout, row)
	return nil
}

func (p *AoC6Sub1Processor) Compute() int {
	occupiedSpots := 0 // starting point is counted
	// keep track of what was visited ([]int can't be a map key)
	visitedSpots := make([][]int, len(p.layout))
	for idx, _ := range visitedSpots {
		visitedSpots[idx] = make([]int, len(p.layout[0]))
	}

	advanceResult := p.advanceGuard()
	for advanceResult != OffGrid {
		if advanceResult == SuccessfulAdvance {
			if visitedSpots[p.currentCoordinates[0]][p.currentCoordinates[1]] == 0 {
				occupiedSpots++
				visitedSpots[p.currentCoordinates[0]][p.currentCoordinates[1]] = 1
			}
		}

		if advanceResult == AdvanceBlocked {
			// change direction
			switch p.currentDirection {
			case GuardFacingUp:
				p.currentDirection = GuardFacingRight
			case GuardFacingRight:
				p.currentDirection = GuardFacingDown
			case GuardFacingDown:
				p.currentDirection = GuardFacingLeft
			case GuardFacingLeft:
				p.currentDirection = GuardFacingUp
			}
		}
		advanceResult = p.advanceGuard()
	}

	return occupiedSpots
}

func (p *AoC6Sub1Processor) advanceGuard() GuardAdvanceResult {
	nextCoords := p.nextLocationForGuard()
	if nextCoords[0] < 0 || nextCoords[0] >= len(p.layout) ||
		nextCoords[1] < 0 || nextCoords[1] >= len(p.layout[0]) {
		return OffGrid
	}

	if p.layout[nextCoords[0]][nextCoords[1]] != "." {
		return AdvanceBlocked
	}

	p.currentCoordinates = nextCoords
	return SuccessfulAdvance
}

func (p *AoC6Sub1Processor) nextLocationForGuard() []int {
	switch p.currentDirection {
	case GuardFacingUp:
		return []int{p.currentCoordinates[0] - 1, p.currentCoordinates[1]}
	case GuardFacingDown:
		return []int{p.currentCoordinates[0] + 1, p.currentCoordinates[1]}
	case GuardFacingLeft:
		return []int{p.currentCoordinates[0], p.currentCoordinates[1] - 1}
	case GuardFacingRight:
		return []int{p.currentCoordinates[0], p.currentCoordinates[1] + 1}
	}
	fmt.Printf("Shouldn't get here: %v\n", p.currentCoordinates)
	return []int{}
}

func aoc6sub1() {
	processor := &AoC6Sub1Processor{[][]string{}, []int{}, GuardFacingUp, -1}
	generateSolution(processor)
}
