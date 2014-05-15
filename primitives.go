package main

import ( 
    "fmt"
    "math"
)

func main(){
    var x, y int = 5, 9
    var f float64 = math.Sqrt(float64(x*x + y*y))

    //No need for 'var' declaration when using :=
    //The type will be infered
    z := int(f)

    fmt.Println(x, y, z)
}
