package main

import "fmt"

func main() {
	states := make(map[string]string)
	states["AK"] = "Alaska"
	states["CA"] = "California"
	states["FL"] = "Florida"

	fmt.Println(states["CA"])

	// Print all states
	for _, value := range states {
		fmt.Println(value)
	}
}
