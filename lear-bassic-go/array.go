package main

import "fmt"

func main() {

	var names [6]string
	names[0] = "Moh"
	names[1] = " "
	names[2] = "Faisal"
	names[3] = " "
	names[4] = "Amir"
	names[2] = "taek"
	fmt.Println(names)
	// cara langsung
	var nitas = [6]string{
		"Moh",
		" ",
		"Faisal",
		" ",
		"Amir",
	}
	nitas[3] = "memek"
	fmt.Println(len(nitas))
	fmt.Println(nitas)

}
