package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	aoc1Sub1()
}

type AoCProcessor interface {
	Compute() int
	ProcessLine(string) error
}

type AoC1Sub1Processor struct {
	lefts, rights []int
}

func (p *AoC1Sub1Processor) ProcessLine(line string) error {
	fields := strings.Split(line, "   ")
	leftValue, err := strconv.Atoi(fields[0])
	if err != nil {
		return err
	}
	rightValue, err := strconv.Atoi(fields[1])
	if err != nil {
		return err
	}
	p.AddValues(leftValue, rightValue)
	return nil
}

func processStdin(processor AoCProcessor) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		err := processor.ProcessLine(scanner.Text())
		if err != nil {
			fmt.Printf("Error when processing line: %v\n", err)
		}
	}
}

func (p *AoC1Sub1Processor) AddValues(left, right int) {
	p.lefts = append(p.lefts, left)
	p.rights = append(p.rights, right)
}

func (p *AoC1Sub1Processor) Compute() int {
	runningTotal := 0
	sort.Ints(p.lefts)
	sort.Ints(p.rights)
	for idx, leftValue := range p.lefts {
		runningTotal += (int)(math.Abs((float64)(p.rights[idx] - leftValue)))

	}
	return runningTotal
}

// https://adventofcode.com/2024/day/1
func aoc1Sub1() {
	processor := &AoC1Sub1Processor{[]int{}, []int{}}
	processStdin(processor)
	fmt.Println(processor.Compute())
}
