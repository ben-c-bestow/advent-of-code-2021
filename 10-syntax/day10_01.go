package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type stack struct {
	top *node
}

type node struct {
	next  *node
	value string
}

func push(s stack, n *node) {
	n.next = s.top
	s.top = n
}

func pop(s stack) string {
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
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	lines := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	firsterrors := make([]string, 0, 128)
	opens := "([{<"
	closes := ")]}>"
	for _, line := range lines {
		chars := strings.Split(line, "")
		s := stack{&node{value: chars[0]}}
		//assumption: all lines start with an open
		for _, char := range chars[1:] {
			if in(opens, char) {
				push(s, &node{value: char})
			} else if in(closes, char) {
				if (char == ")" && s.top.value == "(") ||
					(char == "]" && s.top.value == "[") ||
					(char == "}" && s.top.value == "}") ||
					(char == ">" && s.top.value == "<") {
					_ = pop(s)
				} else {
					firsterrors = append(firsterrors, char)
					break
				}
			}
		}
	}
	total := 0
	for _, e := range firsterrors {
		total += points[e]
	}
	fmt.Println(total)
}
