package main

import "fmt"
import "io/ioutil"
import "strconv"

type battery struct {
	num   int
	name  string
	ready bool
	sys   batmonitor
}
type batmonitor struct {
	dpath   string
	bats    []battery
	mainbat int
}

func eprint(lvl int, a ...interface{}) {
	if lvl > 0 && a != nil && a[0] != nil {
		fmt.Println(a)
	}
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
	acp, e := ioutil.ReadFile(bat.path() + "/capacity")
	acp = acp
	if e != nil {
		eprint(1, e)
		return -1
	}
	i, e := strconv.ParseInt(string(acp[0:len(acp)-1]), 0, 32)
	eprint(1, e)
	return int(i)
}
func (bat *battery) path() string {
	if bat.name != "" {
		return bat.sys.dpath + bat.name
	} else {
		files, e := ioutil.ReadDir(bat.sys.dpath)
		if e != nil {
			eprint(1, e)
			return ""
		}
		bat.name = files[bat.num].Name()
		return bat.sys.dpath + bat.name
	}
}

func main() {
	system := batmonitor{dpath: "/sys/class/power_supply/"}
	system.getbats()
	system.listbats()

}
