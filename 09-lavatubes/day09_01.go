package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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
	total := 0
	for i, line := range lines {
		if len(line) < 2 {
			continue
		}
		chars := strings.Split(line, "")
		for j := 0; j < len(chars); j++ {
			neighbors := make([]int, 0, 4)
			if i > 0 {
				tchars := strings.Split(lines[i-1], "")
				val, err := strconv.Atoi(tchars[j])
				errorDie(err)
				neighbors = append(neighbors, val)
			}
			if i < len(lines)-1 {
				tchars := strings.Split((lines[i+1]), "")
				val, err := strconv.Atoi(tchars[j])
				errorDie(err)
				neighbors = append(neighbors, val)
			}
			if j > 0 {
				val, err := strconv.Atoi(chars[j-1])
				errorDie(err)
				neighbors = append(neighbors, val)
			}
			if j < len(chars)-1 {
				val, err := strconv.Atoi(chars[j+1])
				errorDie(err)
				neighbors = append(neighbors, val)
			}
			lowest := true
			jval, err := strconv.Atoi(chars[j])
			errorDie(err)
			for _, n := range neighbors {
				if n <= jval {
					lowest = false
				}
			}
			if lowest {
				total += (jval + 1)
				fmt.Println(i, j, jval)
			}
		}
	}
	fmt.Println(total)
}
