package main

import "fmt"

//memanggil fungsi goodMorning
//fungsi goodMorning akan melakukan print selamat pagi + name yang didapat dari parameter fungsi
func main() {
	goodMorning("April")
}

// TODO: answer here
func goodMorning(name string) {
	fmt.Println("Selamat pagi ", name)
}
