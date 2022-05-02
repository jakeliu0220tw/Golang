package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	str_ary := strings.Fields(s)

	// iterate array
	for i := range str_ary {
		key := str_ary[i]
		m[key] = m[key] + 1
	}

	return m
}

func exerciseMaps() {
	wc.Test(WordCount)
}
