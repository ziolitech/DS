package misc

import (
	"slices"
	"strings"
)

func findCommonChars(input []string) string {
	if len(input) == 0 {
		return ""
	}

	commonFreqMap := make(map[rune]int, len(input[0]))
	for _, r := range input[0] {
		commonFreqMap[r]++
	}

	for i := 1; i < len(input); i++ {
		if len(commonFreqMap) == 0 {
			return ""
		}

		tempWordMap := make(map[rune]int, len(input[i]))
		for _, r := range input[i] {
			tempWordMap[r]++
		}

		for r, freq := range commonFreqMap {
			if val, ok := tempWordMap[r]; ok {
				if val < freq {
					commonFreqMap[r] = val
				}
			} else {
				delete(commonFreqMap, r)
			}
		}
	}

	return buildString(commonFreqMap)
}

func buildString(freq map[rune]int) string {
	if len(freq) == 0 {
		return ""
	}

	totalLen := 0
	keys := make([]rune, 0, len(freq))
	for r, count := range freq {
		if count > 0 {
			keys = append(keys, r)
			totalLen += count
		}
	}

	// Sort keys so the output string is deterministic across runs
	slices.Sort(keys)

	var sb strings.Builder
	sb.Grow(totalLen)
	for _, r := range keys {
		for range freq[r] {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
