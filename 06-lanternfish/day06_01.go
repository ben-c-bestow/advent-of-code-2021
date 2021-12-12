package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type fish struct {
	daysTillSpawn int
	//probably will want a "directSpawn" pointer next
}

type sea struct {
	f []*fish
}

func nextDay(s *sea) {
	var newFish []*fish
	for _, f := range s.f {
		if f.daysTillSpawn == 0 {
			newFish = append(newFish, &fish{8})
			f.daysTillSpawn = 6
		} else {
			f.daysTillSpawn--
		}
	}
	for _, f := range newFish {
		s.f = append(s.f, f)
	}
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

	valstrs := strings.Split(string(fileBytes), ",")
	s := &sea{}
	s.f = make([]*fish, 0, 2048)
	for _, vs := range valstrs {
		vv := strings.TrimSpace(vs)
		if vv == "" {
			continue
		}
		val, err := strconv.Atoi(vv)
		if err != nil {
			fmt.Println(vs)
			os.Exit(1)
		}
		s.f = append(s.f, &fish{val})
	}
	for i := 0; i < 80; i++ {
		fmt.Println(i)
		nextDay(s)
	}
	fmt.Println(len(s.f))
}
