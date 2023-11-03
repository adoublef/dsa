package maximalproduct

import (
	"math"
	"testing"
)

func TestSolution3(t *testing.T) {
	type testcase struct {
		name string
		path string
		exp  truple
	}

	tt := []testcase{
		{name: "case 10", path: "./assets/3.txt", exp: truple{0, 0, 0}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := solution3(tc.path)
			if !res.equal(tc.exp) {
				t.Errorf("expected %d %d but got %d %d", tc.exp[0], tc.exp[1], res[0], res[1])
			}
		})
	}
}

func solution3(path string) (r *truple, err error) {
	s, err := readInput(path)
	if err != nil {
		return nil, err
	}
	ns, err := parseSlice(s)
	return algo3(ns), err
}

func maxmin3(ns []int) (max *truple, min *tuple) {
	max, min = &truple{math.MinInt, math.MinInt, math.MinInt}, &tuple{math.MaxInt, math.MaxInt}
	for _, n := range ns {
		if n > max[0] {
			max[0], max[1], max[2] = n, max[2], max[1]
		} else if n > max[1] {
			max[1], max[2] = n, max[1]
		} else if n > max[2] {
			max[2] = n
		}
		if n < min[1] {
			min[0], min[1] = min[1], n
		} else if n < min[0] {
			min[0] = n
		}
	}
	return
}

func algo3(ns []int) *truple {
	max, min := maxmin3(ns)
	if minProd := min.product(); minProd*max[2] > max.product() {
		return &truple{min[0], min[1], max[2]}
	}
	return max
}
