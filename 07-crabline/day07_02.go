package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
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
	vals := make([]int, 0, 512)
	for _, vs := range valstrs {
		vst := strings.TrimSpace(vs)
		val, err := strconv.Atoi(vst)
		if err == nil {
			vals = append(vals, val)
		}
	}
	fuel := make([]int, 0, 2048)
	min, max := MinMax(vals)
	for i := min; i < max; i++ {
		total := 0
		for _, v := range vals {
			steps := math.Abs(float64(v - i))
			total += int((math.Pow(steps, 2) + float64(steps)) / 2)
		}
		fuel = append(fuel, total)
	}
	minsteps, _ := MinMax(fuel)
	fmt.Println(minsteps)
}
