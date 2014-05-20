package gotour

import (
    "testing"
    "fmt"
)

func TestMaps(t *testing.T) {
    Carins = make(map[string]Vertex)
    Carins["PDX"] = Vertex{
        45.523204, -122.671324,
    }
    fmt.Println(Carins)

    x := Carins["PDX"]
    if x.Lat != 45.523204 {
        t.Errorf("Lat not saved correctly %q", x.Lat)
    }
}

func TestMapContents(t *testing.T){
    Carins = make(map[string]Vertex)
    Carins["PDX"] = Vertex{
        45.523204, -122.671324,
    }
    
    /*v, ok
    v, ok := Carins["SEA"]
    fmt.Println(v, ok)
    */
    
    _, ok := Carins["SEA"]
    if ok != false {
        t.Errorf("SEA never defined!??! Should not be present")
    }
}

func TestWC(t *testing.T){
    TestWordCount(WordCount)
}

