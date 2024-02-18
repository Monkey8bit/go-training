package main

import (
	"fmt"
	"os"
	"regexp"

	// "regexp"
	"strings"
)

const fiveSymbolsPattern = `\D\d{2}\D{2}`
const fourSymbolsPattern = `\D\d\D{2}`

var regexMap = map[int]string{
	4: fourSymbolsPattern,
	5: fiveSymbolsPattern,
}

func main() {
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		fmt.Println(filename)

		var carNumbersSequense []string

		for _, line := range strings.Split(string(data), "\n")[1:] {
			lineLength := len(line)

			// if lineLength%5 != 0 && lineLength%4 != 0 {
			// 	fmt.Println("-")
			// 	continue
			// }

			if lineLength < 6 {
				var sequence string

				if len(line) == 5 {
					sequence = line[:5]
				} else {
					sequence = line[:4]
				}

				if matched, _ := regexp.Match(regexMap[lineLength], []byte(sequence)); matched {
					carNumbersSequense = append(carNumbersSequense, sequence)
				}
			} else {
				for i := range line {
					fourSymbolsSequence := line[i : 4+i]

					if matched, _ := regexp.Match(regexMap[4], []byte(fourSymbolsSequence)); matched {
						carNumbersSequense = append(carNumbersSequense, fourSymbolsSequence)
					}

					if len(line) < i+5 {
						break
					}

					fiveSymbolsSequence := line[i : 5+i]

					if matched, _ := regexp.Match(regexMap[5], []byte(fiveSymbolsSequence)); matched {
						carNumbersSequense = append(carNumbersSequense, fiveSymbolsSequence)
					}
				}

			}

			if len(carNumbersSequense) > 0 && len(strings.Join(carNumbersSequense, "")) == lineLength {
				fmt.Println(strings.Join(carNumbersSequense, " "))
			} else {
				fmt.Println("-")
			}
			carNumbersSequense = carNumbersSequense[:0]
		}

		fmt.Println()
	}
}

func checkSequence(sequence string, size int) {

}
