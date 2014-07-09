package main

import "fmt"

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println("---")
	anotherEven := makeEvenGenerator()
	fmt.Println(anotherEven())
	fmt.Println(anotherEven())
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return i
	}
}

