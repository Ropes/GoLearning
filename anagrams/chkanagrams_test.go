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

	result = LenCheck(x, z)
	if result == false {
		t.Errorf("Strings of same length but false returned..")
	}
}

func TestRuneSliceSort(t *testing.T) {
	x := "hihi"
	y := "hhii"

	result := SortWord(x)
	if result != y {
		t.Errorf("Word wasn't sorted properly: %s", result)
	}
}

func TestRuneSliceSort2(t *testing.T) {
	x := "watcatw"
	y := "aacttww"
	result := SortWord(x)

	if result != y {
		t.Errorf("Word wasn't sorted properly: %s", result)
	}
}
