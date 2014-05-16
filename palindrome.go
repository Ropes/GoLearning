package main 

func CheckPalindrome(s string) bool {

    var pal = []rune(s)
    var mid = len(pal)/2

    for i, j := 0, len(pal); i < mid; i++ {
        j--
        if j <= i { 
            return true 
        } else if !(pal[i] == pal[j]) { 
            return false 
        }
    }
    return true
}
