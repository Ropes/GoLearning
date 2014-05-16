package palindrome

import (
    "fmt"
)

def CheckPalindrome(s string) bool{
    var pal = []rune(s)
    var mid = len(pal)/2
    for i := 0, j := len(pal)-1; i < mid; i++, j-- {
       if j <= i { return true }
       else if !(pal[i] == pal[j]) { return false }
    }
}
