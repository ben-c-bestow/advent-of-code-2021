package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type cave struct {
	name        string
	small       bool
	connects    []*cave
	routesToEnd [][]*cave
}

type field struct {
	caves []*cave
	start *cave
	end   *cave
}

func addCaves(f *field, a string, b string) {
	acave := caveByName(f, a)
	if acave == nil {
		acave = &cave{name: a}
		f.caves = append(f.caves, acave)
		if string(a[0]) == strings.ToLower(string(a[0])) {
			fmt.Println(a, " small")
			acave.small = true
		} else {
			acave.small = false
		}
		if a == "start" {
			f.start = acave
		} else if a == "end" {
			f.end = acave
		}
	}
	bcave := caveByName(f, b)
	if bcave == nil {
		bcave = &cave{name: b}
		f.caves = append(f.caves, bcave)
		if string(b[0]) == strings.ToLower(string(b[0])) {
			fmt.Println(b, " small")
			bcave.small = true
		} else {
			bcave.small = false
		}
		if b == "start" {
			f.start = bcave
		} else if b == "end" {
			f.end = bcave
		}
	}
	acave.connects = append(acave.connects, bcave)
	bcave.connects = append(bcave.connects, acave)
}

func caveByName(f *field, name string) *cave {
	for _, c := range f.caves {
		if c.name == name {
			return c
		}
	}
	return nil
}

func cavenames(c []*cave) []string {
	n := make([]string, 0, 64)
	for _, cc := range c {
		n = append(n, cc.name)
	}
	return n
}

func getRoutes(f *field) [][]*cave {
	for _, c := range f.end.connects {
		routes := make([][]*cave, 0, 256)
		one_step := []*cave{f.end}
		routes = append(routes, one_step)
		c.routesToEnd = routes
	}
	cs := f.end.connects
	for len(cs) > 0 {
		next_cs := make([]*cave, 0, 1028)
		for _, c := range cs {
			for _, connect := range c.connects {
				addconnect := false
				//c says: connect, you can add c+route for all of my routes as long as they follow rules
				for _, route := range c.routesToEnd {
					if connect.small {
						if inRoute(route, connect.name) {
							continue
						}
					}
					proproute := make([]*cave, 0, 32)
					proproute = append(proproute, c)
					for _, elem := range route {
						proproute = append(proproute, elem)
					}
					routeexists := false
					for _, existingroute := range connect.routesToEnd {
						if routesMatch(proproute, existingroute) {
							routeexists = true
							break
						}
					}
					if routeexists {
						continue
					}
					addconnect = true
					connect.routesToEnd = append(connect.routesToEnd, proproute)
				}
				if addconnect && connect.name != "start" {
					next_cs = append(next_cs, connect)
				}
			}
		}
		cs = next_cs
	}
	return f.start.routesToEnd
}

func in(haystack []string, needle string) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}
	return false
}

func routesMatch(a []*cave, b []*cave) bool {
	anames := cavenames(a)
	bnames := cavenames(b)
	sort.Strings(anames)
	sort.Strings(bnames)
	if len(anames) != len(bnames) {
		return false
	}
	for i, an := range anames {
		if an != bnames[i] {
			return false
		}
	}
	return true
}

func inRoute(c []*cave, name string) bool {
	for _, cc := range c {
		if cc.name == name {
			return true
		}
	}
	return false
}

func errorDie(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
		os.Exit(0)
	}

	fileName := os.Args[1]

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	f := &field{}
	for _, l := range lines {
		line := strings.TrimSpace(l)
		endpoints := strings.Split(line, "-")
		addCaves(f, endpoints[0], endpoints[1])
	}
	finalroutes := getRoutes(f)
	for _, r := range finalroutes {
		fmt.Println(cavenames(r))
	}
	fmt.Println(len(finalroutes))
}
