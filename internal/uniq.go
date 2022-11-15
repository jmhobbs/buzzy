package internal

// Get all the unique letters in a word
func Unique(word string) []byte {
	uniqueLetters := make(map[rune]bool)
	for _, letter := range word {
		uniqueLetters[letter] = true
	}

	letters := make([]byte, 0, len(uniqueLetters))
	for k := range uniqueLetters {
		letters = append(letters, byte(k))
	}

	return []byte(letters)
}
