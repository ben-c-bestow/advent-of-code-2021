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
	easycount := 0
	for _, l := range lines {
		line := strings.TrimSpace(l)
		if len(line) < 1 {
			continue
		}
		chunks := strings.Split(line, " | ")
		//unique := strings.Split(chunks[0], " ")
		data := strings.Split(chunks[1], " ")
		for _, datum := range data {
			if len(datum) == 2 || len(datum) == 3 || len(datum) == 4 || len(datum) == 7 {
				easycount++
			}
		}
	}
	fmt.Println(easycount)
}
