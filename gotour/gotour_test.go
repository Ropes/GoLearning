package gotour

import (
    "testing"
    "fmt"
    "math"
)


func TestSqrt(t *testing.T){
    x := 4.0
    z := Sqrt(x)
    fmt.Println(z)
    if math.Abs(z - 2.0) > 0.0001{
        t.Errorf("Sqrt return value not 2.0!: %E", z)
    }
}
