package asciifiles

import (
	"fmt"
	"os"
	"strings"
)

func Asciiart(text, filename string) string {

	letters := strings.Split(text, "")

	//input := os.Args[1:]

	// goes through each character in the input and checks if its outside bounds of
	// printable character(32-126), if so then print nothing
	// (this is because runes can range to unicode which has way more characters)
	for _, word := range letters[0] {
		if word < 32 || word > 126 {
			return ""
		}
	}

	// This reads the standard.txt file and checks for error
	bytes, err := os.ReadFile("./asciifiles/" + filename + ".txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// splits text file by new line
	newAscii := strings.Split(string(bytes), "\n")

	// creates blank array of runes to be appeneded with ascii equivalent of the users input
	var userInput []rune
	// create Boolean for new line set to false (\n)

	// loop through each character of user input
	for _, character := range letters[0] {

		// appends corresponding asciiart characters of letters in users input into blank rune array
		userInput = append(userInput, character)
	}
	// users input and re-formatted standard.txt file passed as aruguments in pritnart function

	return printArtAscii(userInput, newAscii)
}
func printArtAscii(userInput []rune, Ascii []string) string {
	empty := ""
	// loops through each row of individual ascii character in art file
	for line := 1; line <= 8; line++ {
		// Loop through each character of users input
		for _, character := range userInput {
			//Match users input(rune) with each row of Ascii Art
			skip := (character - 32) * 9
			// print the line from art file at the position specified
			//by calculation of skip to find the corresponding line in to users input
			empty += Ascii[line+int(skip)]

		}

		// Tells the function to skip to the next line before commencing next loop

	}
	return empty
}
