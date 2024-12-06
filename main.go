package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	aoc3sub2()
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

func generateSolution(processor AoCProcessor) {
	processStdin(processor)
	fmt.Println(processor.Compute())
}

// utility function for removing an item from a slice
// unlike slices.Delete, this doesn't simply 0 out the value
// but actually makes a new slice with the item removed
func deleteNthSliceItem[S ~[]E, E any](slice S, idxToDelete int) S {
	newSlice := make(S, 0)
	for idx, item := range slice {
		if idx == idxToDelete {
			continue
		}
		newSlice = append(newSlice, item)
	}
	return newSlice
}
