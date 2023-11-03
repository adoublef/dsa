package highestposition

import (
	"testing"
)

func TestSolution(t *testing.T) {
	type testcase struct {
		name   string
		input  string
		output string
	}

	tt := []testcase{
		{name: "test 1", input: "assets/input/1.txt", output: "assets/output/1.txt"},
		{name: "test 2", input: "assets/input/2.txt", output: "assets/output/2.txt"},
		{name: "test 3", input: "assets/input/3.txt", output: "assets/output/3.txt"},
		{name: "test 4", input: "assets/input/4.txt", output: "assets/output/4.txt"},

		// {name: "test 7", input: "1000 100 15", expected: 0},
		// {name: "test 8", input: "852 105 55", expected: 2},
		// {name: "test 9", input: "625 933 175 17 241", expected: 3},
		{name: "test 10", input: "assets/input/10.txt", output: "assets/output/10.txt"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := problem(tc.input)
			if err != nil {
				t.Fatalf("issue with input: %v", err)
			}
			exp, err := solution(tc.output)
			if err != nil {
				t.Fatalf("issue with output: %v", err)
			}
			if res != exp {
				t.Errorf("expected %v but got %v", exp, res)
			}
		})
	}
}

func algo(ps []int) int {
	var mPos, mVal int
	for i := 0; i < len(ps)-1; i++ {
		if aVal := ps[i]; aVal > mVal {
			mPos, mVal = i, aVal
		}
	}
	var rPos, rVal int
	for j := mPos + 1; j < len(ps)-1; j++ {
		if aVal, bVal := ps[j], ps[j+1]; aVal%10 == 5 && aVal > bVal && aVal > rVal {
			rPos, rVal = j, aVal
		}
	}
	if rPos == 0 || mVal == rVal {
		return rPos
	}
	return rPos + 1
}

func problem(path string) (int, error) {
	s, err := readInput(path)
	if err != nil {
		return 0, err
	}
	ps, err := parseInput(s)
	return algo(ps), err
}

func solution(path string) (int, error) {
	n, err := readOutput(path)
	if err != nil {
		return 0, err
	}
	return parseOutput(n)
}
