package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"collections"
	"digits"
	"output"
)

func loadChars() (map[string]string) {
	// Define output map
	charSmallToBig := make(map[string]string)

	// Define spliting condition
	emptyLine, err := regexp.Compile("(?m)(^\n|\n$)")
	if err != nil {
		panic(err)
	}

	// Split by empty lines
	charachterLines := emptyLine.Split(digits.DigitsString, -1)

	// Find nested building chars
	charsSmall := collections.Unique(strings.Split(digits.DigitsString, ""))

	// Find match between "small" (one char) to big (multiline) representation
	for _, charSmall := range charsSmall {
		for _, charBig := range charachterLines {
			if strings.Contains(charBig, charSmall) {
				charSmallToBig[charSmall] = charBig
				break
			}
		}
	}
	return charSmallToBig
}

func defineShape(chars []string) (int, int) {
	maxHeight := 0
	maxWidth := 0

	for _, c := range chars {
		cl := strings.Split(c, "\n")
		curHeigth := len(cl)
		if curHeigth > maxHeight {
			maxHeight = curHeigth
		}
		for _, l := range cl {
			curWidth := len(l)
			if curWidth > maxWidth {
				maxWidth = curWidth
			}
		}
	}
	return maxHeight, maxWidth
}

func main() {
	// Validate args amount
	if len(os.Args) != 2 {
		fmt.Printf("Wrong amount of args %d\n", len(os.Args) - 1)
		os.Exit(1)
	}

	input := os.Args[1]
	// Validate args as uint64
	inputUint64, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		fmt.Printf("Wrong input number. Must be unsigned integer. Max value: %d\n", ^uint64(0))
		os.Exit(1)
	}
	input = strconv.FormatUint(inputUint64, 10)

	charRepr := loadChars()
	inputCollection := strings.Split(input, "")

	inputCollectionCopy := make([]string, len(inputCollection))
	copy(inputCollectionCopy, inputCollection)
	inputCollectonUnique := collections.Unique(inputCollectionCopy)

	// difference to validate input

	keys := make([]string, 0, len(charRepr))
	values := make([]string, 0, len(charRepr))
	for k, v := range charRepr {
		keys = append(keys, k)
		values = append(values, v)
	}

	diff := collections.Difference(inputCollectonUnique, keys)
	// Verify that input chars match available in program
	if len(diff) > 0 {
		fmt.Printf("Not correct input symbols: %s\n", os.Args[1])
		os.Exit(1)
	}

	height, width := defineShape(values)
	output.Print(inputCollection, charRepr, height, width)
}
