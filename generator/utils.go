package generator

import (
	"math/rand"
	"strings"
	"time"
)

func randomInt(min, max int, r *rand.Rand) int { //
	if min >= max {
		panic("max must be greater than min")
	}
	return r.Intn(max-min) + min
}

func removeCharacters(str, characters string) string {
	return strings.Map(func(r rune) rune {
		if !strings.ContainsRune(characters, r) {
			return r
		}
		return -1
	}, str)
}

func addAtleastOne(str1, str2 []rune) []rune {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randNum := r.Intn(len(str2) - 1)
	randNumIndex := r.Intn(len(str1))
	str1[randNumIndex] = rune(str2[randNum])
	return str1
}
