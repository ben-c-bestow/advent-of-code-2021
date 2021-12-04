package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
	bits := float64(len(lines[1]))
	var digits []string
	reallines := 0
	onescount := make([]int, int(bits))
	for _, line := range lines {
		if len(line) < 5 {
			continue
		}
		reallines++
		digits = strings.Split(line, "")
		for i, digit := range digits {
			if digit == "1" {
				onescount[i]++
			}
		}
	}

	gammabin := make([]string, len(onescount))
	for i, v := range onescount {
		if v > (reallines / 2) {
			gammabin[i] = "1"
		} else {
			gammabin[i] = "0"
		}
	}
	gamma, err := strconv.ParseInt(strings.Join(gammabin, ""), 2, 64)
	maxval := (math.Pow(2.0, bits) - 1)
	fmt.Println("gamma", gamma)
	if err == nil {
		fmt.Println(gamma * (int64(maxval) - gamma))
	} else {
		fmt.Println("error parsing")
	}
}
