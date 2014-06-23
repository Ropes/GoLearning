package anagrams

import (
	"sort"
	"unicode/utf8"
)

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

func LenCheck(x string, y string) bool {
	if len(x) == len(y) {
		return true
	}
	return false
}

func SortWord(word string) string {
	w := []rune(word)
	rs := RuneSlice(w)
	sort.Sort(rs)
	return string(rs)
}

func Anagram1(x string, y string) bool {
	if LenCheck(x, y) {
		//Continue anagram check
		//TODO: map characters

	}
	return false
}
