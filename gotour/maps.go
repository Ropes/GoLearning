package gotour

import (
    "strings"
    "math"
)

type Vertex struct{
    Lat, Lon float64
}

func UselessAbs(v *Vertex) float64 {
    return math.Sqrt(v.Lat * v.Lat + v.Lon * v.Lon)
}

//Init: m = make(map[string]Vertex)
var Carins map[string]Vertex

type MyMap map[string]Vertex

func (mm MyMap) ShowPDX() Vertex {
    return  mm["PDX"]
}

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    m := make(map[string]int)

    for _, w := range words{
        m[w]++
    }
    return m
}
