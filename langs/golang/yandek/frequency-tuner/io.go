package frequencytuner

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) (fts []freqTuner, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	sc.Scan() //first line ignore

	for sc.Scan() {
		line := strings.Fields(sc.Text())
		if len(line) == 1 {
			x, err := strconv.ParseFloat(line[0], 64)
			if err != nil {
				return nil, err
			}
			fts = append(fts, freqTuner{notes: x, status: unknown})
		} else {
			x, err := strconv.ParseFloat(line[0], 64)
			if err != nil {
				return nil, err
			}
			fts = append(fts, freqTuner{notes: x, status: fromString(line[1])})
		}
	}

	return fts, sc.Err()
}

/*
if len([]string) == 1 then unknown
4
554
880 further
440 closer
441 closer
*/
