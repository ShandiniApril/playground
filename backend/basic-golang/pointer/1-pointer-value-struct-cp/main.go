package main

import (
	"fmt"
	//"golang.org/x/tools/go/pointer"
)

//buat struct Rectangle dengan dua atribut yaitu Width dan Length
// tambah dua method :
// SetWidthPointer(width int) dan SetWidthValue(width int)
// SetWidthPointer(width int) untuk mengubah width dengan pointer receiver
// SetWidthValue(width int) untuk mengubah width dengan value

// TODO: answer here
type Rectangle struct {
	Width  int
	Length int
}

func (r *Rectangle) SetWidthPointer(newPointer int) {
	r.Width = newPointer
}

func (r Rectangle) SetWidthValue(newValue int) {
	r.Width = newValue
}

func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20

	fmt.Println("sebelum melakukan set width dengan pointer", r.Width)
	r.SetWidthPointer(30)
	fmt.Println("sesudah melakukan set width dengan pointer", r.Width)
	r.SetWidthValue(70)
	fmt.Println("sesudah melakukan set width dengan value", r.Width)
}
