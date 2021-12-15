package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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
	for x := 0; x < 10; x++ {
		fmt.Println(x)
		nuchain := links[0]
		for i := 0; i < len(links); i++ {
			if i == len(links)-1 {
				break
			}
			key := strings.Join(links[i:i+2], "")
			if insert, ok := rules[key]; ok {
				nuchain = nuchain + insert
			}
			nuchain = nuchain + links[i+1]
		}
		links = strings.Split(nuchain, "")
	}
	freq := make(map[string]int)
	top := ""
	bottom := ""
	for _, link := range links {
		if _, ok := freq[link]; ok {
			freq[link]++
		} else {
			freq[link] = 1
		}
		if top == "" || freq[link] > freq[top] {
			top = link
		}
		if bottom == "" || freq[link] < freq[bottom] {
			bottom = link
		}
	}
	fmt.Println(len(links), "\n", strings.Join(links, ""))
	fmt.Println(top, freq[top], bottom, freq[bottom])
	fmt.Println(freq[top] - freq[bottom])
}
