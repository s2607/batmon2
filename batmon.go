package main

import "fmt"
import "io/ioutil"

type battery struct {
	num   int
	name  string
	path  string
	ready bool
}
type batmonitor struct {
	dpath string
}

func (sys batmonitor) getbats() []battery {
	files, c := ioutil.ReadDir(sys.dpath)
	fmt.Println(c)
	bats := make([]battery, len(files))
	for c, f := range files {
		fmt.Println(c)
		bats[c].name = f.Name()
	}
	return bats
}
func (bat *battery) setpath(sys batmonitor) {
	if bat.name != "" {
		bat.path = sys.dpath + bat.name
	} else {

	}
	bat.ready = true
}

func main() {
	system := batmonitor{dpath: "/sys/class/power_supply/"}
	curbat := battery{name: "BAT1"}
	curbat.setpath(system)

	fmt.Println(curbat.path)
	fmt.Println("hello world")
}
