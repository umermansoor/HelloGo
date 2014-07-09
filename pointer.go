package main

import "fmt"

func main() {
	x := 1
	increment(&x)
	fmt.Println(x)

	y := new(int)
	*y = 2
	increment(y)
	fmt.Println(*y)
}

func increment(x *int) {
	*x = *x + 1
}
