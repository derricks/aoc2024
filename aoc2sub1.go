package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/2
type AoC2Sub1Processor struct {
	// keep track for debugging
	lineIsSafe []bool
}

func (p *AoC2Sub1Processor) ProcessLine(line string) error {
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

	isSortedAsc := slices.IsSortedFunc(intFields, func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	})

	isSortedDesc := slices.IsSortedFunc(intFields, func(a, b int) int {
		if a < b {
			return 1
		} else if a > b {
			return -1
		} else {
			return 0
		}
	})

	isSorted := isSortedAsc || isSortedDesc
	// easy to eliminate
	if !isSorted {
		p.lineIsSafe = append(p.lineIsSafe, false)
		return nil
	}
	// line is sorted
	// somewhat trickier. calculate diff between adjacent values
	for idx := 0; idx < len(intFields)-1; idx++ {
		diff := (int)(math.Abs((float64)(intFields[idx] - intFields[idx+1])))
		if diff < 1 || diff > 3 {
			p.lineIsSafe = append(p.lineIsSafe, false)
			return nil
		}
	}
	p.lineIsSafe = append(p.lineIsSafe, true)

	return nil
}

func (p *AoC2Sub1Processor) Compute() int {
	count := 0
	for _, safe := range p.lineIsSafe {
		if safe {
			count++
		}
	}
	return count
}

func aoc2Sub1() {
	processor := &AoC2Sub1Processor{[]bool{}}
	processStdin(processor)
	fmt.Println(processor.Compute())
}
