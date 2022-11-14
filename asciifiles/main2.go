package asciifiles

import (
	"fmt"
	"os"
	"strings"
)

func Asciiart(text, fontname []string) string {

	text2 := strings.Join(text, "")
	font2 := strings.Join(fontname, "")

	for _, word := range text2 {
		if word > 128 {
			return ""
		}
	}

	for _, word := range text2 {
		if word == '\n' {
			word += '\n'
		}
	}

	// This reads the standard.txt file and checks for error
	bytes, err := os.ReadFile("./asciifiles/" + font2 + ".txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// splits text file by new line
	var newAscii []string

	if font2 == "thinker-toy" {
		newAscii = strings.Split(string(bytes), "\r\n")

	} else {
		newAscii = strings.Split(string(bytes), "\n")

	}
	var userInput []string
	var answer string

	newline := false

	for j := 0; j < len(text); j++ {

		if text[j] == "\r" {
			if len(text) != j+1 && text[j+1] == "\n" {
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
		userInput = append(userInput, text[j])

	}
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
