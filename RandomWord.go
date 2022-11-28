package hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func RandomWord(file string) string {
	var random_word string
	var random_number int
	var tab []string
	text, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(text)
	for scanner.Scan() {
		tab = append(tab, scanner.Text())
	}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	random_number = r.Intn(len(tab))
	random_word = tab[random_number]
	random_number = 0
	return random_word
}
