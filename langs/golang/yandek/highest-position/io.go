package highestposition

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseInput(ss string) (ns []int, err error) {
	for _, s := range strings.Fields(ss) {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ns = append(ns, n)
	}
	return
}

func readInput(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	const maxCapacity int = 400_000 // your required line length
	buf := make([]byte, maxCapacity)
	sc.Buffer(buf, maxCapacity)

	sc.Scan() //first line ignore
	sc.Scan() //input list

	return sc.Text(), sc.Err()
}

func readOutput(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Scan()

	return sc.Text(), sc.Err()
}

func parseOutput(s string) (int, error) {
	return strconv.Atoi(s)
}
