package palindromes

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

func RecChk(s []rune) bool {
    if len(s) <= 1 {
        return true
    } else {
        return s[0] == s[len(s)-1] && RecChk(s[1:len(s)-1])
    }
}

func RecPalindrome(str string) bool {
    s := []rune(str)
    return RecChk(s)
}

