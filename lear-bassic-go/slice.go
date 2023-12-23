package main

import "fmt"

func main() {
	// cara langsung
	var nitas = [6]string{
		"Moh",
		"hammad",
		"Faisal",
		"Abdullah",
		"Amir",
		"Hamka",
	}
	nitas[3] = "mikail"
	//nitas1 := nitas[:]
	/*	fmt.Println(nitas1)
		fmt.Println(len(nitas))
		fmt.Println(nitas)*/
	//nitass := [...]string{`zhenita`, `deliany`, "moh", "faisal", "amir", `daud`, `pradopo`, "hemas"}
	//nitass1 := append(nitas1, "daryanti")
	//fmt.Println(nitas1)
	//fmt.Println(nitass[:])
	//fmt.Println(nitass1)
	//fmt.Println(len(nitass1))
	var numbers = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := numbers[8:10]
	slice1[0] = 16
	slice1[1] = 12
	numbers[8] = 12
	numbers[9] = 16
	slice2 := append(slice1, 90, 100)
	//numbers2 := append(numbers, 80, 90, 100)
	//numbers2[0] = 12
	fromSlice := numbers[:]
	toSlice := make([]int, len(fromSlice), cap(fromSlice))
	copy(fromSlice, toSlice)
	fmt.Println(numbers)
	fmt.Println(slice1)
	fmt.Println(slice2)
	//fmt.Println(toSlice)
}
