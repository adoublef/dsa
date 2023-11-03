package day01

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	is "github.com/stretchr/testify/require"
)

var (
	session string
	input   string
)

/*
--- Day 1: Not Quite Lisp ---

Santa was hoping for a white Christmas, but his weather machine's "snow" function is powered by stars, and he's fresh out! To save Christmas, he needs you to collect fifty stars by December 25th.

Collect stars by helping Santa solve puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

Here's an easy puzzle to warm you up.

Santa is trying to deliver presents in a large apartment building, but he can't find the right floor - the directions he got are a little confusing. He starts on the ground floor (floor 0) and then follows the instructions one character at a time.

An opening parenthesis, (, means he should go up one floor, and a closing parenthesis, ), means he should go down one floor.

The apartment building is very tall, and the basement is very deep; he will never find the top or bottom floors.

For example:

	(()) and ()() both result in floor 0.
	((( and (()(()( both result in floor 3.
	))((((( also results in floor 3.
	()) and ))( both result in floor -1 (the first basement level).
	))) and )())()) both result in floor -3.
	To what floor do the instructions take Santa?
*/
func TestPartA(t *testing.T) {
	t.Parallel()

	type testcase struct {
		name string
		in   string
		exp  int
	}

	tt := []testcase{
		{name: "example 1", in: "(())", exp: 0},
		{name: "example 2", in: "()()", exp: 0},
		{name: "example 3", in: "(((", exp: 3},
		{name: "example 4", in: "(()(()(", exp: 3},
		{name: "example 5", in: "))(((((", exp: 3},
		{name: "example 6", in: "())", exp: -1},
		{name: "example 7", in: "))(", exp: -1},
		{name: "example 8", in: ")))", exp: -3},
		{name: "example 9", in: ")())())", exp: -3},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			is.Equal(t, tc.exp, partA(tc.in))
		})
	}

	t.Logf("partA: %d", partA(input))
}

/*
--- Part Two ---

Now, given the same instructions, find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.

For example:

	) causes him to enter the basement at character position 1.
	()()) causes him to enter the basement at character position 5.
	What is the position of the character that causes Santa to first enter the basement?
*/
func TestPartB(t *testing.T) {
	t.Parallel()

	type testcase struct {
		name string
		in  string
		exp int
	}

	tt := []testcase{
		{name: "example 1", in: ")", exp: 1},
		{name: "example 2", in: "()())", exp: 5},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			is.Equal(t, tc.exp, partB(tc.in))
		})
	}

	t.Logf("partB: %d", partB(input))
}

func partA(s string) int {
	var n int
	for _, r := range s {
		switch r {
		case '(':
			n++
		case ')':
			n--
		}
	}

	return n
}

func partB(s string) int {
	var n int
	for i, r := range s {
		switch r {
		case '(':
			n++
		case ')':
			n--
		}

		if n == -1 {
			return i + 1
		}
	}

	return -1
}

func readInput(year, day int, session string) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// read body into string
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func TestMain(m *testing.M) {
	if session = os.Getenv("AOC_SESSION"); session == "" {
		fmt.Println("warning: AOC_SESSION not set")
	}

	var err error
	input, err = readInput(2015, 1, session)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}
