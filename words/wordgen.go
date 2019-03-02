package words

import (
	"math/rand"
	"time"
)

func RandomWordSimple(maxLength int) string {

	vocals := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	consonances := []string{"a", "e", "i", "o", "u"}
	phones := []string{"bu", "ca", "ce", "ci", "co", "cu", "da", "de", "di", "do", "du", "fa", "fe", "fi", "fo", "fu", "ga", "ge", "gi", "go", "gu", "ha", "he", "hi", "ho", "hu", "ja", "je", "ji", "jo", "ju", "ka", "ke", "ki", "ko", "ku", "la", "le", "li", "lo", "lu", "ma", "me", "mi", "mo", "mu", "na", "ne", "ni", "no", "nu", "pa", "pe", "pi", "po", "pu", "qa", "qe", "qi", "qo", "qu", "ra", "re", "ri", "ro", "ru", "sa", "se", "si", "so", "su", "ta", "te", "ti", "to", "tu", "va", "ve", "vi", "vo", "vu", "wa", "we", "wi", "wo", "wu", "xa", "xe", "xi", "xo", "xu", "ya", "ye", "yi", "yo", "yu", "za", "ze", "zi", "zo", "zu"}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	nbr := r1.Intn(maxLength) + 1

	startVowel := r1.Intn(2)
	endConsonant := r1.Intn(2)

	var word = ""

	for inx := 0; inx <= nbr; inx++ {
		word += phones[r1.Int()%len(phones)]
	}

	if startVowel == 1 {
		word = consonances[r1.Int()%len(consonances)] + word
	}

	if endConsonant == 1 {
		word += vocals[r1.Int()%len(vocals)]
	}

	return word
}
