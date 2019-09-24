package output

import (
	"fmt"

	"strings"
)

const asterics = "*"
const newLine = "\n"
const space = " "

func printBorders(length int) {
	for i := 0; i < length; i++ {
		fmt.Printf(asterics)
	}
}

func printNewLine(amount int) {
	for i := 0; i < amount; i++ {
		fmt.Printf(newLine)
	}
}

func Print(charCollection []string, charRepr map[string]string, height, width int) {
	// define width and height of big charachter
	var charOutput []string
	outputLength := width*len(charCollection) + len(charCollection) - 1
	printBorders(outputLength)

	printNewLine(2)
	for _, ch := range charCollection {
		charOutput = append(charOutput, charRepr[ch])
	}
	for line := 0; line < height; line++ {
		for _, char := range charOutput {
			lines := strings.Split(char, newLine)
			fmt.Printf(lines[line])
			if diff := width - len(lines[line]); diff > 0 {
				for c := 0; c < diff; c++ {
					fmt.Printf(space)
				}
			}
			fmt.Printf(space)
		}
		printNewLine(1)
	}
	printNewLine(1)
	printBorders(outputLength)
	printNewLine(1)
}
