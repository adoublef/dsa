package maximalproduct

type tuple [2]int

func (t tuple) product() int { return t[0] * t[1] }

func (t tuple) equal(other tuple) bool {
	return t[0] == other[0] && t[1] == other[1]
}

type truple [3]int

func (t truple) product() int { return t[0] * t[1] * t[2] }

func (t truple) equal(other truple) bool {
	return t[0] == other[0] && t[1] == other[1] && t[2] == other[2]
}
