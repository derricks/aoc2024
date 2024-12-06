package main

import (
	"math"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/2
type AoC2Sub2Processor struct {
	// keep track for debugging
	lineIsSafe []bool
}

func (p *AoC2Sub2Processor) ProcessLine(line string) error {
	// convert fields to ints
	stringFields := strings.Split(strings.TrimSpace(line), " ")
	intFields := make([]int, len(stringFields))
	for idx, stringValue := range stringFields {
		intValue, err := strconv.Atoi(stringValue)
		if err != nil {
			return err
		}
		intFields[idx] = intValue
	}

	// this is a brute-force solution, but a smarter
	// method would require figuring out exactly which field
	// can be removed to make the whole thing safe, which takes
	// more time than I have to work on this
	if p.levelsAreSafe(intFields) {
		p.lineIsSafe = append(p.lineIsSafe, true)
	} else {
		for idx, _ := range intFields {
			newSlice := deleteNthSliceItem(intFields, idx)
			if p.levelsAreSafe(newSlice) {
				p.lineIsSafe = append(p.lineIsSafe, true)
				return nil
			}
		}
		p.lineIsSafe = append(p.lineIsSafe, false)
	}
	return nil
}

func (p *AoC2Sub2Processor) levelsAreSafe(values []int) bool {

	isAscending := values[1] > values[0]
	isDescending := values[0] > values[1]
	if !isAscending && !isDescending {
		return false
	}

	for idx := 0; idx < len(values)-1; idx++ {
		if values[idx] < values[idx+1] && isDescending {
			return false
		} else if values[idx] > values[idx+1] && isAscending {
			return false
		} else if values[idx] == values[idx+1] {
			return false
		}

		// and finally determine if the diff is in range
		diff := (int)(math.Abs((float64)(values[idx] - values[idx+1])))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func (p *AoC2Sub2Processor) Compute() int {
	count := 0
	for _, safe := range p.lineIsSafe {
		if safe {
			count++
		}
	}
	return count
}

func aoc2Sub2() {
	processor := &AoC2Sub2Processor{[]bool{}}
	generateSolution(processor)
}
