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
        t.Errorf("Lat not saved correctly %E", x.Lat)
    }
}

func TestMapContents(t *testing.T){
    Carins = make(map[string]Vertex)
    Carins["PDX"] = Vertex{
        45.523204, -122.671324,
    }

    _, ok := Carins["SEA"]
    if ok != false {
        t.Errorf("SEA never defined!??! Should not be present")
    }
}

func TestWC(t *testing.T){
    TestWordCount(WordCount)
}

func TestMyType(t *testing.T){
    m := make(MyMap)
    m["PDX"] = Vertex{
        45.523204, -122.671324,
    }
    pdx := m.ShowPDX()
    fmt.Println("PDX Vertex returned from Type Function: %q", pdx)
    if pdx.Lat != 45.523204 {
        t.Errorf("PDX never defined!??! Should be present")
    }
}

func TestVertexInterface(t *testing.T){

    x := &Vertex{-1, -2}
    y := &Vertex{-3, 5}

    var a Vtex
    a = x

    mid := a.Midpoint(y)
    fmt.Println(mid)
    if mid.Lat != -2.0 || mid.Lon != 1.5{
        t.Errorf("Midpoint Vertex incorrect! %q", mid)
    }
}

