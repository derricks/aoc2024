package main

import "testing"

type leftLookTest struct {
	// simpler structure since the row index will always be the same
	grid           [][]string
	startCols      []int
	expectedToFind []bool // there should be one of these for each starting location
}

func TestEastLooker(test *testing.T) {
	tests := []leftLookTest{
		{[][]string{{"S", "A", "M", "X"}}, []int{3}, []bool{true}},
		{[][]string{{"S", "A", "M"}}, []int{2}, []bool{false}},
		{[][]string{{"S", "A", "M", "X", "S", "A", "M", "X"}}, []int{3, 7}, []bool{true, true}},
		{[][]string{{"A", "M", "X"}}, []int{2}, []bool{false}},
	}

	for idx, testCase := range tests {
		for colIdx, col := range testCase.startCols {
			looker := WestLooker(0, col)
			found := looker.findXmasInGrid(testCase.grid)
			if found != testCase.expectedToFind[colIdx] {
				test.Errorf("Test case %d. Expected to find XMAS in %s but did not", idx, testCase.grid)
			}

		}
	}

}
