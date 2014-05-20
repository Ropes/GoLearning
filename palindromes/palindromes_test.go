package palindromes

import (
    "testing"
)

type TData struct { 
    chk string
    want bool 
}
var test_data = []TData {
    {"racecar", true}, 
    {"test", false},
    {"yxxy", true},
    {"101", true},
    {"at", false},
}
func PalindromesIter(t *testing.T) {
    for _, c := range test_data {
        got := CheckPalindrome(c.chk)
        if got != c.want {
            t.Errorf("Palindrome[%q == %q], should be: %q", c.chk, got, c.want)
        }
    }
}

func BenchmarkPalindromesIter(t *testing.B) {
    for i := 0; i < t.N; i++{
        for _, c := range test_data {
            got := CheckPalindrome(c.chk)
            if got != c.want {
                t.Errorf("Palindrome[%q == %q], should be: %q", c.chk, got, c.want)
            }
        }
    }
}

func BenchmarkPalindromesRec(t *testing.B) {
    for i := 0; i < t.N; i++{
        for _, c := range test_data {
            got := RecPalindrome(c.chk)
            if got != c.want {
                t.Errorf("Recursive Palindrome[%q == %q], should be: %q", c.chk, got, c.want)
            }
        }
    }
}
