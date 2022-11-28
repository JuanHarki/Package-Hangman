package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func DispalyAsciiCharacter(tab_str []string, file string) {
	var letter int
	var tab []string
	text, err := os.Open("fichierTxt/" + file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(text)
	for scanner.Scan() {
		tab = append(tab, scanner.Text())
	}
	for k := 0; k < 8; k++ {
		for _, i := range tab_str {
			if i[0] > 96 {
				letter = (int(rune(i[0])) - 97) * 9
			} else {
				letter = -16
			}
			letter_position := 585 + letter
			fmt.Print(tab[letter_position+k])
		}
		fmt.Print("\n")
	}

}
