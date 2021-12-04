package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

type node struct {
	x        int
	y        int
	value    int
	selected bool
	parent   *board
}

type bingo struct {
	group []*node
}

type board struct {
	dim     int
	n       map[coord]*node
	bingoes []*bingo
	winner  bool
}

func boardset(b *board) {
	for i := 0; i < b.dim; i++ {
		horiz := make([]*node, 0, b.dim)
		vert := make([]*node, 0, b.dim)
		for j := 0; j < b.dim; j++ {
			// fmt.Println("bset", i, j)
			horiz = append(horiz, b.n[coord{j, i}])
			vert = append(vert, b.n[coord{i, j}])
		}
		// fmt.Println("h", horiz)
		// fmt.Println("v", vert)
		b.bingoes = append(b.bingoes, &bingo{horiz})
		b.bingoes = append(b.bingoes, &bingo{vert})
		// if i == 0 {
		// 	diagright := make([]*node, 0, b.dim)
		// 	diagleft := make([]*node, 0, b.dim)
		// 	for j := 0; j < b.dim; j++ {
		// 		diagright = append(diagright, b.n[coord{j, j}])
		// 		diagleft = append(diagleft, b.n[coord{b.dim - j - 1, j}])
		// 	}
		// 	b.bingoes = append(b.bingoes, &bingo{diagright})
		// 	b.bingoes = append(b.bingoes, &bingo{diagleft})
		// }
	}
}

func checkForBingo(b *board) int {
	for _, bn := range b.bingoes {
		complete := true
		total := 0
		for _, n := range bn.group {
			total += n.value
			if !n.selected {
				complete = false
				break
			}
		}
		if complete {
			return total
		}
	}
	return -1
}

func valUnmarked(b *board) int {
	total := 0
	for _, space := range b.n {
		if !space.selected {
			total += space.value
		}
	}
	return total
}

func activate(b *board, i int) bool {
	for _, n := range b.n {
		if n.value == i {
			n.selected = true
			return true
			//assuming all values per board are unique
		}
	}
	return false
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

	chunks := strings.Split(string(fileBytes), "\n\n")
	draws := strings.Split(chunks[0], ",")
	var boards []*board
	for _, new_b := range chunks[1:] {
		if len(new_b) < 5 {
			continue
		}
		rows := strings.Split(new_b, "\n")
		row0 := strings.Fields(rows[0])
		thisbe := &board{dim: len(row0)}
		thisbe.winner = false
		thisbe.n = make(map[coord]*node)
		thisbe.bingoes = make([]*bingo, 0, 10)
		//increase to 12 for diagonals
		for j, row := range rows {
			valstrs := strings.Fields(row)
			for i, valstr := range valstrs {
				//new nodes have x, y, value, parent
				value, err := strconv.Atoi(valstr)
				if err != nil {
					fmt.Println("cell parse error")
				}
				thisun := &node{x: i, y: j, parent: thisbe, value: value, selected: false}
				thisbe.n[coord{i, j}] = thisun
			}
		}
		boardset(thisbe)
		boards = append(boards, thisbe)
	}
	for _, draw := range draws {
		fmt.Println("drew", draw)
		val, err := strconv.Atoi(draw)
		if err != nil {
			fmt.Println("drew an invalid value", draw)
		}
		for i, b := range boards {
			if b.winner {
				continue
			}
			if activate(b, val) {
				got := checkForBingo(b)
				if got != -1 {
					unmarked := valUnmarked(b)
					fmt.Println("board number", i+1)
					fmt.Println("score", val*unmarked)
					b.winner = true
				}
			}
		}
	}
}
