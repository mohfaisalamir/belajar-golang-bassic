package main

import "fmt"

func main() {
	fmt.Println("ini String")
	fmt.Println("Moh. Faisal Amir")
	fmt.Println(len("Moh. Faisal Amir"))
	fmt.Println("Moh. Faisal Amir"[0])         //  Bukan M yang ditampilkan melainkan 77 (byte M)
	fmt.Println(string("Moh. Faisal Amir"[0])) //  tambah sintakx "string" ==> string("String_anda"[])
}

// menghitgung jemlah karakter menggunakan len("string")
// mengambil karakter pada posisi yang di tentukan "string"[]
