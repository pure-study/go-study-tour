package main

import (
    "fmt"
    "io"
    "golang.org/x/tour/reader"
)

type MyReader struct{}

// Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
    if b == nil || len(b) == 0 {
        return 0, io.ErrShortBuffer
    }

    l := len(b)
    for i:=0; i<l; i++ {
        b[i] = 'A'
    }
    return l, nil
}

func main() {
    reader.Validate(MyReader{})
    
    my := MyReader{}
    buff := make([]byte, 0)
    fmt.Println(my.Read(buff))
}
