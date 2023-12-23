package main

import "fmt"

func main() {
	var nilai = 70
	var absensi = 21

	var lulusNilai = nilai > 65
	var lulusAbsen = absensi > 20
	var lulus = lulusNilai && lulusAbsen

	fmt.Println(lulus)
}
