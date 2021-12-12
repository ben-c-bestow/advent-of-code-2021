package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

var SEGORDER = []string{
	"top",
	"nw",
	"center",
	"sw",
	"bottom",
	"se",
	"ne",
}

func boolToString(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

func parse(strmap map[int]string, instr string) int {
	for digit, jumble := range strmap {
		jslice := strings.Split(jumble, "")
		islice := strings.Split(instr, "")
		sort.Strings(jslice)
		sort.Strings(islice)
		jsorted := strings.Join(jslice, "")
		isorted := strings.Join(islice, "")
		if jsorted == isorted {
			return digit
		}
	}
	return -1
}

func inANotinB(a string, b string) []string {
	fmt.Println(a, b)
	result := make([]string, 0, 7)
	achars := strings.Split(a, "")
	for _, ac := range achars {
		fmt.Println("\t", b, ac)
		if !in(b, ac) {
			result = append(result, ac)
		}
	}
	fmt.Println("result:", result)
	fmt.Println()
	return result
}

func missingSegs(a string) []string {
	every := "abcdefg"
	return inANotinB(every, a)
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

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
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
	hardcount := 0
	for _, l := range lines {
		line := strings.TrimSpace(l)
		fmt.Println("\n\n\n\n\n")
		fmt.Println(line)
		if len(line) < 1 {
			continue
		}
		chunks := strings.Split(line, " | ")
		unique := strings.Split(chunks[0], " ")

		strkey := make(map[int]string)
		schema := make(map[string]string)
		for _, loc := range SEGORDER {
			schema[loc] = ""
		}
		fivesegs := make([]string, 0, 3)
		sixsegs := make([]string, 0, 3)
		for _, u := range unique {
			if len(u) == 2 {
				strkey[1] = u
			} else if len(u) == 3 {
				strkey[7] = u
			} else if len(u) == 4 {
				strkey[4] = u
			} else if len(u) == 7 {
				strkey[8] = u
			} else if len(u) == 5 {
				fivesegs = append(fivesegs, u)
			} else if len(u) == 6 {
				sixsegs = append(sixsegs, u)
			}
			fmt.Println("len", len(u), u)
		}
		topseg := inANotinB(strkey[7], strkey[1])
		schema["top"] = topseg[0]
		//top is the only seg diff bt 7 and 1
		for _, str := range sixsegs {
			diff4 := inANotinB(strkey[4], str)
			diff7 := inANotinB(strkey[7], str)
			fmt.Println("diff4", diff4)
			if len(diff4) == 0 {
				//9 is the only 6seg glyph that contains all segs from 4
				strkey[9] = str
				fmt.Println("4", strkey[4], "9", strkey[9])
				x := missingSegs(str)
				schema["sw"] = x[0]
			} else if len(diff7) == 0 {
				strkey[0] = str
				//0 and 9 have all segments from 7, we just tested for 9
				x := missingSegs(str)
				schema["center"] = x[0]
			} else {
				strkey[6] = str
			}
		}
		for _, str := range fivesegs {
			threetest := inANotinB(strkey[1], str)
			if in(str, schema["sw"]) {
				strkey[2] = str
				//2 is the only 5seg that uses southwest
				fmt.Println(strkey[7], strkey[2])
				seseg := inANotinB(strkey[7], strkey[2])
				schema["se"] = seseg[0]
				//7 has southeast, 2 does not
				neseg := inANotinB(strkey[1], schema["se"])
				schema["ne"] = neseg[0]
				//whichever seg in 1 is not se is ne
			} else if len(threetest) == 0 {
				strkey[3] = str
				ms3 := strings.Join(missingSegs(strkey[3]), "")
				//missing segs are sw and nw
				nwseg := inANotinB(ms3, schema["sw"])
				schema["nw"] = nwseg[0]
				arch := strkey[7] + ms3
				//these together make an arch shape, 0 without bottom
				botseg := inANotinB(strkey[0], arch)
				schema["bottom"] = botseg[0]
			} else {
				strkey[5] = str
			}
		}

		fmt.Println("\n\n", strkey)
		data := strings.Split(chunks[1], " ")
		fmt.Println(chunks[1])
		for i, datum := range data {
			mult := math.Pow10(3 - i)
			one := parse(strkey, datum)
			fmt.Println(datum, ":", mult, one)
			hardcount += int(mult) * one
		}
	}
	fmt.Println(hardcount)
}
