package main

import "fmt"

func pushData(begin int, end int, in chan bool, out chan int) {
	for i := begin; i <= end; i++ {
		b := <-in
		if b == true {
			out <- i
		}
	}
}

func main() {
	sum := 0
	in := make(chan bool, 100)
	out := make(chan int, 100)
	defer close(in)
	defer close(out)

	go pushData(1, 25, in, out)
	go pushData(26, 50, in, out)
	go pushData(51, 75, in, out)
	go pushData(76, 100, in, out)

	for i := 1; i <= 100; i++ {
		in <- true
		v := <-out
		sum += v
	}

	fmt.Println(sum)
}
