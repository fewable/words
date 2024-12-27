package words_test

import (
	"reflect"
	"testing"

	"github.com/fewable/words"
)

func testRandomElement(arr []string) string {
	return arr[0]
}

func TestWordsBuilder(t *testing.T) {
	type testCase struct {
		name           string
		wordsToAdd     []int
		separator      string
		expectedWords  []string
		expectedString string
	}

	testCases := []testCase{
		{
			name:           "no words",
			wordsToAdd:     []int{},
			separator:      "",
			expectedWords:  []string{},
			expectedString: "",
		},
		{
			name:           "short",
			wordsToAdd:     []int{1},
			separator:      "",
			expectedWords:  []string{"the"},
			expectedString: "the",
		},
		{
			name:           "medium",
			wordsToAdd:     []int{2},
			separator:      "_",
			expectedWords:  []string{"about"},
			expectedString: "about",
		},
		{
			name:           "short short long medium long medium",
			wordsToAdd:     []int{1, 1, 3, 2, 3, 2},
			separator:      "_",
			expectedWords:  []string{"the", "the", "information", "about", "information", "about"},
			expectedString: "the_the_information_about_information_about",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := words.NewBuilder()

			for _, wordType := range tc.wordsToAdd {
				switch wordType {
				case 1:
					b.AddShortWord()
				case 2:
					b.AddMediumWord()
				case 3:
					b.AddLongWord()
				}
			}

			b.WithSeparator(tc.separator)
			b.WithCustomRandomness(testRandomElement)

			words := b.GetWords()
			if !reflect.DeepEqual(words, tc.expectedWords) {
				t.Fatalf("expected words array: %v, got: %v", tc.expectedWords, words)
			}

			str := b.GetString()
			if str != tc.expectedString {
				t.Fatalf("expected string: %s, got: %s", tc.expectedString, str)
			}
		})
	}
}

// sanity check: ensure that the secure/unsecure random element functions actually work
func TestWordsBuilderCheckRandomElement(t *testing.T) {
	type testCase struct {
		name       string
		wordsToAdd []int
		separator  string
		useSecure  bool
	}

	testCases := []testCase{
		{
			name:       "default random",
			wordsToAdd: []int{1, 1, 3, 2, 3, 2},
			separator:  "_",
			useSecure:  false,
		},
		{
			name:       "secure random",
			wordsToAdd: []int{1, 1, 3, 2, 3, 2},
			separator:  "_",
			useSecure:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := words.NewBuilder()

			for _, wordType := range tc.wordsToAdd {
				switch wordType {
				case 1:
					b.AddShortWord()
				case 2:
					b.AddMediumWord()
				case 3:
					b.AddLongWord()
				}
			}

			b.WithSeparator(tc.separator)
			if tc.useSecure {
				b.WithSecureRandomness()
			}

			if len(b.GetWords()) != len(tc.wordsToAdd) {
				t.Fatalf("expected words array length: %d, got: %d", len(tc.wordsToAdd), len(b.GetWords()))
			}
		})
	}
}
