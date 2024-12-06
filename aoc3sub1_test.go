package main

import "testing"

type mulProgramTest struct {
	mulProgram string
	expected   int
}

func TestEvalMulProgram(test *testing.T) {
	tests := []mulProgramTest{
		{"mul(3,4)", 12},
		{"mul(2,4)", 8},
		{"mul(5,5)", 25},
		{"mul(11,8)", 88},
		{"mul(8,5)", 40},
	}
	p := &AoC3Sub1Processor{[]string{}}

	for idx, testCase := range tests {
		actual := p.evalMulProgram(testCase.mulProgram)
		if actual != testCase.expected {
			test.Errorf("Test case %d: expected %d but got %d", idx, testCase.expected, actual)
		}
	}
}
