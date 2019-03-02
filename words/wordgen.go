package words

import (
	"math/rand"
	"time"
)

type WordGenerator struct {
	vocals      []string
	consonances []string
	phones      []string
}

func (this WordGenerator) RandomWord(maxLength int) string {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	nbr := r1.Intn(maxLength) + 1

	startVowel := r1.Intn(2)
	endConsonant := r1.Intn(2)

	var word = ""

	for inx := 0; inx <= nbr; inx++ {
		word += this.phones[r1.Int()%len(this.phones)]
	}

	if startVowel == 1 {
		word = this.consonances[r1.Int()%len(this.consonances)] + word
	}

	if endConsonant == 1 {
		word += this.vocals[r1.Int()%len(this.vocals)]
	}

	return word
}

func New() {

	vocals := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	consonances := []string{"a", "e", "i", "o", "u"}

	lenCons := len(consonances)
	phones := make([]string, len(vocals)*lenCons)

	for inx, v := range vocals {
		for iny, c := range consonances {

			index := lenCons*inx + iny
			phones[index] = v + c
		}
	}

	generator := &WordGenerator{
		vocals:      vocals,
		consonances: consonances,
		phones:      phones,
	}

	return generator
}
