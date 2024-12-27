package main

import (
	"fmt"

	"github.com/fewable/words"
)

func main() {
	// Generate short, medium and long words
	wb := words.NewBuilder().AddShortWord().AddMediumWord().AddLongWord()

	// Add a separator between each word
	wb.WithSeparator("_")

	// Call GetString() to generate a string
	// as many times as you want
	fmt.Println(wb.GetString()) // when_framing_regulation
	fmt.Println(wb.GetString()) // espn_criminal_basically
	fmt.Println(wb.GetString()) // midi_prime_creations

	// Can mutate object later
	wb.WithSeparator(" ")
	wb.AddMediumWord()
	wb.AddShortWord()
	fmt.Println(wb.GetString()) // work mercy appointment declined man

	// OPTIONAL: can use secure randomness for word selection
	wb.WithSecureRandomness()
	fmt.Println(wb.GetString()) // zero evaluate influence boulder char
}
