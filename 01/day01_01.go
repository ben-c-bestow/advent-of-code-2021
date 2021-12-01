package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	var vals = make([]int, len(lines))
	for i, v := range lines {
		if iv, err := strconv.Atoi(v); err == nil {
			vals[i] = iv
		}
	}

	curr := -1
	delta := 0
	tally := 0
	for _, v := range vals {
		if curr == -1 {
			curr = v
			continue
		} else {
			delta = v - curr
			if delta > 0 {
				tally++
			}
			curr = v
		}
	}
	fmt.Println(tally)
}
