package asciifiles

import (
	"fmt"
	"os"
	"strings"
)

func Asciiart(text, filename string) string {

	for _, word := range text {
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

	var userInput []string

	userInput = append(userInput, text)

	user3 := strings.Join(userInput, "")
	return printArtAscii(user3, newAscii)
}

func printArtAscii(userInput string, Ascii []string) string {
	empty := ""
	// loops through each row of individual ascii character in art file
	for line := 1; line <= 8; line++ {

		for _, character := range userInput {

			//Match users input(rune) with each row of Ascii Art
			skip := (character - 32) * 9
			// print the line from art file at the position specified
			//by calculation of skip to find the corresponding line in to users input
			fmt.Print(Ascii[line+int(skip)])

		}
		fmt.Println()
		// Tells the function to skip to the next line before commencing next loop

	}
	return empty

}
