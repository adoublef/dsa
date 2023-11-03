package minesweeper

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) (mines []cooridinate, rows, cols int) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)

	params := []string{}

	for scanner.Scan() {
		params = append(params, scanner.Text())
	}

	r_c_m := strings.Fields(params[0])
	r, _ := strconv.Atoi(r_c_m[0])
	c, _ := strconv.Atoi(r_c_m[1])
	m, _ := strconv.Atoi(r_c_m[2])

	// var mines []cooridinate //

	for i := 1; i <= m; i++ {
		coords := strings.Fields(params[i])
		coord_row, _ := strconv.Atoi(coords[0])
		coord_column, _ := strconv.Atoi(coords[1])
		array := cooridinate{coord_row, coord_column}
		mines = append(mines, array)

	}

	return mines, r, c
}
