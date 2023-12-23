package main

import "fmt"

func main() {
	var name string
	name = "Moh. Faisal Amir"
	name = "Suwardi Suryoningrat"
	fmt.Println(name)
	var numb int
	numb = len("Moh. Faisal Amir")
	fmt.Println(numb)

	var nama = "Moh. Faisal Amir" // bisa lang sung tampa deklarasi Type, tapi jangnan menulis 'var name " maka akan error, tapi "var name "Amir" ',
	fmt.Println(nama)
	var angka = len("Hadji Oemar Said Tjokroaminoto")
	fmt.Println(angka)

	// di go tidak wajib menulis type data(kunci var)
	//asal saat membuat variable langsung menginisialisasi datanya: var = apa_kek
	//agar tidak perlu kata kunci var  kita perlu mengunakan kata kunci ":=" sat menginisialisasikan data pada variable tersebut contoh ;

	namae := "amir" // membuat := baru dengan jenis data yang sama tidak boleh, kecuali beda jenis data
	namae = "usman"
	fmt.Println(namae)
	number := len("amir")
	fmt.Println(number)

	// bikin var bisa skaligus banyak

	var (
		name2 = "zhenita deliany"
		num   = len("Zhenita Delinay")
	)

	fmt.Println(name2)
	fmt.Println(num)

}
