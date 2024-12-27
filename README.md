# fewable/words - generate human readable strings from random words in Go

Uses the top 10,000 English words collected from here: https://github.com/first20hours/google-10000-english

Basically this lets you generate cute strings like `kind_anderson_scholarships` or `true hottest instruments maple her` based on your configuration.

Just call `wb.GetString()`!

## Example

```
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
```

## LICENSE

Read this:

```
Data files are derived from the *Google Web Trillion Word Corpus*, as described by [Thorsten Brants and Alex Franz](http://googleresearch.blogspot.com/2006/08/all-our-n-gram-are-belong-to-you.html), and distributed by the [Linguistic Data Consortium](http://www.ldc.upenn.edu/Catalog/CatalogEntry.jsp?catalogId=LDC2006T13). Subsets of this corpus distributed by [Peter Novig](http://norvig.com/ngrams/). Corpus editing and cleanup by Josh Kaufman.

Educational and personal/research use of this data is permitted under the LDC license, Norvig's MIT license for his contributions, and US fair use doctrine. I do not recommend using this data for commercial purposes without licensing it from the Linguistic Data Consortium.
```

Source: https://github.com/first20hours/google-10000-english/blob/master/LICENSE.md
