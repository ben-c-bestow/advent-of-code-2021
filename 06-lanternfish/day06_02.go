package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type sea struct {
	fish map[int]int
}

func nextDay(s *sea) {
	newFish := make(map[int]int)
	for i := 0; i < 9; i++ {
		newFish[i] = 0
	}
	for days, num := range s.fish {
		if days == 0 {
			newFish[8] += num
			newFish[6] += num
		} else {
			newFish[days-1] += num
		}
	}
	s.fish = newFish
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
	s := &sea{}
	s.fish = make(map[int]int)
	valstrs := strings.Split(string(fileBytes), ",")
	for i := 0; i < 9; i++ {
		s.fish[i] = 0
	}
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
		s.fish[val]++
	}
	for i := 0; i < 256; i++ {
		fmt.Println(i)
		nextDay(s)
	}
	total := 0
	for _, val := range s.fish {
		total += val
	}
	fmt.Println(total)
}
