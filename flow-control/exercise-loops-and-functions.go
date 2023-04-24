package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    fmt.Println("==================")
    z := 1.0
    preZ := x
    for i:=1; math.Abs(preZ - z) >= 0.00000001; i++ {
        preZ = z
        z -= (z * z - x) / (2 * z)  // Newton's method
        fmt.Printf("%v: %v\n", i, z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(10000))
}
