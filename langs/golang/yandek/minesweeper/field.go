package minesweeper

type fieldMap map[cooridinate]rune

type field [][]string

func (f field) String() string {
	var res string
	for i, row := range f {
		if i > 0 {
			res += "\n"
		}
		for _, col := range row {
			res += col
		}
	}
	return res
}

type cooridinate [2]int

func (k cooridinate) outOfBounds(rows, cols int) bool {
	return k[0] < 1 || k[1] < 1 || k[0] > rows || k[1] > cols
}
