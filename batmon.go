package main

import "fmt"
import "io/ioutil"

type battery struct {
	num   int
	name  string
	path  string
	ready bool
	sys   batmonitor
}
type batmonitor struct {
	dpath   string
	bats    []battery
	mainbat int
}

func (sys *batmonitor) getbats() ([]battery, error) {
	files, e := ioutil.ReadDir(sys.dpath)
	if e != nil {
		return nil, e
	}
	sys.bats = make([]battery, len(files))
	for c, f := range files {
		sys.bats[c].name = f.Name()
		sys.bats[c].num = c
		sys.bats[c].sys = *sys
	}
	return sys.bats, nil
}
func (sys batmonitor) listbats() {
	for c := range sys.bats {
		fmt.Println(sys.bats[c])
	}
}
func (bat battery) String() string {
	return fmt.Sprintf("num:%d\t cap:%d\t name:%s", bat.num, bat.capacity(), bat.name)
}
func (bat battery) capacity() int {
	return 100
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
	system.getbats()
	system.listbats()

}
