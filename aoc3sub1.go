package main

import (
	"regexp"
	"strconv"
)

// https://adventofcode.com/2024/day/3
type AoC3Sub1Processor struct {
	mulExtracts []string
}

var mulExtract1Regex = regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")
var digitExtractRegex = regexp.MustCompile("[0-9]+")

const REGEX_FIND_ALL = -1

func (p *AoC3Sub1Processor) ProcessLine(line string) error {
	muls := mulExtract1Regex.FindAllString(line, REGEX_FIND_ALL)
	for _, mulProgram := range muls {
		p.mulExtracts = append(p.mulExtracts, mulProgram)
	}
	return nil
}

func (p *AoC3Sub1Processor) Compute() int {
	runningTotal := 0
	for _, program := range p.mulExtracts {
		runningTotal += p.evalMulProgram(program)
	}
	return runningTotal
}

func (p *AoC3Sub1Processor) evalMulProgram(mulProgram string) int {
	digitStrings := digitExtractRegex.FindAllString(mulProgram, REGEX_FIND_ALL)
	digits := make([]int, 2)
	digits[0], _ = strconv.Atoi(digitStrings[0])
	digits[1], _ = strconv.Atoi(digitStrings[1])
	return digits[0] * digits[1]
}

func aoc3sub1() {
	processor := &AoC3Sub1Processor{[]string{}}
	generateSolution(processor)
}
