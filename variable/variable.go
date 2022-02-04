package main

import "fmt"

func basicVar() {
	var name string

	name = "Farrid Kuntoro"
	fmt.Println(name)

	name = "Farrid Kun"
	fmt.Println(name)
}

func varLangsung() {
	var namaDepan = "Farrid"
	fmt.Println(namaDepan)
	
	var umur = 22
	fmt.Println(umur)
}

func tanpaKataKunciVar() {
	namaBelakang := "Kuntoro"
	fmt.Println(namaBelakang)

	namaBelakang = "Putra"
	fmt.Println(namaBelakang)
}

func multipleVar() {
	var (
		namaDepan = "Farrid"
		namaBelakang = "Kuntoro"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}

func main(){
	basicVar()

	varLangsung()

	tanpaKataKunciVar()

	multipleVar()
}