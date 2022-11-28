package hangman

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(file string) []byte {
	var tab_byte []byte
	text, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(text)
	for scanner.Scan() {
		tab_byte = append(tab_byte, scanner.Bytes()...)
	}

	return tab_byte
}
