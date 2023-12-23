package main

import "fmt"

func main() {
	type NoKTP string

	var ktp NoKTP = "3509110205950005"
	var contoh NoKTP = "2222222"

	fmt.Println(ktp)
	fmt.Println(NoKTP(contoh))
}
