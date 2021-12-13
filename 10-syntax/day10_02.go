package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type stack struct {
	top *node
}

type node struct {
	next  *node
	value string
}

func push(s *stack, n *node) {
	n.next = s.top
	s.top = n
}

func pop(s *stack) string {
	str := s.top.value
	s.top = s.top.next
	return str
}

func errorDie(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
}

func printStack(s *stack) {
	n := s.top
	for n != nil {
		fmt.Print(n.value)
		n = n.next
	}
	fmt.Println()
}

func in(haystack string, needle string) bool {
	post := len(needle) - 1
	hchars := strings.Split(haystack, "")
	for i := 0; i < (len(haystack) - post); i++ {
		hsub := hchars[i : i+1+post]
		hsusbtr := strings.Join(hsub, "")
		if hsusbtr == needle {
			return true
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

	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	lines := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	//firsterrors := make([]string, 0, 128)
	scores := make([]int, 0, 64)
	opens := "([{<"
	closes := ")]}>"
	for _, l := range lines {
		line := strings.TrimSpace(l)
		chars := strings.Split(line, "")
		s := &stack{&node{value: chars[0]}}
		//assumption: all lines start with an open
		illegal := false
		for _, char := range chars[1:] {
			// fmt.Printf("Top %v: Char: %v \n", s.top.value, char)
			if in(opens, char) {
				push(s, &node{value: char})
			} else if in(closes, char) {
				if (char == ")" && s.top.value == "(") ||
					(char == "]" && s.top.value == "[") ||
					(char == "}" && s.top.value == "{") ||
					(char == ">" && s.top.value == "<") {
					_ = pop(s)
				} else {
					//firsterrors = append(firsterrors, char)
					fmt.Println("SKIP")
					illegal = true
					break
				}
			}
		}
		if illegal {
			continue
		}
		//stated: all non-illegal lines are incomplete
		score := 0
		ttop := s.top
		for ttop != nil {
			fmt.Print(ttop.value)
			ttop = ttop.next
		}
		fmt.Println()
		for s.top != nil {
			printStack(s)
			if s.top.value == "(" {
				score = (score * 5) + points[")"]
			} else if s.top.value == "[" {
				score = (score * 5) + points["]"]
			} else if s.top.value == "{" {
				score = (score * 5) + points["}"]
			} else if s.top.value == "<" {
				score = (score * 5) + points[">"]
			} else {
				break
			}
			_ = pop(s)
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	fmt.Println(scores[(len(scores) / 2)])
}
