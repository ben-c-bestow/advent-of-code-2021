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

	lines := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	total := 0
	for i, line := range lines {
		if len(line) < 2 {
			continue
		}
		chars := strings.Split(line, "")
		for j, _ := range chars {
			neighbors := make([]int, 0, 8)
			ichecks := make([]int, 0, 3)
			jchecks := make([]int, 0, 3)
			if i > 0 {
				ichecks = append(ichecks, i-1)
			}
			ichecks = append(ichecks, i)
			if i < len(lines)-1 {
				ichecks = append(ichecks, i+1)
			}
			if j > 0 {
				jchecks = append(jchecks, j-1)
			}
			jchecks = append(jchecks, j)
			if j < len(chars)-1 {
				jchecks = append(jchecks, j+1)
			}
			for _, ic := range ichecks {
				for _, jc := range jchecks {
					if i == ic && j == jc {
						continue
					}
					fmt.Println(ic, jc)
					var tchars []string
					if i != ic {
						tchars = strings.Split(lines[ic], "")
					} else {
						tchars = chars
					}
					val, err := strconv.Atoi(tchars[jc])
					if err != nil {
						fmt.Println(err.Error())
						os.Exit(1)
					}
					neighbors = append(neighbors, val)
				}
			}
			lowest := true
			jval, err := strconv.Atoi(chars[j])
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			for _, n := range neighbors {
				if n < jval {
					lowest = false
				}
			}
			if lowest {
				total += (jval + 1)
			}
		}
	}
	fmt.Println(total)
}
