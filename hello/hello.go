package main

import "fmt"
import "github.com/ropes/palindromes"

func main(){
    uno := "test"
    dos := "racecar"
    tres := "101"

    fmt.Println("Hellow!!! I like Go's install so far!  Very straitforward!")
    fmt.Println(CheckPalindrome(uno))
    fmt.Println(CheckPalindrome(dos))
    fmt.Println(CheckPalindrome(tres))
}
