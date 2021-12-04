package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Readout struct {
	Value string
	Valid bool
}

func valids(readouts []*Readout) []*Readout {
	var result []*Readout
	for _, r := range readouts {
		if r.Valid {
			result = append(result, r)
		}
	}
	fmt.Println()
	return result
}

func is_final(readouts []*Readout) bool {
	var firstval = readouts[0].Value
	for _, r := range readouts {
		if r.Value != firstval {
			return false
		}
	}
	return true
}

func findFilter(chars []string, high bool) string {
	ones := 0
	for _, char := range chars {
		if char == "1" {
			ones++
		}
	}
	majority_ones := ones > (len(chars) / 2)
	if ones == (len(chars) / 2) {
		majority_ones = high
	}
	fmt.Println(ones, len(chars), majority_ones == high)
	if majority_ones == high {
		return "1"
	} else {
		return "0"
	}
}

func filterReadout(readouts []*Readout, high bool) string {
	firstline := readouts[0].Value
	to_filter := valids(readouts)
	chars := make([]string, len(to_filter))
	for i := 0; i < len(firstline); i++ {
		chars = nil
		fmt.Println("remaining: ", len(to_filter))
		for _, r := range to_filter {
			word := strings.Split(r.Value, "")
			chars = append(chars, word[i])
		}
		keychar := findFilter(chars, high)
		fmt.Println(high, i, keychar)
		for _, r := range to_filter {
			word := strings.Split(r.Value, "")
			fmt.Println(r.Value, word[i])
			if word[i] != keychar {
				r.Valid = false
			}
		}
		to_filter = valids(to_filter)
		if is_final(to_filter) {
			return to_filter[0].Value
		}
	}
	fmt.Println("error")
	return "Error"
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

	lines := strings.Split(string(fileBytes), "\n")
	var oreads []*Readout
	var creads []*Readout
	for _, line := range lines {
		if len(line) > 1 {
			o := &Readout{line, true}
			c := &Readout{line, true}
			oreads = append(oreads, o)
			creads = append(creads, c)
		}
	}
	oxstring := filterReadout(oreads, true)
	co2string := filterReadout(creads, false)
	oxval, err := strconv.ParseInt(oxstring, 2, 64)
	co2val, err := strconv.ParseInt(co2string, 2, 64)
	fmt.Println(oxval, co2val, oxval*co2val)
}
