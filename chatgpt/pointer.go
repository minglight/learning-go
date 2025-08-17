package main

import "fmt"

func pointerDemo() {
	type Foo struct {
		x int
	}
	var x = &Foo{}
	var y *int // Fixed: should be *int, not &int
	var z *Foo // z 是指向 Foo 的指標類型

	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("z=", z)
}
