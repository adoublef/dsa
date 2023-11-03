package frequencytuner

import (
	"math"
	"testing"
)

func TestSolution(t *testing.T) {
	freqTuners, err := readFile("assets/1.txt")
	if err != nil {
		t.Error(err)
	}

	fr := freqRange{30, 40_00}

	if algo(freqTuners, &fr); fr[0] != 531 || fr[1] != 660 {
		t.Errorf("got %v, want %v", fr, freqRange{531, 660})
	}
}

func TestAlgo(t *testing.T) {
	init := freqRange{30, 40_00}
	algo([]freqTuner{
		{notes: 400, status: unknown},
		{notes: 220, status: closer},
		{notes: 300, status: farther},
	}, &init)
	_ = init
}

func algo(as []freqTuner, fr *freqRange) {
	for i := 0; i < len(as)-1; i++ {
		switch curr, mid := as[i+1], (as[i].notes+as[i+1].notes)/2.0; curr.status {
		case closer:
			if curr.notes > mid {
				fr[0] = math.Max(fr[0], mid)
			} else {
				fr[1] = math.Min(fr[1], mid)
			}
		case farther:
			if curr.notes < mid {
				fr[0] = math.Max(fr[0], mid)
			} else {
				fr[1] = math.Min(fr[1], mid)
			}
		}
	}
}
