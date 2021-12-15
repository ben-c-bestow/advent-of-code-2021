package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func errorDie(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
}

func coordToString(c *coord) string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}

func stringToCoord(s string) *coord {
	xy := strings.Split(s, ",")
	x, err := strconv.Atoi(xy[0])
	errorDie(err)
	y, err := strconv.Atoi(xy[1])
	errorDie(err)
	return &coord{x, y}
}

func printCoords(cs []coord) {
	for _, c := range cs {
		fmt.Printf("%d,%d; ", c.x, c.y)
	}
	fmt.Println()
	return
}

func keys(slice map[string]int) []string {
	ks := make([]string, 0, len(slice))
	for k := range slice {
		ks = append(ks, k)
	}
	return ks
}

func cEq(a *coord, b *coord) bool {
	if a.x == b.x && a.y == b.y {
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

	chunks := strings.Split(strings.TrimSpace(string(fileBytes)), "\n\n")
	hits := strings.Split(strings.TrimSpace(chunks[0]), "\n")
	folds := strings.Split(strings.TrimSpace(chunks[1]), "\n")
	grid := make(map[string]int)
	for _, h := range hits {
		hit := strings.TrimSpace(h)
		if len(hit) < 3 {
			continue
		}
		grid[hit] = 1
	}
	for i, f := range folds {
		fold := strings.TrimSpace(f)
		eq := strings.Split(fold, " ")[2]
		parts := strings.Split(eq, "=")
		axis := parts[0]
		value, err := strconv.Atoi(parts[1])
		errorDie(err)
		nugrid := make(map[string]int)
		for _, key := range keys(grid) {
			kcoord := stringToCoord(key)
			if axis == "x" {
				if kcoord.x > value {
					nux := value - (kcoord.x - value)
					nucoord := &coord{nux, kcoord.y}
					nugrid[coordToString(nucoord)] = 1
				} else {
					nugrid[key] = 1
				}
			} else {
				if kcoord.y > value {
					nuy := value - (kcoord.y - value)
					nucoord := &coord{kcoord.x, nuy}
					nugrid[coordToString(nucoord)] = 1
				} else {
					nugrid[key] = 1
				}
			}
		}
		grid = nugrid
		if i == 0 {
			fmt.Println(len(grid))
		}
	}
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			propkey := fmt.Sprintf("%d,%d", j, i)
			if _, ok := grid[propkey]; ok {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
