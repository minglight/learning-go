package main

import "fmt"

func failedUpdate(px *int) {
	a := 23
	px = &a
	fmt.Println("failedUpdate *px=", *px)
	fmt.Println("failedUpdate px=", px)
}

func update(px *int) {
	*px = 23
	fmt.Println("update *px=", *px)
	fmt.Println("update px=", px)
}

func main() {
	x := 1
	fmt.Println("original x=", x)
	failedUpdate(&x)
	fmt.Println("after failedUpdate, x=", x)
	update(&x)
	fmt.Println("after update, x=", x)

	for i := 0; i < 10; i++ {
		fmt.Print("=")
	}
	fmt.Println()

	var y *int
	fmt.Println("original y=", y)
	failedUpdate(y)
	fmt.Println("after failedUpdate, y=", y)
	update(y)
	fmt.Println("after update, y=", y)

}
