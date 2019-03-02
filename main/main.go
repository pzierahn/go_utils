package main

import (
	"../words"
	"fmt"
	"strings"
)

// 	"github.com/SSSaaS/sssa-golang"
// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//
// func RandStringRunes(n int) string {
//
// 	b := make([]rune, n)
//
// 	for inx := range b {
// 		b[ inx ] = letterRunes[rand.Intn(len(letterRunes))]
// 	}
//
// 	return string(b)
// }
//
// func sssaas() {
//
// 	text := RandStringRunes(1024 * 4)
// 	fmt.Println("msg =", text)
//
// 	parts, _ := sssa.Create(3, 9, text)
//
// 	for inx, part := range parts {
// 		fmt.Println("inx=" + strconv.Itoa(inx) + " part=" + part)
// 	}
//
// 	msg, _ := sssa.Combine(parts[:3])
//
// 	fmt.Println("msg =", msg)
// }

func main() {

	count := 0
	genWords := 10

	for {
		word := words.RandomWord(5)

		// if len(word) > 7 {
		// 	continue
		// }

		if !strings.HasPrefix(word, "qu") {
			continue
		}

		// if words.Levenshtein(word, "zierahn") > 3 {
		// 	continue
		// }

		// if Levenshtein(word, "sidious") > 3 {
		// 	continue
		// }

		// if Levenshtein(word, "skywalker") > 3 {
		// 	continue
		// }

		// if !strings.Contains(word, "x") {
		// 	continue
		// }

		fmt.Println(word)
		count++

		if count >= genWords {
			break
		}
	}
}
