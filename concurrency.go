package main

import "fmt"

func loop10() {
	for i := 0; i<10; i++ {
		fmt.Println(i)
	}
}

func main() {
	for i:=0; i<10; i++ {
		go loop10()
	}
	fmt.Println("main")
	var input string
	fmt.Scanln(&input)
}
