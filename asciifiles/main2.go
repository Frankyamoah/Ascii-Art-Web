package asciifiles

import (
	"fmt"
	"os"
	"strings"
)

func Asciiart(text, fontname []string) string {

	//turn the input text into a string
	text2 := strings.Join(text, "")
	//range over each character in string to see if printable
	for _, word := range text2 {
		if word > 128 {
			return ""
		}
	}

	//turn chosen fontname to string
	font2 := strings.Join(fontname, "")
	// This reads the file of fontname given and checks for error
	bytes, err := os.ReadFile("./asciifiles/" + font2 + ".txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	//varibale to store the font/banner files after they've been split
	var newAscii []string
	//have to treat thinkertoy different to standard and shadow
	if font2 == "thinker-toy" {
		//for some reason thinkertoy needs to be split by '\r\n' otherwise will print gibberish
		newAscii = strings.Split(string(bytes), "\r\n")

	} else {
		//cannot be split by '\r\n' otherwise will print gibberish
		newAscii = strings.Split(string(bytes), "\n")

	}
	//to append each character to slice
	var userInput []string

	//to concatenate with printasciiart function when there is a new line
	//this is because for every \n encountered the string array will be
	//wiped to begin the line again without having stored the result before the \n
	var answer string

	newline := false

	for i := 0; i < len(text); i++ {

		//if "\r\n" encountered and not end of string then newline=true and move on
		if text[i] == "\r" {
			if len(text) != i+1 && text[i+1] == "\n" {
				newline = true
				continue

			}
		}
		if newline {
			newline = false
			answer += printArtAscii(userInput, newAscii)
			userInput = []string{}
			continue

		}
		userInput = append(userInput, text[i])

	}
	//concatenate answer(outputs that have been stored everytime newline encountered)
	//with whatever is after the final newline character
	answer += printArtAscii(userInput, newAscii)
	return answer
}

func printArtAscii(userInput []string, Ascii []string) string {
	empty := ""
	// loops through each row of individual ascii character in art file
	for line := 1; line <= 8; line++ {

		for _, word := range userInput {
			for _, character := range word {

				skip := (character - 32) * 9

				empty += Ascii[line+int(skip)]
			}

		}
		empty += "\n"
		//fmt.Println()
	}
	return empty

}
