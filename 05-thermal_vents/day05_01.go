package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	file  int
	value int
}

type column struct {
	rank    int
	members []*cell
}

type field struct {
	c []*column
}

type coord struct {
	x int
	y int
}

func punchHoleAt(f *field, c *coord) {
	for _, col := range f.c {
		if col.rank == c.x {
			for _, mem := range col.members {
				if mem.file == c.y {
					mem.value++
					return
				}
			}
			col.members = append(col.members, &cell{c.y, 1})
			return
		}
	}
	newcol := &column{rank: c.x}
	newcol.members = make([]*cell, 0, 1000)
	newcol.members = append(newcol.members, &cell{c.y, 1})
	f.c = append(f.c, newcol)
}

func strToCoord(cstr string) *coord {
	parts := strings.Split(cstr, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil
	}
	return &coord{x, y}
}

func drawline(f *field, linestr string) {
	fmt.Println(linestr)
	endpointstrs := strings.Split(linestr, " -> ")
	alpha := strToCoord(endpointstrs[0])
	beta := strToCoord(endpointstrs[1])
	if alpha.x == beta.x {
		var start, end int
		if alpha.y > beta.y {
			start = beta.y
			end = alpha.y
		} else {
			start = alpha.y
			end = beta.y
		}
		for j := start; j < end+1; j++ {
			// fmt.Print(alpha.x, j, " ")
			punchHoleAt(f, &coord{alpha.x, j})
		}
		// fmt.Println()
	} else if alpha.y == beta.y {
		var start, end int
		if alpha.x > beta.x {
			start = beta.x
			end = alpha.x
		} else {
			start = alpha.x
			end = beta.x
		}
		for i := start; i < end+1; i++ {
			// fmt.Print(i, alpha.y, " ")
			punchHoleAt(f, &coord{i, alpha.y})
		}
		// fmt.Println()
	}
}

func multiples(f *field) int {
	total := 0
	for _, col := range f.c {
		for _, mem := range col.members {
			if mem.value > 1 {
				total++
			}
		}
	}
	return total
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
	f := &field{make([]*column, 0, 1000)}
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		drawline(f, line)
	}
	fmt.Println(multiples(f))
}
