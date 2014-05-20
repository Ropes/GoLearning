package palindromes

import (
    "testing"
)

type TData struct { 
    chk string
    want bool 
}
func TestPalindromes(t *testing.T) {
    var test_data = []TData {
        {"racecar", true}, 
        {"test", false},
        {"yxxy", true},
        {"101", true},
        {"at", false},
    }

    for _, c := range test_data {
        got := CheckPalindrome(c.chk)
        if got != c.want {
            t.Errorf("Palindrome[%q == %q], should be: %q", c.chk, got, c.want)
        }
    }

    for _, c := range test_data {
        got := RecPalindrome(c.chk)
        if got != c.want {
            t.Errorf("Recursive Palindrome[%q == %q], should be: %q", c.chk, got, c.want)
        }
    }
}
