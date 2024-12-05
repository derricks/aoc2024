package main

import (
	"fmt"
	"strconv"
	"strings"
)

type AoC1Sub2Processor struct {
	// keeps track of how many of each value was found in the right
	rightCounts map[int]int
	// keep track of the lefts
	leftList []int
}

func (p *AoC1Sub2Processor) ProcessLine(line string) error {
	fields := strings.Split(line, "   ")
	leftValue, err := strconv.Atoi(fields[0])
	if err != nil {
		return err
	}
	rightValue, err := strconv.Atoi(fields[1])
	if err != nil {
		return err
	}

	p.leftList = append(p.leftList, leftValue)

	if _, found := p.rightCounts[rightValue]; !found {
		p.rightCounts[rightValue] = 0
	}
	p.rightCounts[rightValue] += 1
	return nil
}

func (p *AoC1Sub2Processor) Compute() int {
	runningTotal := 0
	for _, left := range p.leftList {
		if count, found := p.rightCounts[left]; found {
			runningTotal += (left * count)
		}
	}
	return runningTotal
}

// https://adventofcode.com/2024/day/1#part2
func aoc1Sub2() {
	processor := &AoC1Sub2Processor{make(map[int]int), []int{}}
	processStdin(processor)
	fmt.Println(processor.Compute())
}
