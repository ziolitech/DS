package single

import "unicode/utf8"

// Given a string - find the first unique character
// Approach-1 : 2 pass .. in one pass store the freq and 2nd pass
// - iterate and find the first unique char. (Simple)

// Approach-2 : Keep on assuming the first is the unique element ..until
// it breaks that rule .. on every element addition to map - we check if
// the unicity is still maintained .. if not we keep on jumping char by char
// until we find a unique char .. and update the index of that.. at the end of
// the main loop .. idx would have contained the first unique char .. if no
// unique char is found .. we return -1.

// returns lets say index of first unique char

func findFirstUniqueChar(input string) int {
	if len(input) == 0 {
		return -1
	}

	// idx - represents the index of the first unique character.
	idx := -1
	freqMap := make(map[rune]int)

	for byteIndex, r := range input {
		freqMap[r]++

		if idx == -1 {
			// assuming that it is the first non repeating rune.
			idx = byteIndex
		}

		// Check the rune currently at our candidate byte index (idx), NOT r.
		for idx <= byteIndex {
			// returns the candidate rune starting idx itself.
			// MAIN LINE : converts a char starting index to rune and tells it width also
			// based on width that we find the byteIndex of the next rune
			candidateRune, width := utf8.DecodeRuneInString(input[idx:])
			if freqMap[candidateRune] > 1 {
				// Current candidate is a duplicate; advance idx by its byte width
				// leverage width of the rune to find the byteIndex of the next rune.
				idx += width
			} else {
				// Found a valid unique candidate so far; stop advancing
				break
			}
		}

		// If idx passed the current loop position i, no unique char exists yet
		if idx > byteIndex {
			idx = -1
		}
	}

	return idx
}

type charPos struct {
	r   rune
	idx int
}

func findFirstUniqueChar2(input string) int {
	freqMap := make(map[rune]int)
	var q []charPos

	for i, r := range input {
		freqMap[r]++

		// Only push to queue if it's the first time we see this rune
		if freqMap[r] == 1 {
			q = append(q, charPos{r: r, idx: i})
		}

		// Drop duplicates from the front of the queue
		for len(q) > 0 && freqMap[q[0].r] > 1 {
			q = q[1:]
		}
	}

	if len(q) == 0 {
		return -1
	}
	return q[0].idx
}
