package maximalproduct

import (
	"math"
	"testing"
)

func TestSolution(t *testing.T) {
	type testcase struct {
		name string
		path string
		exp  tuple
	}

	tt := []testcase{
		{name: "case 10", path: "./assets/2.txt", exp: tuple{29793, 29710}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := solution(tc.path)
			if !res.equal(tc.exp) {
				t.Errorf("expected %d %d but got %d %d", tc.exp[0], tc.exp[1], res[0], res[1])
			}
		})
	}
}

func solution(path string) (r *tuple, err error) {
	s, err := readInput(path)
	if err != nil {
		return nil, err
	}
	ns, err := parseSlice(s)
	return algo(ns), err
}

func TestAlgo(t *testing.T) {
	n1 := algo([]int{-4, 3, -5, 2, 5})
	if n1[0] != -4 || n1[1] != -5 {
		t.Errorf("expected -4 -5 but got %d %d", n1[0], n1[1])
	}

	n2 := algo([]int{4, 3, 5, 2, 5})
	if n2[0] != 5 || n2[1] != 5 {
		t.Errorf("expected 5 5 but got %d %d", n2[0], n2[1])
	}
}

func maxmin(ns []int) (max, min *tuple) {
	max, min = &tuple{math.MinInt, math.MinInt}, &tuple{math.MaxInt, math.MaxInt}
	for _, n := range ns {
		if n > max[0] {
			max[0], max[1] = n, max[0]
		} else if n > max[1] {
			max[1] = n
		}
		if n < min[1] {
			min[0], min[1] = min[1], n
		} else if n < min[0] {
			min[0] = n
		}
	}
	return
}

// at minumum there is 2 numbers
func algo(ns []int) *tuple {
	max, min := maxmin(ns)
	if minProd := min.product(); minProd > max.product() {
		return min
	}
	return max
}
