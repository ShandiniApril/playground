package main

import "fmt"

func main() {
	//fungsi yang mengembalikan nilai pangkat dua dari parameter yang diterima
	//contoh input parameter yang diterima (3)
	//maka fungsi akan mengembalikan 9

	// TODO: answer here
	func square (x int) {
		fmt.Println(x * x)
	} (3)

	square()
	fmt.Println(square)
}
