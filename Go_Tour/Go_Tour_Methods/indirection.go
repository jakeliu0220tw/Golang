package main

import "fmt"

func indirection() {
	v := Vertex{3, 4}
	v.ScaleMethod2(2)
	fmt.Println(v)
	ScaleFunc(&v, 10)
	fmt.Println(v)

	p := &Vertex{3, 4}
	p.ScaleMethod2(2)
	fmt.Println(*p)
	ScaleFunc(p, 10)
	fmt.Println(*p)
}

func indirection2() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(FuncAbs(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(FuncAbs(*p))
}
