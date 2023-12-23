package main

import "fmt"

func main() {
	var name1 = "Amir"
	var name2 = "Ami"
	var comparison bool = name2 != name1
	fmt.Println(comparison)
	var nilaiAkhir = 90
	var absensi = 20
	var lulusNilai = nilaiAkhir > 60
	var lulusAbsen = absensi > 20

	var lulus bool = lulusNilai && lulusAbsen
	fmt.Println(lulus)
}
