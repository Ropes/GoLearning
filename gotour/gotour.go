package gotour

import "fmt"

func Pic(dx, dy int) [][]uint8 {
    fmt.Println(dx, dy)
    cols := make([][]uint8,         dy)

    for i, _ := range cols {
        cols[i] = make([]uint8, dx)
    }

    for x := range make([]int, dx) {
        for y := range make([]int, dy) {
            cols[y][x] = uint8((x + y) / 2)
        }
    }

    return cols
}

func Fibonacci() func() int {
    nm1 := 1
    nm2 := 0
    return func() int {
                nm1, nm2 = nm1 + nm2, nm1
                return nm1
            }
}
