package main

import "fmt"

func main() {
	const name = 1 // const jika gak dipakek tidak error, beda dengan var
	const Lastname = "amir"
	//error
	/*	name = 1
		Lastname = "amir"*/
	fmt.Println(name, Lastname)

	//konversi dari integer ke integer
	var nilai32 int32 = 129
	var nilsi8 int8 = int8(nilai32)
	var nilai64 int64 = int64(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilsi8)
	fmt.Println(nilai64)

	//konversi dari integer ke string atau sebaliknya
	var name1 = "Moh. Faisal Amir"
	var M = name1[12]
	var mString = string(M)
	fmt.Println(name1)
	fmt.Println(M)
	fmt.Println(mString)

}
