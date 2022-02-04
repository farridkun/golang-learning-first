package main

import "fmt"

func main() {
	type noKTP string
	type status bool

	var ktpFarrid noKTP = "123456789"
	var statusPerkawinan status = true

	fmt.Println(ktpFarrid)
	fmt.Println(noKTP("1111111111"))
	fmt.Println(statusPerkawinan)
}