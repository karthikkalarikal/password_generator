package generator

import (
	"math/rand"
	"time"
)

const (
	lowerLetters       = "abcdefghijklmnopqrstuvwxyz"
	upperLetters       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberSet          = "0123456789"
	symbolSet          = "!$%^&*()_+{}:@[];'#<>?,./|\\-=?"
	similarChars       = "il1Lo0O"
	symbolAmbiguousSet = "<>[](){}:;'/|\\,"
)

type Config struct {
	Length          uint
	UpperCase       bool
	Numbers         bool
	Symbols         bool
	SimilarChars    bool
	SymbolAmbigious bool
}
type Generate struct {
	Config *Config
}

func New(config *Config) *Generate {
	return &Generate{Config: config}
}
func (g *Generate) GeneratePassword() string {
	var charset string
	var password []rune
	temp := 0
	charset += lowerLetters
	if g.Config.UpperCase {
		charset += upperLetters
		if g.Config.SimilarChars {
			charset = removeCharacters(charset, similarChars)
		}
	}
	if g.Config.Numbers {
		charset += numberSet
		if g.Config.SimilarChars {
			charset = removeCharacters(charset, similarChars)
		}
	}
	if g.Config.Symbols {
		charset += symbolSet
		if g.Config.SymbolAmbigious {
			charset = removeCharacters(charset, symbolAmbiguousSet)
		}
	}

	for len(password) < int(g.Config.Length) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// max := big.NewInt(int64(len(charset)))
		randNum := r.Intn(len(charset) - 1)
		if temp != randNum {
			password = append(password, []rune(charset)[randNum])
			temp = randNum
		} else {
			continue
		}

	}
	// if g.Config.UpperCase {

	// 	password = addAtleastOne(password, []rune(upperLetters))
	// }
	// if g.Config.Numbers {

	// 	password = addAtleastOne(password, []rune(numberSet))
	// }
	// if g.Config.Symbols {

	// 	password = addAtleastOne(password, []rune(symbolSet))
	// }

	return string(password)
}
