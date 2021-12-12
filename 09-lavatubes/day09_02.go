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
	field := make([][]uint8, len(lines))
	for i := 0; i < len(lines); i++ {
		field[i] = make([]uint8, len(lines[0]))
		chars := strings.Split(lines[i], "")
		for j, char := range chars {
			val, err := strconv.Atoi(char)
			errorDie(err)
			field[i][j] = uint8(val)
		}
	}
	lowpoints := make([]coord, 0, 1024)
	basins := make(map[coord][]coord)
	for i := 0; i < len(field); i++ {
		fmt.Println(i)
		for j, cell := range field[i] {
			if cell == 9 {
				continue
			}
			ii := i
			jj := j
			this_path := make([]coord, 0, 32)
			for {
				neighbors := make([]coord, 0, 4)
				if ii > 0 {
					neighbors = append(neighbors, coord{ii - 1, jj})
				}
				if ii < len(field)-1 {
					neighbors = append(neighbors, coord{ii + 1, jj})
				}
				if jj > 0 {
					neighbors = append(neighbors, coord{ii, jj - 1})
				}
				if jj < len(field[0])-1 {
					neighbors = append(neighbors, coord{ii, jj + 1})
				}
				this_coord := coord{ii, jj}
				lowest := this_coord
				this_path = append(this_path, this_coord)
				for _, n := range neighbors {
					if field[n.row][n.col] < field[lowest.row][lowest.col] {
						lowest = n
					}
				}
				if cEq(&lowest, &this_coord) {
					if this_basin, ok := basins[this_coord]; ok {
						for _, ppoint := range this_path {
							found := false
							for _, bpoint := range this_basin {
								if cEq(&bpoint, &ppoint) {
									found = true
									break
								}
							}
							if !found {
								this_basin = append(this_basin, ppoint)
							}
						}
						basins[this_coord] = this_basin
					} else {
						basins[this_coord] = this_path
						lowpoints = append(lowpoints, this_coord)
					}
					fmt.Printf("lowpoint %d,%d:  ", this_coord.col, this_coord.row)
					printCoords(this_path)
					break
				}
				ii = lowest.row
				jj = lowest.col
			}
		}
	}
	topthree := make([]int, 3)
	for _, basin := range basins {
		if len(basin) > topthree[0] {
			topthree[2] = topthree[1]
			topthree[1] = topthree[0]
			topthree[0] = len(basin)
		} else if len(basin) > topthree[1] {
			topthree[2] = topthree[1]
			topthree[1] = len(basin)
		} else if len(basin) > topthree[2] {
			topthree[2] = len(basin)
		}
	}
	fmt.Println(topthree, topthree[0]*topthree[1]*topthree[2])
	fmt.Println(total)
}
