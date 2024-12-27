package words

import (
	"crypto/rand"
	_ "embed"
	"math/big"
	mathrand "math/rand"
	"strings"
)

//go:embed google-10000-english-usa-no-swears-short.txt
var shortWordsContent string

//go:embed google-10000-english-usa-no-swears-medium.txt
var mediumWordsContent string

//go:embed google-10000-english-usa-no-swears-long.txt
var longWordsContent string

var shortWords = strings.Split(shortWordsContent, "\n")
var mediumWords = strings.Split(mediumWordsContent, "\n")
var longWords = strings.Split(longWordsContent, "\n")

type RandomElement func([]string) string

func (b *Builder) WithCustomRandomness(fn RandomElement) *Builder {
	b.randomElement = fn
	return b
}

func (b *Builder) WithSecureRandomness() *Builder {
	b.randomElement = secureRandomElement
	return b
}

func unsecureRandomElement(arr []string) string {
	index := mathrand.Intn(len(arr))
	return arr[index]
}

func secureRandomElement(arr []string) string {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(arr))))
	if err != nil {
		panic(err)
	}
	return arr[int(n.Int64())]
}

type Builder struct {
	ordering      []int // 1: short, 2: medium, 3: long
	separator     string
	randomElement RandomElement
}

func NewBuilder() *Builder {
	return &Builder{
		ordering:      []int{},
		randomElement: unsecureRandomElement,
	}
}

func (b *Builder) AddShortWord() *Builder {
	b.ordering = append(b.ordering, 1)
	return b
}

func (b *Builder) AddMediumWord() *Builder {
	b.ordering = append(b.ordering, 2)
	return b
}

func (b *Builder) AddLongWord() *Builder {
	b.ordering = append(b.ordering, 3)
	return b
}

func (b *Builder) WithSeparator(s string) *Builder {
	b.separator = s
	return b
}

func (b *Builder) GetWords() []string {
	words := []string{}

	for _, v := range b.ordering {
		var wordList []string
		switch v {
		case 1:
			wordList = shortWords
		case 2:
			wordList = mediumWords
		case 3:
			wordList = longWords
		}

		word := b.randomElement(wordList)

		words = append(words, word)
	}

	return words
}

func (b *Builder) GetString() string {
	return strings.Join(b.GetWords(), b.separator)
}
