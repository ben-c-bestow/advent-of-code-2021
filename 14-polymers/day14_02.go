package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type node struct {
	symbol  string
	counted bool
	next    *node
}

func tally(freq map[string]int, head *node) map[string]int {
	pos := head
	freqout := freq
	for pos.next != nil {
		if !pos.counted {
			if _, ok := freq[pos.symbol]; ok {
				freqout[pos.symbol]++
			} else {
				freqout[pos.symbol] = 1
			}
			pos.counted = true
		}
		pos = pos.next
	}
	return freqout
}

func printList(head *node) {
	pos := head
	for pos != nil {
		fmt.Print(pos.symbol)
		pos = pos.next
	}
	fmt.Println()
}

func listToSlice(head *node) []string {
	s := make([]string, 0, 2048)
	pos := head
	for pos.next != nil {
		s = append(s, pos.symbol)
		pos = pos.next
	}
	return s
}

func errorDie(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
	os.Exit(1)
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
	chain := strings.TrimSpace(chunks[0])
	ru := strings.Split(strings.TrimSpace(chunks[1]), "\n")
	rules := make(map[string]string)
	for _, r := range ru {
		rule := strings.TrimSpace(r)
		parts := strings.Split(rule, " -> ")
		rules[parts[0]] = parts[1]
	}

	links := strings.Split(chain, "")
	freq := make(map[string]int)
	jumps := make(map[string]*node)
	for z := 0; z < 2; z++ {
		nulinks := make([]string, 0, 4196)
		for a := 0; a < len(links)-1; a++ {
			seed := links[a] + links[a+1]
			_, jok := jumps[seed]
			if jok && z == 1 {
				// fmt.Println(seed)
				// printList(jumps[seed])
				freq = tally(freq, jumps[seed])
				if a == len(links)-2 {
					special := &node{symbol: links[len(links)-1]}
					special.next = &node{symbol: "$"}
					freq = tally(freq, special)
				}
				continue
			}
			head := &node{symbol: links[a], counted: false}
			pos := head
			pos.next = &node{symbol: links[a+1], counted: false}
			pos = pos.next

			for x := 0; x < 10; x++ {
				pos = head
				for pos.next != nil {
					nxt := pos.next
					key := pos.symbol + nxt.symbol
					if insert, ok := rules[key]; ok {
						innode := &node{symbol: insert, counted: false}
						innode.next = nxt
						pos.next = innode
					}
					pos = nxt
				}
				freq = tally(freq, head)
				if a == len(links)-2 {
					special := &node{symbol: pos.symbol}
					special.next = &node{}
					freq = tally(freq, special)
				}
			}
			if z == 0 {
				fmt.Println(freq)
				if _, ok := jumps[seed]; !ok {
					jumps[seed] = &*head
					pos = jumps[seed]
					for pos != nil {
						pos.counted = false
						pos = pos.next
					}
				}
				// fmt.Println(seed, "->", strings.Join(listToSlice(head), ""))
				for _, char := range listToSlice(head) {
					nulinks = append(nulinks, char)
				}
				if a == len(links)-2 {
					nulinks = append(nulinks, links[a+1])
				}
				fmt.Println(len(strings.Join(nulinks, "")))

			}
		}
		links = nulinks
	}

	top := ""
	bottom := ""
	for char, f := range freq {
		if top == "" || f > freq[top] {
			top = char
		}
		if bottom == "" || f < freq[bottom] {
			bottom = char
		}
	}
	total := 0
	for _, f := range freq {
		total += f
	}
	fmt.Println(total)
	fmt.Println(top, freq[top], bottom, freq[bottom])
	fmt.Println(freq[top] - freq[bottom])
}
