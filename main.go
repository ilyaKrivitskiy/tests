package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	lex := strings.Split(s, " ")
	for i := range lex {
		_, ok := m[lex[i]]
		if ok {
			m[lex[i]]++
		} else {
			m[lex[i]] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
