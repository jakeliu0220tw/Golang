package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	fmt.Println("dx =", dx, "dy =", dy)

	s := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		s2 := make([]uint8, dx)
		s[i] = s2
		for j := 0; j < dx; j++ {
			s2[j] = uint8(i * j)
		}
	}
	return s
}

func exerciseSlices() {
	pic.Show(Pic)
}
