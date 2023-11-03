package frequencytuner

type freqRange [2]float64

type status int

// from string to status
func fromString(s string) status {
	switch s {
	case "closer":
		return closer
	case "further":
		return farther
	default:
		return unknown
	}
}

const (
	unknown status = iota
	closer
	farther
)

type freqTuner struct {
	notes  float64
	status status
}
