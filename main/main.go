package main

import (
	"../words"
	"fmt"
	"github.com/golang-collections/collections/set"
	"strings"
)

func main() {

	count := 0
	genWords := 0

	newWords := set.New()

	for {
		// word := words.RandomWord(5)
		word := words.RandomWordMinMax(7, 8)
		// word := words.RandomWordSimple(10)

		// if len(word) > 7 {
		// 	continue
		// }

		// if !strings.HasPrefix(word, "x") {
		// 	continue
		// }

		// if words.Levenshtein(word, "zierahn") > 3 {
		// 	continue
		// }

		// if words.Levenshtein(word, "zierahn") > 3 {
		// 	continue
		// }

		// if Levenshtein(word, "sidious") > 3 {
		// 	continue
		// }

		// if Levenshtein(word, "skywalker") > 3 {
		// 	continue
		// }

		if !strings.Contains(word, "oo") {
			continue
		}

		if newWords.Has(word) {
			continue
		}

		newWords.Insert(word)

		fmt.Println(word)
		count++

		if count >= genWords {
			break
		}
	}
}
