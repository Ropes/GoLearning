package gotour

import (
	"math"
	"strings"
)

type Vtex interface {
	UselessAbs() float64
	Midpoint(y *Vertex) Vertex
}

type Vertex struct {
	Lat, Lon float64
}

func (v Vertex) UselessAbs() float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Lon*v.Lon)
}

func (v *Vertex) Midpoint(y *Vertex) Vertex {
	return Vertex{(v.Lat + y.Lat) / 2.0, (v.Lon + y.Lon) / 2.0}
}

//Init: m = make(map[string]Vertex)
var Carins map[string]Vertex

type MyMap map[string]Vertex

func (mm MyMap) ShowPDX() Vertex {
	return mm["PDX"]
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)

	for _, w := range words {
		m[w]++
	}
	return m
}
