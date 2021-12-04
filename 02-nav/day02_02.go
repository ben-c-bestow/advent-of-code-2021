package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

	lines := strings.Split(string(fileBytes), "\n")

	curr_depth := 0
	curr_xdist := 0
	curr_aim := 0
	cmd := ""
	cval := 0
	for _, line := range lines {
		cmd = ""
		cval = 0
		fmt.Sscan(line, &cmd, &cval)
		fmt.Println(cmd, cval)
		if cmd == "forward" {
			curr_xdist += cval
			curr_depth += (curr_aim * cval)
		} else if cmd == "up" {
			curr_aim -= cval
		} else if cmd == "down" {
			curr_aim += cval
		} else {
			continue
		}
	}
	fmt.Println(curr_depth * curr_xdist)
}
