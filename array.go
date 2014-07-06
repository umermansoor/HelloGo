package main

import "fmt"

func main() {
	var x[5]int
	x[4] = 300
	fmt.Println(x)

	var y [5]float64
	y[0] = 100
	y[1] = 95
	y[2] = 90
	y[3] = 95
	y[4] = 95

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += y[i]
	}

	fmt.Println(total / float64(len(y)))

	for k,v := range y {
		fmt.Println("y[", k, "]", "=", v)
	}

	//var r []float64
	r := make([]float64, 5)
	
	for i := 0; i < 5; i++ {
		r[i] = float64(i+2)
	}

	fmt.Println(r[0:3])
}
