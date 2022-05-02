package main

import "fmt"

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	s3[3] = 99

	return s1[3], s2[3], s3[3]
}

func nolinked() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99

	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, 6)
	s1[3] = 99

	return s1[3], s2[3]
}

func capNoLinked() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99

	return s1[3], s2[3]
}

func copyNoLinked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)
	s1[3] = 99

	return s1[3], s2[3], copied
}

func appendNoLinked() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...)
	s1[3] = 99

	return s1[3], s2[3]
}

func main() {
	l1, l2, l3 := linked()
	fmt.Println("linked: ", l1, l2, l3)

	nl1, nl2 := nolinked()
	fmt.Println("nolinked: ", nl1, nl2)

	cl1, cl2 := capLinked()
	fmt.Println("capLinked: ", cl1, cl2)

	ncl1, ncl2 := capNoLinked()
	fmt.Println("capNoLinked: ", ncl1, ncl2)

	cpn1, cpn2, cpn3 := copyNoLinked()
	fmt.Println("copyNoLinked: ", cpn1, cpn2, cpn3)

	anl1, anl2 := appendNoLinked()
	fmt.Println("appendNoLinked: ", anl1, anl2)
}
