package anagrams

import (
	"io/ioutil"
	"sort"
	"strings"
	"unicode/utf8"
)

//RuneSlice type attaches functions to allow sorting of runes
type RuneSlice []rune

func (rs RuneSlice) Len() int {
	return utf8.RuneCountInString(string(rs))
}

func (rs RuneSlice) Less(i, j int) bool {
	return rs[i] < rs[j]
}

func (rs RuneSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

//LenCheck returns true if strings are of equal length, else false
func LenCheck(x string, y string) bool {
	if len(x) == len(y) {
		return true
	}
	return false
}

//SortWord takes a string and returns its characters sorted as a string
func SortWord(word string) string {
	w := []rune(word)
	rs := RuneSlice(w)
	sort.Sort(rs)
	return string(rs)
}

//Anagram1 sorts the characters of two given strings and returns true if they match
func Anagram1(x string, y string) bool {
	if LenCheck(x, y) {
		//Continue anagram check
		//TODO: map characters
		amapX := SortWord(x)
		amapY := SortWord(y)
		if amapX == amapY {
			return true
		}
	}
	return false
}

func ReadSystemWords() ([]string, error) {
	path := "/usr/share/dict/words"
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
		return nil, err
	}
	return strings.Split(string(contents), "\n"), nil
}

/*
func AnagramList(words []string) map[string][]string {
	anagrams = make(map[string][]string)
	for _, w := range words {

	}
}
*/
