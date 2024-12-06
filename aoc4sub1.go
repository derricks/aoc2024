package main

import "strings"

type AoC4Sub1Processor struct {
	wordSearch [][]string // indexed by row then column
}

type XmasLooker map[string][]int // for a given letter, where should you expect to find it in a grid of row/column

func (x XmasLooker) findXmasInGrid(grid [][]string) bool {
	for _, letter := range []string{"X", "M", "A", "S"} {
		row, col := x[letter][0], x[letter][1]
		if row < 0 || row >= len(grid) {
			return false
		}

		if col < 0 || col >= len(grid[0]) {
			return false
		}

		if grid[row][col] != letter {
			return false
		}
	}
	return true
}

func WestLooker(row, col int) XmasLooker {
	return map[string][]int{
		"X": {row, col},
		"M": {row, col - 1},
		"A": {row, col - 2},
		"S": {row, col - 3},
	}
}

func EastLooker(row, col int) XmasLooker {
	return map[string][]int{
		"X": {row, col},
		"M": {row, col + 1},
		"A": {row, col + 2},
		"S": {row, col + 3},
	}
}

func NorthLooker(row, col int) XmasLooker {
	return map[string][]int{
		"X": {row, col},
		"M": {row - 1, col},
		"A": {row - 2, col},
		"S": {row - 3, col},
	}
}

func SouthLooker(row, col int) XmasLooker {
	return map[string][]int{
		"X": {row, col},
		"M": {row + 1, col},
		"A": {row + 2, col},
		"S": {row + 3, col},
	}
}

func NorthwestLooker(row, col int) XmasLooker {
	return map[string][]int{
		"X": {row, col},
		"M": {row - 1, col - 1},
		"A": {row - 2, col - 2},
		"S": {row - 3, col - 3},
	}
}

func SouthwestLooker(row, col int) XmasLooker {
	return map[string][]int{
		"X": {row, col},
		"M": {row + 1, col - 1},
		"A": {row + 2, col - 2},
		"S": {row + 3, col - 3},
	}
}

func SoutheastLooker(row, col int) XmasLooker {
	return map[string][]int{
		"X": {row, col},
		"M": {row + 1, col + 1},
		"A": {row + 2, col + 2},
		"S": {row + 3, col + 3},
	}
}

func NortheastLooker(row, col int) XmasLooker {
	return map[string][]int{
		"X": {row, col},
		"M": {row - 1, col + 1},
		"A": {row - 2, col + 2},
		"S": {row - 3, col + 3},
	}
}

func (p *AoC4Sub1Processor) ProcessLine(line string) error {
	p.wordSearch = append(p.wordSearch, strings.Split(line, ""))
	return nil
}

func (p *AoC4Sub1Processor) Compute() int {

	runningTotal := 0
	for rowIdx, row := range p.wordSearch {
		for colIdx, letter := range row {
			if letter == "X" {
				lookers := []XmasLooker{
					NorthwestLooker(rowIdx, colIdx),
					WestLooker(rowIdx, colIdx),
					SouthwestLooker(rowIdx, colIdx),
					SouthLooker(rowIdx, colIdx),
					SoutheastLooker(rowIdx, colIdx),
					EastLooker(rowIdx, colIdx),
					NortheastLooker(rowIdx, colIdx),
					NorthLooker(rowIdx, colIdx),
				}

				for _, looker := range lookers {
					if looker.findXmasInGrid(p.wordSearch) {
						runningTotal++
					}
				}
			}
		}
	}
	return runningTotal
}

func aoc4sub1() {
	processor := &AoC4Sub1Processor{make([][]string, 0)}
	generateSolution(processor)
}
