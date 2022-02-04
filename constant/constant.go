package main

import "fmt"

func basicConstant() {
	const namaDepan string = "Farrid"
	const namaBelakang = "Kuntoro"
	const umur = 22
	
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
	fmt.Println(umur)
}

func simpleContant() {
	const (
		namaDepan string = "Farrid"
		namaBelakang = "Kuntoro"
		umur = 22
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
	fmt.Println(umur)
}

func main() {
	basicConstant()

	simpleContant()
}