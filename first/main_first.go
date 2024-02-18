package main

import (
	// "bufio"
	"fmt"
	"os"
	"strings"
)

var shipsValidateMap = map[string]uint8{
	"1": 4,
	"2": 3,
	"3": 2,
	"4": 1,
}

func main() {
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		fmt.Println(filename)

		for _, line := range strings.Split(string(data), "\n")[1:] {
			fmt.Println(checkValidity(strings.Split(line, " ")))
		}

		fmt.Println()
	}
}

func checkValidity(sequence []string) string {
	validityDict := make(map[string]uint8)

	for _, num := range sequence {
		validityDict[num]++
		if !checkNumber(num, validityDict[num]) {
			return "NO"
		}
	}

	return "YES"
}

func checkNumber(num string, count uint8) bool {
	return shipsValidateMap[num] < count
}
