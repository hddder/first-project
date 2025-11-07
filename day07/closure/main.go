package main

import "fmt"

func addr() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := addr()
	fmt.Println(ret(1))
}
