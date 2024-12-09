package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type AoC5Sub2Processor struct {
	// for reuse of tested methods
	sub1Processor      *AoC5Sub1Processor
	invalidLines       []string
	beforeRequirements map[string][]string
}

func (p *AoC5Sub2Processor) ProcessLine(line string) error {
	// if it's a page ordering line, pass it through
	if PAGE_ORDERING_REGEX.MatchString(line) {
		// capture the before requirements
		instructions := strings.Split(line, "|")
		before, after := instructions[0], instructions[1]
		if _, found := p.beforeRequirements[after]; found {
			p.beforeRequirements[after] = append(p.beforeRequirements[after], before)
		} else {
			p.beforeRequirements[after] = []string{before}
		}
		return p.sub1Processor.ProcessLine(line)
	}
	// only keep the _invalid_ lines because that's all we care about
	// this is somewhat inefficient, because we'll split here and then again
	// when we compute, but not a big deal in the grand scheme of things
	if !p.sub1Processor.isUpdateValid(strings.Split(line, ",")) {
		p.invalidLines = append(p.invalidLines, line)
	}

	return nil
}

type updateOrderNode struct {
	value       string
	downstreams []updateOrderNode
}

// corrects a set of fields in accordance with the ordering rules
func (p *AoC5Sub2Processor) fixFields(fields []string) []string {
	slices.SortFunc(fields, func(a, b string) int {
		// find either a or b
		if afters, foundBefore := p.sub1Processor.afterRequirements[a]; foundBefore {
			// if a has instructions for b being after, then return 1
			if slices.Contains(afters, b) {
				return -1
			}
		}
		if befores, foundAfter := p.beforeRequirements[a]; foundAfter {
			// b should be before a
			if slices.Contains(befores, b) {
				return 1
			}

		}
		// and then cover the other case where b has the info
		if afters, foundBefore := p.sub1Processor.afterRequirements[b]; foundBefore {
			if slices.Contains(afters, a) {
				// a should be after b
				return 1
			}
		}
		if befores, foundAfter := p.beforeRequirements[b]; foundAfter {
			if slices.Contains(befores, a) {
				// ashould be before b
				return -1
			}
		}

		fmt.Printf("Error comparing %v, %v\n", a, b)
		return 0

	})
	return fields
}

func (p *AoC5Sub2Processor) Compute() int {
	runningTotal := 0
	for _, invalidLine := range p.invalidLines {
		newLine := p.fixFields(strings.Split(invalidLine, ","))
		middleValue, _ := strconv.Atoi(newLine[len(newLine)/2])
		runningTotal += middleValue
	}
	return runningTotal
}

func aoc5sub2() {
	processor := &AoC5Sub2Processor{&AoC5Sub1Processor{
		make(map[string][]string),
		make([]string, 0),
	}, []string{},
		make(map[string][]string),
	}
	generateSolution(processor)
}
