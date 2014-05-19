package palindromes

import "testing"

func TestPalindromes(t *testing.T) {
    /*j
    var test_data = []struct {
        chk string, want bool
    }
    test_data[0] = {"racecar", true}
        {"test", false},
        {"yxxy", true},
        {"101", true},
        {"at", false},
    */

    /*
    for _, c := range test_data {
        got := CheckPalindrome(c.chk)
        if got != c.want {
            t.Errorf("Palindrome[%q == %q], should be: %q", c.chk, got, c.want)
        }
    }
    */
    type X struct { 
        chk string
        want bool 
    }
    test := X{"yxxy", true}
    got := CheckPalindrome(test.chk)
    if got != test.want {
        t.Errorf("Palindrome[%q == %q], should be: %q", test.chk, got, test.want)
    }
}
