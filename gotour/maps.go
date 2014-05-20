package gotour

import "strings"

type Vertex struct{
    Lat, Lon float64
}

//Init: m = make(map[string]Vertex)
var Carins map[string]Vertex

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    m := make(map[string]int)

    for _, w := range words{
        m[w]++
    }
    return m
}
