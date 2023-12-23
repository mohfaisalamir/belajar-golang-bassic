package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = b * a

	c += 90
	fmt.Println(c)
	c += 16
	fmt.Println(c)
}
