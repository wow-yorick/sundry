package other

import (
	"fmt"
	"unicode/utf8"
)

func buyao() {
	// s := []int{1, 2, 3}
	// fmt.Println(s) //[1 2 3]

	// copy(s, []int{4, 5, 6, 7, 8, 9})
	// fmt.Println(s) //[4 5 6]
	// bytes := []byte("hello world")
	// fmt.Println(string(bytes))
	// copy(bytes, "ha ha")
	// fmt.Println(string(bytes))

	dt := []byte("   sadfsdf asf asd fsdf sadf asdfadsfa asfd")
	skipSpaces(dt)
}

func skipSpaces(data []byte) {
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSpace(r) {
			break
		}
	}
	fmt.Println(start)
}

func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}
