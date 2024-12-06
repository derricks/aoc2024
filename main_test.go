package main

import (
	"fmt"
	"slices"
	"testing"
)

type deleteTest struct {
	startSlice    []int
	toDelete      int
	expectedSlice []int
}

func TestDeleteNthSliceItem(test *testing.T) {
	tests := []deleteTest{
		{[]int{1, 2, 7, 8, 9}, 1, []int{1, 7, 8, 9}},
		{[]int{1, 2, 7, 8, 9}, 0, []int{2, 7, 8, 9}},
		{[]int{1, 2, 7, 8, 9}, 4, []int{1, 2, 7, 8}},
	}

	for idx, testCase := range tests {
		newSlice := deleteNthSliceItem(testCase.startSlice, testCase.toDelete)
		if !slices.Equal(newSlice, testCase.expectedSlice) {
			test.Error(fmt.Sprintf("Test case: %d, expected %v, got %v", idx, testCase.expectedSlice, newSlice))
		}
	}

}
