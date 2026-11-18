package main

import "fmt"

func testRuneAndString() {
	s := "Hello, 世界"

	s1 := "G🚀0"

	for i, r := range s1 {
		// r is of type rune (int32)
		fmt.Printf("Byte Index: %d | Char: %c | Rune Value: %d | Type: %T\n", i, r, r, r)
	}

	// BAD: Iterates over individual bytes (garbles '世界')
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
		// if i == len(s)-1 {
		// ATTENTION : THIS IS INVALID .. you cannot modify a string They are immutable.
		// 	s[i] = 'h'
		// }
	}

	// GOOD: Decodes complete runes
	for index, r := range s {
		// index jumps by the byte width of each rune (0, 1, 2, 3, 4, 5, 6, 7, 10)
		fmt.Printf("Byte index %d: %c (Unicode: %U)\n", index, r, r)
	}
}

func main() {
	testRuneAndString()
}
