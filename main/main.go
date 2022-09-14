package main

import (
	"fmt"
	"github.com/pzierahn/go_utils/words"
)

func main() {

	count := 0
	genWords := 40

	newWords := make(map[string]bool)

	for {
		//word := words.RandomWord(10)
		word := words.RandomWordMinMax(4, 7)
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

		if words.Levenshtein(word, "word") > 2 {
			continue
		}

		// if !strings.Contains(word, "oo") {
		// 	continue
		// }

		if newWords[word] {
			continue
		}

		newWords[word] = true

		fmt.Println(word)
		count++

		if count >= genWords {
			break
		}
	}
}
