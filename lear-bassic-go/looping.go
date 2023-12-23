package main

import "fmt"

func showUniqueNumb(len int) {
	for I := 0; I <= len; I++ {
		if I%2 != 0 && I%5 != 0 && I%8 != 0 {
			fmt.Println(I)
		}
	}
}

func main() {
	showUniqueNumb(20)
}
