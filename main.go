package main

import "fmt"

func main() {
	//server.Run()
	b := struct{ A, B, exp int }{
		1, 2, 3,
	}
	sliceTest := []int{1, 2, 3}
	fmt.Println(sliceTest)
	fmt.Println(b)
}
