package asciifiles

import (
	"fmt"
	"os"
	"strings"
)

func Asciiart(text, fontname string) string {

	for _, word := range text {
		if word < 32 || word > 126 {
			return ""
		}
	}

	// This reads the standard.txt file and checks for error
	bytes, err := os.ReadFile("./asciifiles/" + fontname + ".txt")
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

			skip := (character - 32) * 9

			empty += Ascii[line+int(skip)]

		}
		empty += "\n"
		//	fmt.Println()
	}
	return empty

}
