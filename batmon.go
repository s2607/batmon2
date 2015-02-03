package main

import "fmt"

type battery struct {
	num      int
	name     string
	path     string
	capacity int
}

func (bat *battery) open() {
	bat.path = "/sys/" + bat.name
}

func main() {
	curbat := battery{name: "BAT1"}
	curbat.open()

	fmt.Println(curbat.path)
	fmt.Println("hello world")
}
