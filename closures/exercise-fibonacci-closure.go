package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    i, j := 0, 1
    return func() int {
        ret := i
        i, j = j, i + j
        return ret
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 20; i++ {
        fmt.Println(f())
    }
}
