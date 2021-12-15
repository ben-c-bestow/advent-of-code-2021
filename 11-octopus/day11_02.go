package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	row int
	col int
}

func errorDie(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
}

func printCoords(cs []coord) {
	for _, c := range cs {
		fmt.Printf("%d,%d; ", c.col, c.row)
	}
	fmt.Println()
	return
}

func cEq(a *coord, b *coord) bool {
	if a.col == b.col && a.row == b.row {
		return true
	} else {
		return false
	}
}

func noTens(field [][]int) bool {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			if field[i][j] > 9 {
				return false
			}
		}
	}
	return true
}

func turn(field [][]int) ([][]int, int) {
	flashes := 0
	flashlocs := make([]*coord, 0, 64)
	nufield := make([][]int, len(field))
	for i := 0; i < len(field); i++ {
		nufield[i] = make([]int, len(field[0]))
		for j := 0; j < len(field[0]); j++ {
			nufield[i][j] = field[i][j] + 1
		}
	}
	for !noTens(nufield) {
		// fieldPrint(delta)
		for i := 0; i < len(field); i++ {
			for j := 0; j < len(field[0]); j++ {
				if nufield[i][j] > 9 {
					flashes++
					flashlocs = append(flashlocs, &coord{i, j})
					field[i][j] = 0
					istart := i
					iend := i
					jstart := j
					jend := j
					if i > 0 {
						istart = i - 1
					}
					if i < len(field)-1 {
						iend = i + 1
					}
					if j > 0 {
						jstart = j - 1
					}
					if j < len(field[0])-1 {
						jend = j + 1
					}
					for ii := istart; ii <= iend; ii++ {
						for jj := jstart; jj <= jend; jj++ {
							if ii == i && jj == j {
								continue
							}
							nufield[ii][jj]++
						}
					}
				}
			}
		}
		for _, fl := range flashlocs {
			nufield[fl.row][fl.col] = 0
		}
	}
	return nufield, flashes
}

func fieldPrint(field [][]int) {
	for _, line := range field {
		for _, cell := range line {
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}
	fmt.Println()
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
	field := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		field[i] = make([]int, len(lines[0]))
		chars := strings.Split(lines[i], "")
		for j, char := range chars {
			val, err := strconv.Atoi(char)
			errorDie(err)
			field[i][j] = int(val)
		}
	}
	var flashes int
	a := 0
	for flashes < (len(field) * len(field[0])) {
		field, flashes = turn(field)
		total += flashes
		a++
	}
	fmt.Println(a)
}
