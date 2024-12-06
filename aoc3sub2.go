package main

import (
	"regexp"
	"slices"
	"strings"
)

// https://adventofcode.com/2024/day/3
type AoC3Sub2Processor struct {
	enabled bool
	// mostly so we can reuse its tested evalMulProgram method
	// as well as Compute. So mul programs get added to its field
	// but we write our own ProcessLine to push into that field
	sub1Processer *AoC3Sub1Processor
}

var DO_COMMAND_REGEX = regexp.MustCompile("do\\(\\)")
var DONT_COMMAND_REGEX = regexp.MustCompile("don't\\(\\)")

type MulCommandType int

const (
	DoCommand   MulCommandType = 1
	DontCommand MulCommandType = 2
	MulCommand  MulCommandType = 3
)

func (p *AoC3Sub2Processor) ProcessLine(line string) error {
	// we're going to mutate the line as we go, so make a copy
	lineCopy := strings.Clone(line)
	doLoc := DO_COMMAND_REGEX.FindStringIndex(lineCopy)
	dontLoc := DONT_COMMAND_REGEX.FindStringIndex(lineCopy)
	mulLoc := mulExtract1Regex.FindStringIndex(lineCopy)

	for doLoc != nil || dontLoc != nil || mulLoc != nil {
		validLocs := make([][]int, 0)
		if doLoc != nil {
			validLocs = append(validLocs, doLoc)
		}

		if dontLoc != nil {
			validLocs = append(validLocs, dontLoc)
		}

		if mulLoc != nil {
			validLocs = append(validLocs, mulLoc)
		}

		// sort so that the soonest loc is at the top
		slices.SortFunc(validLocs, func(a, b []int) int {
			if a[0] < b[0] {
				return -1
			}

			if a[0] > b[0] {
				return 1
			}

			return 0
		})

		subString := lineCopy[validLocs[0][0]:validLocs[0][1]]
		p.processCommand(p.commandFromString(subString), subString)

		// advance past the first command
		lineCopy = lineCopy[validLocs[0][1]:]
		doLoc = DO_COMMAND_REGEX.FindStringIndex(lineCopy)
		dontLoc = DONT_COMMAND_REGEX.FindStringIndex(lineCopy)
		mulLoc = mulExtract1Regex.FindStringIndex(lineCopy)
	}
	return nil
}

func (p *AoC3Sub2Processor) processCommand(command MulCommandType, commandText string) {
	switch {
	case command == DoCommand:
		p.enabled = true
	case command == DontCommand:
		p.enabled = false
	case command == MulCommand:
		if p.enabled {
			p.sub1Processer.mulExtracts = append(p.sub1Processer.mulExtracts, commandText)
		}
	}
}

func (p *AoC3Sub2Processor) commandFromString(cmd string) MulCommandType {
	if strings.HasPrefix(cmd, "don") {
		return DontCommand
	}

	if strings.HasPrefix(cmd, "do") {
		return DoCommand
	}
	return MulCommand
}

func (p *AoC3Sub2Processor) Compute() int {
	return p.sub1Processer.Compute()
}

func aoc3sub2() {
	processor := &AoC3Sub2Processor{true, &AoC3Sub1Processor{[]string{}}}
	generateSolution(processor)
}
