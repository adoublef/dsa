package fakebin

import (
	"testing"
)

func FakeBin(x string) (s string) {
	for _, ch := range x {
		if ch < '5' {
			s += "0"
		} else {
			s += "1"
		}
	}
	return
}

// 54 -> 6
// 50 -> 2
func TestFakeBin(t *testing.T) {
	hello := "620382"
	if FakeBin(hello) != "100010" {
		t.Errorf("expected `100010` but got %v", FakeBin(hello))
	}
}
