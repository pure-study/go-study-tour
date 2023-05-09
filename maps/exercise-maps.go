package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	for _, word := range strings.Fields(s) {
		ret[word] += 1
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
