package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	m := make(map[string]int)
	for i := range arr {
		if _, ok := m[arr[i]]; ok == true {
			m[arr[i]] += 1
		} else {
			m[arr[i]] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}