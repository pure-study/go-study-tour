package main

import (
    "fmt"
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    //io.Copy(os.Stdout, &r)
    b, _ := io.ReadAll(r)
    rot13 := string(b)
    fmt.Println(rot13)
    
    // do the same operation again
    r2 := rot13Reader{strings.NewReader(rot13)}
    io.Copy(os.Stdout, &r2)
}

func (mr rot13Reader) Read(b []byte) (int, error) {
    if b == nil || len(b) == 0 {
        return 0, io.ErrShortBuffer
    }
    
    readBytes, err := mr.r.Read(b)
    if err != nil {
        return 0, err
    }
    
    for i:=0; i<readBytes; i++ {
        b[i] = rotate13Alphabet(b[i])
    }
    return readBytes, nil
}

func rotate13Alphabet(c byte) byte {
    ret := c
    switch {
    case c >= 'A' && c <= 'Z':
        ret = (c - 'A' + 13) % 26 + 'A'
    case c >= 'a' && c <= 'z':
        ret = (c - 'a' + 13) % 26 + 'a'
    }
    return ret
}
