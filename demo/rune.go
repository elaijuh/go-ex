package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "世界"

	fmt.Printf("lenth: %d\n", len(s1))                          // 6, Chinese character occupies 3 bytes
	fmt.Printf("rune length: %d\n", utf8.RuneCountInString(s1)) // 2

	// increase index by number of bytes the character occupies
	for i := 0; i < len(s1); {
		r, size := utf8.DecodeRuneInString(s1[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size // i += 3
	}

	// print single quote character and it's unicode decimal representation
	s2 := "hello, 世界"
	for i, r := range s2 {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// hex of bytes
	fmt.Printf("% x\n", s1)

	// hex of unicode of the character
	fmt.Printf("%x\n", []rune(s1))

}
