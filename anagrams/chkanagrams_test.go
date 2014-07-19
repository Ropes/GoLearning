package anagrams

import "testing"

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

func TestAnagram1(t *testing.T) {
	x := "heater"
	y := "reheat"

	result := Anagram1(x, y)
	if result == false {
		t.Error("Anagram not detected correctly")
	}
}

func TestReadWords(t *testing.T) {
	words, err := ReadSystemWords()
	if err != nil {
		t.Log("No error reading word list")
	}
	t.Log(words[:5])
	if len(words) < 90000 {
		t.Errorf("Words not read from dictionary successfully")
	}
}

func TestAnagramList(t *testing.T) {
	words, err := ReadSystemWords()
	if err != nil {
		t.Log("No error reading word list")
	}
	anagrams := AnagramList(words)
	if len(anagrams) < 5000 {
		t.Log(anagrams["acr"])
		t.Errorf("Number of anagram combinations dubiously low for number of words..")
	}
	//fmt.Println(anagrams["aflt"])
}
