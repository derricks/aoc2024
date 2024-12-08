package main

import "strings"

type CrossedMas map[string][][]int

func (m CrossedMas) findCrossedMas(wordSearch [][]string) bool {
	for letter, coordList := range m {
		// for each set of coordinates, every one has to be the given letter
		for _, coords := range coordList {
			row := coords[0]
			col := coords[1]
			// out of bounds checks
			if row < 0 || row >= len(wordSearch) {
				return false
			}
			if col < 0 || col >= len(wordSearch[0]) {
				return false
			}

			if wordSearch[row][col] != letter {
				return false
			}
		}
	}
	return true
}

// each of these represents a specific set of coordinates offset from a given
// row, column that define a possible arrangement of crossed MAS's
// as defined by the upper left M
func MsOnLeft(row, col int) CrossedMas {
	return map[string][][]int{
		// M . S
		// . A .
		// M . S
		"M": {[]int{row, col}, []int{row + 2, col}},
		"A": {[]int{row + 1, col + 1}},
		"S": {[]int{row, col + 2}, {row + 2, col + 2}},
	}
}

// as defined by upper right M
func MsOnRight(row, col int) CrossedMas {
	// S . M
	// . A .
	// S . M
	return map[string][][]int{
		"M": {[]int{row, col}, []int{row + 2, col}},
		"A": {[]int{row + 1, col - 1}},
		"S": {[]int{row, col - 2}, []int{row + 2, col - 2}},
	}
}

// as measured by the upper left
func MsOnTop(row, col int) CrossedMas {
	// M . M
	// . A .
	// S . S
	return map[string][][]int{
		"M": {[]int{row, col}, []int{row, col + 2}},
		"A": {[]int{row + 1, col + 1}},
		"S": {[]int{row + 2, col}, []int{row + 2, col + 2}},
	}

}

// as measured by bottom left
func MsOnBottom(row, col int) CrossedMas {
	// S . S
	// . A .
	// M . M
	return map[string][][]int{
		"M": {[]int{row, col}, []int{row, col + 2}},
		"A": {[]int{row - 1, col + 1}},
		"S": {[]int{row - 2, col}, []int{row - 2, col + 2}},
	}
}

type AoC4Sub2Processor struct {
	wordSearch [][]string
}

func (p *AoC4Sub2Processor) ProcessLine(line string) error {
	p.wordSearch = append(p.wordSearch, strings.Split(line, ""))
	return nil
}

func (p *AoC4Sub2Processor) Compute() int {
	count := 0
	for rowIndex, row := range p.wordSearch {
		for colIndex, _ := range row {
			finders := []CrossedMas{
				MsOnLeft(rowIndex, colIndex),
				MsOnRight(rowIndex, colIndex),
				MsOnTop(rowIndex, colIndex),
				MsOnBottom(rowIndex, colIndex),
			}

			for _, finder := range finders {
				if finder.findCrossedMas(p.wordSearch) {
					count++
				}
			}

		}
	}
	return count
}

func aoc4sub2() {
	processor := &AoC4Sub2Processor{make([][]string, 0)}
	generateSolution(processor)
}
