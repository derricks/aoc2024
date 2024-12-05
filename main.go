package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	aoc1Sub2()
}

type AoCProcessor interface {
	Compute() int
	ProcessLine(string) error
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
