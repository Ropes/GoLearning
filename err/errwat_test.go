package err

import (
    "testing"
    "fmt"
)

func TestErrFunc(t *testing.T){
    if err := failWhale(); err != nil{
        fmt.Println(err)
    }
}
