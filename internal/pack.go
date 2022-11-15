package internal

// Take a word, and pack it into a bit field
// for all the unique, lowercase letters in the
// word
func pack(word string) [4]byte {
	var out int32 = 0

	for _, letter := range word {
		// discard out of range characters
		if letter < 97 || letter > 122 {
			continue
		}
		out = out | 1<<(letter-97)
	}

	return [4]byte{
		byte(out >> 24),
		byte(out >> 16),
		byte(out >> 8),
		byte(out),
	}
}
