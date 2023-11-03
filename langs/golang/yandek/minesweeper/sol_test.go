package minesweeper

import "testing"

func TestSolution(t *testing.T) {
	mines, rows, cols := readInput("./assets/1.txt")

	res := algo(mines, rows, cols)

	if res.String() != "*2*21\n1212*\n12243\n*3***\n13*5*" {
		t.Errorf("Expected:\n*2*21\n1212*\n12243\n*3***\n13*5*\n\nGot:\n%s", res.String())
	}
}

func TestAlgo(t *testing.T) {
	mines := []cooridinate{
		{1, 3},
		{2, 5},
		{4, 1},
		{4, 3},
		{4, 4},
		{4, 5},
		{5, 3},
		{5, 5},
	}

	rows := 5
	cols := 5

	res := algo(mines, rows, cols)
	_ = res
}

func algo(mines []cooridinate, rows, cols int) field {
	fm := initMap(mines, rows, cols)
	return makeField(fm, rows, cols)
}

func TestCreateMap(t *testing.T) {
	// mines := []cooridinate{
	// 	{1, 1},
	// 	{2, 2},
	// }

	// createMap(mines, 3, 2)

	// mines := []cooridinate{
	// 	{1, 3},
	// 	{2, 1},
	// 	{4, 2},
	// 	{4, 4},
	// }

	// createMap(mines, 4, 4)

	mines := []cooridinate{
		{1, 3},
		{2, 5},
		{4, 1},
		{4, 3},
		{4, 4},
		{4, 5},
		{5, 3},
		{5, 5},
	}

	initMap(mines, 5, 5)
}

func initMap(mines []cooridinate, rows, cols int) (fm fieldMap) {
	fm = make(fieldMap)

	offsets := []cooridinate{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1},
		{0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	for _, mine := range mines {
		fm[mine] = '*'

		for _, o := range offsets {
			if c := (cooridinate{mine[0] + o[0], mine[1] + o[1]}); !c.outOfBounds(rows, cols) {
				if _, ok := fm[c]; !ok {
					fm[c] = '1'
				} else {
					if fm[c] != '*' {
						fm[c]++
					}
				}
			}
		}
	}

	return
}

func makeField(f fieldMap, rows, cols int) field {
	var res field

	for i := 0; i < rows; i++ {
		res = append(res, make([]string, cols))
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			k := cooridinate{i + 1, j + 1} // 1-indexed
			if _, ok := f[k]; ok {
				res[i][j] = string(f[k])
			} else {
				res[i][j] = "0"
			}
		}
	}

	return res
}
