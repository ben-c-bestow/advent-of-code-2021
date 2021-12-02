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

	tally := 0
	curr := 0
	prev := 0
	for i, v := range vals {
		if i < 3 {
			continue
		}
		curr = v + vals[i-1] + vals[i-2]
		prev = vals[i-1] + vals[i-2] + vals[i-3]
		if curr > prev {
			tally++
		}
	}
	fmt.Println(tally)
}
