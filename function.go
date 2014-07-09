package main

import "fmt"

func main() {
	//not_implemented()
	xs := []float64{98, 96, 97, 95}
	fmt.Println(average(xs))
	x,y := two_values()
	fmt.Println(x, y)
	closure1()
	fh := closure2()
	fmt.Println(fh())
}

func not_implemented() {
	panic("not implemented")
}

func two_values() (int, int) {
	return 3,4
}

//create a closure

func closure1() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1,1))

}

func closure2() func() int{

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	return increment
}

func average (x []float64) float64 {
	total := 0.0
	for _, v := range x {
		total += v
	}

	return total / float64(len(x))
}
