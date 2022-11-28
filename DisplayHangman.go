package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func DisplayHangman(compteur int) {
	var tab []string
	first_line := (compteur - 1) * 8
	text, err := os.Open("fichierTxt/hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(text)
	for scanner.Scan() {
		tab = append(tab, scanner.Text())
	}
	for i := first_line; i < first_line+8; i++ {
		fmt.Println(tab[i])
	}
}
