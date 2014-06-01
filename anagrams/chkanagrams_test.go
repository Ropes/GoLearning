package anagrams

import(
    "testing"
)


func TestLengthChk(t* testing.T){
    x := "hihi"
    y := "hihii"

    result := Anagram1(x, y)
    if result == true{
        t.Errorf("Strings of different length but true returned")
    }
}
