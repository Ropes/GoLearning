package anagrams

import (
	"testing"
)

func TestLengthChk(t *testing.T) {
	x := "hihi"
	y := "hihii"
	z := "ihih"

	result := LenCheck(x, y)
	if result == true {
		t.Errorf("Strings of different length but true returned")
	}
}
