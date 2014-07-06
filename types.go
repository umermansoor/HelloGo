package main

import "fmt"

func main() {
	var phrase string = "go away duck"
	phrase2 := "alternate variable declaration & assignment" 
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("1 + 1 = ", 1.0 + 1.0)
	fmt.Println(len("Hello world"))
	fmt.Println(`Hello "World"`)
	fmt.Println("Hello World"[0])
	fmt.Println("Hello " + "World")
	fmt.Println(true && true)
	fmt.Println(!true)
	fmt.Println(phrase)
	phrase = "got any lemonade"
	fmt.Println(phrase)
	fmt.Println(phrase2)
}
